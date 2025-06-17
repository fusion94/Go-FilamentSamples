package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name       string
		configPath string
		wantErr    bool
	}{
		{
			name:       "empty config path",
			configPath: "",
			wantErr:    false,
		},
		{
			name:       "nonexistent config file",
			configPath: "/nonexistent/config.json",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := LoadConfig(tt.configPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if config == nil {
					t.Error("LoadConfig() returned nil config")
				}
				if config.MaxWorkers != runtime.NumCPU() {
					t.Errorf("Expected MaxWorkers to be %d, got %d", runtime.NumCPU(), config.MaxWorkers)
				}
			}
		})
	}
}

func TestConfig_LoadFromFile(t *testing.T) {
	tempDir := t.TempDir()
	configFile := filepath.Join(tempDir, "config.json")

	// Test valid config file
	validConfig := Config{
		CSVFile:    "test.csv",
		OutputDir:  "output",
		MaxWorkers: 8,
		Verbose:    true,
	}

	data, err := json.MarshalIndent(validConfig, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(configFile, data, 0644); err != nil {
		t.Fatal(err)
	}

	config := &Config{}
	err = config.loadFromFile(configFile)
	if err != nil {
		t.Errorf("loadFromFile() error = %v", err)
	}

	if config.CSVFile != validConfig.CSVFile {
		t.Errorf("Expected CSVFile %s, got %s", validConfig.CSVFile, config.CSVFile)
	}
	if config.MaxWorkers != validConfig.MaxWorkers {
		t.Errorf("Expected MaxWorkers %d, got %d", validConfig.MaxWorkers, config.MaxWorkers)
	}
	if config.Verbose != validConfig.Verbose {
		t.Errorf("Expected Verbose %v, got %v", validConfig.Verbose, config.Verbose)
	}
}

func TestConfig_LoadFromFile_InvalidJSON(t *testing.T) {
	tempDir := t.TempDir()
	configFile := filepath.Join(tempDir, "invalid.json")

	invalidJSON := `{"csv_file": "test.csv", "max_workers": "invalid"}`
	if err := os.WriteFile(configFile, []byte(invalidJSON), 0644); err != nil {
		t.Fatal(err)
	}

	config := &Config{}
	err := config.loadFromFile(configFile)
	if err == nil {
		t.Error("loadFromFile() should return error for invalid JSON")
	}
}

func TestConfig_SaveToFile(t *testing.T) {
	tempDir := t.TempDir()
	configFile := filepath.Join(tempDir, "config.json")

	config := &Config{
		CSVFile:    "test.csv",
		OutputDir:  "output",
		MaxWorkers: 4,
		Verbose:    false,
		DryRun:     true,
	}

	err := config.SaveToFile(configFile)
	if err != nil {
		t.Errorf("SaveToFile() error = %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		t.Error("Config file was not created")
	}

	// Load and verify content
	loadedConfig := &Config{}
	err = loadedConfig.loadFromFile(configFile)
	if err != nil {
		t.Errorf("Failed to load saved config: %v", err)
	}

	if loadedConfig.CSVFile != config.CSVFile {
		t.Errorf("Expected CSVFile %s, got %s", config.CSVFile, loadedConfig.CSVFile)
	}
	if loadedConfig.MaxWorkers != config.MaxWorkers {
		t.Errorf("Expected MaxWorkers %d, got %d", config.MaxWorkers, loadedConfig.MaxWorkers)
	}
	if loadedConfig.DryRun != config.DryRun {
		t.Errorf("Expected DryRun %v, got %v", config.DryRun, loadedConfig.DryRun)
	}
}

func TestConfig_SaveToFile_CreateDirectory(t *testing.T) {
	tempDir := t.TempDir()
	nestedDir := filepath.Join(tempDir, "nested", "deep", "directory")
	configFile := filepath.Join(nestedDir, "config.json")

	config := &Config{
		CSVFile: "test.csv",
	}

	err := config.SaveToFile(configFile)
	if err != nil {
		t.Errorf("SaveToFile() should create directories: %v", err)
	}

	// Verify directory and file exist
	if _, err := os.Stat(nestedDir); os.IsNotExist(err) {
		t.Error("Nested directory was not created")
	}
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		t.Error("Config file was not created")
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
				MaxWorkers: 4,
			},
			wantErr: false,
		},
		{
			name: "missing CSV file",
			config: Config{
				MaxWorkers: 4,
			},
			wantErr: true,
		},
		{
			name: "zero workers",
			config: Config{
				CSVFile:    "test.csv",
				MaxWorkers: 0,
			},
			wantErr: false,
		},
		{
			name: "negative workers",
			config: Config{
				CSVFile:    "test.csv",
				MaxWorkers: -5,
			},
			wantErr: false,
		},
		{
			name: "too many workers",
			config: Config{
				CSVFile:    "test.csv",
				MaxWorkers: 100,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalWorkers := tt.config.MaxWorkers
			err := tt.config.Validate()
			
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Check worker count corrections
			if originalWorkers < 1 && tt.config.MaxWorkers != 1 {
				t.Errorf("Workers should be corrected to 1, got %d", tt.config.MaxWorkers)
			}
			if originalWorkers > 32 && tt.config.MaxWorkers != 32 {
				t.Errorf("Workers should be capped at 32, got %d", tt.config.MaxWorkers)
			}
		})
	}
}

