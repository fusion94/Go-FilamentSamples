package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/guntharp/go-filamentsamples/pkg/models"
)

type Parser struct {
	SkipHeader bool
}

func NewParser() *Parser {
	return &Parser{
		SkipHeader: true,
	}
}

func (p *Parser) ParseFile(filename string) ([]*models.FilamentSample, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	return p.Parse(file)
}

func (p *Parser) Parse(reader io.Reader) ([]*models.FilamentSample, error) {
	csvReader := csv.NewReader(reader)
	csvReader.Comment = '#'
	csvReader.TrimLeadingSpace = true

	var samples []*models.FilamentSample
	lineNum := 0

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to parse CSV at line %d: %w", lineNum+1, err)
		}

		lineNum++

		if p.SkipHeader && lineNum == 1 {
			if p.isHeaderRow(record) {
				continue
			}
		}

		if len(record) == 0 || strings.TrimSpace(record[0]) == "" {
			continue
		}

		sample, err := p.parseRecord(record, lineNum)
		if err != nil {
			return nil, fmt.Errorf("error at line %d: %w", lineNum, err)
		}

		samples = append(samples, sample)
	}

	return samples, nil
}

func (p *Parser) isHeaderRow(record []string) bool {
	if len(record) == 0 {
		return false
	}
	
	header := strings.ToLower(strings.TrimSpace(record[0]))
	return header == "brand" || header == "manufacturer"
}

func (p *Parser) parseRecord(record []string, lineNum int) (*models.FilamentSample, error) {
	if len(record) < 5 {
		return nil, fmt.Errorf("insufficient columns, expected at least 5, got %d", len(record))
	}

	sample := &models.FilamentSample{
		Brand:      strings.TrimSpace(record[0]),
		Type:       strings.TrimSpace(record[1]),
		Color:      strings.TrimSpace(record[2]),
		TempHotend: strings.TrimSpace(record[3]),
		TempBed:    strings.TrimSpace(record[4]),
	}

	if len(record) > 5 && strings.TrimSpace(record[5]) != "" {
		sample.BrandSize = strings.TrimSpace(record[5])
	}
	if len(record) > 6 && strings.TrimSpace(record[6]) != "" {
		sample.TypeSize = strings.TrimSpace(record[6])
	}
	if len(record) > 7 && strings.TrimSpace(record[7]) != "" {
		sample.ColorSize = strings.TrimSpace(record[7])
	}

	if err := sample.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	return sample, nil
}