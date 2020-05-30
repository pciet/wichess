import { layoutSelector } from '../layout.js'
import { Pieces, NoKind } from '../pieceDefs.js'
import { Characteristics } from '../pieceCharacteristics.js'
import { orientationString } from '../piece.js'
import { Board } from '../game.js'

import { promoting } from './promotion.js'

let selectedPiece

export function showSelectedPieceCharacteristics(addrIndex) {
    selectedPiece = Board[addrIndex]
}

export function unshowSelectedPieceCharacteristics() {
    selectedPiece = undefined
}

export function writePieceCharacteristics(kind, orientation = undefined) {
    if (promoting === true) {
        return
    }
    if (kind === NoKind) {
        if (selectedPiece === undefined) {
            document.querySelector('#descriptionmargin').innerHTML = ''
        } else {
            writePieceCharacteristics(selectedPiece.kind, selectedPiece.orientation)
        }
        return
    }
    let t = '<div id="description">'
    t += `
<div id="piecemargin">
    <div></div>
    <div id="piece">`+Pieces[kind].name+`</div>
    <div></div>
</div>
<div id="orientation"><div></div><div>`+orientationString(orientation)+`</div><div></div></div>
<div><div></div>`
    const cs = Pieces[kind].characteristics
    if (cs[0] !== undefined) {
        t += '<div class="characteristic">'+Characteristics[cs[0]]+'</div>'
        if (cs[1] !== undefined) {
            t += '<div class="characteristic">'+Characteristics[cs[1]]+'</div>'
        }
    }
    t += '<div></div></div>'
    layoutSelector('#descriptionmargin', t)
}
