#version 3.7;

#include "header.inc"
#include "camera.inc"
#include "finishes.inc"
#include "piece.inc"

object {
    merge {
        object { Rook }
        object {
            Ghost
            scale 1.6
        }
    }
    texture { WhiteT }
}
