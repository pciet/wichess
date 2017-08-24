#declare Neck = difference {
    box { 0.5, -0.5 }
    torus { 0.9, 0.77 }
    translate <0, 0.5, 0>
}

#declare Top = difference {
    object { Neck }
    cylinder { <0,0.9,0>, <0,1.1,0>, 0.24 }
    #local i = 0;
    #while (i < 4)
        box {
            <-0.04, 0.9, -1>, <0.04, 1.1, 1>
            rotate 45*i
        }
        #local i = i + 1;
    #end
}

#declare Base = cylinder {
    <0,-0.1,0>, <0,0,0>, 0.35
}

#declare Rook = merge {
    object { Base }
    object { Top }
}
