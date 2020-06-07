import { button } from '../button.js'

import { chessBoard } from './board.js'

// ct stands for "centered text", a div structure compatible with layout.js that can be used
// to vertically and horizontally center text in a styled box. Use the padding of #[id]margin
// for margin, text is in #[id]text, and any classes in the argument are applied to #[id].
// The click handler should be applied to #[id] to create a button.
function ct(id, classes = '', inline = false, noselect = true, text = '') {
    let t = '<div '
    if (inline === true) {
        t += 'class="inline" '
    }
    t += `id="`+id+`margin">
    <div `
    if (classes !== '') {
    t += 'class="'+classes+'" '
    }
    t += `id="`+id+`">
        <div></div>
    <div `
    if (noselect === true) {
        t += 'class="noselect" '
    }
    t += `id="`+id+`text">`
    if (text !== '') {
        t += text
    }
    return t + `</div>
        <div></div>
    </div>
</div>
`
}

export const players = ct('blackname', 'playername', false, false) + `
<div>
    <div></div>
    <div class="noselect" id="against">against</div>
    <div></div>
</div>
` + ct('whitename', 'playername', false, false)

const landscapeBar = `
<div id="navigation">
    ` + ct('backconcede', 'navbutton', true) + ct('ack', 'navbutton', true) + `
</div>
<div id="playernames">` + players + `</div>
<div id="controls">
    ` + ct('showmoves', 'control', true) + `
    <div class="inline">
        ` + ct('swapinterface', 'control') + ct('mute', 'control') + `
    </div>
</div>
<div id="statusbox">
    <div class="statusverticalmargin"></div>
    ` + ct('status') + `
    <div class="statusverticalmargin"></div>
</div>
`

// Promotion temporarily replaces #playernames with the choice buttons.
export const promotion = `
<div class="inline">
    ` + ct('queenprom', 'promotebutton', false, true, 'Queen') + 
    ct('rookprom', 'promotebutton', false, true, 'Rook') + `
</div>
<div class="inline">
    ` + ct('knightprom', 'promotebutton', false, true, 'Knight') +
    ct('bishopprom', 'promotebutton', false, true, 'Bishop') + `
</div>
`

export const landscape = `
<div class="inline">` + landscapeBar + `</div>
<div class="inline" id="board">` + chessBoard() + `</div>`

export const reverseLandscape = `
<div class="inline" id="board">` + chessBoard() + `</div>
<div class="inline">` + landscapeBar + `</div>`

function floatingLandscape(sideClassName) {
    return `
<div class="inline">
    <div class="` + sideClassName + ` inline"></div>
    <div class="inline" id="floatingbar">
    ` + landscapeBar + `
    </div>
    <div class="` + sideClassName + ` inline"></div>
</div>
<div class="inline" id="board">` + chessBoard() + `</div>`
}

function reverseFloatingLandscape(sideClassName) {
    return `
<div class="inline" id="board">` + chessBoard() + `</div>
<div class="inline">
    <div class="` + sideClassName + ` inline"></div>
    <div class="inline" id="floatingbar">
    ` + landscapeBar + `
    </div>
    <div class="` + sideClassName + ` inline"></div>
</div>`
}

export const landscapeFloating = floatingLandscape('floatingside')
export const landscapeWideFloating = floatingLandscape('widefloatingside')
export const landscapeVeryWideFloating = floatingLandscape('verywidefloatingside')

export const reverseLandscapeFloating = reverseFloatingLandscape('floatingside')
export const reverseLandscapeWideFloating = reverseFloatingLandscape('widefloatingside')
export const reverseLandscapeVeryWideFloating = 
    reverseFloatingLandscape('verywidefloatinglandscape')

export const square = `
<div id="squaretop">
    <div class="inline" id="description"></div>
    <div class="inline">
        <div id="squarebuttonspacer"></div>
        <div>
            <div class="inline" id="squarenav">
            ` + button('navbutton', 'ack', '&#x2713;', true) + 
                button('navbutton', 'back', '&#8592;', true) + `
            </div>
            <div class="inline" id="mute"></div>
        </div>
        <div id="squarebuttonspacer"></div>
    </div>
</div>
<div id="boardrow">
    <div class="inline">
        <div id="players"></div>
        <div>
            <div></div>
            <div id="condition"></div>
            <div></div>
        </div>
    </div>
    <div class="inline" id="board">` + chessBoard() + `</div>
</div>
`

export const portrait = `
<div>
    <div id="mute"></div>
` + button('navbutton', 'ack', '&#x2713;', true) + 
    button('navbutton', 'back', '&#8592;', true) + `
</div>
<div>
    <div class="inline" id="players"></div>
    <div class="inline" id="description"></div>
</div>
<div>
    <div class="inline" id="condition"></div>
</div>
<div id="board">` + chessBoard() + `
</div>
`

export const unsupportedWindowDimension = `
<div class="inline"></div>
<div class="inline" id="unsupported">
    <div></div>
    <div>This window is set to an unsupported dimension.<br>If this is a desktop window then 
    resize it.</div>
    <div></div>
</div>
<div class="inline"></div>
`
