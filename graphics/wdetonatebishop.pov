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
            scale 1.1
            translate <0,0.3,0>
        }
    }
    texture { WhiteT }
}
