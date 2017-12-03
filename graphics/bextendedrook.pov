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
            translate <0,1.12,0>
            scale 0.9
        }
    }
    texture { BlackT }
}
