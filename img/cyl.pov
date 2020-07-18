// This cylinder object shows the maximum size a piece can be.

#include "board.inc"

#declare BoundingCylinder = object {
    merge {
        cylinder {
            <0,0,0>, <0,0,3>, 3
        }
        cylinder {
            <0,0,2.9>, <0,0,5>, 2.2
        }
        cylinder {
            <0,0,4.9>, <0,0,7>, 1.2
        }
        cylinder {
            <0,0,6.9>, <0,0,8>, 0.8
        }
        cylinder {
            <0,0,7.9>, <0,0,8.6>, 0.5
        }
    }
    texture {
        pigment { rgb 1 }
    }
}

Piece(BoundingCylinder)
