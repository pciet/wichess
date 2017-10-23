#version 3.7;

#include "header.inc"
#include "camera.inc"
#include "finishes.inc"
#include "piece.inc"

object {
    merge {
        object { Rook }
        object {
            Detonate
            scale 1.6
            translate <0,-0.45,0>
        }
    }
    texture { BlackT }
}
