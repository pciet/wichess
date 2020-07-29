#include "knight.inc"

#declare appropriate = merge {
    object { knight }
    #local i = 0;
    #while (i < 4)
    torus {
        1.7,0.25
        rotate <90,0,0>
        translate <0,0,0.5+i>
        material { TrimMaterial }
    }
    #local i = i + 1;
    #end
    material { PieceMaterial }
}
