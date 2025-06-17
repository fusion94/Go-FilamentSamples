package generator

import (
	"fmt"
	"sync"

	"github.com/guntharp/go-filamentsamples/pkg/models"
)

// MockExecutor is a mock implementation of the OpenSCAD executor
type MockExecutor struct {
	GenerateSTLFunc   func(outputPath string, args []string) error
	CheckAvailableFunc func() error
	GetVersionFunc     func() (string, error)
	callCount         int
	mu                sync.Mutex
}

func (m *MockExecutor) GenerateSTL(outputPath string, args []string) error {
	m.mu.Lock()
	m.callCount++
	m.mu.Unlock()
	
	if m.GenerateSTLFunc != nil {
		return m.GenerateSTLFunc(outputPath, args)
	}
	return nil
}

func (m *MockExecutor) CheckAvailable() error {
	if m.CheckAvailableFunc != nil {
		return m.CheckAvailableFunc()
	}
	return nil
}

func (m *MockExecutor) GetVersion() (string, error) {
	if m.GetVersionFunc != nil {
		return m.GetVersionFunc()
	}
	return "Mock OpenSCAD 1.0.0", nil
}

func (m *MockExecutor) GetCallCount() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.callCount
}

// MockParser is a mock implementation of the CSV parser
type MockParser struct {
	ParseFileFunc func(filename string) ([]*models.FilamentSample, error)
}

func (m *MockParser) ParseFile(filename string) ([]*models.FilamentSample, error) {
	if m.ParseFileFunc != nil {
		return m.ParseFileFunc(filename)
	}
	return []*models.FilamentSample{}, nil
}

// Helper function to create test samples
func createTestSamples(count int) []*models.FilamentSample {
	samples := make([]*models.FilamentSample, count)
	for i := 0; i < count; i++ {
		samples[i] = &models.FilamentSample{
			Brand:      fmt.Sprintf("Brand%d", i),
			Type:       "PLA",
			Color:      fmt.Sprintf("Color%d", i),
			TempHotend: "200-220",
			TempBed:    "60",
		}
	}
	return samples
}