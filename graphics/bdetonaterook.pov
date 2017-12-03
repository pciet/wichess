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
            scale 1.5
            translate <0,-0.22,0>
            scale <1.1,0,1.1>
        }
    }
    texture { BlackT }
}
