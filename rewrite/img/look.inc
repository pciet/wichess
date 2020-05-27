// look.inc is a POV-Ray scene for showing a detailed full piece picture. The piece geometry
// follows the inclusion of this file.

global_settings {
    assumed_gamma 1.0
}

camera {
    location <3,10.5,8.2>
    up <0,0,1>
    right <1,0,0>
    sky <0,0,1>
    look_at <0,0,4.1>
}

light_source {
    <-4,6,9.5>, rgb <0.97,0.98,0.97>
}

sphere {
    <0,0,0>, 500
    texture {
        pigment { rgb <0.98,0.99,0.97> }
        finish {
            diffuse 0.99
            specular 0.3
            roughness 0.05
            metallic
        }
    }
}

box {
    <-3.5,-3.5,-10>,<3.5,3.5,-0.0001>
    texture {
        pigment { rgb <0.8,0.7,0.6> }
        finish {
            ambient 0
            diffuse 0.8
            specular 0.1
        }
    }
    no_shadow
}
