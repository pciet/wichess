#declare Neck = difference {
    box { 0.5, -0.5 }
    torus { 0.9, 0.77 }
    translate <0,0.5,0>
}

#declare Base = cylinder {
    <0,0,0> <0,0.1,0>, 0.3
}

#declare Cap = cylinder {
    <0,0.95,0>, <0,0.97,0>, 0.35
}

#declare Top = merge {
    object { Cap }
    box {
        <0.03, 1, 0.04>
        <-0.03, 1.15, -0.04>
    }
    box {
        <-0.03001, 1.05, -0.15>
        <0.03001, 1.1, 0.15>
    }
    rotate <0,45,0>
}

#declare King = merge {
    object { Top }
    object { Neck }
    object { Base }
}
