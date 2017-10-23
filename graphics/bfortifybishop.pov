#version 3.7;

#include "header.inc"
#include "camera.inc"
#include "finishes.inc"
#include "piece.inc"

object {
    merge {
        object { Bishop }
        object {
            Fortify
            translate <0,0.3,0>
            scale 1.2
        }
    }
    texture { BlackT }
}
