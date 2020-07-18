#include "rook.inc"

#declare impossible = merge {
    difference {
        object { rook }

        #local i = 0;
        #while (i < 4)
        box {
            <-0.5,-0.5,-1>,<0.5,0.5,10>
            translate <2.6,0,0>
            rotate <0,0,i*360/4>
        }
        #local i = i + 1;
        #end

        #local i = 0;
        #while (i < 4)
        box {
            <-0.5,-0.5,4.5>,<0.5,0.5,10>
            translate <1.1,0,0>
            rotate <0,0,45 + (i*360/4)>
        }
        #local i = i + 1;
        #end
    }

    rotate <0,0,30>

    material { PieceMaterial }
}
