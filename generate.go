package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func getOpenSCADPath() string {
	switch runtime.GOOS {
	case "windows":
		return `C:\Program Files\OpenSCAD\openscad.exe`
	case "linux":
		return "openscad"
	case "darwin":
		defaultPath := "/Applications/OpenSCAD.app/Contents/MacOS/OpenSCAD"
		if _, err := os.Stat(defaultPath); err == nil {
			return defaultPath
		}
		userPath := os.ExpandEnv("$HOME/Applications/OpenSCAD.app/Contents/MacOS/OpenSCAD")
		if _, err := os.Stat(userPath); err == nil {
			return userPath
		}
		_, err := exec.LookPath("openscad")
		if err != nil {
			fmt.Println("OpenSCAD not found. Please install or add it to your PATH.")
			os.Exit(1)
		}
		return "openscad"
	default:
		fmt.Printf("Unsupported platform: %s\n", runtime.GOOS)
		os.Exit(1)
		return ""
	}
}

func genSamples(filePath string) {
	outputDir := filepath.Join(filepath.Dir(filePath), "stl")
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		fmt.Println("Failed to create output directory:", err)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed to open CSV file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		if len(record) == 0 || strings.HasPrefix(record[0], "#") || strings.TrimSpace(record[0]) == "" {
			continue // Skip comments or empty lines
		}

		filename := filepath.Join(outputDir, strings.Join(record, "_")+".stl")
		args := []string{
			"-o", filename,
			"-D", fmt.Sprintf(`BRAND="%s"`, record[0]),
			"-D", fmt.Sprintf(`TYPE="%s"`, record[1]),
			"-D", fmt.Sprintf(`COLOR="%s"`, record[2]),
			"-D", fmt.Sprintf(`TEMP_HOTEND="%s"`, record[3]),
			"-D", fmt.Sprintf(`TEMP_BED="%s"`, record[4]),
		}

		if len(record) > 5 && record[5] != "" {
			args = append(args, "-D", fmt.Sprintf("BRAND_SIZE=%s", record[5]))
		}
		if len(record) > 6 && record[6] != "" {
			args = append(args, "-D", fmt.Sprintf("TYPE_SIZE=%s", record[6]))
		}
		if len(record) > 7 && record[7] != "" {
			args = append(args, "-D", fmt.Sprintf("COLOR_SIZE=%s", record[7]))
		}

		// Get the OpenSCAD path
		openscadPath := getOpenSCADPath()
		args = append(args, filepath.Join(filepath.Dir(filePath), "FilamentSamples.scad"))
		fmt.Println("Running:", openscadPath, strings.Join(args, " "))

		cmd := exec.Command(openscadPath, args...)
		if err := cmd.Run(); err != nil {
			fmt.Println("Failed to run OpenSCAD:", err)
			return
		}
	}
}

func main() {
	csvFile := "samples.csv"
	if len(os.Args) > 1 {
		csvFile = os.Args[1]
	}
	fmt.Println("Using CSV file:", csvFile)

	genSamples(csvFile)
}
