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

#declare Steal = merge {
    sphere {
        <0,1.23,0>, 0.07
    }
    box {
        <0.1,1.12,-0.3>, <0.2,1.2,0.3>
    }
    box {
        <-0.1,1.12,-0.3>, <-0.2,1.2,0.3>
    }
    object { LongBody }
    object { Cap }
}
