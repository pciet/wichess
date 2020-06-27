#include "pawn.inc"

#declare Prism = box {
    <-1.5,-0.3,0>,<1.5,0.3,0.6>
    material { TrimMaterial }
}

// form is a POV-Ray keyword

#declare formpawn = merge {
    object { pawn }
    #local i = 0;
    #while (i < 4)
        union {
            object {
                Prism
                translate <0,-0.9,0>
            }
            object {
                Prism
                translate <0,0.9,0>
            }
            scale <1.8,1,1>
            rotate <0,0,i*180/4>
            translate <0,0,0.4+(0.2*i)>
        }
    #local i = i + 1;
    #end
    object {
        Prism
        translate <0,-0.7,5.39999>
    }
    object {
        Prism
        translate <0,0.7,5.39999>
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
