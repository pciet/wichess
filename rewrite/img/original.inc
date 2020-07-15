#include "bishop.inc"

#declare BigBox = box {
    <-2.8,-0.3,-0.0001>,<2.8,0.3,3>
}

#declare Box = box {
    <-2.5,-0.05,-0.0001>,<2.5,0.05,3>
}

#declare original = difference {
    merge {
        object { bishop }
        difference {
            merge {
                #local i = 0;
                #while (i < 4)
                object {
                    BigBox
                    rotate <0,0,(i*180/4)+25>
                }
                #local i = i + 1;
                #end
            }
            torus {
                1, 0.25
                rotate <90,0,0>
                scale <2.2,2.2,1>
                translate <0,0,2.8>
            }
            material { TrimMaterial }
        }
        difference {
            merge {
                #local i = 0;
                #while (i < 3)
                object {
                    Box
                    scale <1,3,1>
                    rotate <0,0,(i*180/3)-25>
                }
                #local i = i + 1;
                #end
            }
            torus {
                1.2, 0.6
                rotate <90,0,0>
                scale 2
                translate <0,0,2.4>
            }
            scale <0.6,0.6,0.1>
            translate <0,0,5.1>
            material { TrimMaterial }
        }
    }
    box {
        <-1,-1,-1>,<1,1,1>
        rotate <10,0,0>
        translate <0,0,7.6>
    }
    material { PieceMaterial }
}
