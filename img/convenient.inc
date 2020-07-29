#include "bishop.inc"

#declare convenient = difference {
    merge {
        object { bishop }
        torus {
            2,0.3
            rotate <90,0,0>
            translate <0,0,1.4>
            material { TrimMaterial }
        }
        #local i = 0;
        #while (i < 4)
        torus {
            1,0.2
            rotate <90,0,0>
            translate <0,0,2.3+(i/1.5)>
            material { TrimMaterial }
        }
        #local i = i + 1;
        #end
    }
    torus {
        3,0.4
        rotate <90,0,0>
        translate <0,0,1.4>
    }
    torus {
        3.1,0.3
        rotate <90,0,0>
        translate <0,0,1.1>
    }
    torus {
        0.4,0.15
        rotate <90,0,0>
        translate <0,0,7>
    }
    material { PieceMaterial }
}
