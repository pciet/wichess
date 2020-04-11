#declare queen = merge {
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
        <4,0.4>,<2.7,0.4999>,<2.4,1.50001><1,3>
        rotate <90,0,0>
    }
    cylinder {
        <0,0,0.49999>, <0,0,1.5>, 2.2
    }
    lathe {
        cubic_spline
        4,
        <8,0.8>,<1.8,1.4999>,<0.7,7>,<1,7.45>
        rotate <90,0,0>
    }
    lathe {
        cubic_spline
        5,
        <0,6>,<0.2,6.1>,<1.1,6.5>,<0.7,6.8>,<1,7>
        rotate <90,0,0>
    }
    lathe {
        cubic_spline
        5,
        <0,6.5>,<0.2,6.6>,<1.3,7>,<0.2,7.3>,<0,7.5>
        rotate <90,0,0>
    }
    #local i = 0;
    #while (i < 17)
        sphere {
            <0,0,7>, 0.15
            translate <1.2,0,0>
            rotate <0,0,7>
            rotate <0,0,(360/17)*i>
            material { TrimMaterial }
        }
    #local i = i + 1;
    #end
    cone {
        <0,0,7>, 0.7, <0,0,8>, 0.8
    }
    #local i = 0;
    #while (i < 13)
        sphere {
            <0,0,8>,0.15
            translate <0.7,0,0>
            rotate <0,0,(360/13)*i>
            material { TrimMaterial }
        }
    #local i = i + 1;
    #end
    #local i = 0;
    #while (i < 15)
        sphere {
            <0,0,8>,0.1
            translate <0.4,0,0>
            rotate <0,0,10>
            rotate <0,0,(360/15)*i>
            material { TrimMaterial }
        }
    #local i = i + 1;
    #end
    lathe {
        cubic_spline
        5,
        <0.3,7.9>,<0.2,8><0.1,8.6><0.05,8.8><0.01,9>
        rotate <90,0,0>
    }
    difference {
        sphere {
            <0,0,8.8>,0.25
            material { TrimMaterial }
        }
        torus {
            0.27,0.1
            rotate <90,0,0>
            translate <0,0,8.8>
        }
    }
    material { PieceMaterial }
}
