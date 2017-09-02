#include "piece.inc"

#declare Cap = cylinder {
    <0,0.6,0>, <0,0.65,0>, 0.28
}

#declare Pawn = merge {
    object { ShortBody }
    object { Cap }
    sphere {
        0, 0.1
        translate <0,0.65,0>
    }
}

