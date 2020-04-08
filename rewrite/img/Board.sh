#!/bin/bash

# Board.sh constructs and renders the POV-Ray scene for the
# piece named in the single script argument.
# A temporary file called boardrendertemp.pov is used.
# The output is two image files wboard.png and bboard.png.

if [[ ! -f ./$1.inc ]]
then
    echo "no piece file $1.inc"
    exit 1
fi

POV="boardrendertemp.pov"

if [[ -f ./$POV ]]
then
    echo "remove $POV to use Board.sh"
    exit 1
fi

DIM=2048
PREFIX=""

for p in "White" "Black"
do
    echo '#version 3.7;
#include "board.inc"
#include "materials.inc"
#declare TrimMaterial = '$p'TrimMaterial
#declare PieceMaterial = '$p'Material
#include "'$1'.inc"
Piece('$1')' > $POV

    if [ "$p" == "White" ]
    then
        PREFIX="w"
    else
        PREFIX="b"
    fi
    
    povray +I$POV +H$DIM +W$DIM Quality=8 +FN +A +O"$PREFIX"board.png

    rm $POV
done
