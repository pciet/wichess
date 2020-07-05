import { ct } from '../layout_ct.js'
import { NoKind } from '../pieceDefs.js'

import { armyImage } from './army.js'
import { collectionImage, pickImage } from './collection.js'

// ct(id, classes, inline, noselect, text)

export const landscape = `
<div id="top">
` + ct('name', '', true, false, window.Name) +
    ct('title', '', true, false, 'WISCONSIN CHESS') + `
    <div id="topspacer" class="inline"></div>
    <a class="inline" href="/quit">` + ct('quit', '', false, true, 'Quit') + `</a>
</div>
<div>
    <div id="left" class="inline">
        ` + ct('details', '', false, true, 'Piece Details') + `
        <div id="picksmargin">
            <img id="leftpick" class="inline pick" src="` + pickImage(window.LeftPiece) + `">
            <img id="rightpick" class="inline pick" src="` + pickImage(window.RightPiece) + `">
        </div>
        <p id="pickdesc">The two above pieces can be added to the army. If you then complete a
            game used ones are added to your collection and replaced with another random piece.</p>
        ` + ct('match', '', false, true, 'Match') + `
    </div>
    <div class="inline">
        <p id="armydesc">Below is your army that will be put onto the game board. Click on pieces 
            in the collection below that or to the left then at where you want them in the army to
            customize. 
            Pieces must be the same basic kind as the regular chess piece they are replacing.</p>
        <div id="army">` + army() + `</div>
        <div id="collection">` + collection() + `</div>
    </div>
</div>
`

function army() {
    let t = ''
    for (let i = 0; i < 2; i++) {
        t += '<div id="army'+i+'">'
        for (let j = 0; j < 8; j++) {
            const index = (8*i) + j
            t += '<img id="a'+index+'" class="inline armycell noselect" src="' + armyImage(index) +
                '">'
        }
        t += '</div>'
    }
    return t
}

function collection() {
    let t = ''
    for (let i = 0; i < 3; i++) {
        t += '<div>'
        for (let j = 0; j < 7; j++) {
            const index = (7*i) + j
            if (window.Collection[index] === NoKind) {
                t += '<div id="c'+index+'" class="inline collectioncell noselect"></div>'
            } else {
                t += '<img id="c'+index+'" class="inline collectioncell noselect" src="' + 
                    collectionImage(index) + '">'
            }
        }
        t += '</div>'
    }
    return t
}
