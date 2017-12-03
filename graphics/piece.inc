#declare KingBase = merge {
    cylinder {
        <0,0,0>, <0,0.1,0>, 0.9
    }
    torus {
        0.9, 0.04
        translate <0,0.05,0>
    }
    cylinder {
        <0,0.1,0>, <0,0.2,0>, 0.8
    }
    torus {
        0.8, 0.04
        translate <0,0.15,0>
    }
    cylinder {
        <0,0.2,0>, <0,0.3,0>, 0.7
    }
    torus {
        0.7, 0.04
        translate <0,0.25,0>
    }
    translate <0,-1,0>
}

#declare QueenBase = merge {
    cylinder {
        <0,0,0>, <0,0.1,0>, 0.85
    }
    torus {
        0.85, 0.04
        translate <0,0.05,0>
    }
    cylinder {
        <0,0.1,0>, <0,0.2,0>, 0.75
    }
    torus {
        0.75, 0.04
        translate <0,0.15,0>
    }
    cylinder {
        <0,0.2,0>, <0,0.3,0>, 0.65
    }
    torus {
        0.65, 0.04
        translate <0,0.25,0>
    }
    translate <0,-1,0>
}

#declare SpecialBase = merge {
    cylinder {
        <0,0,0>, <0,0.09,0>, 0.85
    }
    torus {
        0.85, 0.04
        translate <0,0.0435,0>
    }
    cylinder {
        <0,0.09,0>, <0,0.18,0>, 0.75
    }
    torus {
        0.75, 0.035
        translate <0,0.1435,0>
    }
    translate <0,-1,0>
}

#declare PawnBase = merge {
    cylinder {
        <0,0,0>, <0,0.08,0>, 0.8
    }
    torus {
        0.8, 0.035
        translate <0,0.04,0>
    }
    cylinder {
        <0,0.08,0>, <0,0.16,0>, 0.7
    }
    torus {
        0.7, 0.03
        translate <0,0.12,0>
    }
    translate <0,-1,0>
}

#declare KingNeck = cylinder {
    <0,-0.7,0>, <0,0.8,0>, 0.55
}

#declare QueenNeck = cylinder {
    <0,-0.7,0>, <0,0.78,0>, 0.5
}

#declare SpecialNeck = merge {
    torus {
        0.45, 0.04
        translate <0,-0.82,0>
    }
    cylinder {
        <0,-0.82,0>, <0,0.5,0>, 0.45
    }
}

#declare PawnNeck = merge {
    torus {
        0.45, 0.04
        translate <0,-0.87,0>
    }
    cylinder {
        <0,-0.84,0>, <0,0.6,0>, 0.45
    }
}

#declare KingCap = merge {
    cylinder {
        <0,0.8,0>, <0,0.85,0>, 0.6
    }
    torus {
        0.6, 0.02
        translate <0,0.835,0>
    }
    cylinder {
        <0,0.85,0>, <0,0.9,0>, 0.55
    }
    torus {
        0.55, 0.02
        translate <0,0.875,0>
    }
    cylinder {
        <0,0.9,0>, <0,0.95,0>, 0.5
    }
    torus {
        0.5, 0.02
        translate <0,0.935,0>
    }
    merge {
        box {
            <0.03, 0.95, 0.04>
            <-0.03, 1.25, -0.04>
        }
        box {
            <-0.03001, 1.115, -0.15>
            <0.03001, 1.175, 0.15>
        }
        scale 1.8
        translate <0,-0.9,0>
    }
    cylinder {
        <0,0.95,0>, <0,0.98,0>, 0.2
    }
    torus {
        0.2, 0.01
        translate <0,0.965,0>
    }
    rotate <0,45,0>
}

#declare King = merge {
    object { KingBase }
    object { KingNeck }
    object { KingCap }
}

#declare QueenCap = merge {
    cylinder {
        <0,0.78,0>, <0,0.83,0>, 0.58
    }
    torus {
        0.58, 0.02
        translate <0,0.81,0>
    }
    cylinder {
        <0,0.83,0>, <0,0.88,0>, 0.53
    }
    torus {
        0.53, 0.02
        translate <0,0.865,0>
    }
    cylinder {
        <0,0.88,0>, <0,0.93,0>, 0.48
    }
    torus {
        0.48, 0.02
        translate <0,0.9,0>
    }
    cylinder {
        <0,0.93,0>, <0,0.95,0>, 0.2
    }
    torus {
        0.2, 0.01
        translate <0,0.94,0>
    }
    #local i = 0;
    #while (i < 30)
        sphere {
            0, 0.04
            translate <0.42,0.9415,0>
            rotate y*360/30*i
        }
        #local i = i + 1;
    #end
    sphere {
        <0,1,0>, 0.12
    }
}

#declare Queen = merge {
    object { QueenBase }
    object { QueenNeck }
    object { QueenCap }
}

#declare RookCap =  merge {
    cylinder {
        <0,0.5,0>, <0,0.6,0>, 0.5
    }
    torus {
        0.5, 0.03
        translate <0,0.55,0>
    }
    cylinder {
        <0,0.6,0>, <0,0.65,0>, 0.4
    }
    torus {
        0.4, 0.015
        translate <0,0.635,0>
    }
    difference {
        cylinder { <0,0.64,0>, <0,1,0>, 0.45 }
        #local i = 0;
        #while (i < 2)
            box {
                <0.1,0.78,-1>, <-0.1,1.11,1>
                rotate <0,90*i,0>
            }
            #local i = i + 1;
        #end
        cylinder { <0,0.7,0>, <0,1.11,0>, 0.25 }
    }
}

