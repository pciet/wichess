#declare WhiteTrimMaterial = material {
    texture {
        pigment { rgbf <0.8,0.8,0.97,0.8> }
        finish {
            diffuse albedo 0.95
            specular 0
            roughness 0.04
        }
    }
    interior {
        ior 1.4
        dispersion 1.08
        caustics 0.5
    }
}

#declare WhiteMaterial = material {
    texture {
        pigment { rgbf <0.8,0.8,0.97,0.5> }
        finish {
            diffuse albedo 0.7
            specular 0.7
            roughness 0.1
        }
    }
    interior {
        ior 1.5
        dispersion 1.08
        caustics 0.5
    }
}

#declare BlackTrimMaterial = material {
    texture {
        pigment { rgb <0.7,0.3,0.05> }
        normal {
            dents
            scale 0.08
        }
        finish {
            diffuse albedo 0.2
            specular 0.98
            metallic
        }
    }
}

#declare BlackMaterial = material {
    texture {
        pigment { 
            gradient z
            color_map {
                [0 rgb <0.6,0.3,0.1>]
                [0.5 rgb <0.4,0.4,0.3>]
                [1 rgb <0.6,0.3,0.1>]
            }
            scale 4
        }
        normal {
            dents
            scale 0.08
        }
        finish {
            ambient 0.1
            diffuse albedo 0.5
            specular 0.7
            roughness 0.1
            reflection 0.1
            metallic
        }
    }
}

#declare LookMaterial = material {
    texture {
        pigment { 
            gradient z
            color_map {
                [0 rgbf <0.1,0.2,0.4,0.8>]
                [0.5 rgbf <0.4,0.4,0.3,1>]
                [1 rgbf <0.1,0.2,0.5,0.8>]
            }
            scale 20
        }
        finish {
            ambient 0.1
            diffuse albedo 0.55
            specular 0.7
            roughness 0.1
            reflection 0.05
        }
    }
    interior {
        ior 1.8
        dispersion 1.08
        caustics 0.5
    }
}

#declare LookTrimMaterial = LookMaterial
