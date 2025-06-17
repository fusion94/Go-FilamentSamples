package openscad

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type Executor struct {
	OpenSCADPath string
	ScadFile     string
}

func NewExecutor(scadFile string) (*Executor, error) {
	path, err := findOpenSCADPath()
	if err != nil {
		return nil, err
	}

	return &Executor{
		OpenSCADPath: path,
		ScadFile:     scadFile,
	}, nil
}

func (e *Executor) GenerateSTL(outputPath string, args []string) error {
	cmdArgs := []string{"-o", outputPath}
	cmdArgs = append(cmdArgs, args...)
	cmdArgs = append(cmdArgs, e.ScadFile)

	cmd := exec.Command(e.OpenSCADPath, cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("OpenSCAD execution failed: %w", err)
	}

	return nil
}

func findOpenSCADPath() (string, error) {
	switch runtime.GOOS {
	case "windows":
		path := `C:\Program Files\OpenSCAD\openscad.exe`
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
		return "", fmt.Errorf("OpenSCAD not found at %s", path)

	case "linux":
		path, err := exec.LookPath("openscad")
		if err != nil {
			return "", fmt.Errorf("OpenSCAD not found in PATH")
		}
		return path, nil

	case "darwin":
		paths := []string{
			"/Applications/OpenSCAD.app/Contents/MacOS/OpenSCAD",
			os.ExpandEnv("$HOME/Applications/OpenSCAD.app/Contents/MacOS/OpenSCAD"),
		}

		for _, path := range paths {
			if _, err := os.Stat(path); err == nil {
				return path, nil
			}
		}

		path, err := exec.LookPath("openscad")
		if err != nil {
			return "", fmt.Errorf("OpenSCAD not found in standard locations or PATH")
		}
		return path, nil

	default:
		return "", fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
}

func (e *Executor) CheckAvailable() error {
	if _, err := os.Stat(e.OpenSCADPath); err != nil {
		return fmt.Errorf("OpenSCAD not found at %s: %w", e.OpenSCADPath, err)
	}
	return nil
}

func (e *Executor) GetVersion() (string, error) {
	cmd := exec.Command(e.OpenSCADPath, "--version")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get OpenSCAD version: %w", err)
	}
	return string(output), nil
}