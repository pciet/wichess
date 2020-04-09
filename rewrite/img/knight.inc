/*************************************************************************************
 *  Ted's Chess Pieces Include File   (Chesssets version 1.2)
 *
 *  File: teds_Chess_Pieces.inc
 *
 *  Original Chess Pieces by Ted Fisher, 1997 (Blue Knight Board)
 *  pieces are on the origin, sitting on y = 0
 *
 *  Author: James Garner ( jkgarner@charter.com )
 *
 *************************************************************************************
 *
 *  LEGAL NOTICE:
 *
 *  You may use this file to generate your scenes or distribute it freely.
 *  Please remember to give credit where credit is due; to not do so is both
 *  illeagal and immoral.  I have extensivly borrowed from other users and expect to
 *  be borrowed from, but credit is expected where deserved.
 *
 *  You may modify this file or any others included.  Remeber before doing so,
 *  I do not guarentee these files will work on you machine, nor cause undesired
 *  consuquences. (Use at your own risk.) I can not guarentee any modifications you
 *  choose to make.
 *  
 *  If you choose to modify or add to this work, please do the following:
 *  1. document what you did (why and how would be apprecieated by the next guy)
 *  2. please let me know what you have done (send you new source)
 *  3. leave this notice intact   
 *
 *************************************************************************************
 *
 *  Include File Declares the following objects:
 *    Teds_Pawn
 *    Teds_Rook
 *    Teds_Knight
 *    Teds_Bishop
 *    Teds_Queen
 *    Teds_King
 *    Teds_PIECESIZE
 *
 ************************************************************************************/

/****************************************************************************************
 * Changes for https://github.com/pciet/wichess by Matthew Juran, 2017:
 *   - isolated Teds_Knight from the rest of the set and changed declared name to Knight
 *   - replaced sor base with cylinders + torus pattern used with the rest of the set
 *     (sor is better looking but harder to maintain and may be added in the future)
 *   - set translation, scaling, and rotation for my set specifics
 *
 * This knight.inc is used in wichess/graphics/<w/b><type>knight.pov files where final
 * rotation and texture is set and combined with the same lighting applied to all pieces.
 * The generated images are at wichess/web/img/ and a viewer page is created at 
 * wichess/graphics/img/ when rendering a new set with wichess/graphics/generate.
 *
 * wichess/graphics/generate.go shows how the camera is moved to generate 65 images.
 *
 * Carving a knight out from a cylinder is a mind-bender, good job to the original authors.
 *
 * Issues with povray 3.7:
 *   - two surfaces on the head of the knight are equal which renders a noise pattern
 ****************************************************************************************/

/****************************************************************************************
 * Changes for https://github.com/pciet/wichess by Matthew Juran, 2020:
 *   - reintroduced original sor base
 *   - resized and reoriented (Z is up)
 *   - renamed declared name to knight
 *   - added material reference
 ****************************************************************************************/

/*************************************************************************************
 *  Declare the Knight
 *************************************************************************************/

//sor{
#declare Knight_Base = sor {
    9,
    <0, -.01>,
    <0, 0>,
    <1.5, .01>,
    <1.5, .49>,
    <0, .5>,
    <1.25, .51>,
    <1.25, .99>,
    <0, 1>,
    <0, 1.01>
}

#declare Knight_Body = cylinder{
    <0,1,0>,
    <0,2,0>,
    1.1
}

#declare Knight_Underside_Curve = difference{
    cylinder{
        <1.5,2,0>,
        <-1.5,2,0>,
        2.5
    }

    merge{
        cylinder{
            <1.6,2,0>,
            <-1.6,2,0>,
            1.1
        }

        box{
            <1.6,2,4.1>,
            <-1.6,7,-4.1>
        }
    }
}

#declare Knight_Mid_Low = merge{
    object{Knight_Body
           translate<0,.99,0>
    }
    object{Knight_Body
           translate<0,1.98,0>
    }
    object{Knight_Body
           translate<0,2.97,0>
    }
    object{Knight_Body
           translate<0,3.96,0>
    }

    difference{
        object{Knight_Body}

        object{Knight_Underside_Curve}
    }
}

#declare Knight_Side_Curve = cylinder{
    <8,5,4>,
    <8,5,-4>,
    7.65
}

#declare Knight_Box = box{
    <-3,-3,-3>,
    <3,3,3>
}

#declare Knight_Neck_Cut = merge{
    cylinder{
        <.25,0,0>,
        <-.25,0,0>,
        .1
    }

    box{
        <.25,-.1,-1>,
        <-.25,.1,0>
        rotate<-15,0,0>
    }

    box{
        <.25,-.1,-1>,
        <-.25,.1,0>
        rotate<15,0,0>
    }
}

