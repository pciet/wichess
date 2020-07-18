#include "pawn.inc"

#declare cross = merge {
    box {
        <-1.1,-0.25,5.5999>,<1.1,0.25,5.78>
    }
    box {
        <-1.1,-0.25,5.5999>,<1.1,0.25,5.781>
        rotate <0,0,90>
    }
    box {
        <-0.1,-0.1,5.5998>,<0.1,0.1,5.70001>
    }
}

#declare war = merge {
    difference {
        object { pawn }
        box {
            <-2,-2,5.8>,<2,2,8>
        }
    }
    object {
        cross
        scale <2.5,2.5,2>
        translate <0,0,-10>
        rotate <0,0,20>
        material { TrimBMaterial }
    }
#local i = 0;
#while (i < 8)
    object {
        cross
        rotate <0,0,40>
        translate <0,0,-1+(0.18*i)>
        scale <1.5,1.5,1>
        scale <1/((i*0.05)+1),1/((i*0.05)+1),1>
        material { TrimMaterial }
    }
#local i = i + 1;
#end
    material { PieceMaterial }
}
