#declare king = merge {
    difference { // base
        union {
            cylinder {
                <0,0,-0.00001>,<0,0,0.5>,2.9
            }
            torus {
                2.9, 0.1
                rotate <90,0,0>
                scale <1,1,5>
            }
            cylinder {
                <0,0,0.49999>, <0,0,1.5>, 2.6
            }
            material { TrimMaterial }
        }
        #local i = 0;
        #while (i < 3)
            box {
                <-1,-1,-1>,<1,1,1>
                rotate <-35,0,0>
                scale <20,1,1>
                translate <0,3,1>
                rotate <0,0,i*360/3>
            }
        #local i = i + 1;
        #end
        scale <1.1,1.1,1>
    }
    lathe {
        cubic_spline
        4,
        <8,0.8>,<2,1.4999>,<0.6,5.6999>,<0.9,7>
        rotate <90,0,0>
    }
    difference {
        lathe {
            cubic_spline
            5,
            <0.2,4>,<0.2,6>,<1.6,7>,<0,7.3>,<0,7.3>
            rotate <90,0,0>
            translate <0,0,-1.5>
            scale <1.1,1.1,1>
            material { TrimBMaterial }
        }
        #local i = 0;
        #while (i < 10)
            box {
                <-0.5,-0.5,-1>,<0.5,0.5,1>
                rotate <0,0,45>
                translate <1.9,0,6.1>
                rotate <0,0,i*360/10>
            }
        #local i = i + 1;
        #end
    }
    #local i = 0;
    #while (i < 10)
        lathe {
            cubic_spline
            4,
            <2,0>,<0.3,0>,<0,2>,<0,3>
            rotate <90,0,0>
            translate <0.5,0,5.7>
            rotate <0,0,i*360/10>
            material { TrimMaterial }
        }
    #local i = i + 1;
    #end
    rotate <0,0,10>
    material { PieceMaterial }
}
