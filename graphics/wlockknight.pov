#version 3.7;

#include "header.inc"
#include "camera.inc"
#include "finishes.inc"
#include "piece.inc"
#include "knight.inc"

object {
    merge {
        object { 
            Lock
            translate <0,0,-0.1>
        }
        object { Knight }
    }
    texture { WhiteT }
}
