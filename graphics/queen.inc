#declare Neck = difference {
    box { 0.5, -0.5 }
    torus { 0.9, 0.77 }
    translate <0,0.5,0>
}

#declare Base = cylinder {
    <0,0,0>, <0,0.1,0>, 0.3
}

#declare Cap = cylinder {
    <0,0.95,0>, <0,0.97,0>, 0.35
}

#declare Top = merge {
    object { Cap }
    #local i = 0;
    #while (i < 8)
        sphere {
            0, 0.07
            translate <0.17, 1.01, 0>
            rotate y*360/8*i
        }
        #local i = i + 1;
    #end
}

#declare Queen = merge {
    object { Neck }
    object { Base }
    object { Top }
}
