import { removeNewlines } from '../layout.js'
import { layoutElement } from '../layoutElement.js'

export function writePlayersIndicator() {
    // TODO: clocks are part of this div
    const p = removeNewlines(`
<div class="name" id="blackname">`+GameInformation.Black.Name+`</div>
<div class="name" id="whitename">`+GameInformation.White.Name+`</div>`)
    const e = document.querySelector('#players')
    e.innerHTML = p
    layoutElement(e)
}
