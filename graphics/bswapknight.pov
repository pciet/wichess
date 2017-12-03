#version 3.7;

#include "header.inc"
#include "camera.inc"
#include "finishes.inc"
#include "piece.inc"
#include "knight.inc"

object {
    merge {
        object { 
            Swap
            scale 0.5
            translate <0,-0.3,0>
        }
        object { Knight }
    }
    rotate <0,180,0>
    texture { BlackT }
}
