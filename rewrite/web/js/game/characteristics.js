import { layoutElement } from '../layoutElement.js'

export function writePieceCharacteristics(kind) {
    // TODO
    const d = document.querySelector('#description')
    d.innerHTML = '<div>Piece Description</div><div>' + kind + '</div>'
    layoutElement(d)
}
