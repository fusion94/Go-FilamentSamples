package generator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/guntharp/go-filamentsamples/internal/csv"
	"github.com/guntharp/go-filamentsamples/internal/openscad"
	"github.com/guntharp/go-filamentsamples/pkg/models"
)

type Config struct {
	CSVFile      string
	OutputDir    string
	ScadFile     string
	MaxWorkers   int
	Verbose      bool
	DryRun       bool
}

func (c *Config) Validate() error {
	if c.CSVFile == "" {
		return fmt.Errorf("CSVFile is required")
	}

	if c.MaxWorkers < 1 {
		c.MaxWorkers = 1
	}

	if c.MaxWorkers > 32 {
		c.MaxWorkers = 32
	}

	return nil
}

type Generator struct {
	config   *Config
	executor *openscad.Executor
	parser   *csv.Parser
	logger   *log.Logger
}

type GenerationResult struct {
	Sample *models.FilamentSample
	Error  error
}

func NewGenerator(config *Config) (*Generator, error) {
	scadPath := config.ScadFile
	if scadPath == "" {
		scadPath = filepath.Join(filepath.Dir(config.CSVFile), "FilamentSamples.scad")
	}

	executor, err := openscad.NewExecutor(scadPath)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize OpenSCAD executor: %w", err)
	}

	parser := csv.NewParser()

	logger := log.New(os.Stdout, "", log.LstdFlags)
	if !config.Verbose {
		logger.SetOutput(os.Stderr)
	}

	return &Generator{
		config:   config,
		executor: executor,
		parser:   parser,
		logger:   logger,
	}, nil
}

func (g *Generator) Generate() error {
	if err := g.executor.CheckAvailable(); err != nil {
		return fmt.Errorf("OpenSCAD check failed: %w", err)
	}

	if g.config.Verbose {
		version, _ := g.executor.GetVersion()
		g.logger.Printf("Using OpenSCAD: %s", version)
	}

	samples, err := g.parser.ParseFile(g.config.CSVFile)
	if err != nil {
		return fmt.Errorf("failed to parse CSV file: %w", err)
	}

	g.logger.Printf("Found %d filament samples to process", len(samples))

	if err := os.MkdirAll(g.config.OutputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	if g.config.DryRun {
		g.logger.Println("Dry run mode - no files will be generated")
		for _, sample := range samples {
			g.logger.Printf("Would generate: %s", sample.Filename())
		}
		return nil
	}

	return g.processParallel(samples)
}

func (g *Generator) processParallel(samples []*models.FilamentSample) error {
	maxWorkers := g.config.MaxWorkers
	if maxWorkers <= 0 {
		maxWorkers = 4
	}

	jobs := make(chan *models.FilamentSample, len(samples))
	results := make(chan GenerationResult, len(samples))

	var wg sync.WaitGroup

	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go g.worker(jobs, results, &wg)
	}

	for _, sample := range samples {
		jobs <- sample
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	var errors []error
	processed := 0

	for result := range results {
		processed++
		if result.Error != nil {
			errors = append(errors, result.Error)
			g.logger.Printf("Failed to generate %s: %v", result.Sample.Filename(), result.Error)
		} else if g.config.Verbose {
			g.logger.Printf("Generated %s (%d/%d)", result.Sample.Filename(), processed, len(samples))
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("generation completed with %d errors", len(errors))
	}

	g.logger.Printf("Successfully generated %d STL files", processed)
	return nil
}

func (g *Generator) worker(jobs <-chan *models.FilamentSample, results chan<- GenerationResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for sample := range jobs {
		err := g.generateSample(sample)
		results <- GenerationResult{
			Sample: sample,
			Error:  err,
		}
	}
}

func (g *Generator) generateSample(sample *models.FilamentSample) error {
	outputPath := filepath.Join(g.config.OutputDir, sample.Filename())
	
	args := sample.OpenSCADArgs()
	
	if g.config.Verbose {
		g.logger.Printf("Generating %s", sample.Filename())
	}

	return g.executor.GenerateSTL(outputPath, args)
}