#declare Knight_Head_Cut = merge{
    difference{
        merge{
            box{
                <1,.75,.7071067811865>,
                <-1,-.75,-.7071067811865>
                translate<0,.7071067811865,.75>
            }
            cylinder{
                <-1,0,0>,
                <1,0,0>,
                1
            }
        }

        merge{
            box{
                <1.1,1.5,1>,
                <-1.1,-1.5,-1>
                rotate<-45,0,0>
                translate<0,1.25,2>
            }
            cylinder{
                <1.1,0,2>,
                <-1.1,0,2>,
                1
           }
        }
    }

    difference{
        box{
            <1,2,2>,
            <-1,-2,-2>
            translate<0,3,0>
        }
        merge{
            cylinder{
                <1.1,0,0>,
                <-1.1,0,0>,
                1.5
            }
            cylinder{
                <1.1,0,1.5>,
                <-1.1,0,1.5>,
                2
            }
        }
        rotate<30,0,0>
        translate<0,-.45,.55>
    }
}

#declare Knight_Box_Cut = intersection{
    box{
        <1,1,2>,
        <-1,-1,-2>
        rotate<60,0,0>
    }
    box{
        <1,1,2>,
        <-1,-1,-2>
        rotate<-60,0,0>
    }
}

#declare Knight_Mane_Cut = merge{
    merge{
        object{Knight_Box_Cut
            scale<3,.25,.25>
            rotate<0,0,65>
        }
        object{Knight_Box_Cut
            scale<3,.25,.25>
            rotate<0,0,-65>
        }
        rotate<30,0,0>
        translate<0,6.5,1.3>
    }
    merge{
        object{Knight_Box_Cut
            scale<3,.25,.25>
            rotate<0,0,65>
        }
        object{Knight_Box_Cut
            scale<3,.25,.25>
            rotate<0,0,-65>
        }
        rotate<30,0,0>
        translate<0,6.2,1.6>
    }
    merge{
        object{Knight_Box_Cut
            scale<3,.25,.25>
            rotate<0,0,70>
        }
        object{Knight_Box_Cut
            scale<3,.25,.25>
            rotate<0,0,-70>
        }
        rotate<35,0,0>
        translate<0,5.8,2>
    }
    merge{
        object{Knight_Box_Cut
            scale<3,.25,.25>
            rotate<0,0,67>
        }
        object{Knight_Box_Cut
            scale<3,.25,.25>
            rotate<0,0,-67>
        }
        rotate<35,0,0>
        translate<0,5.05,1.9>
    }
    merge{
        object{Knight_Box_Cut
            scale<3,.25,.25>
            rotate<0,0,60>
        }
        object{Knight_Box_Cut
            scale<3,.25,.25>
            rotate<0,0,-60>
        }
        rotate<35,0,0>
        translate<0,4.2,1.8>
    }
}

#declare Knight_Face_Cut = merge{
    cylinder{
        <0,1,0>,
        <0,-1,0>,
        1
        rotate<-15,0,15>
        translate<1.25,5.35,-.8>
    }
    cylinder{
        <0,1,0>,
        <0,-1,0>,
        1
        rotate<-15,0,-15>
        translate<-1.25,5.35,-.8>
    }
    cylinder{
        <0,1,0>,
        <0,-1,0>,
        1
        rotate<-15,0,15>
        translate<-1.5,4.9,-.8>
    }
    cylinder{
        <0,1,0>,
        <0,-1,0>,
        1
        rotate<-15,0,-15>
        translate<1.5,4.9,-.8>
    }
    object{Knight_Box_Cut
        scale .25
        rotate<0,0,30>
        translate<.4,5.6,0>
    }
    object{Knight_Box_Cut
        scale .25
        rotate<0,180,-30>
        translate<-.4,5.6,0>
    }
    object{Knight_Box_Cut
        scale<.5,.25,.25>
        rotate<-10,90,0>
        translate<0,6,.5>
    }
    object{Knight_Box_Cut
        scale<.5,.25,.25>
        rotate<90,0,0>
        translate<0,4.45,-1.5>
    }
    union{
        sphere{<.2,0,0>, .1}
        sphere{<-.2,0,0>, .1}
        translate<0,4.8,-1>
    }
    object{Knight_Mane_Cut}
}

#declare Knight_Eye_Cut = difference{
    sphere{
        <0,0,0>, .2
    }

    sphere{
        <.2,0,0>, .2
    }
    translate<-.3,4.95,0>
}

#declare Knight_Mid = difference{
    object{Knight_Mid_Low}

    object{Knight_Side_Curve}
    object{Knight_Side_Curve
        rotate<0,180,0>
    }
    object{Knight_Box
        rotate<0,-40,-22>
        translate<-2.75,2,-3>
    }
    object{Knight_Box
        rotate<0,40,22>
        translate<2.75,2,-3>
    }
    object{Knight_Neck_Cut
        scale 3
        rotate<-30,0,0>
        translate<0,4,-.4>

    }
    object{Knight_Head_Cut
        scale 1
        rotate<45,0,0>
        translate<0,6,-1>
    }
    object{Knight_Face_Cut}
    object{Knight_Eye_Cut
        translate<0,0,.05>
    }
    object{Knight_Eye_Cut
        rotate<0,180,0>
        translate<0,0,.05>
    }
}

/*#declare Teds_Knight = merge{*/
#declare knight = merge {
    object{Knight_Base}
    object{Knight_Mid}
    scale 1.3
    rotate <90,0,0>
    material { PieceMaterial }
}
