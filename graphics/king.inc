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
    box {
        <0.03, 1.1, 0.04>
        <-0.03, 1.25, -0.04>
    }
    box {
        <-0.03001, 1.15, -0.15>
        <0.03001, 1.2, 0.15>
    }
    rotate <0,45,0>
}

#declare King = merge {
    object { Top }
    object { LongBody }
}
