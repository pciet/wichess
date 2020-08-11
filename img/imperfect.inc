#include "pawn.inc"

#declare Box = box {
    <-3,-0.2,-0.0001>,<3,0.2,2>
}

#declare imperfect = merge {
    difference {
        object { pawn }
        box {
            <-1,-1,-1>,<1,1,1>
            scale 2
            rotate <45,0,90>
            translate <2,0,7>
        }
        material { TrimMaterial }
    }
    object {
        Box
        material { TrimMaterial }
        rotate <0,0,-30>
    }
    object {
        Box
        rotate <0,0,60>
        material { TrimMaterial }
    }
    material { PieceMaterial }
}
