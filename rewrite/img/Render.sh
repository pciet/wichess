#!/usr/bin/env bash

# For each piece render with the black and white textures on the board and split the image 
# with the cut program into the web/img folder. The look image is also rendered.

# TODO: catch errors, especially from bad povray installation

CUT="cut"

if [[ ! -f ./$CUT ]]
then
    go build -o $CUT cut.go
fi

IMGDIR="../web/img"

mkdir $IMGDIR

PIECES=("pawn" "bishop" "knight" "rook" "queen" "king" "war" "constructive" "formpawn" "confined" "original" "irrelevant")

for p in "${PIECES[@]}"
do
    ./Board.sh $p
    ./$CUT bboard.png $IMGDIR b$p
    ./$CUT wboard.png $IMGDIR w$p
    rm bboard.png wboard.png
    ./Look.sh $p
    mv look.png "$IMGDIR"/look_"$p".png
    ./Pick.sh $p
    mv pick.png "$IMGDIR"/pick_"$p".png
    ./Take.sh $p
    mv btake.png "$IMGDIR"/take_b"$p".png
    mv wtake.png "$IMGDIR"/take_w"$p".png
done

./Board.sh empty
./$CUT bboard.png $IMGDIR empty
rm bboard.png wboard.png

rm $CUT
