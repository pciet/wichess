#include "knight.inc"

#declare brilliant = merge {
    object { knight }
    sphere {
        <0,0,2.3>,1.7
        material { TrimMaterial }
    }
    sphere {
        <0,0,3.3>,1
        material { TrimMaterial }
    }
    sphere {
        <0,0,6.6>,0.5
        material { TrimMaterial }
    }
    material { PieceMaterial }
}
