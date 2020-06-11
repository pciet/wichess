const handednessKey = 'hand'

// false for normal, true for reversed
export let handedness = false

export function initializeHandedness() {
    if (window.localStorage.getItem(handednessKey) === 'true') { swapHandedness() }
}

export function swapHandedness() {
    if (handedness === false) {
        handedness = true
        window.localStorage.setItem(handednessKey, true)
    } else {
        handedness = false
        window.localStorage.setItem(handednessKey, false)
    }
}

