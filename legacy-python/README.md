# 3D Print Filament Sample Generator

|        Samples        |   Samples Box (by Zahg)   |
| :-------------------: | :-----------------------: |
| ![](docs/samples.png) | ![](docs/samples_box.png) |

> [!IMPORTANT]
> This was the original code written in Python

We've all encountered dozens of 3D print filament samples, but most require you
to either request new ones from the creator or use a specialized commercial CAD
tool to adjust the parameters and change the text.

This generator, however, is built on OpenSCAD, a completely free 'programmable'
CAD tool. Now, OpenSCAD can be intimidating for most users since it involves
programming your CAD model. While that's usually not how I work either, for this
type of sample creation, it's the perfect approach.

So, why is this useful? Simple: the included `gen_samples.py` script reads from a
`samples.csv` file where you can list the samples you want. No coding skills
needed!

## How To

- make sure you have OpenSCAD and Python available on your machine (see below
  for OS-specific prerequisite install instructions)
- make sure you have the 'Liberation Sans' Font available, it comes with
  LibreOffice or you can get it from here:
  <https://www.1001freefonts.com/de/liberation-sans.font>
- edit `samples.csv` and put in the filaments you like to generate samples for
- if you are not on Windows or your OpenSCAD install is not in the `C:\Program
Files` standard path, edit `gen_samples.py` and at the top put in the path to
  the executable for the variable `OPENSCAD`, or just uncomment the `openscad`
  entry if it is accessible anywhere on your machine because it is part of your
  PATH variable
- run `gen_samples.py`, if your Python install is in your PATH or `.py` files
  are linked you should even be able to just double-click it, if not get a
  command line shell and run `python gen_samples.py` (on Linux you might have to
  use python3 instead)

## Using VS Code

Another convenient method that works across systems is using VS Code. If you've
cloned this repo and opened it in VS Code, ensure you have the Python extension
installed. From there, you can easily edit the `.csv` file, open the
`gen_samples.py` script, and run it directly using the play button in the top
right corner. Of course, you'll still need to have Python and OpenSCAD
installed, but VS Code should automatically detect Python if it's already set up
on your system.

## Using the stl-Models

The output will be generated under the `stl`-Folder.
If you pull the models into your Slicer make sure to select 0.2mm line
thickness(1), the actual Filament profile(2) you want to print with.
The print time (3) will vary according to your print speed settings for that
filament, but should be in the ballpark of 20-30min.

![slicer settings](docs/slicer_settings.png)

There should be approx 5g of filament being used (1) (depending on density of
course) Also make sure after slicing that the 'thickness staircase'(2) is visible
(especially the single layer all the way to the right), there is no non-solid
infill areas (3) and the letters look proper:

![slicer sliced](docs/slicer_sliced.png)

You can then proceed with printing! Have fun!

## Prerequisite Installation

### Windows

The by far easiest method nowadays is to use `winget`, just get a Powershell or
a cmd-prompt and run:

```sh
winget install OpenSCAD.OpenSCAD
winget install Python.Python.3.11
```

This should work on all recently supported versions of Windows 10 and 11.

Alternatively go to the OpenSCAD and Python webpages and download the latest
version from there.

### Linux

Any reasonably recent distribution will have a Python3 version available in the
PATH as `python3`, for very recent ones `python` is also at version 3 already.
You can check with `python --version`.

For OpenSCAD I recommend checking if your standard package manager has it
available (like `apt`, `dnf`, `pacman` etc.). Usually the distros software
centers also now search Flatpak or AppImage repositories, so I recommend you
just search for OpenSCAD there and fetch it this way.

### MacOS

If you are running homebrew you should be able to get both Python 3 and OpenSCAD
via `brew`, otherwise go to the webpages of OpenSCAD and Python and get the
newest Mac distribution from there.

## Ackknowledgements

I derived this OpenSCAD model from blazerat over at printables:
<https://www.printables.com/de/model/356074-filament-sample-card/files>

I really loved the temperature idea, but I like the form factor and Box better
offered by Seemomster
<https://www.printables.com/de/model/228249-filament-samples-42-materials> and
originally Zahg
(<https://www.printables.com/de/model/16322-filament-sample-with-box/files>)

So I basically turned the samples by blazerat into the form factor of Zahgs
samples and box, with added info, while maintaining the 5-step thickness
staircase.

## Your Support

If you like what you see, you can leave me a like and a comment here:

https://makerworld.com/en/models/16866#profileId-24328