#declare Rook = merge {
    object { SpecialBase }
    object { SpecialNeck }
    object { RookCap }
}

#declare BishopCap = merge {
    cylinder {
        <0,0.5,0>, <0,0.6,0>, 0.5
    }
    torus {
        0.5, 0.03
        translate <0,0.55,0>
    }
    cylinder {
        <0,0.6,0>, <0,0.65,0>, 0.4
    }
    torus {
        0.4, 0.015
        translate <0,0.635,0>
    }
    difference {
        sphere {
            <0,0.73,0> 0.35
            scale <1,1.1,1>
        }
        box {
            <-1,-0.2,-0.025>, <1,0.2,0.025>
            rotate -x*23
            rotate y*90
            translate <-0.15,1,0.15>
        }
    }
    sphere {
        <0,1.15,0>, 0.06
    }
    torus {
        0.06, 0.01
        translate <0,1.15,0>
    }
}

#declare Bishop = merge {
    object { SpecialBase }
    object { SpecialNeck }
    object { BishopCap }
}

#declare PawnCap = merge {
    cylinder {
        <0,0.6,0>, <0,0.68,0>, 0.5
    }
    torus {
        0.5, 0.03
        translate <0,0.64,0>
    }
    sphere {
        <0,0.8,0>, 0.35
    }
}

#declare Pawn = merge {
    object { PawnBase }
    object { PawnNeck }
    object { PawnCap }
}

#declare Swap = merge {
    #local i = 0;
    #while (i < 48)
        box {
            <-0.1,0,-0.1>, <0.1,1.3,0.1>
            rotate <i*11.25,i*5,0>
            translate <0,0.4,0>
        }
        #local i = i + 1;
    #end
}

#declare SwapPawn = merge {
    object { Pawn }
    object {
        Swap
        scale 0.7
        translate <0,-0.5,0>
    }
}

#declare Lock = merge {
    #local i = 0;
    #while (i < 6)
        box {
            <-0.1,-2,-0.1>,<0.1,2,0.1>
            rotate <90,90,0>
            translate <0,i/1.2,0>
        }
        #local i = i + 1;
    #end
    scale 0.4
    translate <0,-0.8,0>
}

#declare LockPawn = merge {
    object { Pawn }
    object { Lock }
}

#declare Recon = merge {
    box {
        <-0.18,-1.2,-0.18>, <0.18,1.2,0.18>
        rotate <90,0,0>
        translate <0.5,5,0>
    }
    box {
        <-0.18,-1.2,-0.18>, <0.2,1.2,0.18>
        rotate <90,0,0>
        translate <-0.5,5,0>
    }
    scale 0.35
    translate <0,-0.78,0>
    rotate <0,15,0>
}

#declare ReconPawn = merge {
    object { Pawn }
    object { Recon }
}

#declare Detonate = merge {
    box {
        <0.3,0.82,0.08>, <-0.3,0.8,-0.08>
        rotate <0,-45,0>
    }
    box {
        <0.3,0.821,0.08>, <-0.3,0.8,-0.08>
        rotate <0,45,0>
    }
}

#declare DetonatePawn = merge {
    object { Pawn }
    object {
        Detonate
        scale 1.1
        translate <0,0.245,0>
    }
}

#declare Ghost = merge {
    #local i = 0;
    #while (i < 3)
        torus {
            0.25, 0.04
            translate <0,i*0.18,0>
        }
        #local i = i + 1;
    #end
}

#declare GhostPawn = merge {
    object { Pawn }
    object { Ghost }
}

#declare Guard = merge {
    #local i = 0;
    #while (i < 4)
        box {
            <0.02,0.7,-0.5>, <-0.02,1.1-(i/100),0.5>
            rotate 45*i
        }
        #local i = i + 2;
    #end
}

#declare GuardPawn = merge {
    object { Pawn }
    object {
        Guard
        translate <0,0.05,0>
    }
}

#declare Rally = merge {
    #local i = 0;
    #while (i < 32)
        box {
            <0.02,0.4,-0.5>, <-0.02,0.6-(i/100),0.5>
            rotate <2*i,11.25*i,0>
        }
        #local i = i + 1;
    #end
}

#declare RallyPawn = merge {
    object { Pawn }
    object {
        Rally
        scale 1.3
    }
}

#declare Fortify = merge {
    #local i = 0;
    #while (i < 4)
        box {
            <0.02,-0.6,-0.02>, <-0.02,0.6-(i/100),0.02>
            rotate <90,45*i,0>
        }
        #local i = i + 1;
    #end
    scale <1.5,0,1.5>
}

#declare FortifyPawn = merge {
    object { Pawn }
    object { Fortify }
}

#declare Extended = merge {
    box {
        <-1,-1,-1>, <1,1,1>
    }
    scale 0.2
}

#declare ExtendedPawn = merge {
    object { Pawn }
    object {
        Extended
        translate <0,1,0>
    }
}
