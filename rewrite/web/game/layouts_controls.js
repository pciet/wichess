import { resetLayouts, layoutPage, gameDone } from '../game.js'
import { layoutSelector } from '../layout.js'

import { ct } from './layouts_ct.js'
import { swapHandedness, handedness } from './layouts_handedness.js'
import { swapWhitespace } from './layouts_whitespace.js'
import { swapOrientation } from './layouts_orientation.js'
import { addNavigationClickHandlers, navigationLayout } from './layouts_navigation.js'
import { unshowPreviousMove, unshowMoveablePieces, 
    addShowMovesHandler, addShowPreviousMoveHandler } from './moves.js'
import { toggleMute, setMuteIcon, muted } from './audio.js'

export let optionControlsShown = false

export function addControlsClickHandlers() {
    if (optionControlsShown === false) {
        addShowMovesHandler()
        addShowPreviousMoveHandler()
        document.querySelector('#showoptions').onclick = showOptions
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

function showOptions() {
    unshowPreviousMove()
    unshowMoveablePieces()
    optionControlsShown = true
    layoutSelector('#controls', controlsLayout(handedness))
    setMuteIcon(muted())
    addControlsClickHandlers()
    layoutSelector('#navigation', navigationLayout())
    addNavigationClickHandlers()
}

const showOptionsButton = `
<div class="inline">` +
    ct('showoptions', 'control', false, true, '&#x2022;') + `
</div>`

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

export function controlsLayout(reversed = false) {
    let t = ''
    if (optionControlsShown === false) {
        if (reversed === true) {
            t += gameControls + showOptionsButton
        } else {
            t += showOptionsButton + gameControls
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
