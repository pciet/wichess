import { modes } from './indexDefinitions.js'
import { button } from './indexLayoutElements.js'

export const indexLandscape = `
<div class="inline modebox" id="landscapemode">
`+ button('modebutton', 'computermode', 'COMPUTER', false) + button('modebutton', 'friendmode', 'FRIEND', false) + button('modebutton', 'timedmode', 'TIMED', false) +`
</div>
<div class="inline">
    <div id="army"></div>
    <div id="modeoptions"></div>
</div>
`
export const indexPortrait = `
<div>
    <div id="army"></div>
    <div id="modeoptions"></div>
</div>
<div class="modebox" id="portraitmode">
`+ button('modebutton', 'computermode', 'COMPUTER', true) + button('modebutton', 'friendmode', 'FRIEND', true) + button('modebutton', 'timedmode', 'TIMED', true) +`
</div>
`

export const indexSkinnyPortrait = `
<div>
    <div id="skinnyheader"></div>
    <div id="army"></div>
    <div id="modeoptions"></div>
</div>
<div class="modebox" id="portraitmode">
`+ button('modebutton', 'computermode', 'COMPUTER', true) + button('modebutton', 'friendmode', 'FRIEND', true) + button('modebutton', 'timedmode', 'TIMED', true) +`
</div>
`
