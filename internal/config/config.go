package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	CSVFile      string `json:"csv_file"`
	OutputDir    string `json:"output_dir"`
	ScadFile     string `json:"scad_file"`
	MaxWorkers   int    `json:"max_workers"`
	Verbose      bool   `json:"verbose"`
	DryRun       bool   `json:"dry_run"`
	OpenSCADPath string `json:"openscad_path"`
}

func LoadConfig(configPath string) (*Config, error) {
	config := &Config{
		MaxWorkers: runtime.NumCPU(),
	}

	if configPath != "" {
		if err := config.loadFromFile(configPath); err != nil {
			return nil, fmt.Errorf("failed to load config from %s: %w", configPath, err)
		}
	}

	return config, nil
}

func (c *Config) loadFromFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, c)
}

func (c *Config) SaveToFile(path string) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func (c *Config) Validate() error {
	if c.CSVFile == "" {
		return fmt.Errorf("csv_file is required")
	}

	if c.MaxWorkers < 1 {
		c.MaxWorkers = 1
	}

	if c.MaxWorkers > 32 {
		c.MaxWorkers = 32
	}

	return nil
}

func DefaultConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "filament-samples.json"
	}
	return filepath.Join(home, ".config", "filament-samples", "config.json")
}

func ExampleConfig() *Config {
	return &Config{
		CSVFile:    "samples.csv",
		OutputDir:  "stl",
		ScadFile:   "FilamentSamples.scad",
		MaxWorkers: 4,
		Verbose:    false,
		DryRun:     false,
	}
}