#include "rook.inc"

#declare simple = difference {
    object { rook }

    cylinder {
        <0,0,5.1>,<0,0,10>,2
    }

    scale <1,1,1.1>

    material { PieceMaterial }
}
