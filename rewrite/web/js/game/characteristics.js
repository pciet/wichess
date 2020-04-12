import { layoutSelector } from '../layout.js'
import { Pieces, NoKind } from '../pieceDefs.js'
import { Characteristics } from '../pieceCharacteristics.js'
import { promoting } from './promotion.js'

export function writePieceCharacteristics(kind) {
    if (promoting === true) {
        return
    }
    let t = ''
    if (kind != NoKind) {
        t += '<div id="piece">'+Pieces[kind].name+'</div>'

        const cs = Pieces[kind].characteristics
        if (cs[0] !== undefined) {
            t += '<div class="characteristic">'+Characteristics[cs[0]].name+'</div>'
        }
        if (cs[1] !== undefined) {
            t += '<div class="characteristic">'+Characteristics[cs[1]].name+'</div>'
        }
    }
    layoutSelector('#description', t)
}
