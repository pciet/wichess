#version 3.7;

#include "header.inc"
#include "camera.inc"
#include "finishes.inc"
#include "piece.inc"

object {
    merge {
        object { Rook }
        object {
            Recon
            translate <0,0.2,0>
            rotate <0,90,0>
        }
    }
    texture { BlackT }
}
