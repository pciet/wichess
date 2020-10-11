import { ct } from '../layout_ct.js'
import { NoKind } from '../pieceDefs.js'
import { collectionImage, pickImage } from '../collection.js'

// ct(id, classes, inline, noselect, text)

export function portrait() { 
    return `
<div id="portraitbuttons">
    ` + ct('ack', '', true, true, '&#x2713;') + ct('details', '', true, true, 'Piece') + `
</div>
<div id="collection">` + collection() + `</div>
` + picks
}

export function landscape() {
    return `
<div>
    <div class="inline">
        <div></div>
        ` + ct('ack', '', false, true, '&#x2713;') + `
        <div></div>
    </div>
    <div id="collection" class="inline landscapecollection">` + collection() + `</div>
    <div class="inline">
        <div></div>
        ` + ct('details', '', false, true, 'Piece') + `
        <div></div>
    </div>
</div>
` + picks
}

const picks = `
<div>
    <div class="inline">
        <div class="rewardtitle">Left Pick</div>
        <div class="rewardmargin">
            <img id="left" class="rewardcell noselect" src="`+pickImage(window.LeftPiece)+`">
        </div>
    </div>
    <div class="inline">
        <div class="rewardtitle">Right Pick</div>
        <div class="rewardmargin">
            <img id="right" class="rewardcell noselect" src="`+pickImage(window.RightPiece)+`">
        </div>
    </div>
    <div class="inline">
        <div class="rewardtitle">Reward</div>
        <div class="rewardmargin">
            <img id="reward" class="rewardcell noselect" src="`+pickImage(window.RewardPiece)+`">
        </div>
    </div>
</div>
`

function collection() {
    let t = ''
    for (let i = 0; i < 3; i++) {
        t += '<div>'
        for (let j = 0; j < 7; j++) {
            const index = (7*i) + j
            t += '<div id="cc'+index+'" class="inline collectionbox">'
            t += '<img id="cb'+index+'" class="collectioncell noselect" src="'+
                collectionImage(index) + '">'
            t += '<img id="c'+index+'" class="collectioncell noselect piecehidden">'
            t += '</div>'
        }
        t += '</div>'
    }
    return t
}

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
