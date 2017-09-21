#include "piece.inc"

#declare Cap = difference {
    sphere {
        0, 0.3
        scale <1,1.1,1>
        translate <0, 0.9, 0>
    }
    box {
        <-1,-0.2,-0.025>, <1,0.2,0.025>
        rotate -x*30
        rotate y*90
        translate <0, 1.21, 0>
    }
}

#declare Ghost = merge {
    sphere {
        <0,1.23,0>, 0.07
    }
    #local i = 0;
    #while (i < 3)
        torus {
            0.2,0.1
            translate <0,i*0.2,0>
        }
        #local i = i + 1;
    #end
    object { LongBody }
    object { Cap }
}
