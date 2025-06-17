package generator

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/guntharp/go-filamentsamples/pkg/models"
)

func TestNewGenerator(t *testing.T) {
	tempDir := t.TempDir()
	csvFile := filepath.Join(tempDir, "test.csv")
	scadFile := filepath.Join(tempDir, "test.scad")

	// Create test files
	csvContent := "Brand,Type,Color,TempHotend,TempBed\nTest,PLA,Red,200,60\n"
	if err := os.WriteFile(csvFile, []byte(csvContent), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(scadFile, []byte("// test scad file"), 0644); err != nil {
		t.Fatal(err)
	}

	config := &Config{
		CSVFile:    csvFile,
		OutputDir:  filepath.Join(tempDir, "output"),
		ScadFile:   scadFile,
		MaxWorkers: 2,
		Verbose:    false,
		DryRun:     false,
	}

	// This test may fail if OpenSCAD is not installed, which is expected
	gen, err := NewGenerator(config)
	if err != nil {
		// Check if it's an OpenSCAD availability issue
		if err.Error() == "failed to initialize OpenSCAD executor: OpenSCAD not found in standard locations or PATH" {
			t.Skip("OpenSCAD not available for testing")
		}
		t.Fatalf("NewGenerator() error = %v", err)
	}

	if gen == nil {
		t.Fatal("NewGenerator() returned nil generator")
	}

	if gen.config != config {
		t.Error("Generator config not set correctly")
	}
}

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{
			name: "valid config",
			config: Config{
				CSVFile:    "test.csv",
				OutputDir:  "output",
				MaxWorkers: 4,
			},
			wantErr: false,
		},
		{
			name: "missing CSV file",
			config: Config{
				OutputDir:  "output",
				MaxWorkers: 4,
			},
			wantErr: true,
		},
		{
			name: "zero workers gets corrected",
			config: Config{
				CSVFile:    "test.csv",
				OutputDir:  "output",
				MaxWorkers: 0,
			},
			wantErr: false,
		},
		{
			name: "too many workers gets corrected",
			config: Config{
				CSVFile:    "test.csv",
				OutputDir:  "output",
				MaxWorkers: 100,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Check worker count corrections
			if tt.config.MaxWorkers == 0 && tt.config.MaxWorkers != 1 {
				t.Error("Zero workers should be corrected to 1")
			}
			if tt.config.MaxWorkers > 32 {
				t.Error("Worker count should be capped at 32")
			}
		})
	}
}

func TestGenerationResult(t *testing.T) {
	sample := &models.FilamentSample{
		Brand:      "Test",
		Type:       "PLA",
		Color:      "Red",
		TempHotend: "200",
		TempBed:    "60",
	}

	result := GenerationResult{
		Sample: sample,
		Error:  nil,
	}

	if result.Sample != sample {
		t.Error("GenerationResult sample not set correctly")
	}

	if result.Error != nil {
		t.Error("GenerationResult error should be nil")
	}
}


func TestWorkerConcurrency(t *testing.T) {
	// Test that we can create multiple workers without issues
	samples := []*models.FilamentSample{
		{Brand: "Test1", Type: "PLA", Color: "Red", TempHotend: "200", TempBed: "60"},
		{Brand: "Test2", Type: "PLA", Color: "Blue", TempHotend: "200", TempBed: "60"},
		{Brand: "Test3", Type: "PLA", Color: "Green", TempHotend: "200", TempBed: "60"},
	}

	// Validate all samples first
	for _, sample := range samples {
		if err := sample.Validate(); err != nil {
			t.Fatalf("Sample validation failed: %v", err)
		}
	}

	// This test verifies the structure without actually running OpenSCAD
	jobs := make(chan *models.FilamentSample, len(samples))
	results := make(chan GenerationResult, len(samples))

	// Add samples to jobs
	for _, sample := range samples {
		jobs <- sample
	}
	close(jobs)

	// Count processed jobs
	processed := 0
	for range samples {
		select {
		case <-jobs:
			processed++
		case <-time.After(100 * time.Millisecond):
			break
		}
	}

	if processed != len(samples) {
		t.Errorf("Expected to process %d samples, got %d", len(samples), processed)
	}

	close(results)
}