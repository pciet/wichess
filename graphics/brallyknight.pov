#version 3.7;

#include "header.inc"
#include "camera.inc"
#include "finishes.inc"
#include "piece.inc"
#include "knight.inc"

object {
    merge {
        object { 
            Rally
            scale 1.1
            translate <0.1,0,0>
            rotate <0,170,0>
            rotate <0,90,0>
        }
        object { Knight }
    }
    rotate <0,180,0>
    texture { BlackT }
}
