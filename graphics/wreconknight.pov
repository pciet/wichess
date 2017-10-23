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
            scale 1.1
            rotate <0,-20,0>
        }
        object { Knight }
    }
    texture { WhiteT }
}
