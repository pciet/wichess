#include "knight.inc"

#declare Box = box {
    <-3,-0.2,-0.0001>,<3,0.2,2>
}

#declare constructive = merge {
    object { knight }
    object {
        Box
        material { TrimMaterial }
    }
    object {
        Box
        rotate <0,0,90>
        material { TrimMaterial }
    }
}
