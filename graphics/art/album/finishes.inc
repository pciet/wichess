#include "colors.inc"
#include "textures.inc"
#include "stones1.inc"

#declare WhiteT = texture {
    T_Stone15 scale 0.1
    finish {
        specular 0.75
        roughness 0.001
    }
}

#declare BlackT = texture {
    DMFDarkOak scale 0.1
    finish {
        specular 0.3
        roughness 0.01
    }
}
