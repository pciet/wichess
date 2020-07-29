#include "rook.inc"

#declare warprook = difference {
    object { rook }

    sor {
        4,
        <5,3>,<1.1,3.5>,<1.5,6>,<1.2,10>
        rotate <90,0,0>
    }
    
    torus {
        1.8,0.1
        rotate <90,0,0>
        translate <0,0,4.9>
    }

    scale <1,0.8,1>
    rotate <0,0,20>

    material { PieceMaterial }
}
