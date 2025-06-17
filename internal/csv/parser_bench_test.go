package csv

import (
	"strings"
	"testing"
)

func BenchmarkParser_Parse_SmallDataset(b *testing.B) {
	csvData := `Brand,Type,Color,TempHotend,TempBed
Test Brand,PLA,Red,200-220,60
Test Brand,PETG,Blue,240-260,70
Test Brand,ABS,Green,240-270,80`

	parser := NewParser()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := strings.NewReader(csvData)
		_, err := parser.Parse(reader)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkParser_Parse_LargeDataset(b *testing.B) {
	// Generate a larger CSV dataset
	var csvBuilder strings.Builder
	csvBuilder.WriteString("Brand,Type,Color,TempHotend,TempBed\n")
	
	materials := []string{"PLA", "PETG", "ABS", "TPU", "WOOD", "METAL"}
	colors := []string{"Red", "Blue", "Green", "Yellow", "Black", "White", "Purple", "Orange"}
	brands := []string{"Brand A", "Brand B", "Brand C", "Brand D"}
	
	for _, brand := range brands {
		for _, material := range materials {
			for _, color := range colors {
				csvBuilder.WriteString(brand + "," + material + "," + color + ",200-220,60\n")
			}
		}
	}
	
	csvData := csvBuilder.String()
	parser := NewParser()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := strings.NewReader(csvData)
		_, err := parser.Parse(reader)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkParser_Parse_WithComments(b *testing.B) {
	csvData := `# This is a comment
Brand,Type,Color,TempHotend,TempBed
# Another comment
Test Brand,PLA,Red,200-220,60
# Yet another comment
Test Brand,PETG,Blue,240-260,70
# Final comment
Test Brand,ABS,Green,240-270,80`

	parser := NewParser()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := strings.NewReader(csvData)
		_, err := parser.Parse(reader)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkParser_Parse_HeaderDetection(b *testing.B) {
	csvData := `Brand,Type,Color,TempHotend,TempBed
Test Brand,PLA,Red,200-220,60
Test Brand,PETG,Blue,240-260,70`

	parser := NewParser()
	parser.SkipHeader = true
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := strings.NewReader(csvData)
		_, err := parser.Parse(reader)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkParser_ParseRecord(b *testing.B) {
	parser := NewParser()
	record := []string{"Test Brand", "PLA", "Red", "200-220", "60", "12", "8", "10"}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := parser.parseRecord(record, 1)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkParser_ParseRecord_Minimal(b *testing.B) {
	parser := NewParser()
	record := []string{"Test Brand", "PLA", "Red", "200-220", "60"}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := parser.parseRecord(record, 1)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkParser_IsHeaderRow(b *testing.B) {
	parser := NewParser()
	records := [][]string{
		{"Brand", "Type", "Color", "TempHotend", "TempBed"},
		{"Manufacturer", "Type", "Color", "TempHotend", "TempBed"},
		{"Test Brand", "PLA", "Red", "200-220", "60"},
		{"", "", "", "", ""},
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, record := range records {
			_ = parser.isHeaderRow(record)
		}
	}
}