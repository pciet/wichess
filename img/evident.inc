#include "pawn.inc"

#declare evident = merge {
    difference {
        object { pawn }
        box {
            <-1,-1,-1>,<1,1,1>
            rotate <45,45,0>
            translate <0,0,7>
        }
        #local i = 0;
        #while (i < 4)
        box {
            <-1,-1,-1><1,1,1>
            rotate <45,45,0>
            translate <2,0,2>
            rotate <0,0,90*i>
        }
        #local i = i + 1;
        #end
    }

    material { PieceMaterial }
}
