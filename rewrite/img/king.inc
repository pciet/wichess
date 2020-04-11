#declare king = merge {
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
        <4,0.4>,<2.7,0.4999>,<2.2,1.50001><1,3>
        rotate <90,0,0>
    }
    cylinder {
        <0,0,0.49999>, <0,0,1.5>, 2.2
    }
    lathe {
        cubic_spline
        4,
        <8,0.8>,<2,1.4999>,<0.8,7>,<0.8,7.45>
        rotate <90,0,0>
    }
    lathe {
        cubic_spline
        5,
        <0,6>,<0.2,6.2>,<1.4,6.6>,<0.2,7.2>,<0,7.8>
        rotate <90,0,0>
    }
    lathe {
        cubic_spline
        5,
        <0,6.5>,<0.2,6.6>,<1.35,7>,<0.2,7.3>,<0,7.5>
        rotate <90,0,0>
        material { TrimMaterial }
    }
    lathe {
        cubic_spline
        5,
        <0.6,6.9>,<0.7,6.9999>,<0.9,8>,<0,8.1>,<0,8.5>
        rotate <90,0,0>
        material { TrimMaterial }
    }
    box {
        <-0.125,-0.125,8><0.125,0.125,8.9>
        material { TrimMaterial }
    }
    box {
        <-0.1251,-0.1251,-0.5><0.1251,0.1251,0.5>
        rotate <90,0,45>
        translate <0,0,8.5>
        material { TrimMaterial }
    }
    sphere {
        <0,0,8.8>,0.25
        material { TrimMaterial }
    }
    material { PieceMaterial }
}
