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
            scale 1.3
            translate <0,0,0.1>
        }
        object { Knight }
    }
    rotate <0,180,0>
    texture { BlackT }
}
