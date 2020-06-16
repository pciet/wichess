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
        <4,0.4>,<2.6,0.4999>,<2.2,1.30001>,<1,3>
        rotate <90,0,0>
    }
    cylinder {
        <0,0,0.49999>, <0,0,1.3>, 2.2
    }
    lathe { // neck
        cubic_spline
        4,
        <8,0.8>,<2.3,1.1>,<1.2,4.9>,<3,7.45>
        rotate <90,0,0>
    }
    merge { // disk
        lathe { // bottom of disk trim
            cubic_spline
            4,
            <1.5,3>,<1.1,4>,<1.6,4.89>,<3,5>
            rotate <90,0,0>
        }
        lathe { // disk
            cubic_spline
            5,
            <0,4.7>,<0.2,4.8>,<1.8,4.9>,<0.2,5.3>,<0,6>
            rotate <90,0,0>
        }
        lathe { // top of disk trim
            cubic_spline
            4,
            <4,4.5>,<1.6,4.9>,<1.4,5.1>,<1.37,9>
            rotate <90,0,0>
        }
        material { TrimMaterial }
    }
    merge { // cap
        cylinder {
            <0,0,4.89999>,<0,0,5.30001>,1.40001
        }
        intersection { //battlement
            difference {
                cylinder {
                    <0,0,4.89999>,<0,0,6>,1.4
                }
                #local i = 0;
                #while (i < 5)
                    box {
                        <-3.5,-0.13,-2.7>,<3.5,0.13,2>
                        rotate <0,0,i*180/5>
                        translate <0,0,8>
                    }
                    #local i = i+1;
                #end
                cylinder {
                    <0,0,5.29999>,<0,0,6.1>,0.8
                }
            }
            torus {
                1.15,0.4
                rotate <90,0,0>
                scale <1,1,1.1>
                translate <0,0,5.5>
            }
        }
    }
    material { PieceMaterial }
    rotate <0,0,30>
}
