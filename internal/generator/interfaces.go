package generator

import "github.com/guntharp/go-filamentsamples/pkg/models"

// Executor defines the interface for OpenSCAD operations
type Executor interface {
	GenerateSTL(outputPath string, args []string) error
	CheckAvailable() error
	GetVersion() (string, error)
}

// Parser defines the interface for CSV parsing operations
type Parser interface {
	ParseFile(filename string) ([]*models.FilamentSample, error)
}