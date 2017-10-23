#version 3.7;

#include "header.inc"
#include "camera.inc"
#include "finishes.inc"
#include "piece.inc"

object {
    merge {
        object { Rook }
        object {
            Lock
            scale 1.3
            rotate <0,90,0>
        }
    }
    texture { BlackT }
}
