#!/usr/bin/env python3

import subprocess
import csv
import os
import platform
import sys

if platform.system() == 'Windows':
    OPENSCAD = 'C:\Program Files\OpenSCAD\openscad.exe'
elif platform.system() == 'Linux':
    OPENSCAD = 'openscad'
else:
    OPENSCAD = 'openscad' #'/Applications/OpenSCAD.app/Contents/MacOS/OpenSCAD'

MYDIR = os.path.dirname(os.path.realpath(__file__))


def gen_samples(file=f"{MYDIR}/samples.csv"):
    os.makedirs(f"{MYDIR}/stl", exist_ok=True)
    with open(file, '+r') as f:
        data = csv.reader(f)

        for l in data:
            if len(l) > 0:
                print("Processing:", l)
                filename = "stl/" + "_".join(l) + ".stl"
                subprocess.run(
                    [OPENSCAD,
                    '-o', filename,
                    '-D', f'BRAND="{l[0]}"',
                    '-D', f'TYPE="{l[1]}"',
                    '-D', f'COLOR="{l[2]}"',
                    '-D', f'TEMP_HOTEND="{l[3]}"',
                    '-D', f'TEMP_BED="{l[4]}"',
                    f'{MYDIR}/FilamentSamples.scad',
                    ], check=True)


if __name__ == "__main__":
    if len(sys.argv) < 1:
        print("You did not pass any .csv file, using samples.csv")
        gen_samples()
    else:
        swatch_file = sys.argv[1]
        print("Using swatch file: " + swatch_file)
        gen_samples(swatch_file)
