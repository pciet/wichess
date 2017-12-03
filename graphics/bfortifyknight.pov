#version 3.7;

#include "header.inc"
#include "camera.inc"
#include "finishes.inc"
#include "piece.inc"
#include "knight.inc"

object {
    merge {
        object { 
            Fortify
        }
        object { Knight }
    }
    rotate <0,180,0>
    texture { BlackT }
}
