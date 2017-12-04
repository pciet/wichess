#version 3.7;

#include "header.inc"
#include "camera.inc"
#include "finishes.inc"
#include "piece.inc"
#include "knight.inc"

object {
    merge {
        object { 
            Extended
            translate <0,1.45,-0.1>
            scale 0.75
        }
        object { Knight }
    }
    rotate <0,180,0>
    texture { BlackT }
}
