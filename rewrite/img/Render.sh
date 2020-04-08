#!/bin/bash

# For each piece render with the black and white textures
# on the board and split the image with the cut program
# into the web/img folder. The look image is also rendered.

CUT="cut"

if [[ ! -f ./$CUT ]]
then
    go build -o $CUT cut.go
fi

IMGDIR="../web/img"

mkdir $IMGDIR

PIECES=("pawn bishop")

for p in "${PIECES[@]}"
do
    ./Board.sh $p
    ./$CUT bboard.png $IMGDIR b$p
    ./$CUT wboard.png $IMGDIR w$p
    rm bboard.png wboard.png
    ./Look.sh $p
    mv look.png "$IMGDIR"/look_"$p".png
done
