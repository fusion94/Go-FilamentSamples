package models

import (
	"testing"
)

func BenchmarkFilamentSample_Validate(b *testing.B) {
	sample := FilamentSample{
		Brand:      "Test Brand",
		Type:       "PLA",
		Color:      "Red",
		TempHotend: "200-220",
		TempBed:    "60",
		BrandSize:  "12",
		TypeSize:   "8",
		ColorSize:  "10",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sample.Validate()
	}
}

func BenchmarkFilamentSample_ValidateInvalid(b *testing.B) {
	sample := FilamentSample{
		Brand:      "", // Invalid - empty brand
		Type:       "PLA",
		Color:      "Red",
		TempHotend: "200-220",
		TempBed:    "60",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sample.Validate()
	}
}

func BenchmarkFilamentSample_Filename(b *testing.B) {
	sample := FilamentSample{
		Brand:      "Test Brand",
		Type:       "PLA",
		Color:      "Red",
		TempHotend: "200-220",
		TempBed:    "60",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sample.Filename()
	}
}

func BenchmarkFilamentSample_OpenSCADArgs(b *testing.B) {
	sample := FilamentSample{
		Brand:      "Test Brand",
		Type:       "PLA",
		Color:      "Red",
		TempHotend: "200-220",
		TempBed:    "60",
		BrandSize:  "12",
		TypeSize:   "8",
		ColorSize:  "10",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sample.OpenSCADArgs()
	}
}

func BenchmarkFilamentSample_OpenSCADArgsMinimal(b *testing.B) {
	sample := FilamentSample{
		Brand:      "Test Brand",
		Type:       "PLA",
		Color:      "Red",
		TempHotend: "200-220",
		TempBed:    "60",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sample.OpenSCADArgs()
	}
}

func BenchmarkTemperatureValidation_Range(b *testing.B) {
	sample := FilamentSample{}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sample.validateTemperature("200-220")
	}
}

func BenchmarkTemperatureValidation_Single(b *testing.B) {
	sample := FilamentSample{}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sample.validateTemperature("200")
	}
}

func BenchmarkTemperatureValidation_Invalid(b *testing.B) {
	sample := FilamentSample{}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sample.validateTemperature("invalid")
	}
}