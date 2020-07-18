import { resetLayouts, layoutPage, gameDone, selectedPiece } from '../game.js'
import { layoutSelector } from '../layout.js'
import { Pieces } from '../pieceDefs.js'
import { ct } from '../layout_ct.js'

import { swapHandedness, handedness } from './layouts_handedness.js'
import { swapWhitespace } from './layouts_whitespace.js'
import { swapOrientation } from './layouts_orientation.js'
import { addNavigationClickHandlers, navigationLayout } from './layouts_navigation.js'
import { unshowPreviousMove, unshowMoveablePieces, 
    addShowMovesHandler, addShowPreviousMoveHandler } from './moves.js'
import { toggleMute, setMuteIcon, muted } from './audio.js'

export let optionControlsShown = false

// TODO: click handlers should be a separate group of files from layout

export function addControlsClickHandlers() {
    if (optionControlsShown === false) {
        addShowMovesHandler()
        addShowPreviousMoveHandler()
        document.querySelector('#showoptions').onclick = showOptions
        document.querySelector('#piececard').onclick = () => {
            if (selectedPiece === undefined) { return }
            window.open('/details?p=' + Pieces[selectedPiece])
        }
        return
    }

    document.querySelector('#showoptions').onclick = hideOptions

    document.querySelector('#swapinterface').onclick = () => {
        swapHandedness()
        resetLayouts()
        layoutPage()
    }

    document.querySelector('#whitespace').onclick = () => {
        swapWhitespace()
        resetLayouts()
        layoutPage()
    }

    document.querySelector('#swapplayers').onclick = () => {
        swapOrientation()
        resetLayouts()
        layoutPage()
    }

    document.querySelector('#mute').onclick = toggleMute 
}

export function hideOptions() {
    optionControlsShown = false
    layoutSelector('#controls', controlsLayout(handedness))
    addControlsClickHandlers()
    layoutSelector('#navigation', navigationLayout())
    addNavigationClickHandlers()
}

export function showOptions() {
    if (optionControlsShown === false) {
        unshowPreviousMove()
        unshowMoveablePieces()
        optionControlsShown = true
    }
    layoutSelector('#controls', controlsLayout(handedness))
    setMuteIcon(muted())
    addControlsClickHandlers()
    layoutSelector('#navigation', navigationLayout())
    addNavigationClickHandlers()
}

const showOptionsButton = `
<div id="showoptionsdiv" class="inline">` +
    ct('showoptions', 'control', false, true, '&#x2022;') + `
</div>`

function gameControls(reversed = false) {
    let t = '<div class="inline">'

    const card = ct('piececard', 'control', true, true, '?')
    const board = '<div class="inline">' +
        ct('showmoves', 'control', false, true, '&#x2318;') +
        ct('showprev', 'control', false, true, '&#x21BA;') + '</div>'

    if (reversed === true) {
        return t + board + card + '</div>'
    }

    return t + card + board + '</div>'
}

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

export function controlsLayout(reversed = false) {
    let t = ''
    if (optionControlsShown === false) {
        if (reversed === true) {
            t += gameControls(reversed) + showOptionsButton
        } else {
            t += showOptionsButton + gameControls(reversed)
        }
    } else {
        if (reversed === true) {
            t += interfaceControls + showOptionsButton
        } else {
            t += showOptionsButton + interfaceControls
        }
    }
    return t
}
