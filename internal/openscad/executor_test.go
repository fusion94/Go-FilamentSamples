package openscad

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestNewExecutor(t *testing.T) {
	tempDir := t.TempDir()
	scadFile := filepath.Join(tempDir, "test.scad")
	
	if err := os.WriteFile(scadFile, []byte("// test scad file"), 0644); err != nil {
		t.Fatal(err)
	}

	// This test may fail if OpenSCAD is not installed
	executor, err := NewExecutor(scadFile)
	if err != nil {
		// Check if it's an OpenSCAD availability issue
		if runtime.GOOS == "darwin" && err.Error() == "OpenSCAD not found in standard locations or PATH" {
			t.Skip("OpenSCAD not available for testing")
		}
		t.Fatalf("NewExecutor() error = %v", err)
	}

	if executor == nil {
		t.Fatal("NewExecutor() returned nil executor")
	}

	if executor.ScadFile != scadFile {
		t.Errorf("Expected ScadFile %s, got %s", scadFile, executor.ScadFile)
	}

	if executor.OpenSCADPath == "" {
		t.Error("OpenSCADPath should not be empty")
	}
}

func TestFindOpenSCADPath(t *testing.T) {
	path, err := findOpenSCADPath()
	
	// This test is platform dependent
	switch runtime.GOOS {
	case "windows":
		if err != nil {
			// On Windows, we expect a specific error if OpenSCAD is not found
			expectedPath := `C:\Program Files\OpenSCAD\openscad.exe`
			if err.Error() != "OpenSCAD not found at "+expectedPath {
				t.Errorf("Unexpected error on Windows: %v", err)
			}
		} else if path == "" {
			t.Error("Windows path should not be empty if no error")
		}
	case "linux":
		if err != nil && err.Error() != "OpenSCAD not found in PATH" {
			t.Errorf("Unexpected error on Linux: %v", err)
		}
	case "darwin":
		if err != nil && err.Error() != "OpenSCAD not found in standard locations or PATH" {
			t.Errorf("Unexpected error on macOS: %v", err)
		}
	default:
		if err == nil {
			t.Error("Should return error for unsupported platform")
		}
	}
}

func TestExecutor_CheckAvailable(t *testing.T) {
	tempDir := t.TempDir()
	scadFile := filepath.Join(tempDir, "test.scad")
	
	if err := os.WriteFile(scadFile, []byte("// test scad file"), 0644); err != nil {
		t.Fatal(err)
	}

	// Test with invalid OpenSCAD path
	executor := &Executor{
		OpenSCADPath: "/nonexistent/path/openscad",
		ScadFile:     scadFile,
	}

	err := executor.CheckAvailable()
	if err == nil {
		t.Error("CheckAvailable() should return error for nonexistent path")
	}

	// Test with existing file (not actually OpenSCAD, but exists)
	existingFile := filepath.Join(tempDir, "fake_openscad")
	if err := os.WriteFile(existingFile, []byte("fake"), 0755); err != nil {
		t.Fatal(err)
	}

	executor.OpenSCADPath = existingFile
	err = executor.CheckAvailable()
	if err != nil {
		t.Errorf("CheckAvailable() should not return error for existing file: %v", err)
	}
}

func TestExecutor_GenerateSTL(t *testing.T) {
	tempDir := t.TempDir()
	scadFile := filepath.Join(tempDir, "test.scad")
	outputFile := filepath.Join(tempDir, "output.stl")
	
	if err := os.WriteFile(scadFile, []byte("cube([10,10,10]);"), 0644); err != nil {
		t.Fatal(err)
	}

	// Create a fake OpenSCAD executable that just creates an empty output file
	fakeOpenSCAD := filepath.Join(tempDir, "fake_openscad")
	var scriptContent string
	if runtime.GOOS == "windows" {
		fakeOpenSCAD += ".bat"
		scriptContent = "@echo off\necho. > %2\n"
	} else {
		scriptContent = "#!/bin/bash\ntouch \"$2\"\n"
	}
	
	if err := os.WriteFile(fakeOpenSCAD, []byte(scriptContent), 0755); err != nil {
		t.Fatal(err)
	}

	executor := &Executor{
		OpenSCADPath: fakeOpenSCAD,
		ScadFile:     scadFile,
	}

	args := []string{"-D", "TEST=1"}
	err := executor.GenerateSTL(outputFile, args)
	if err != nil {
		t.Errorf("GenerateSTL() error = %v", err)
	}

	// Check if output file was created
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Error("Output STL file was not created")
	}
}

func TestExecutor_GetVersion(t *testing.T) {
	tempDir := t.TempDir()
	scadFile := filepath.Join(tempDir, "test.scad")
	
	if err := os.WriteFile(scadFile, []byte("// test scad file"), 0644); err != nil {
		t.Fatal(err)
	}

	// Create a fake OpenSCAD that returns version info
	fakeOpenSCAD := filepath.Join(tempDir, "fake_openscad")
	var scriptContent string
	if runtime.GOOS == "windows" {
		fakeOpenSCAD += ".bat"
		scriptContent = "@echo off\necho OpenSCAD 2021.01\n"
	} else {
		scriptContent = "#!/bin/bash\necho \"OpenSCAD 2021.01\"\n"
	}
	
	if err := os.WriteFile(fakeOpenSCAD, []byte(scriptContent), 0755); err != nil {
		t.Fatal(err)
	}

	executor := &Executor{
		OpenSCADPath: fakeOpenSCAD,
		ScadFile:     scadFile,
	}

	version, err := executor.GetVersion()
	if err != nil {
		t.Errorf("GetVersion() error = %v", err)
	}

	if version == "" {
		t.Error("GetVersion() returned empty version")
	}

	if runtime.GOOS != "windows" && version != "OpenSCAD 2021.01\n" {
		t.Errorf("Expected version 'OpenSCAD 2021.01\\n', got %q", version)
	}
}

func TestPlatformSpecificPaths(t *testing.T) {
	// Test that we have platform-specific paths defined
	switch runtime.GOOS {
	case "windows":
		// Should check C:\Program Files\OpenSCAD\openscad.exe
	case "linux":
		// Should use PATH lookup
	case "darwin":
		// Should check /Applications and ~/Applications
	default:
		// Should return unsupported platform error
		_, err := findOpenSCADPath()
		if err == nil || err.Error() != "unsupported platform: "+runtime.GOOS {
			t.Errorf("Expected unsupported platform error, got: %v", err)
		}
	}
}