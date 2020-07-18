#!/usr/bin/env bash

if [ $# -eq 0 ]
then
    echo "Take.sh [piece name] [optional short]"
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

POV="takerendertemp.pov"

if [[ -f ./$POV ]]
then
    echo "remove $POV to use Take.sh"
    exit 1
fi

DIM=512
PREFIX=""

for p in "White" "Black"
do
    echo '#version 3.7;
#include "take.inc"
#include "materials.inc"
#declare TrimMaterial = '$p'TrimMaterial
#declare TrimBMaterial = '$p'TrimBMaterial
#declare PieceMaterial = '$p'Material
#declare PlainMaterial = Plain'$p'Material
#include "'$1'.inc"
object { '$1' }' > $POV

    if [ "$p" == "White" ]
    then
        PREFIX="w"
    else
        PREFIX="b"
    fi

    if [ "$SHORT" = true ]
    then
        povray +I$POV +H$DIM +W$DIM Quality=5 +FN +UA +O"$PREFIX"take.png
    else
        povray +I$POV +H$DIM +W$DIM Quality=8 +FN +UA +A +O"$PREFIX"take.png
    fi

    rm $POV
done
