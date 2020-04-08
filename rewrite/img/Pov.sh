#!/bin/bash

# Input file is first argument, output is second.
# PNG output.
povray +I$1 +H2048 +W2048 Quality=8 +FN +A +O$2 
