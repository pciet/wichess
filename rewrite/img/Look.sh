#!/bin/bash

#Look.sh constructs and renders the piece looks.
#A temporary file lookrendertemp.pov is used.
#The output is look.png.

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

povray +I$POV +H$DIM +W$DIM Quality=8 +FN +A +Olook.png

rm $POV
