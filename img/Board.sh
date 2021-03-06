#!/usr/bin/env bash

# Board.sh constructs and renders the POV-Ray scene of the piece named in the input argument, 
# where the piece is placed on all 64 squares of the chess board defined in board.inc.
# If a second argument "short" is included then a fast render is done.
# Two boards are rendered, one for each the white and black materials.
# A temporary file called boardrendertemp.pov is used.
# The output is two image files wboard.png and bboard.png.

if [ $# -eq 0 ]
then
    echo "Board.sh [piece name] [optional short]"
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

POV="boardrendertemp.pov"

if [[ -f ./$POV ]]
then
    echo "remove $POV to use Board.sh"
    exit 1
fi

DIM=3072
PREFIX=""

for p in "White" "Black"
do
    echo '#version 3.7;
#include "board.inc"
#include "materials.inc"
#declare TrimMaterial = '$p'TrimMaterial
#declare TrimBMaterial = '$p'TrimBMaterial
#declare PieceMaterial = '$p'Material
#declare PlainMaterial = Plain'$p'Material
#include "'$1'.inc"
Piece('$1')' > $POV

    if [ "$p" == "White" ]
    then
        PREFIX="w"
    else
        PREFIX="b"
    fi
    
    if [ "$SHORT" = true ]
    then
        povray +I$POV +H512 +W512 Quality=5 +FN +O"$PREFIX"board.png
    else
        # anti-aliasing, the +A argument, adds significant time to rendering and has been 
        # removed for now to speed this up
        povray +I$POV +H$DIM +W$DIM Quality=8 +FN +O"$PREFIX"board.png
    fi

    rm $POV
done
