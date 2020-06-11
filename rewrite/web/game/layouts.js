import { button } from '../button.js'
import { Orientation } from '../piece.js'

import { chessBoard } from './board.js'
import { whitespace } from './layouts_whitespace.js'
import { handedness } from './layouts_handedness.js'
import { orientation } from './layouts_orientation.js'

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

export function players() {
    let top = 'blackname'
    let bottom = 'whitename'
    if (orientation === Orientation.BLACK) {
        top = 'whitename'
        bottom = 'blackname'
    }
    return ct(top, 'playername', false, false) + `
<div>
    <div></div>
    <div class="noselect" id="against">against</div>
    <div></div>
</div>
` + ct(bottom, 'playername', false, false)
}

function landscapeBar(reversed = false) {
    let t = `
<div id="toptakes"></div>
<div id="playernames">` + players() + `</div>
<div id="selectedpiece"></div>
<div id="navigation">
    <div class="inline"></div>
    ` + ct('ack', '', true, true, '&#x2713;') + `
    <div class="inline"></div>
</div>
<div id="statusbox">
    <div class="statusverticalmargin"></div>
    ` + ct('status') + `
    <div class="statusverticalmargin"></div>
</div>
<div id="controls">`

    const gameControls = `
    <div class="inline">` +
        ct('showmoves', 'control', false, true, '&#x2318;') +
        ct('showprev', 'control', false, true, '&#x21BA;') + `
    </div>`

    const interfaceControls = `
    <div class="inline">
        <div>
            ` + ct('swapinterface', 'control', true, true, '&#x2194;') + 
            ct('swapplayers', 'control', true, true, '&#x2195;') + `
        </div>
        <div>
            ` + ct('mute', 'control', true) + 
            ct('whitespace', 'control', true, true, '&#x21F2;') + `
        </div>
    </div>`

    if (reversed === true) {
        t += gameControls + interfaceControls
    } else {
        t += interfaceControls + gameControls
    }
    t += `
</div>
<div id="bottomtakes"></div>
`
    return t
}

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

function board(reversed = false) {
    const box = `<div id="boardbox" class="inline">` + chessBoard() + `</div>`

    let t = `
<div class="inline" id="board">`

    if (whitespace === false) {
        return t + box + '</div>'
    }

    t += `
    <div class="boardvertspace"></div>
    <div>
    `
    const spacer = `<div class="boardhorzspace inline"></div>`

    if (reversed === false) {
        t += box + spacer
    } else {
        t += spacer + box
    }
    return t + `   
    </div>
    <div class="boardvertspace"></div>
</div>`
}

export function landscape() {
    if (handedness === false) {
        return `<div class="inline">` + landscapeBar() + `</div>` + board()
    }
    return board(true) + `<div class="inline">` + landscapeBar(true) + `</div>`
}

function floatingLandscape(sideClassName) {
    return `
<div class="inline">
    <div class="` + sideClassName + ` inline"></div>
    <div class="inline" id="floatingbar">
    ` + landscapeBar() + `
    </div>
    <div class="` + sideClassName + ` inline"></div>
</div>` + board()
}

function reverseFloatingLandscape(sideClassName) {
    return board(true) + `
<div class="inline">
    <div class="` + sideClassName + ` inline"></div>
    <div class="inline" id="floatingbar">
    ` + landscapeBar(true) + `
    </div>
    <div class="` + sideClassName + ` inline"></div>
</div>`
}

export function landscapeFloating() {
    if (handedness === false) {
        return floatingLandscape('floatingside')
    }
    return reverseFloatingLandscape('floatingside')
}

export function landscapeWideFloating() {
    if (handedness === false) {
        return floatingLandscape('widefloatingside')
    }
    return reverseFloatingLandscape('widefloatingside')
}

export function landscapeVeryWideFloating() {
    if (handedness === false) {
        return floatingLandscape('verywidefloatingside')
    }
    return reverseFloatingLandscape('verywidefloatinglandscape')
}

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
