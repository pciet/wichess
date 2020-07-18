// This white trim looks good on the pawn but isn't always good for other pieces because the
// shape can be lost at a distance.
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

#declare WhiteTrimBMaterial = material {
    texture {
        pigment { rgb <0.38,0.44,0.78> }
        normal {
            waves
            scale 0.03
        }
        finish {
            diffuse 0.7
            specular 0.6
            roughness 0.01
        }
    }
}

// transparency/interior was removed from WhiteMaterial because it multiplied rendering time 
// too much.

// reflection doesn't seem to add a big amount of extra rendering time.

#declare WhiteMaterial = material {
    texture {
        pigment { rgb <0.6,0.6,0.97> }
        normal {
            waves
            scale 0.1
        }
        finish {
            ambient 0.15
            diffuse albedo 0.7
            specular 0.2
            roughness 0.8
            reflection 0.4
        }
    }
}

#declare PlainWhiteMaterial = material {
    texture {
        pigment { rgb <0.6,0.6,0.97> }
        finish {
            ambient 0.15
            diffuse albedo 0.7
            specular 0.2
            roughness 0.8
            reflection 0.4
        }
    }
}

#declare BlackTrimMaterial = material {
    texture {
        pigment {
            gradient z
            color_map {
                [0 rgb <0.7,0.3,0.05>]
                [0.5 rgb <0.7,0.3,0.1>]
                [1 rgb <0.8,0.3,0.05>]
            }
            scale 0.5
        }
        normal {
            dents
            scale 0.05
        }
        finish {
            ambient 0.15
            diffuse albedo 0.7
            specular 0.99
            roughness 0.05
            metallic
        }
    }
}

#declare BlackTrimBMaterial = BlackTrimMaterial

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
            ambient 0.03
            diffuse albedo 0.8
            specular 0.7
            roughness 0.1
            reflection 0.2
            metallic
        }
    }
}

#declare PlainBlackMaterial = BlackMaterial

#declare LookMaterial = material {
    texture {
        pigment { 
            gradient z
            color_map {
                [0 rgbf <0.5,0.7,1, 0.3>]
                [0.7 rgbf <1,0.4,0.5, 0.4>]
                [1 rgbf <0.5,0.7,1, 0.3>]
            }
            scale 6
        }
        normal {
            wrinkles
            scale 0.01
        }
        finish {
            ambient 0
            diffuse albedo 0.3
            specular 0.7
            roughness 0.1
        }
    }
    interior {
        ior 1.8
        dispersion 1.08
        caustics 0.5
    }
}

#declare LookTrimMaterial = LookMaterial
#declare LookTrimBMaterial = LookMaterial
#declare PlainLookMaterial = LookMaterial
