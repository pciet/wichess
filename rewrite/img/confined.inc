#include "pawn.inc"

#declare Box = box {
    <-8,-5,-1>,<8,-1,10>
}
#declare confined = merge {
    difference {
        object { pawn }
        object { Box }
        object {
            Box
            translate <0,6.5,0>
        }
        object {
            Box
            translate <0,4,7>
            material { TrimMaterial }
        }
    }
    rotate <0,0,45>
    material { PieceMaterial }
}
