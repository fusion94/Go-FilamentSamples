package generator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/guntharp/go-filamentsamples/pkg/models"
)

func BenchmarkConfig_Validate(b *testing.B) {
	config := Config{
		CSVFile:    "test.csv",
		OutputDir:  "output",
		MaxWorkers: 4,
		Verbose:    false,
		DryRun:     false,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = config.Validate()
	}
}

func BenchmarkGenerationResult_Creation(b *testing.B) {
	sample := &models.FilamentSample{
		Brand:      "Test Brand",
		Type:       "PLA",
		Color:      "Red",
		TempHotend: "200-220",
		TempBed:    "60",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := GenerationResult{
			Sample: sample,
			Error:  nil,
		}
		_ = result
	}
}

func BenchmarkWorkerChannelOperations(b *testing.B) {
	// Benchmark channel operations similar to worker pattern
	samples := make([]*models.FilamentSample, 100)
	for i := range samples {
		samples[i] = &models.FilamentSample{
			Brand:      "Test Brand",
			Type:       "PLA",
			Color:      "Red",
			TempHotend: "200-220",
			TempBed:    "60",
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jobs := make(chan *models.FilamentSample, len(samples))
		results := make(chan GenerationResult, len(samples))

		// Fill jobs channel
		for _, sample := range samples {
			jobs <- sample
		}
		close(jobs)

		// Simulate processing
		go func() {
			for sample := range jobs {
				results <- GenerationResult{
					Sample: sample,
					Error:  nil,
				}
			}
			close(results)
		}()

		// Collect results
		processed := 0
		for range results {
			processed++
		}
	}
}

func BenchmarkSampleProcessing_Sequential(b *testing.B) {
	samples := make([]*models.FilamentSample, 10)
	for i := range samples {
		samples[i] = &models.FilamentSample{
			Brand:      "Test Brand",
			Type:       "PLA",
			Color:      "Red",
			TempHotend: "200-220",
			TempBed:    "60",
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, sample := range samples {
			_ = sample.Filename()
			_ = sample.OpenSCADArgs()
		}
	}
}

func BenchmarkSampleProcessing_Concurrent(b *testing.B) {
	samples := make([]*models.FilamentSample, 10)
	for i := range samples {
		samples[i] = &models.FilamentSample{
			Brand:      "Test Brand",
			Type:       "PLA",
			Color:      "Red",
			TempHotend: "200-220",
			TempBed:    "60",
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jobs := make(chan *models.FilamentSample, len(samples))
		done := make(chan bool)

		// Producer
		go func() {
			for _, sample := range samples {
				jobs <- sample
			}
			close(jobs)
		}()

		// Consumer
		go func() {
			for sample := range jobs {
				_ = sample.Filename()
				_ = sample.OpenSCADArgs()
			}
			done <- true
		}()

		<-done
	}
}

func BenchmarkFileOperations(b *testing.B) {
	tempDir := b.TempDir()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Simulate file operations similar to generator
		outputDir := filepath.Join(tempDir, "output")
		err := os.MkdirAll(outputDir, 0755)
		if err != nil {
			b.Fatal(err)
		}
		
		filename := filepath.Join(outputDir, "test.stl")
		_, err = os.Stat(filename)
		// It's expected that the file doesn't exist
		_ = err
	}
}

func BenchmarkConfig_WorkerValidation(b *testing.B) {
	configs := []Config{
		{CSVFile: "test.csv", MaxWorkers: 0},
		{CSVFile: "test.csv", MaxWorkers: -5},
		{CSVFile: "test.csv", MaxWorkers: 100},
		{CSVFile: "test.csv", MaxWorkers: 4},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, config := range configs {
			configCopy := config // Avoid modifying the original
			_ = configCopy.Validate()
		}
	}
}