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

#declare Rook = merge {
    object { LongBody }
    object { Top }
}
