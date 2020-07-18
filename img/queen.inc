#declare queen = merge {
    difference { // base
        merge {
            merge {
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
            cylinder {
                <0,0,0.49999>, <0,0,1.5>, 2.2
            }
        }
        #local i = 0;
        #while (i < 4)
            box {
                <-1,-1,-1>,<1,1,1>
                scale <1,8,1>
                rotate <0,45,0>
                translate <3,0,1>
                rotate <0,0,i*360/4>
            }
        #local i = i + 1;
        #end
        scale <1.2,1.2,1>
        translate <0,0,0.4>
    }
    lathe { // neck
        cubic_spline
        4,
        <8,0.8>,<1.8,1.4999>,<0.9,6>,<3,7.45>
        rotate <90,0,0>
    }
    lathe { // disk
        cubic_spline
        5,
        <0,6.5>,<0.2,6.6>,<1.7,7>,<0.2,7.3>,<0,7.5>
        rotate <90,0,0>
        translate <0,0,-1>
        material { TrimBMaterial }
    }
    #local i = 0; // crown
    #while (i < 10)
        sphere {
            <0,0,7.1>, 0.3
            translate <1.3,0,0>
            rotate <0,0,7>
            rotate <0,0,(360/10)*i>
            translate <0,0,-1>
            material { TrimBMaterial }
        }
    #local i = i + 1;
    #end
    sphere {
        <0,0,7.5>,0.8
        translate <0,0,-1>
        material { TrimMaterial }
    }
    rotate <0,0,30>
    material { PieceMaterial }
}
