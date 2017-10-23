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
            scale 1.3
            translate <-0.2,-0.3,0>
            rotate <0,170,0>
            rotate <0,90,0>
        }
        object { Knight }
    }
    texture { WhiteT }
}
