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
            translate <0,0.07,0>
        }
    }
    texture { WhiteT }
}
