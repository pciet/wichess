#declare pawn = merge {
    cylinder {
        <0,0,-0.001>, <0,0,0.5>, 2.8 
    }
    difference {
        torus {
            2.8, 0.25
            rotate <90,0,0>
            translate <0,0,0.25>
            material { TrimMaterial }
        }
        cylinder {
            <0,0,0.4>,<0,0,1>,8
        }
    }
    lathe {
        quadratic_spline
        3,
        <3.1,-0.01>,<3,0.29>,<2.801,0.5001>
        rotate <90,0,0>
    }
    cylinder {
        <0,0,0.49>, <0,0,1.25>, 2
    }
    lathe {
        quadratic_spline
        3,
        <4,0.4>,<2.5,0.7><1.999,1.2>
        rotate <90,0,0>
        translate <0,0,-0.2001>
    }
    lathe {
        quadratic_spline
        3,
        <3,-0.0001>,<1.3,1>,<0.3,4.6>
        rotate <90,0,0>
    }
    sphere {
        <0,0,5.5>, 1
        material { TrimMaterial }
    }
    material { PieceMaterial }
}
