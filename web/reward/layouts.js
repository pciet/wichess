import { ct } from '../layout_ct.js'
import { NoKind } from '../pieceDefs.js'
import { collectionImage, pickImage } from '../collection.js'

// ct(id, classes, inline, noselect, text)

export const landscape = `
<div>
    <div class="inline">
        <div></div>
        ` + ct('ack', '', false, true, '&#x2713;') + `
        <div></div>
    </div>
    <div id="collection" class="inline">` + collection() + `</div>
    <div class="inline">
        <div></div>
        ` + ct('details', '', false, true, 'Piece') + `
        <div></div>
    </div>
</div>
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
            t += '<img id="c'+index+'" class="inline collectioncell noselect" src="' + 
                collectionImage(index) + '">'
        }
        t += '</div>'
    }
    return t
}
