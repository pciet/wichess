#declare WhiteT = texture {
    pigment { rgb <1,0.999,0.98> }
    finish {
        diffuse 0.6
        specular 0.5
        roughness 0.1
    }
}

#declare BlackT = texture {
    pigment { rgb <0.3,0.13,0.12> }
    finish {
        diffuse 0.6
        specular 0.5
        roughness 0.1
        reflection 0.04
    }
}
