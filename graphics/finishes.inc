#declare WhiteT = texture {
    pigment { rgb <1,1,1> }
    finish {
        ambient 0.1
        diffuse 0.2
        reflection 0.1
        refraction 1
        ior 1.5
        specular 1
        roughness 0.001
    }
}

#declare BlackT = texture {
    pigment { rgb <0.9,0.25,0.1> }
    finish {
        ambient 0.1
        diffuse 0.2
        reflection 0
        refraction 0.5
        ior 1.5
        specular 1
        roughness 0.001
    }
}
