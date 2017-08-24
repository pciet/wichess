#declare Neck = difference {
    box { 0.5, -0.5 }
    torus { 0.9, 0.77 }
    translate <0,0.5,0>
}

#declare Cap = difference {
    sphere {
        0, 0.17
        scale <1,1.05,1>
        translate <0, 1, 0>
    }
    box {
        <-1,-0.15,-0.025>, <1,0.15,0.025>
        rotate x*12
        rotate y*90
        translate <0, 1.21, 0.075>
    }
}

#declare Base = cylinder {
    <0,-0.1,0>, <0,0,0>, 0.35
}

#declare Bishop = merge {
    object { Base }
    object { Neck }
    object { Cap }
}
