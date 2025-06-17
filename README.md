# 3D Print Filament Sample Generator

[![Go](https://github.com/fusion94/Go-FilamentSamples/actions/workflows/go.yml/badge.svg)](https://github.com/fusion94/Go-FilamentSamples/actions/workflows/go.yml)
[![Spell Checking](https://github.com/fusion94/Go-FilamentSamples/actions/workflows/SpellCheck.yml/badge.svg)](https://github.com/fusion94/Go-FilamentSamples/actions/workflows/SpellCheck.yml)

|        Samples        |   Samples Box (by Zahg)   |
| :-------------------: | :-----------------------: |
| ![](docs/samples.png) | ![](docs/samples_box.png) |

We've all encountered dozens of 3D print filament samples, but most require you
to either request new ones from the creator or use a specialized commercial CAD
tool to adjust the parameters and change the text.

This generator, however, is built on OpenSCAD, a completely free 'programmable'
CAD tool. Now, OpenSCAD can be intimidating for most users since it involves
programming your CAD model. While that's usually not how I work either, for this
type of sample creation, it's the perfect approach.

## Prerequisites

1. **Go Installed:**

Ensure you have Go installed on your machine. You can download it from the
official Go website. Follow the installation instructions for your operating
system.

2. **OpenSCAD Installed:**

The application depends on OpenSCAD to generate STL files. Make sure it is installed and accessible in your system's PATH. You can download OpenSCAD from the OpenSCAD website.

- Depending on your operating system, the installation path will vary:
  - **Windows:** Typically `C:\Program Files\OpenSCAD\openscad.exe`.
  - **Linux:** Can be installed via your package manager (e.g., `apt`, `yum`, etc.).
  - **macOS:** Installed through a direct download or Homebrew (`brew install openscad`).

3. CSV File:

Create a CSV file with the necessary parameters. The format of the CSV file
should include the columns expected by the script, typically:

```
BRAND,TYPE,COLOR,TEMP_HOTEND,TEMP_BED,BRAND_SIZE,TYPE_SIZE,COLOR_SIZE
```

Ensure this file is placed in the same directory as the Go application or
provide its path as a command-line argument.

4. Directory Structure:

The Go application will create an stl directory in the same location as the CSV
file to store the generated STL files.

## Building and Running

### Build the Application

Using Make (recommended):

```bash
make build
```

Or using Go directly:

```bash
go build -o filament-samples ./cmd/filament-samples
```

### Running the Application

Basic usage:

```bash
./filament-samples                    # Uses samples.csv in current directory
./filament-samples -csv myfile.csv    # Use custom CSV file
./filament-samples -help              # Show all options
```

Advanced options:

```bash
# Use custom settings
./filament-samples -csv samples.csv -output ./stl -workers 8 -verbose

# Dry run to see what would be generated
./filament-samples -dry-run

# Show version
./filament-samples -version
```

### Command Line Options

- `-csv string`: Path to CSV file (default: "samples.csv")
- `-output string`: Output directory for STL files (default: "stl/" relative to CSV file)
- `-scad string`: Path to OpenSCAD file (default: "FilamentSamples.scad" relative to CSV file)
- `-workers int`: Maximum concurrent workers (default: number of CPU cores)
- `-verbose`: Enable verbose logging
- `-dry-run`: Show what would be generated without creating files
- `-version`: Show version information
- `-help`: Show help information

## Testing and Development

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Generate detailed coverage report
make coverage-report
```

### Coverage Reports

The project includes automated coverage reporting:

- **HTML Report**: `coverage.html` - Interactive coverage visualization
- **Markdown Report**: `COVERAGE_REPORT.md` - Detailed coverage breakdown
- **Terminal Output**: Real-time coverage percentages during test runs

### Development Commands

```bash
# Build the application
make build

# Run tests with verbose output
make test-verbose

# Clean build artifacts
make clean

# Run linting and formatting
make lint
make fmt
```

### Troubleshooting

- **OpenSCAD Not Found**: If you see an error about OpenSCAD not being found, ensure it is correctly installed and accessible via the command line.
- **CSV Format Errors**: Make sure the CSV file does not contain empty lines or improperly formatted lines, as they will be skipped.
- **Permissions**: Ensure you have the necessary permissions to create directories and
  files in the specified output location.

## Acknowledgements

This was originally a fork of Markus Krause's
[FilamentSamples](https://github.com/markusdd/FilamentSamples) that was written
in python. I forked that repository and made significant changes to it including a
complete rewrite and basic unit tests. You can find that in the
[legacy-python](legacy-python/) directory.

I have since rewritten that completely in Golang.

## Recent Improvements (v2.0)

The application has been significantly improved with the following enhancements:

### Code Quality & Organization

- **Proper Go project structure** with separate packages for different concerns
- **Comprehensive error handling** with proper error propagation
- **Extensive unit test suite** with **80.3% overall coverage** 🟢
- **Structured logging** with configurable verbosity levels
- **Mock interfaces** for external dependencies enabling better testing

### Test Coverage Results

| Package | Coverage | Status |
|---------|----------|--------|
| **pkg/models** | 100.0% | 🟢 Excellent |
| **internal/config** | 96.4% | 🟢 Excellent |
| **internal/generator** | 95.8% | 🟢 Excellent |
| **internal/csv** | 95.5% | 🟢 Excellent |
| **internal/openscad** | 74.4% | 🟡 Good |
| **Overall Project** | **80.3%** | 🟢 **Excellent** |

### Testing Infrastructure

- **Comprehensive unit tests** for all core functionality
- **Integration tests** for CLI workflows
- **Benchmark tests** for performance validation
- **Mock implementations** for OpenSCAD and CSV parsing
- **Edge case coverage** for validation and error handling
- **Concurrent processing tests** with proper synchronization
- **Platform-specific path testing** for cross-platform compatibility

### Performance & Reliability

- **Concurrent STL generation** using worker pools for faster processing
- **CSV validation** with proper field checking and temperature range validation
- **Robust OpenSCAD path detection** across different platforms
- **Better file handling** with proper cleanup and error recovery
- **Dependency injection** for better testability and modularity

### User Experience

- **Rich command-line interface** with comprehensive flag support
- **Dry-run mode** to preview what will be generated
- **Verbose logging** for troubleshooting
- **Progress tracking** for large datasets
- **Improved Makefile** with multiple build targets including coverage reports

### Technical Features

- **Header row detection** in CSV files
- **Comment support** in CSV files (lines starting with #)
- **Configurable worker count** for optimal performance
- **Cross-platform compatibility** improvements
- **Better module naming** following Go conventions
- **Interface-based design** for better testability and extensibility

### Development Tools

- **Automated coverage reporting** with HTML and markdown output
- **Comprehensive Makefile** with build, test, and coverage targets
- **Structured project layout** following Go best practices
- **Proper .gitignore** for Go projects

This was derived this OpenSCAD model from blazerat over at printables:
<https://www.printables.com/de/model/356074-filament-sample-card/files>

## Your Support

If you like what you see, you can leave me a like and a comment here:

https://makerworld.com/en/models/16866#profileId-24328
