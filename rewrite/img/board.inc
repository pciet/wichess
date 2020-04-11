// board.inc is a complete POV-Ray scene, where the Piece macro
// is used by a .pov file to add a piece to all 64 squares
// on the board for rendering.

global_settings {
    assumed_gamma 1.0
}

camera {
    location <40,40,80>
    up <0,0,1>
    right <1,0,0>
    look_at <40,40,0>
}

light_source {
    <250,200,1900>, rgb <0.97,0.98,0.97>
}

box {
    <-2000,-2000,-2000>, <2000,2000,2000>
    texture {
        pigment { rgb <0.99,0.98,0.97> }
    }
}

// Adding texture/normal to the square background
// multiplies the filesize almost 5x. Ripples and
// bozo normals looked good though.

#declare OddTexture = texture {
    pigment { rgb <0.02,0.05,0.01> }
    finish {
        ambient 0
        diffuse albedo 0.2
        specular 0.05
    }
}

#declare EvenTexture = texture {
    pigment { rgb <0.01,0.04,0.03> }
    finish {
        ambient 0
        diffuse albedo 0.3
        specular 0.02
    }
}

#macro BoardBox(X, Y, EvenSquare)
box {
    <X,Y,0>, <X+10,Y+10,-1>
    #if (EvenSquare = 0)
        texture { OddTexture }
    #else
        texture { EvenTexture }
    #end
}
#end

#local i = 0;
#while (i < 8)
    #local j = 0;
    #while (j < 8)
        #local xaddr = i*10;
        #local yaddr = j*10;
        #if (mod(j,2) = 0)
            #if (mod(i,2) = 0)
                BoardBox(xaddr, yaddr, 0)
            #else
                BoardBox(xaddr, yaddr, 1)
            #end
        #else
            #if (mod(i,2) = 0)
                BoardBox(xaddr, yaddr, 1)
            #else
                BoardBox(xaddr, yaddr, 0)
            #end
        #end
    #local j = j + 1;
    #end
#local i = i +1;
#end

#macro Piece(P)
    #local i = 0;
    #while (i < 8)
        #local j = 0;
        #while (j < 8)
            object {
                P
                translate <(i*10)+5,(j*10)+5,0>
                no_reflection
            }
            #local j = j +1;
        #end
        #local i = i + 1;
    #end
#end
