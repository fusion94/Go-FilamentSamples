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
    OPENSCAD = '/Applications/OpenSCAD.app/Contents/MacOS/OpenSCAD'

MYDIR = os.path.dirname(os.path.realpath(__file__))


def gen_samples(file=f"{MYDIR}/samples.csv"):
    os.makedirs(f"{MYDIR}/stl", exist_ok=True)
    with open(file, '+r') as f:
        data = csv.reader(f)

        for l in data:
            if len(l) > 0:
                print("Processing:", l)
                if l[0].startswith("#"):
                    print("Line is a comment, skipping...")
                    continue
                if not l[0].strip():
                    print("Line is empty, skipping...")
                    continue
                filename = "stl/" + "_".join(l) + ".stl"
                args = [OPENSCAD]  # process name
                outfile = ['-o', filename]
                brand = ['-D', f'BRAND="{l[0]}"']
                type = ['-D', f'TYPE="{l[1]}"']
                color = ['-D', f'COLOR="{l[2]}"']
                noztemp = ['-D', f'TEMP_HOTEND="{l[3]}"']
                bedtemp = ['-D', f'TEMP_BED="{l[4]}"']
                # extra parameters
                brand_font_size = None
                material_font_size = None
                color_font_size = None
                if len(l) > 5:
                    print("brand size: " + l[5])
                    if l[5]:
                        brand_font_size = ['-D', f'BRAND_SIZE={l[5]}']
                if len(l) > 6:
                    if l[6]:
                        material_font_size = ['-D', f'TYPE_SIZE={l[6]}']
                if len(l) > 7 and l[7]:
                    brand_font_size = ['-D', f'COLOR_SIZE={l[7]}']
                infile = f'{MYDIR}/FilamentSamples.scad'
                args.extend(outfile)
                args.extend(brand)
                args.extend(type)
                args.extend(color)
                args.extend(noztemp)
                args.extend(bedtemp)
                if brand_font_size:
                    print("adding optional parameter brand_size: " + l[5])
                    args.extend(brand_font_size)
                if material_font_size:
                    print("adding optional parameter type_size: " + l[6])
                    args.extend(material_font_size)
                if color_font_size:
                    print("adding optional parameter color_size: " + l[7])
                    args.extend(color_font_size)
                args.append(infile)  # this MUST be last param
                print("Calling OpenSCAD with params: ", args)
                subprocess.run(args, check=True)


def find_mac_openscad():
    global OPENSCAD
    # on mac verify we have openscad
    if platform.system() == 'Darwin':
        found = False
        if not os.path.exists(OPENSCAD):
            print("OpenSCAD is not in default /Applications")
            if not os.path.exists(os.path.expanduser('~') + OPENSCAD):
                print("OpenSCAD is not in user's ~/Applications")
                print("Trying to see if there is an openscad command in your path...")
                args = ['openscad', '-v']
                try:
                    proc = subprocess.run(args, capture_output=True, check=True)
                    if proc.returncode == 0:
                        found = True
                        print("openscad version found in path: ", proc.stderr.decode())
                        OPENSCAD = 'openscad'
                except Exception:
                    print("Could not find any openscad command in your path... aborting")
            else:
                print("OpenSCAD found in user's application folder")
                OPENSCAD = os.path.expanduser('~') + OPENSCAD
                found = True
        else:
            print("OpenSCAD found in system application folder")
            found = True
        if not found:
            print("Could not find OpenSCAD binary, please make sure that you have the OpenSCAD app in /applications or "
                  "~/Applications or an 'openscad' command in your path")
            sys.exit(1)


if __name__ == "__main__":
    find_mac_openscad()

    if len(sys.argv) < 2:
        print("You did not pass any .csv file, using samples.csv")
        gen_samples()
    else:
        swatch_file = sys.argv[1]
        print("Using swatch file: " + swatch_file)
        gen_samples(swatch_file)
