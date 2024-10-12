#!/usr/bin/env python3

import subprocess
import csv
import os
import platform
import sys

# Set OpenSCAD path based on the OS
def get_openscad_path():
    if platform.system() == 'Windows':
        return r'C:\Program Files\OpenSCAD\openscad.exe'
    elif platform.system() == 'Linux':
        return 'openscad'
    elif platform.system() == 'Darwin':  # macOS
        default_path = '/Applications/OpenSCAD.app/Contents/MacOS/OpenSCAD'
        if os.path.exists(default_path):
            return default_path
        user_path = os.path.expanduser('~/Applications/OpenSCAD.app/Contents/MacOS/OpenSCAD')
        if os.path.exists(user_path):
            return user_path
        try:
            subprocess.run(['openscad', '-v'], capture_output=True, check=True)
            return 'openscad'
        except subprocess.CalledProcessError:
            print("OpenSCAD not found. Please install or add it to your PATH.")
            sys.exit(1)
    else:
        print(f"Unsupported platform: {platform.system()}")
        sys.exit(1)

# Generate STL files based on CSV input
def gen_samples(file_path=f"{os.path.dirname(os.path.realpath(__file__))}/samples.csv"):
    output_dir = f"{os.path.dirname(os.path.realpath(__file__))}/stl"
    os.makedirs(output_dir, exist_ok=True)
    
    with open(file_path, 'r') as f:
        data = csv.reader(f)
        for line in data:
            if not line or line[0].startswith("#") or not line[0].strip():
                continue  # Skip comments or empty lines

            filename = f"{output_dir}/{'_'.join(line)}.stl"
            args = [
                OPENSCAD, '-o', filename, '-D', f'BRAND="{line[0]}"',
                '-D', f'TYPE="{line[1]}"', '-D', f'COLOR="{line[2]}"',
                '-D', f'TEMP_HOTEND="{line[3]}"', '-D', f'TEMP_BED="{line[4]}"'
            ]

            if len(line) > 5 and line[5]:  # Optional font size parameters
                args.extend(['-D', f'BRAND_SIZE={line[5]}'])
            if len(line) > 6 and line[6]:
                args.extend(['-D', f'TYPE_SIZE={line[6]}'])
            if len(line) > 7 and line[7]:
                args.extend(['-D', f'COLOR_SIZE={line[7]}'])

            args.append(f'{os.path.dirname(os.path.realpath(__file__))}/FilamentSamples.scad')
            print("Running:", " ".join(args))
            subprocess.run(args, check=True)

if __name__ == "__main__":
    OPENSCAD = get_openscad_path()

    csv_file = sys.argv[1] if len(sys.argv) > 1 else f"{os.path.dirname(os.path.realpath(__file__))}/samples.csv"
    print(f"Using CSV file: {csv_file}")
    gen_samples(csv_file)
