#version 3.7;

#include "header.inc"
#include "camera.inc"
#include "finishes.inc"
#include "piece.inc"

object {
    merge {
        object { Rook }
        object {
            Extended
            scale 1.3
            translate <0,-0.2,0>
        }
    }
    texture { WhiteT }
}