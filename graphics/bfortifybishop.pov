#version 3.7;

#include "header.inc"
#include "camera.inc"
#include "finishes.inc"
#include "piece.inc"

object {
    merge {
        object { Bishop }
        object {
            Fortify
        }
    }
    texture { BlackT }
}
