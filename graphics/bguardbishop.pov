#version 3.7;

#include "header.inc"
#include "camera.inc"
#include "finishes.inc"
#include "piece.inc"

object {
    merge {
        object { Bishop }
        object {
            Guard
            scale 0.8
            translate <0,0.2,0>
        }
    }
    texture { BlackT }
}
