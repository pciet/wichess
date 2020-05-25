#!/usr/bin/env bash

# Look.sh constructs and renders the piece look image, which is a side view used in the 
# army picker. The first argument is the name of the piece, and an optional "short" second 
# argument will cause a fast render. A temporary file lookrendertemp.pov is used.
# The output is look.png.

if [ $# -eq 0 ]
then
    echo "Look.sh [piece name] [optional short]"
    exit 1
fi

SHORT=false

if [ $# -eq 2 ]
then
    SHORT=true
fi

if [[ ! -f ./$1.inc ]]
then
    echo "no piece file $1.inc"
    exit 1
fi

POV="lookrendertemp.pov"

if [[ -f ./$POV ]]
then
    echo "remove $POV to use Look.sh"
    exit 1
fi

DIM=1024

echo '#version 3.7;
#include "look.inc"
#include "materials.inc"
#declare TrimMaterial = LookTrimMaterial
#declare PieceMaterial = LookMaterial
#include "'$1'.inc"
object { '$1' }' > $POV

if [ "$SHORT" = true ]
then
    povray +I$POV +H512 +W512 Quality=5 +FN +Olook.png
else
    povray +I$POV +H$DIM +W$DIM Quality=8 +FN +Olook.png
fi

rm $POV
