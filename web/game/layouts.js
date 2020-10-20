import { button } from '../button.js'
import { Orientation, pieceTakeImageName } from '../piece.js'
import { NoKind } from '../pieceDefs.js'
import { ct } from '../layout_ct.js'

import { chessBoard } from './board.js'
import { hasComputerPlayer } from './players.js'
import { captureList, spaceCaptureImages } from './captures.js'
import { whitespace } from './layouts_whitespace.js'
import { handedness } from './layouts_handedness.js'
import { orientation } from './layouts_orientation.js'
import { controlsLayout } from './layouts_controls.js'
import { navigationLayout } from './layouts_navigation.js'

const activePlayerChar = '&#x394;'

export function players() {
    let top = 'blackname'
    let bottom = 'whitename'
    if (orientation === Orientation.BLACK) {
        top = 'whitename'
        bottom = 'blackname'
    }
    let names = ct(top, 'playername', false, false) + `
    <div class="noselect vcenter" id="against">against</div>
` + ct(bottom, 'playername', false, false)

    if (hasComputerPlayer() === true) {
        return names
    }

    return `
<div class="inline" id="activeindicators">
    ` + ct(top + 'active', 'activeindicator', false, true, activePlayerChar) +
        `<div></div>` +
        ct(bottom + 'active', 'activeindicator', false, true, activePlayerChar) + `
</div>
<div class="inline">` + names + `
</div>`
}

export function topTakes() {
    if (orientation === Orientation.BLACK) {
        return takes(Orientation.WHITE)
    }
    return takes(Orientation.BLACK)
}

export function bottomTakes() { return takes(orientation) }

function takes(o) {
    const l = captureList(o)
    let prefix = 'w'
    if (o === Orientation.BLACK) {
        prefix = 'b'
    }
    let t = ''
    for (let i = 0; i < 15; i++) {
        const k = l[i]
        if (k === NoKind) {
            break
        }
        t += takeImg(prefix, i, o, k)
    }
    return t
}

function takeImg(prefix, index, or, k) {
    return '<img id="t'+prefix+index+
        '" class="takeimg" src="/web/img/'+pieceTakeImageName(or, k)+'">'
}

export function appendTakePieceImage(or, k, index) {
    // or is the orientation of the captured piece
    let takes, prefix, imgOr
    if (or === Orientation.WHITE) {
        prefix = 'b'
        imgOr = Orientation.BLACK
        if (orientation === Orientation.WHITE) {
            // black took and is top
            takes = 'toptakes'
        } else {
            takes = 'bottomtakes'
        }
    } else {
        prefix = 'w'
        imgOr = Orientation.WHITE
        if (orientation === Orientation.BLACK) {
            takes = 'toptakes'
        } else {
            takes = 'bottomtakes'
        }
    }
    let t = document.querySelector('#'+takes)
    t.innerHTML = t.innerHTML + takeImg(prefix, index, imgOr, k)
    spaceCaptureImages()
}

function landscapeBar(reversed = false) {
    let t = `
<div id="toptakes">` + topTakes() + `</div>
<div id="topspacer"></div>
<div id="playernames">` + players() + `</div>
<div id="navigation">` + navigationLayout() + `</div>
<div id="statusbox">
    <div class="statusverticalmargin"></div>
    ` + ct('status') + `
    <div class="statusverticalmargin"></div>
</div>
<div id="bottomspacer"></div>`

    t += '<div id="controls">' + controlsLayout(reversed) + `</div>
<div id="bottomtakes">` + bottomTakes() + `</div>
`
    return t
}

function squareBar(reversed = false) {
    return `
<div id="squareprom"></div>
<div id="navigation">` + navigationLayout() + `</div>
<div id="statusbox">
    <div class="statusverticalmargin"></div>
    ` + ct('status') + `
    <div class="statusverticalmargin"></div>
</div>
<div id="controls">` + controlsLayout(reversed) + `</div><div></div>`
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

    let t = `<div class="inline" id="board">`

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

function squareBoard() {
    let top = 'blackname'
    let bottom = 'whitename'
    if (orientation === Orientation.BLACK) {
        top = 'whitename'
        bottom = 'blackname'
    }
    const hasComputer = hasComputerPlayer()

    let t = `<div class="squareboardbar">`
    if (hasComputer === false) {
        t += ct(top + 'active', 'activeindicator', true, true, activePlayerChar)
    }
    t += ct(top, 'playername', true, false) + `
        <div class="inline" id="toptakes">` + topTakes() + '</div></div>'

    t += `<div id="board"><div id="boardbox" class="inline">` + chessBoard() + `</div></div>`

    t += `<div class="squareboardbar">`
    if (hasComputer === false) {
        t += ct(bottom + 'active', 'activeindicator', true, true, activePlayerChar)
    }
    t += ct(bottom, 'playername', true, false) + `
        <div class="inline" id="bottomtakes">` + bottomTakes() + '</div></div>'

    return t
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

export function square() {
    if (handedness === false) {
        return `<div class="inline" id="squarebar">` + squareBar() + `</div>` + 
            `<div class="inline">` + squareBoard() + `</div>`
    }
    return `<div class="inline">` + squareBoard() + `</div><div class="inline" id="squarebar">` + 
        squareBar(true) + `</div>`
}

export function portrait() {
    return `
<div id="portraittop">
    <div class="inline" id="statusbox">
        ` + ct('status') + `
    </div>
    <div class="inline" id="portraitnavigation">` + navigationLayout(true) + `</div>
    <div class="inline" id="portraitcontrols">` + controlsLayout(false, true) + `</div>
</div>
` + squareBoard()
}

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
