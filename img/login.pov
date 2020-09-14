#version 3.7;

// login.pov is the artistic teaser image shown on the login webpage.

global_settings {
    assumed_gamma 1.0
}

camera {
    location <0,-20,7>
    direction y
    right x
    up z
}

sky_sphere { pigment { rgb 1 } }

light_source {
    <-60,-30,40>, rgb 1
}

merge {
    box {
        <-1,-1,-30>,<1,1,3000>
        rotate <0,0,50>
        translate <-5,0,33>
    }
    sphere {
        <-5,1,2>,3.5

    }
    merge {
        box {
            <-1,-1,-6>,<1,1,6>
        }
        box {
            <0.999,-1,5>,<4,1,6>
            rotate <0,0,90>
        }
        scale 0.8
        rotate <90,0,110>
        translate <0,1,2>
    }
    texture {
        pigment {
            rgbf <0.1,0.4,0.7,0.9>
        }
        finish {
            reflection 0.2
            ambient 0
            diffuse albedo 0.3
            specular 0.7
        }
    }
    interior {
        ior 2.4
    }
}
