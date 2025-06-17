package generator

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/guntharp/go-filamentsamples/pkg/models"
)

func TestGenerator_Generate(t *testing.T) {
	tempDir := t.TempDir()
	csvFile := filepath.Join(tempDir, "test.csv")
	outputDir := filepath.Join(tempDir, "output")

	// Create test CSV file with valid temperature ranges
	csvContent := "Brand,Type,Color,TempHotend,TempBed\nTest,PLA,Red,200-220,60-65\n"
	if err := os.WriteFile(csvFile, []byte(csvContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Create mock executor
	mockExecutor := &MockExecutor{
		CheckAvailableFunc: func() error { return nil },
		GetVersionFunc:     func() (string, error) { return "1.0.0", nil },
		GenerateSTLFunc:    func(outputPath string, args []string) error { return nil },
	}

	// Create mock parser
	mockParser := &MockParser{
		ParseFileFunc: func(filename string) ([]*models.FilamentSample, error) {
			return []*models.FilamentSample{
				{
					Brand:      "Test",
					Type:       "PLA",
					Color:      "Red",
					TempHotend: "200-220",
					TempBed:    "60-65",
				},
			}, nil
		},
	}

	// Create generator with mocked components
	gen := &Generator{
		config: &Config{
			CSVFile:    csvFile,
			OutputDir:  outputDir,
			MaxWorkers: 2,
			Verbose:    true,
			DryRun:     false,
		},
		executor: mockExecutor,
		parser:   mockParser,
		logger:   log.New(io.Discard, "", 0),
	}

	// Test successful generation
	err := gen.Generate()
	if err != nil {
		t.Errorf("Generate() error = %v, want nil", err)
	}

	// Verify output directory was created
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		t.Error("Output directory was not created")
	}
}

func TestGenerator_Generate_DryRun(t *testing.T) {
	tempDir := t.TempDir()
	csvFile := filepath.Join(tempDir, "test.csv")
	outputDir := filepath.Join(tempDir, "output")

	// Create test CSV file with valid temperature values
	csvContent := "Brand,Type,Color,TempHotend,TempBed\nTest,PLA,Red,200-220,60-65\nTest,PETG,Blue,240-260,70-75\n"
	if err := os.WriteFile(csvFile, []byte(csvContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Create mock executor
	mockExecutor := &MockExecutor{
		CheckAvailableFunc: func() error { return nil },
		GetVersionFunc:     func() (string, error) { return "1.0.0", nil },
		GenerateSTLFunc: func(outputPath string, args []string) error {
			return errors.New("should not be called in dry run")
		},
	}

	// Create mock parser
	mockParser := &MockParser{
		ParseFileFunc: func(filename string) ([]*models.FilamentSample, error) {
			return []*models.FilamentSample{
				{
					Brand:      "Test",
					Type:       "PLA",
					Color:      "Red",
					TempHotend: "200-220",
					TempBed:    "60-65",
				},
				{
					Brand:      "Test",
					Type:       "PETG",
					Color:      "Blue",
					TempHotend: "240-260",
					TempBed:    "70-75",
				},
			}, nil
		},
	}

	// Create generator with dry run enabled
	gen := &Generator{
		config: &Config{
			CSVFile:    csvFile,
			OutputDir:  outputDir,
			MaxWorkers: 2,
			Verbose:    false,
			DryRun:     true,
		},
		executor: mockExecutor,
		parser:   mockParser,
		logger:   log.New(io.Discard, "", 0),
	}

	// Test dry run
	err := gen.Generate()
	if err != nil {
		t.Errorf("Generate() with dry run error = %v, want nil", err)
	}

	// Verify no STL generation was attempted
	if mockExecutor.GetCallCount() > 0 {
		t.Error("GenerateSTL should not be called in dry run mode")
	}
}

func TestGenerator_Generate_Errors(t *testing.T) {
	tempDir := t.TempDir()
	csvFile := filepath.Join(tempDir, "test.csv")

	tests := []struct {
		name        string
		setupFunc   func() *Generator
		wantErrMsg  string
	}{
		{
			name: "OpenSCAD not available",
			setupFunc: func() *Generator {
				return &Generator{
					config: &Config{
						CSVFile:   csvFile,
						OutputDir: tempDir,
					},
					executor: &MockExecutor{
						CheckAvailableFunc: func() error { return errors.New("OpenSCAD not found") },
					},
					parser: &MockParser{
						ParseFileFunc: func(filename string) ([]*models.FilamentSample, error) {
							return nil, nil // Won't be called since executor check fails first
						},
					},
					logger: log.New(io.Discard, "", 0),
				}
			},
			wantErrMsg: "OpenSCAD check failed",
		},
		{
			name: "CSV parse error",
			setupFunc: func() *Generator {
				return &Generator{
					config: &Config{
						CSVFile:   csvFile,
						OutputDir: tempDir,
					},
					executor: &MockExecutor{
						CheckAvailableFunc: func() error { return nil },
					},
					parser: &MockParser{
						ParseFileFunc: func(filename string) ([]*models.FilamentSample, error) {
							return nil, errors.New("invalid CSV format")
						},
					},
					logger: log.New(io.Discard, "", 0),
				}
			},
			wantErrMsg: "failed to parse CSV file",
		},
		{
			name: "output directory creation failure",
			setupFunc: func() *Generator {
				return &Generator{
					config: &Config{
						CSVFile:   csvFile,
						OutputDir: "/invalid/path/that/cannot/be/created",
					},
					executor: &MockExecutor{
						CheckAvailableFunc: func() error { return nil },
					},
					parser: &MockParser{
						ParseFileFunc: func(filename string) ([]*models.FilamentSample, error) {
							return []*models.FilamentSample{
								{
									Brand:      "Test",
									Type:       "PLA",
									Color:      "Red",
									TempHotend: "200-220",
									TempBed:    "60-65",
								},
							}, nil
						},
					},
					logger: log.New(io.Discard, "", 0),
				}
			},
			wantErrMsg: "failed to create output directory",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := tt.setupFunc()
			err := gen.Generate()
			if err == nil {
				t.Error("Expected error, got nil")
			}
			if !contains(err.Error(), tt.wantErrMsg) {
				t.Errorf("Error = %v, want error containing %v", err, tt.wantErrMsg)
			}
		})
	}
}

func TestGenerator_processParallel(t *testing.T) {
	// Create test samples
	samples := createTestSamples(10)

	// Track generation calls
	var generatedCount int32
	mockExecutor := &MockExecutor{
		GenerateSTLFunc: func(outputPath string, args []string) error {
			atomic.AddInt32(&generatedCount, 1)
			time.Sleep(10 * time.Millisecond) // Simulate work
			return nil
		},
	}

	gen := &Generator{
		config: &Config{
			OutputDir:  t.TempDir(),
			MaxWorkers: 3,
			Verbose:    false,
		},
		executor: mockExecutor,
		logger:   log.New(io.Discard, "", 0),
	}

	// Process samples
	err := gen.processParallel(samples)
	if err != nil {
		t.Errorf("processParallel() error = %v", err)
	}

	// Verify all samples were processed
	if int(generatedCount) != len(samples) {
		t.Errorf("Expected %d samples to be processed, got %d", len(samples), generatedCount)
	}
}

func TestGenerator_processParallel_WithErrors(t *testing.T) {
	// Create test samples
	samples := createTestSamples(5)

	// Mock executor that fails for some samples
	var callCount int32
	mockExecutor := &MockExecutor{
		GenerateSTLFunc: func(outputPath string, args []string) error {
			count := atomic.AddInt32(&callCount, 1)
			if count == 2 || count == 4 {
				return fmt.Errorf("generation failed for call %d", count)
			}
			return nil
		},
	}

	gen := &Generator{
		config: &Config{
			OutputDir:  t.TempDir(),
			MaxWorkers: 2,
			Verbose:    false,
		},
		executor: mockExecutor,
		logger:   log.New(io.Discard, "", 0),
	}

	// Process samples
	err := gen.processParallel(samples)
	if err == nil {
		t.Error("Expected error when some samples fail")
	}
	if !contains(err.Error(), "generation completed with 2 errors") {
		t.Errorf("Unexpected error message: %v", err)
	}
}

func TestGenerator_worker(t *testing.T) {
	jobs := make(chan *models.FilamentSample, 3)
	results := make(chan GenerationResult, 3)

	// Add test samples
	samples := createTestSamples(3)
	for _, sample := range samples {
		jobs <- sample
	}
	close(jobs)

	// Create generator with mock
	var processedCount int32
	gen := &Generator{
		config: &Config{
			OutputDir: t.TempDir(),
		},
		executor: &MockExecutor{
			GenerateSTLFunc: func(outputPath string, args []string) error {
				atomic.AddInt32(&processedCount, 1)
				return nil
			},
		},
		logger: log.New(io.Discard, "", 0),
	}

	// Run worker
	var wg sync.WaitGroup
	wg.Add(1)
	go gen.worker(jobs, results, &wg)
	wg.Wait()
	close(results)

	// Verify results
	resultCount := 0
	for result := range results {
		resultCount++
		if result.Error != nil {
			t.Errorf("Unexpected error: %v", result.Error)
		}
	}

	if resultCount != len(samples) {
		t.Errorf("Expected %d results, got %d", len(samples), resultCount)
	}
	if int(processedCount) != len(samples) {
		t.Errorf("Expected %d samples to be processed, got %d", len(samples), processedCount)
	}
}

func TestGenerator_generateSample(t *testing.T) {
	sample := &models.FilamentSample{
		Brand:      "Test Brand",
		Type:       "PLA",
		Color:      "Red",
		TempHotend: "200",
		TempBed:    "60",
	}

	tests := []struct {
		name    string
		verbose bool
		genErr  error
		wantErr bool
	}{
		{
			name:    "successful generation",
			verbose: false,
			genErr:  nil,
			wantErr: false,
		},
		{
			name:    "successful generation verbose",
			verbose: true,
			genErr:  nil,
			wantErr: false,
		},
		{
			name:    "generation failure",
			verbose: false,
			genErr:  errors.New("OpenSCAD failed"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := &Generator{
				config: &Config{
					OutputDir: t.TempDir(),
					Verbose:   tt.verbose,
				},
				executor: &MockExecutor{
					GenerateSTLFunc: func(outputPath string, args []string) error {
						return tt.genErr
					},
				},
				logger: log.New(io.Discard, "", 0),
			}

			err := gen.generateSample(sample)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateSample() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Helper function
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 || s[0:len(substr)] == substr || contains(s[1:], substr))
}