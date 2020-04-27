import { layoutSelector } from '../layout.js'

export function writePlayersIndicator() {
    layoutSelector('#players', `
<div>
    <div></div>
    <div class="name" id="blackname">`+GameInformation.Black.Name+`</div>
    <div></div>
</div>
<div id="versus">against</div>
<div>
    <div></div>
    <div class="name" id="whitename">`+GameInformation.White.Name+`</div>
    <div></div>
</div>
`)
}
