import { chessBoard } from './board.js'
import { button } from '../button.js'

const landscapeBar = `
<div id="players">
    <div>
        <div class="name inline" id="blackname">Black</div>
        <div class="clock inline" id="blackclock">5:00</div>
    </div>
    <div>
        <div class="name inline" id="whitename">White</div>
        <div class="clock inline" id="whiteclock">5:00</div>
    </div>
</div>
<div>
    <div></div>
    <div id="condition"></div>
    <div></div>
</div>
<div id="controls">
    <div class="inline">
        ` + button('optionbutton', 'layoutbutton', 'L', false) + button('optionbutton', 'mute', 'M', false) + `
    </div>
    <div class="inline" id="buttons">
        ` + button('navbutton', 'ack', 'Acknowledge', false) + button('navbutton', 'back', 'Back', false) + `
    </div>
    <div class="inline" id="description">
        <div id="piece">Super Rook</div>
        <div class="characteristic">Characteristic 1</div>
        <div class="characteristic">Characteristic 2</div>
    </div>
</div>
`

export const promotion = `
<div class="inline">
` + button('promotebutton', 'queenprom', 'Queen', false) + button('promotebutton', 'rookprom', 'Rook', false) + `
</div>
<div class="inline">
` + button('promotebutton', 'knightprom', 'Knight', false) + button('promotebutton', 'bishopprom', 'Bishop', false) + `
</div>
`

export const landscape = `
<div class="inline">
` + landscapeBar + `
</div>
<div class="inline" id="board">` + chessBoard() + `
</div>
`

function floatingLandscape(sideClassName) {
    return `
    <div class="inline">
        <div class="` + sideClassName + ` inline"></div>
        <div class="inline" id="floatingbar">
        ` + landscapeBar + `
        </div>
        <div class="` + sideClassName + ` inline"></div>
    </div>
    <div class="inline" id="board">` + chessBoard() + `
    </div>        
    `
}

export const landscapeFloating = floatingLandscape('floatingside')
export const landscapeWideFloating = floatingLandscape('widefloatingside')
export const landscapeVeryWideFloating = floatingLandscape('verywidefloatingside')

export const square = `
<div id="squaretop">
    <div class="inline" id="description">
        <div></div>
        <div id="piece">Super Rook</div>
        <div class="characteristic">Characteristic 1</div>
        <div class="characteristic">Characteristic 2</div>
        <div></div>
    </div>
    <div class="inline">
        <div id="squarebuttonspacer"></div>
        <div>
            <div class="inline" id="squarenav">
            ` + button('navbutton', 'ack', 'Acknowledge', true) + button('navbutton', 'back', 'Back', true) + `
            </div>
            <div class="inline">
            ` + button('optionbutton', 'layoutbutton', 'L', false) + button('optionbutton', 'mute', 'M', false) + `
            </div>
        </div>
        <div id="squarebuttonspacer"></div>
    </div>
</div>
<div id="boardrow">
    <div class="inline">
        <div id="players">
            <div>
                <div class="name inline" id="blackname">Black</div>
                <div class="clock inline" id="blackclock">5:00</div>
            </div>
            <div>
                <div class="name inline" id="whitename">White</div>
                <div class="clock inline" id="whiteclock">5:00</div>
            </div>
        </div>
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
` + button('optionbutton', 'layoutbutton', 'L', true) + button('optionbutton', 'mute', 'M', true) + button('navbutton', 'ack', 'Acknowledge', true) + button('navbutton', 'back', 'Back', true) + `
</div>
<div>
    <div class="portraitname inline" id="blackname">Black</div>
    <div class="clock inline" id="blackclock">5:00</div>
    <div class="inline" id="portraitdesc">
        <div id="piece">Super Rook</div>
        <div>
            <div class="characteristic inline">Characteristic 1</div>
            <div class="characteristic inline">Characteristic 2</div>
        </div>
    </div>
</div>
<div>
    <div class="portraitname inline" id="whitename">White</div>
    <div class="clock inline" id="blackclock">5:00</div>
    <div class="inline" id="condition"></div>
</div>
<div id="board">` + chessBoard() + `
</div>
`

export const unsupportedWindowDimension = `
<div class="inline"></div>
<div class="inline" id="unsupported">
    <div></div>
    <div>This window is set to an unsupported dimension.<br>If this is a desktop window then resize it.</div>
    <div></div>
</div>
<div class="inline"></div>
`
