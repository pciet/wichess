#!/bin/bash

# Assumes povray 3.7 is installed.

# Input file is first argument, output is second. Renders at 256x256 to a png.
povray +I$1 +H256 +W256 Quality=11 +FN +O$2 
