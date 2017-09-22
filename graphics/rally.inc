#include "piece.inc"

#declare Top = difference {
    cylinder { <0,0.8,0>, <0,1.1,0>, 0.35 }
    #local i = 0;
    #while (i < 4)
        box {
            <0.05, 0.9, -1>, <-0.05, 1.15, 1>
            rotate 45*i
        }
        #local i = i + 1;
    #end
    cylinder { <0,0.9,0>, <0,1.2,0>, 0.2 }
}

#declare Rally = merge {
    #local i = 0;
    #while (i < 32)
        box {
            <0.02, 0.4, -0.5>, <-0.02, 0.6-(i/100), 0.5>
            rotate <2*i,11.25*i,0>
        }
        #local i = i + 1;
    #end
    object { LongBody }
    object { Top }
}
