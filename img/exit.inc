#include "bishop.inc"

#declare exit = merge {
    object { bishop }
    
    #local i = 0;
    #while (i < 4)
    box {
        <-0.2,-0.2,0>,<0.2,0.2,5>
        translate <2.2,0,0>
        rotate <0,0,i*(360/4)>
        material { TrimMaterial }
    }
    #local i = i + 1;
    #end

    rotate <0,0,45>

    material { PieceMaterial }
}
