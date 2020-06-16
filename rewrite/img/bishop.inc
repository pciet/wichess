#declare bishop = merge {
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
    #local i = 0;
    #while (i < 10)
    intersection {
        box {
            <-0.6,-0.6,-0.5>,<0.5,0.5,0.99999>
            translate <2.5,0,0.5>
            rotate <0,0,i*360/10>
        }
        torus {
            2.5, 0.8
            rotate <90,0,0>
            translate <0,0,0.8>
        }
        material { TrimBMaterial }
    }
    #local i = i + 1;
    #end
    lathe {
        cubic_spline
        4,
        <4,0.4>,<2.4,0.4999>,<2.2,1.50001><1,3>
        rotate <90,0,0>
        material { TrimBMaterial }
    }
    cylinder {
        <0,0,0.49999>, <0,0,1.5>, 2.2
        material { TrimBMaterial }
    }
    lathe {
        cubic_spline
        4,
        <8,0.8>,<1.8,1.4999>,<0.7,5.5>,<0.5,7.45>
        rotate <90,0,0>
    }
    lathe { // disk
        cubic_spline
        5,
        <0,5>,<0.2,5.1>,<1.2,5.5>,<0.2,5.8>,<0.2,6>
        rotate <90,0,0>
        translate <0,0,-0.5>
        material { TrimMaterial }
    }
    difference {
        ovus {
            1, 0.5
            rotate <90,0,0>
            translate <0,0,5.8>
        }
        box {
            <-4,-4,-0.2>,<4,4,0.2>
            rotate <0,60,0>
            translate <-2,0,9.5>
        }
    }
    sphere {
        <0,0,7.4>,0.2
    }
    material { PieceMaterial }
}
