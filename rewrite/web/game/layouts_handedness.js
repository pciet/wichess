import { removeLayout, addLayout } from '../layout.js'
import { layoutPage, landscapeRatio, floatingLandscapeRatio,
    wideFloatingLandscapeRatio, veryWideFloatingLandscapeRatio } from '../game.js'

import * as layouts from './layouts.js'

const handednessKey = 'hand'

// false for normal, true for reversed
let handedness = false

export function initializeHandedness() {
    if (window.localStorage.getItem(handednessKey) === 'true') { swapHandedness() }
}

export function swapHandedness() {
    removeLayout(landscapeRatio)
    removeLayout(floatingLandscapeRatio)
    removeLayout(wideFloatingLandscapeRatio)
    removeLayout(veryWideFloatingLandscapeRatio)

    if (handedness === false) {
        handedness = true
        window.localStorage.setItem(handednessKey, true)
        addLayout(landscapeRatio, layouts.reverseLandscape)
        addLayout(floatingLandscapeRatio, layouts.reverseLandscapeFloating)
        addLayout(wideFloatingLandscapeRatio, layouts.reverseLandscapeWideFloating)
        addLayout(veryWideFloatingLandscapeRatio, layouts.reverseLandscapeVeryWideFloating)
    } else {
        handedness = false
        window.localStorage.setItem(handednessKey, false)
        addLayout(landscapeRatio, layouts.landscape)
        addLayout(floatingLandscapeRatio, layouts.landscapeFloating)
        addLayout(wideFloatingLandscapeRatio, layouts.landscapeWideFloating)
        addLayout(veryWideFloatingLandscapeRatio, layouts.landscapeVeryWideFloating)
    }

    layoutPage()
}

