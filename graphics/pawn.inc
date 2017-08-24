#declare Neck = difference {
    box { 0.4, -0.4 }
    torus { 0.9, 0.75 }
    translate <0, 0.5, 0>
}

#declare Base = cylinder {
    <0,0,0>, <0,0.1,0>, 0.3
}

#declare Cap = cylinder {
    <0,0.9,0>, <0,0.95,0>, 0.3
}

#declare Pawn = merge {
    object { Base }
    object { Neck }
    object { Cap }
    sphere {
        0, 0.1
        translate <0,1,0>
    }
}

