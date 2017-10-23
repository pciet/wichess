#version 3.7;

#include "header.inc"
#include "camera.inc"
#include "finishes.inc"
#include "piece.inc"
#include "knight.inc"

object {
    merge {
        object { 
            Detonate
            scale 1.2
            rotate <0,-20,0>
            translate <0,-0.15,-0.15>
        }
        object { Knight }
    }
    rotate <0,180,0>
    texture { BlackT }
}
