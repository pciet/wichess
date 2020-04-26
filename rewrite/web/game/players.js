import { layoutSelector } from '../layout.js'

export function writePlayersIndicator() {
    // TODO: clocks are part of this div
    layoutSelector('#players', `
<div class="name" id="blackname">`+GameInformation.Black.Name+`</div>
<div class="name" id="whitename">`+GameInformation.White.Name+`</div>`)
}
