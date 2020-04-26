import { interiorDimensions } from '../layout.js'

export function sizePickCells() {
    const dims = interiorDimensions(document.querySelector('#picks'))
    document.querySelector('#leftpick').style.width = dims.height + 'px'
    document.querySelector('#rightpick').style.width = dims.height + 'px'
}
