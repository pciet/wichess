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
            scale 1.1
            translate <0.1,0,0>
            rotate <0,170,0>
            rotate <0,90,0>
        }
        object { Knight }
    }
    texture { WhiteT }
}
