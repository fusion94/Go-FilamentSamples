package models

import (
	"testing"
)

func TestFilamentSample_Validate(t *testing.T) {
	tests := []struct {
		name    string
		sample  FilamentSample
		wantErr bool
	}{
		{
			name: "valid sample",
			sample: FilamentSample{
				Brand:      "Test Brand",
				Type:       "PLA",
				Color:      "Red",
				TempHotend: "200-220",
				TempBed:    "60",
			},
			wantErr: false,
		},
		{
			name: "missing brand",
			sample: FilamentSample{
				Type:       "PLA",
				Color:      "Red",
				TempHotend: "200-220",
				TempBed:    "60",
			},
			wantErr: true,
		},
		{
			name: "invalid temperature range",
			sample: FilamentSample{
				Brand:      "Test Brand",
				Type:       "PLA",
				Color:      "Red",
				TempHotend: "220-200",
				TempBed:    "60",
			},
			wantErr: true,
		},
		{
			name: "invalid temperature format",
			sample: FilamentSample{
				Brand:      "Test Brand",
				Type:       "PLA",
				Color:      "Red",
				TempHotend: "invalid",
				TempBed:    "60",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.sample.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("FilamentSample.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFilamentSample_Filename(t *testing.T) {
	sample := FilamentSample{
		Brand:      "Test Brand",
		Type:       "PLA",
		Color:      "Red",
		TempHotend: "200-220",
		TempBed:    "60",
	}

	expected := "Test Brand_PLA_Red_200-220_60.stl"
	if got := sample.Filename(); got != expected {
		t.Errorf("FilamentSample.Filename() = %v, want %v", got, expected)
	}
}

func TestFilamentSample_OpenSCADArgs(t *testing.T) {
	sample := FilamentSample{
		Brand:      "Test Brand",
		Type:       "PLA",
		Color:      "Red",
		TempHotend: "200-220",
		TempBed:    "60",
		BrandSize:  "12",
	}

	args := sample.OpenSCADArgs()
	
	expectedArgs := []string{
		"-D", `BRAND="Test Brand"`,
		"-D", `TYPE="PLA"`,
		"-D", `COLOR="Red"`,
		"-D", `TEMP_HOTEND="200-220"`,
		"-D", `TEMP_BED="60"`,
		"-D", "BRAND_SIZE=12",
	}

	if len(args) != len(expectedArgs) {
		t.Errorf("Expected %d args, got %d", len(expectedArgs), len(args))
		return
	}

	for i, arg := range args {
		if arg != expectedArgs[i] {
			t.Errorf("Arg %d: expected %q, got %q", i, expectedArgs[i], arg)
		}
	}
}