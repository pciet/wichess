#version 3.7;

#include "header.inc"
#include "camera.inc"
#include "finishes.inc"
#include "piece.inc"
#include "knight.inc"

object {
    merge {
        object { 
            Guard
            scale 0.9
        }
        object { Knight }
    }
    rotate <0,180,0>
    texture { BlackT }
}
