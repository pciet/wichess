#include "rook.inc"

#declare irrelevant = merge {
    difference {
        object { rook }

        #local i = 0;
        #while (i < 6)
        torus {
            1.6,0.5
            scale <1,1,2>
            rotate <90,0,0>
            translate <0,0,1>
            translate <0,0,0.6*i>
        }
        #local i = i + 1;
        #end
        
        torus {
            0.9,0.3
            rotate <90,0,0>
            scale <1.5,1.5,1>
            translate <0,0,5.8>
        }

        torus {
            1.4,0.3
            rotate <90,0,0>
            scale <1.5,1.5,1>
            translate <0,0,4.9>
        }
    }

    sphere {
        <0,0,5.5>,1
        material { TrimMaterial }
    }

    rotate <0,0,45>

    material { PieceMaterial }
}
