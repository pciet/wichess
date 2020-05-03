#!/bin/bash

# Pick.sh makes the army picker view of pieces.
# The first arg is the piece name, and "short" is
# optionally second for a fast rough render.

# A temporary file pickrendertemp.pov is used.
# The output is pick.png.

if [ $# -eq 0 ]
then
    echo "Pick.sh [piece name] [optional short]"
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

POV="pickrendertemp.pov"

if [[ -f ./$POV ]]
then
    echo "remove $POV to use Pick.sh"
    exit 1
fi

DIM=512

echo '#version 3.7;
#include "pick.inc"
#include "'$1'.inc"
object { '$1' }' > $POV

if [ "$SHORT" = true ]
then
    povray +I$POV +H256 +W256 Quality=5 +FN +Opick.png
else
    povray +I$POV +H$DIM +W$DIM Quality=8 +FN +Opick.png
fi

rm $POV
