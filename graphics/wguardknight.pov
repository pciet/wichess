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
            translate <0,0.2,0>
        }
        object { Knight }
    }
    texture { WhiteT }
}
