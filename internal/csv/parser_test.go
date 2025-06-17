package csv

import (
	"strings"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name     string
		csvData  string
		wantLen  int
		wantErr  bool
		skipHeader bool
	}{
		{
			name: "valid CSV with header",
			csvData: `Brand,Type,Color,TempHotend,TempBed
Test Brand,PLA,Red,200-220,60
Test Brand,PETG,Blue,240-260,70`,
			wantLen: 2,
			wantErr: false,
			skipHeader: true,
		},
		{
			name: "valid CSV without header",
			csvData: `Test Brand,PLA,Red,200-220,60
Test Brand,PETG,Blue,240-260,70`,
			wantLen: 2,
			wantErr: false,
			skipHeader: false,
		},
		{
			name: "CSV with comments",
			csvData: `# This is a comment
Test Brand,PLA,Red,200-220,60
# Another comment
Test Brand,PETG,Blue,240-260,70`,
			wantLen: 2,
			wantErr: false,
			skipHeader: false,
		},
		{
			name:    "invalid CSV - insufficient columns",
			csvData: `Test Brand,PLA,Red`,
			wantLen: 0,
			wantErr: true,
			skipHeader: false,
		},
		{
			name:    "invalid CSV - bad temperature",
			csvData: `Test Brand,PLA,Red,invalid,60`,
			wantLen: 0,
			wantErr: true,
			skipHeader: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := NewParser()
			parser.SkipHeader = tt.skipHeader
			
			reader := strings.NewReader(tt.csvData)
			samples, err := parser.Parse(reader)
			
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
			if len(samples) != tt.wantLen {
				t.Errorf("Parser.Parse() returned %d samples, want %d", len(samples), tt.wantLen)
			}
		})
	}
}

func TestParser_isHeaderRow(t *testing.T) {
	parser := NewParser()
	
	tests := []struct {
		name   string
		record []string
		want   bool
	}{
		{
			name:   "brand header",
			record: []string{"Brand", "Type", "Color"},
			want:   true,
		},
		{
			name:   "manufacturer header",
			record: []string{"Manufacturer", "Type", "Color"},
			want:   true,
		},
		{
			name:   "data row",
			record: []string{"Test Brand", "PLA", "Red"},
			want:   false,
		},
		{
			name:   "empty record",
			record: []string{},
			want:   false,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parser.isHeaderRow(tt.record); got != tt.want {
				t.Errorf("Parser.isHeaderRow() = %v, want %v", got, tt.want)
			}
		})
	}
}