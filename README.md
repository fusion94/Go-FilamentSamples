# 3D Print Filament Sample Generator

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

## Running the Application

1. Build the Go Application:

- Open a terminal (or command prompt) in the directory containing `generate_main.go`.
- Run the following command to build the application:
  - `go build -o stl_generator generate_main.go`
- This will create an executable named `stl_generator` (or `stl_generator.exe` on Windows).

2. Run the Application:

- Execute the application, optionally providing a path to the CSV file:
  ``` bash 
  ./stl_generator samples.csv  # On Unix/Linux/macOS
  stl_generator.exe samples.csv  # On Windows
  ```
- If no CSV file is provided, it will default to `samples.csv` in the same
  directory.

### Troubleshooting

- **OpenSCAD Not Found**: If you see an error about OpenSCAD not being found, ensure it is correctly installed and accessible via the command line.
- **CSV Format Errors**: Make sure the CSV file does not contain empty lines or improperly formatted lines, as they will be skipped.
- **Permissions**: Ensure you have the necessary permissions to create directories and
files in the specified output location.

## Ackknowledgements

This was originally a fork of Markus Krause's
[FilamentSamples](https://github.com/markusdd/FilamentSamples) that was written
in python. I forked that repp and made significant changes to it including a
complete rewrite and basic unit tests. You can find that in the
[legacy-python](legacy-python/) directory.

I have since rewitten that completely in Golang.

This was derived this OpenSCAD model from blazerat over at printables:
<https://www.printables.com/de/model/356074-filament-sample-card/files>

## Your Support

If you like what you see, you can leave me a like and a comment here:

https://makerworld.com/en/models/16866#profileId-24328
