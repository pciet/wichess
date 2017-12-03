#version 3.7;

#include "header.inc"
#include "camera.inc"
#include "finishes.inc"
#include "piece.inc"

object {
    merge {
        object { Bishop }
        object {
            Swap
            scale 0.7
            translate <0,-0.5,0>
        }
    }
    texture { WhiteT }
}
