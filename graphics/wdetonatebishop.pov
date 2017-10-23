#version 3.7;

#include "header.inc"
#include "camera.inc"
#include "finishes.inc"
#include "piece.inc"

object {
    merge {
        object { Bishop }
        object {
            Detonate
            translate <0,0.7,0>
            scale 0.8
        }
    }
    texture { WhiteT }
}
