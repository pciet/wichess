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
        <8,0.8>,<1.8,1.4999>,<0.6,6.2>,<0.5,7.45>
        rotate <90,0,0>
    }
    difference {
        ovus {
            1.2, 0.5
            rotate <90,0,0>
            translate <0,0,6.44>
            material { TrimMaterial }
        }
        box {
            <-4,-4,-0.2>,<4,4,0.2>
            rotate <0,60,0>
            translate <-1.8,0,9.8>
        }
    }
    lathe {
        cubic_spline
        5,
        <0,5>,<0.2,5.1>,<1.5,5.5>,<0.2,5.8>,<0,6>
        rotate <90,0,0>
    }
    sphere {
        <0,0,7.9>,0.35
        material { TrimMaterial }
    }
    material { PieceMaterial }
}