func TestDefaultConfigPath(t *testing.T) {
	path := DefaultConfigPath()
	if path == "" {
		t.Error("DefaultConfigPath() should not return empty string")
	}

	// Should contain the expected path components
	if !filepath.IsAbs(path) && path != "filament-samples.json" {
		t.Errorf("Expected absolute path or fallback, got %s", path)
	}

	// Test fallback when home directory is not available
	// This is hard to test reliably across platforms
	if filepath.Base(path) != "config.json" && filepath.Base(path) != "filament-samples.json" {
		t.Errorf("Expected config.json or filament-samples.json filename, got %s", filepath.Base(path))
	}
}

func TestExampleConfig(t *testing.T) {
	config := ExampleConfig()
	if config == nil {
		t.Fatal("ExampleConfig() returned nil")
	}

	if config.CSVFile != "samples.csv" {
		t.Errorf("Expected CSVFile 'samples.csv', got %s", config.CSVFile)
	}
	if config.OutputDir != "stl" {
		t.Errorf("Expected OutputDir 'stl', got %s", config.OutputDir)
	}
	if config.ScadFile != "FilamentSamples.scad" {
		t.Errorf("Expected ScadFile 'FilamentSamples.scad', got %s", config.ScadFile)
	}
	if config.MaxWorkers != 4 {
		t.Errorf("Expected MaxWorkers 4, got %d", config.MaxWorkers)
	}
	if config.Verbose != false {
		t.Errorf("Expected Verbose false, got %v", config.Verbose)
	}
	if config.DryRun != false {
		t.Errorf("Expected DryRun false, got %v", config.DryRun)
	}
}

func TestConfig_JSONMarshaling(t *testing.T) {
	original := &Config{
		CSVFile:      "test.csv",
		OutputDir:    "output",
		ScadFile:     "test.scad",
		MaxWorkers:   8,
		Verbose:      true,
		DryRun:       false,
		OpenSCADPath: "/usr/bin/openscad",
	}

	// Marshal to JSON
	data, err := json.Marshal(original)
	if err != nil {
		t.Errorf("JSON Marshal error: %v", err)
	}

	// Unmarshal from JSON
	restored := &Config{}
	err = json.Unmarshal(data, restored)
	if err != nil {
		t.Errorf("JSON Unmarshal error: %v", err)
	}

	// Compare fields
	if restored.CSVFile != original.CSVFile {
		t.Errorf("CSVFile mismatch: expected %s, got %s", original.CSVFile, restored.CSVFile)
	}
	if restored.MaxWorkers != original.MaxWorkers {
		t.Errorf("MaxWorkers mismatch: expected %d, got %d", original.MaxWorkers, restored.MaxWorkers)
	}
	if restored.Verbose != original.Verbose {
		t.Errorf("Verbose mismatch: expected %v, got %v", original.Verbose, restored.Verbose)
	}
	if restored.OpenSCADPath != original.OpenSCADPath {
		t.Errorf("OpenSCADPath mismatch: expected %s, got %s", original.OpenSCADPath, restored.OpenSCADPath)
	}
}