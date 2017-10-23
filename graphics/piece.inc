#declare KingBase = merge {
    cylinder {
        <0,0,0>, <0,0.1,0>, 0.9
    }
    cylinder {
        <0,0.1,0>, <0,0.2,0>, 0.8
    }
    translate <0,-1,0>
}

#declare QueenBase = merge {
    cylinder {
        <0,0,0>, <0,0.1,0>, 0.85
    }
    cylinder {
        <0,0.1,0>, <0,0.2,0>, 0.75
    }
    translate <0,-1,0>
}

#declare SpecialBase = merge {
    cylinder {
        <0,0,0>, <0,0.09,0>, 0.75
    }
    cylinder {
        <0,0.09,0>, <0,0.18,0>, 0.65
    }
    translate <0,-1,0>
}

#declare PawnBase = merge {
    cylinder {
        <0,0,0>, <0,0.08,0>, 0.7
    }
    cylinder {
        <0,0.08,0>, <0,0.16,0>, 0.6
    }
    translate <0,-1,0>
}

#declare KingNeck = cylinder {
    <0,-0.8,0>, <0,0.8,0>, 0.25
}

#declare QueenNeck = cylinder {
    <0,-0.8,0>, <0,0.78,0>, 0.25
}

#declare SpecialNeck = cylinder {
    <0,-0.82,0>, <0,0.68,0>, 0.22
}

#declare PawnNeck = cylinder {
    <0,-0.84,0>, <0,0.6,0>, 0.2
}

#declare KingCap = merge {
    torus {
        0.26, 0.04
        translate <0,0.8,0>
    }
    torus {
        0.26, 0.02
        translate <0,0.85,0>
    }
    cylinder {
        <0,1,0>, <0,0.85,0>, 0.28
    }
    box {
        <0.03, 1, 0.04>
        <-0.03, 1.25, -0.04>
    }
    box {
        <-0.03001, 1.125, -0.15>
        <0.03001, 1.175, 0.15>
    }
    rotate <0,45,0>
}

#declare King = merge {
    object { KingBase }
    object { KingNeck }
    object { KingCap }
}

#declare QueenCap = merge {
    torus {
        0.26, 0.04
        translate <0,0.78,0>
    }
    cylinder {
        <0,0.78,0>, <0,0.93,0>, 0.28
    }
    #local i = 0;
    #while (i < 30)
        sphere {
            0, 0.03
            translate <0.26,0.9315,0>
            rotate y*360/30*i
        }
        #local i = i + 1;
    #end
    sphere {
        <0,0.965,0>, 0.07
    }
}

#declare Queen = merge {
    object { QueenBase }
    object { QueenNeck }
    object { QueenCap }
}

#declare RookCap = difference {
    cylinder { <0,0.68,0>, <0,0.98,0>, 0.35 }
    #local i = 0;
    #while (i < 4)
        box {
            <0.05,0.78,-1>, <-0.05,1.03,1>
            rotate 45*i
        }
        #local i = i + 1;
    #end
    cylinder { <0,0.78,0>, <0,1.08,0>, 0.2 }
}

#declare Rook = merge {
    object { SpecialBase }
    object { SpecialNeck }
    object { RookCap }
}

#declare BishopCap = merge {
    torus {
        0.23, 0.04
        translate <0,0.68,0>
    }
    difference {
        sphere {
            <0,0.83,0> 0.3
            scale <1,1.1,1>
        }
        box {
            <-1,-0.2,-0.025>, <1,0.2,0.025>
            rotate -x*30
            rotate y*90
            translate <-0.2,1.25,0.2>
        }
    }
    sphere {
        <0,1.3,0>, 0.07
    }
}

#declare Bishop = merge {
    object { SpecialBase }
    object { SpecialNeck }
    object { BishopCap }
}

#declare PawnCap = merge {
    torus {
        0.2, 0.04
        translate <0,0.6,0>
    }
    sphere {
        <0,0.7,0>, 0.2
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
    scale 0.35
}

#declare SwapPawn = merge {
    object { Pawn }
    object { Swap }
}

#declare Lock = merge {
    merge {
        box {
            <-0.1,0,-0.1>, <0.1,1.2,0.1>
            rotate <45,0,45>
        }
        box {
            <-0.1,0,-0.1>, <0.1,1.2,0.1>
            rotate <45,45,90>
        }
        box {
            <-0.1,0,-0.1>, <0.1,1.2,0.1>
            rotate <-45,0,-45>
        }
        box {
            <-0.1,0,-0.1>, <0.1,1.2,0.1>
            rotate <-45,-45,-90>
        }
        translate <0,5,0>
    }
    #local i = 0;
    #while (i < 6)
        box {
            <-0.1,-1.2,-0.1>,<0.1,1.2,0.1>
            rotate <90,90,0>
            translate <0,i,0>
        }
        #local i = i + 1;
    #end
    scale 0.35
    translate <0,-1.15,0>
    rotate <0,60,0>
}

#declare LockPawn = merge {
    object { Pawn }
    object { Lock }
}

#declare Recon = merge {
    box {
        <-0.1,-1.3,-0.1>, <0.1,1.3,0.1>
        rotate <90,0,0>
        translate <0.3,5,0>
    }
    box {
        <-0.1,-1.3,-0.1>, <0.1,1.3,0.1>
        rotate <90,0,0>
        translate <-0.3,5,0>
    }
    scale 0.35
    translate <0,-1,0>
    rotate <0,15,0>
}

#declare ReconPawn = merge {
    object { Pawn }
    object { Recon }
}

#declare Detonate = merge {
    box {
        <0.3,0.82,0.05>, <-0.3,0.8,-0.05>
        rotate <0,-45,0>
    }
    box {
        <0.3,0.821,0.05>, <-0.3,0.8,-0.05>
        rotate <0,45,0>
    }
    scale 0.8
    translate <0,0.238,0>
}

#declare DetonatePawn = merge {
    object { Pawn }
    object { Detonate }
}

#declare Ghost = merge {
    #local i = 0;
    #while (i < 3)
        torus {
            0.18, 0.1
            translate <0,i*0.12,0>
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
            <0.02,0.9,-0.5>, <-0.02,1.1-(i/100),0.5>
            rotate 45*i
        }
        #local i = i + 2;
    #end
    scale 0.7
}

#declare GuardPawn = merge {
    object { Pawn }
    object { Guard }
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
    scale 0.75
}

#declare RallyPawn = merge {
    object { Pawn }
    object { Rally }
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
}

#declare FortifyPawn = merge {
    object { Pawn }
    object { Fortify }
}

#declare Extended = merge {
    box {
        <-1,-1,-1>, <1,1,1>
    }
    scale 0.1
    translate <0,0.9,0>
}

#declare ExtendedPawn = merge {
    object { Pawn }
    object { Extended }
}
