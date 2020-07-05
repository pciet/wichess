// pick.inc is used to generate the images used on the index page in the army configuration.

#include "materials.inc"

#declare TrimMaterial = LookTrimMaterial
#declare TrimBMaterial = LookTrimBMaterial
#declare PieceMaterial = LookMaterial
#declare PlainMaterial = PlainLookMaterial

#declare Background = pigment {
    rgb <0.008,0.003,0.005>
}

global_settings {
    assumed_gamma 1.0
    ambient_light 0
}

camera {
    location <0,-5,11>
    up <0,0,1>
    right <1,0,0>
    sky <0,0,1>
    look_at <0,0,3.5>
}

light_source {
    <-4,-6,50>, rgb <1,0.98,0.99>
}

light_source {
    <7,-5,40>, rgb <1,0.98,0.99>
}

sky_sphere {
    pigment { Background }
}

box {
    <-50,-50,-5000>,<50,50,-0.99999>
    texture {
        pigment {
            rgb <0.05,0.05,0.05>
        }
        normal {
            waves
            scale 0.2
        }
        finish {
            ambient 0
            diffuse albedo 0.1
            specular 0.05
            roughness 0.01
            reflection 0.2
        }
    }
}


#local i = 0;
#while (i < 4)
merge {
    box {
        <0,-5,-1>,<1,5,7>
    }
    box {
        <-0.5,-5,-1>,<0,-4.5,7>
    }
    translate <5,0,0>
    rotate <0,0,90*i>
    texture {
        pigment {
            bozo
            color_map {
                [0.2 rgb 0][1 rgb 1]
            }
            scale 0.01
        }
        finish {
            ambient 0
            diffuse albedo 0.1
            specular 0.7
            roughness 0.001
        }
    }
}
#local i = i + 1;
#end

