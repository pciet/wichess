#version 3.7;

#include "header.inc"
#include "camera.inc"
#include "finishes.inc"
#include "piece.inc"
#include "knight.inc"

object {
    merge {
        object { 
            Recon
            scale 0.9
            rotate <0,-20,0>
        }
        object { Knight }
    }
    rotate <0,180,0>
    texture { BlackT }
}
