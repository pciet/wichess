import { removeNewlines } from '../layout.js'
import { layoutElement } from '../layoutElement.js'
import { Kind, kindCharacteristics, kindName } from '../piece.js'

export function writePieceCharacteristics(kind) {
    const d = document.querySelector('#description')
    let t = ''
    if (kind != Kind.NO_KIND) {
        const cs = kindCharacteristics(kind)
        t += '<div id="piece">'+kindName(kind)+'</div>'
        if (cs[0] !== undefined) {
            t += '<div class="characteristic">'+cs[0]+'</div>'
        }
        if (cs[1] !== undefined) {
            t += '<div class="characteristic">'+cs[1]+'</div>'
        }
    }
    d.innerHTML = removeNewlines(t)
    layoutElement(d)
}
