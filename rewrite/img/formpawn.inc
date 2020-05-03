#include "pawn.inc"

#declare Prism = box {
    <-1.5,-0.3,5.5999>,<1.5,0.3,6.2>
    translate <0,0,-0.2>
    material { TrimMaterial }
}

// form is a POV-Ray keyword

#declare formpawn = merge {
    object { pawn }
    object {
        Prism
        translate <0,-0.68,0>
    }
    object {
        Prism
        translate <0,0.68,0>
    }
    lathe {
        quadratic_spline
        4,
        <0,0>,<0.1,3>,<0.7,4.5>,<1.5,5.5>
        rotate <90,0,0>
    }
    material { PieceMaterial }
    rotate <0,0,85>
}
