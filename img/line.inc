#include "knight.inc"

#declare line = merge {
    object { knight }

    #local i = 0;
    #while (i < 5)
    difference {
        sphere {
            <0,0,0>,0.75
        }
        box {
            <-1,-1,-1>,<1,1,1>
            translate <0,0,1.7>
        }
        scale <1,1,3.8>
        translate <2,0,0>
        rotate <0,0,i*360/5>
        material { TrimMaterial }
    }
    #local i = i + 1;
    #end
    material { PieceMaterial }
}
