// look.inc is a POV-Ray scene for showing a detailed full piece picture. The piece geometry
// follows the inclusion of this file.

global_settings {
    assumed_gamma 1.0
}

camera {
    location <3,10,7>
    up <0,0,1>
    right <1,0,0>
    sky <0,0,1>
    look_at <0,0,3>
}

light_source {
    <-4,6,9.5>, rgb <0.97,0.98,0.97>
}

sky_sphere { pigment { rgb 1 } }

box {
    <-3.5,-3.5,-10>,<3.5,3.5,-0.0001>
    texture {
        pigment { rgb <0.75,0.7,0.5> }
        finish {
            ambient 0
            diffuse 0.8
            specular 0.1
        }
        normal {
            wood
            scale 0.3
        }
    }
    no_shadow
}
