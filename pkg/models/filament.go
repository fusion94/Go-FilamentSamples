package models

import (
	"errors"
	"strconv"
	"strings"
)

type FilamentSample struct {
	Brand       string
	Type        string
	Color       string
	TempHotend  string
	TempBed     string
	BrandSize   string
	TypeSize    string
	ColorSize   string
}

func (f *FilamentSample) Validate() error {
	if f.Brand == "" {
		return errors.New("brand is required")
	}
	if f.Type == "" {
		return errors.New("type is required")
	}
	if f.Color == "" {
		return errors.New("color is required")
	}
	if f.TempHotend == "" {
		return errors.New("hotend temperature is required")
	}
	if f.TempBed == "" {
		return errors.New("bed temperature is required")
	}
	
	if err := f.validateTemperature(f.TempHotend); err != nil {
		return err
	}
	if err := f.validateTemperature(f.TempBed); err != nil {
		return err
	}
	
	return nil
}

func (f *FilamentSample) validateTemperature(temp string) error {
	if strings.Contains(temp, "-") {
		parts := strings.Split(temp, "-")
		if len(parts) != 2 {
			return errors.New("invalid temperature range format")
		}
		
		min, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return errors.New("invalid minimum temperature")
		}
		
		max, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return errors.New("invalid maximum temperature")
		}
		
		if min >= max {
			return errors.New("minimum temperature must be less than maximum")
		}
	} else {
		_, err := strconv.Atoi(strings.TrimSpace(temp))
		if err != nil {
			return errors.New("invalid temperature value")
		}
	}
	
	return nil
}

func (f *FilamentSample) Filename() string {
	parts := []string{f.Brand, f.Type, f.Color, f.TempHotend, f.TempBed}
	return strings.Join(parts, "_") + ".stl"
}

func (f *FilamentSample) OpenSCADArgs() []string {
	args := []string{
		"-D", `BRAND="` + f.Brand + `"`,
		"-D", `TYPE="` + f.Type + `"`,
		"-D", `COLOR="` + f.Color + `"`,
		"-D", `TEMP_HOTEND="` + f.TempHotend + `"`,
		"-D", `TEMP_BED="` + f.TempBed + `"`,
	}
	
	if f.BrandSize != "" {
		args = append(args, "-D", "BRAND_SIZE="+f.BrandSize)
	}
	if f.TypeSize != "" {
		args = append(args, "-D", "TYPE_SIZE="+f.TypeSize)
	}
	if f.ColorSize != "" {
		args = append(args, "-D", "COLOR_SIZE="+f.ColorSize)
	}
	
	return args
}