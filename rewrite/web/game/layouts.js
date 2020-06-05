import { button } from '../button.js'

import { chessBoard } from './board.js'

const landscapeBar = `
<div id="players"></div>
<div id="conditionmargin">
    <div class="conditionspacer"></div>
    <div id="condition"></div>
    <div class="conditionspacer"></div>
</div>
<div id="controls">
    <div class="inline">
        <div>
            <div></div>
            <div id="mute"></div>
            <div></div>
        </div>
        <div id="backmargin">
            <div id="back">
                <div></div>
                <div id="backtext"></div>
                <div></div>
            </div>
        </div>
    </div>
    <div class="inline" id="buttons">
        <div></div>
        <div id="ackmargin">
            <div id="ack">
                <div></div>
                <div id="acktext">&#x2713;</div>
                <div></div>
            </div>
        </div>
    </div>
    <div class="inline" id="descriptiondiv">
        <div id="descriptionmargin"></div>
        <div id="descriptionbottomspacer"></div>
    </div>
</div>
`

const fatLandscapeBar = `
<div id="players"></div>
<div id="fatconditionmargin">
    <div class="conditionspacer"></div>
    <div id="condition"></div>
    <div class="conditionspacer"></div>
</div>
<div id="controls">
    <div id="fatlandscapedescription">
        <div id="descriptionmargin"></div>
    </div>
    <div>
        <div class="inline">
            <div>
                <div></div>
                <div id="mute"></div>
                <div></div>
            </div>
            <div id="backmargin">
                <div id="back">
                    <div></div>
                    <div id="backtext"></div>
                    <div></div>
                </div>
            </div>
        </div>
        <div class="inline" id="buttons">
            <div></div>
            <div id="ackmargin">
                <div id="ack">
                    <div></div>
                    <div id="acktext">&#x2713;</div>
                    <div></div>
                </div>
            </div>
        </div>
    </div>
</div>
`

export const promotion = `
<div class="inline">
    <div class="promotemargin">
        <div class="promotebutton" id="queenprom">
            <div></div>
            <div class="promotetext">Queen</div>
            <div></div>
        </div>
    </div>
    <div class="promotemargin">
        <div class="promotebutton" id="rookprom">
            <div></div>
            <div class="promotetext">Rook</div>
            <div></div>
        </div>
    </div>
</div>
<div class="inline">
    <div class="promotemargin">
        <div class="promotebutton" id="knightprom">
            <div></div>
            <div class="promotetext">Knight</div>
            <div></div>
        </div>
    </div>
    <div class="promotemargin">
        <div class="promotebutton" id="bishopprom">
            <div></div>
            <div class="promotetext">Bishop</div>
            <div></div>
        </div>
    </div>
</div>
`

export const fatLandscape = `
<div class="inline">
` + fatLandscapeBar + `
</div>
<div class="inline" id="board">` + chessBoard() + `
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
