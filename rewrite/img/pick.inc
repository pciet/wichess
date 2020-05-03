// pick.inc is used to generate the images used
// on the index page in the army configuration.

#include "materials.inc"

#declare TrimMaterial = LookTrimMaterial
#declare PieceMaterial = LookMaterial

#declare Background = pigment {
    rgb <0.008,0.003,0.005>
}

global_settings {
    assumed_gamma 1.0
    ambient_light 0
}

camera {
    location <0,-4,11.7>
    up <0,0,1>
    right <1,0,0>
    sky <0,0,1>
    look_at <0,0,5>
}

light_source {
    <-4,-6,9.5>, rgb <1,0.98,0.99>
}

light_source {
    <10,2,20>, rgb <0.97,0.98,0.99>
}

sky_sphere {
    pigment { Background }
}

box {
    <-5,-5,-5000>,<5,5,-0.99999>
    texture {
        pigment {
            rgb <1,0.9,0.7>
        }
        finish {
            diffuse albedo 0.7
            specular 0
        }
    }
}

