import { ct } from '../layout_ct.js'
import { NoKind } from '../pieceDefs.js'
import { collectionImage, pickImage } from '../collection.js'

import { armyImage } from './army.js'

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
        ` + ct('pickstitle', '', false, true, 'Random Picks') + `
        <div id="picksmargin">
            <img id="leftpick" class="inline pick" src="` + pickImage(window.LeftPiece) + `">
            <img id="rightpick" class="inline pick" src="` + pickImage(window.RightPiece) + `">
        </div>
        <div class="leftspacer"></div> ` +
        ct('details', '', false, true, 'Piece Details') + `
        <div class="leftspacer"></div>
        ` + ct('match', '', false, true, 'Match') + `
    </div>
    <div class="inline">
        ` + ct('armytitle', '', false, true, 'Army') + `
        <div id="army">` + army() + `</div>
        ` + ct('collectiontitle', '', false, true, 'Collection') + `
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
            t += '<img id="c'+index+'" class="inline collectioncell noselect" src="' + 
                collectionImage(index) + '">'
        }
        t += '</div>'
    }
    return t
}
