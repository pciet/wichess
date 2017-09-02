#include "piece.inc"

#declare Cap = merge {
    cylinder {
        <0,0.9,0>, <0,1.1,0>, 0.28
    }
    cylinder {
        <0,0.8,0>, <0,0.85,0>, 0.33
    }
    cylinder {
        <0,0.85,0>, <0,0.9,0>, 0.3
    }
}

#declare Top = merge {
    object { Cap }
    #local i = 0;
    #while (i < 30)
        sphere {
            0, 0.03
            translate <0.26, 1.11, 0>
            rotate y*360/30*i
        }
        #local i = i + 1;
    #end
}

#declare Queen = merge {
    sphere {
        <0,1.13,0>, 0.07
    }
    object { LongBody }
    object { Top }
}
