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
			name: "valid sample with single temperatures",
			sample: FilamentSample{
				Brand:      "Test Brand",
				Type:       "PLA",
				Color:      "Red",
				TempHotend: "200",
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
			name: "missing type",
			sample: FilamentSample{
				Brand:      "Test Brand",
				Color:      "Red",
				TempHotend: "200-220",
				TempBed:    "60",
			},
			wantErr: true,
		},
		{
			name: "missing color",
			sample: FilamentSample{
				Brand:      "Test Brand",
				Type:       "PLA",
				TempHotend: "200-220",
				TempBed:    "60",
			},
			wantErr: true,
		},
		{
			name: "missing hotend temp",
			sample: FilamentSample{
				Brand:   "Test Brand",
				Type:    "PLA",
				Color:   "Red",
				TempBed: "60",
			},
			wantErr: true,
		},
		{
			name: "missing bed temp",
			sample: FilamentSample{
				Brand:      "Test Brand",
				Type:       "PLA",
				Color:      "Red",
				TempHotend: "200-220",
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
		{
			name: "invalid bed temperature",
			sample: FilamentSample{
				Brand:      "Test Brand",
				Type:       "PLA",
				Color:      "Red",
				TempHotend: "200-220",
				TempBed:    "invalid",
			},
			wantErr: true,
		},
		{
			name: "invalid range format - too many parts",
			sample: FilamentSample{
				Brand:      "Test Brand",
				Type:       "PLA",
				Color:      "Red",
				TempHotend: "200-220-240",
				TempBed:    "60",
			},
			wantErr: true,
		},
		{
			name: "invalid range - non-numeric min",
			sample: FilamentSample{
				Brand:      "Test Brand",
				Type:       "PLA",
				Color:      "Red",
				TempHotend: "abc-220",
				TempBed:    "60",
			},
			wantErr: true,
		},
		{
			name: "invalid range - non-numeric max",
			sample: FilamentSample{
				Brand:      "Test Brand",
				Type:       "PLA",
				Color:      "Red",
				TempHotend: "200-xyz",
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
	tests := []struct {
		name         string
		sample       FilamentSample
		expectedArgs []string
	}{
		{
			name: "minimal args",
			sample: FilamentSample{
				Brand:      "Test Brand",
				Type:       "PLA",
				Color:      "Red",
				TempHotend: "200-220",
				TempBed:    "60",
			},
			expectedArgs: []string{
				"-D", `BRAND="Test Brand"`,
				"-D", `TYPE="PLA"`,
				"-D", `COLOR="Red"`,
				"-D", `TEMP_HOTEND="200-220"`,
				"-D", `TEMP_BED="60"`,
			},
		},
		{
			name: "with brand size",
			sample: FilamentSample{
				Brand:      "Test Brand",
				Type:       "PLA",
				Color:      "Red",
				TempHotend: "200-220",
				TempBed:    "60",
				BrandSize:  "12",
			},
			expectedArgs: []string{
				"-D", `BRAND="Test Brand"`,
				"-D", `TYPE="PLA"`,
				"-D", `COLOR="Red"`,
				"-D", `TEMP_HOTEND="200-220"`,
				"-D", `TEMP_BED="60"`,
				"-D", "BRAND_SIZE=12",
			},
		},
		{
			name: "with type size",
			sample: FilamentSample{
				Brand:      "Test Brand",
				Type:       "PLA",
				Color:      "Red",
				TempHotend: "200-220",
				TempBed:    "60",
				TypeSize:   "8",
			},
			expectedArgs: []string{
				"-D", `BRAND="Test Brand"`,
				"-D", `TYPE="PLA"`,
				"-D", `COLOR="Red"`,
				"-D", `TEMP_HOTEND="200-220"`,
				"-D", `TEMP_BED="60"`,
				"-D", "TYPE_SIZE=8",
			},
		},
		{
			name: "with color size",
			sample: FilamentSample{
				Brand:      "Test Brand",
				Type:       "PLA",
				Color:      "Red",
				TempHotend: "200-220",
				TempBed:    "60",
				ColorSize:  "10",
			},
			expectedArgs: []string{
				"-D", `BRAND="Test Brand"`,
				"-D", `TYPE="PLA"`,
				"-D", `COLOR="Red"`,
				"-D", `TEMP_HOTEND="200-220"`,
				"-D", `TEMP_BED="60"`,
				"-D", "COLOR_SIZE=10",
			},
		},
		{
			name: "with all sizes",
			sample: FilamentSample{
				Brand:      "Test Brand",
				Type:       "PLA",
				Color:      "Red",
				TempHotend: "200-220",
				TempBed:    "60",
				BrandSize:  "12",
				TypeSize:   "8",
				ColorSize:  "10",
			},
			expectedArgs: []string{
				"-D", `BRAND="Test Brand"`,
				"-D", `TYPE="PLA"`,
				"-D", `COLOR="Red"`,
				"-D", `TEMP_HOTEND="200-220"`,
				"-D", `TEMP_BED="60"`,
				"-D", "BRAND_SIZE=12",
				"-D", "TYPE_SIZE=8",
				"-D", "COLOR_SIZE=10",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := tt.sample.OpenSCADArgs()

			if len(args) != len(tt.expectedArgs) {
				t.Errorf("Expected %d args, got %d", len(tt.expectedArgs), len(args))
				return
			}

			for i, arg := range args {
				if arg != tt.expectedArgs[i] {
					t.Errorf("Arg %d: expected %q, got %q", i, tt.expectedArgs[i], arg)
				}
			}
		})
	}
}