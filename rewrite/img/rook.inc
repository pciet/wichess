#declare rook = merge {
    union {
        cylinder {
            <0,0,-0.00001>,<0,0,0.5>,2.9
        }
        torus {
            2.9, 0.1
            rotate <90,0,0>
            scale <1,1,5>
        }
        material { TrimMaterial }
    }
    lathe {
        cubic_spline
        4,
        <4,0.4>,<2.4,0.4999>,<2.2,1.50001><1,3>
        rotate <90,0,0>
    }
    cylinder {
        <0,0,0.49999>, <0,0,1.5>, 2.2
    }
    lathe {
        cubic_spline
        4,
        <8,0.8>,<2,1.4999>,<1.2,5.5>,<2,7.45>
        rotate <90,0,0>
    }
    lathe {
        cubic_spline
        5,
        <0,5>,<0.2,5.1>,<1.5,5.3>,<0.2,5.5>,<0,6>
        rotate <90,0,0>
    }
    difference {
        cylinder {
            <0,0,5.2999>,<0,0,7>,1.4
        }
        #local i = 0;
        #while (i < 2)
            box {
                <-3,-0.3,-2>,<3,0.3,2>
                rotate <0,0,90*i>
                translate <0,0,8.2>
            }
            #local i = i+1;
        #end
        cylinder {
            <0,0,5.5>,<0,0,8>,0.8
        }
    }
    material { PieceMaterial }
    rotate <0,0,30>
}
