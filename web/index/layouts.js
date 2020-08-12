import { ct } from '../layout_ct.js'
import { NoKind } from '../pieceDefs.js'
import { collectionImage, pickImage } from '../collection.js'

import { armyImage } from './army.js'

// ct(id, classes, inline, noselect, text)

const top = `
<div id="top">
` + ct('name', '', true, false, window.Name) +
    ct('title', '', true, false, 'WISCONSIN CHESS') + `
    <a class="inline" id="quita" href="/quit">` + ct('quit', '', false, true, 'Quit') + `</a>
</div>
`

export const landscape = top + `
<div>
    <div id="left" class="inline">
        ` + ct('pickstitle', '', false, true, 'Random Picks') + `
        <div id="picksmargin">
            <img id="leftpick" class="inline pick" src="` + pickImage(window.LeftPiece) + `">
            <img id="rightpick" class="inline pick" src="` + pickImage(window.RightPiece) + `">
        </div>
        <div class="leftspacer"></div> ` +
        ct('details', '', false, true, 'Piece Details') + 
        ct('rules', '', false, true, 'Rules Overview') + `
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

export const square = top + `
<div>
    <div>
        ` + ct('squarearmytitle', '', false, true, 'Army') + `
        <div id="army">` + army() + `</div>
    </div>
    <div id="squarepicks">
        ` + ct('squarepickstitle', '', true, true, 'Random Picks &rarr;') + `
        <div class="inline" id="squarepicksmargin">
            <img id="leftpick" class="inline pick" src="` + pickImage(window.LeftPiece) + `">
            <img id="rightpick" class="inline pick" src="` + pickImage(window.RightPiece) + `">
        </div>
        <div class="inline"></div>
    </div>
    <div>
        ` + ct('squarecollectiontitle', '', false, true, 'Collection') + `
        <div id="collection">` + collection() + `</div>
    </div>
    <div id="squarebuttons">
        ` + ct('squaredetails', '', true, true, 'Piece Details') +
            ct('squarerules', '', true, true, 'Rules Overview') +
            ct('squarematch', '', true, true, 'Match') + `
    </div>
</div>
`

export const unsupported = `
<div class="inline"></div>
<div class="inline" id="unsupported">
    <div></div>
    <div>This window is set to an unsupported dimension.<br>If this is a desktop window then 
    resize it.</div>
    <div></div>
</div>
<div class="inline"></div>
`

function army() {
    let t = ''
    for (let i = 0; i < 2; i++) {
        t += '<div id="army'+i+'">'
        for (let j = 0; j < 8; j++) {
            const index = (8*i) + j
            t += '<div id="ac'+index+'" class="inline armybox">'
            t += '<img id="ab'+index+'" class="armycell noselect" src="' + armyImage(index) + '">'
            t += '<img id="a'+index+'" class="armycell noselect piecehidden">'
            t += '</div>'
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
