import { layoutSelector } from '../layout.js'
import { Pieces, NoKind } from '../pieceDefs.js'
import { Characteristics } from '../pieceCharacteristics.js'

import { promoting } from './promotion.js'

export function writePieceCharacteristics(kind) {
    if (promoting === true) {
        return
    }
    if (kind === NoKind) {
        document.querySelector('#descriptionmargin').innerHTML = ''
        return
    }
    let t = '<div id="description">'
    t += `
<div>
    <div></div>
    <div id="piece">`+Pieces[kind].name+`</div>
    <div></div>
</div>
<div>`
    const cs = Pieces[kind].characteristics
    if (cs[0] !== undefined) {
        t += '<div class="characteristic">'+Characteristics[cs[0]]+'</div>'
        if (cs[1] !== undefined) {
            t += '<div class="characteristic">'+Characteristics[cs[1]]+'</div>'
        }
    }
    t += '</div>'
    layoutSelector('#descriptionmargin', t)
}
