import { Orientation } from '../piece.js'

export let orientation

export function initializeOrientation(o) {
    orientation = o
}

export function swapOrientation() {
    if (orientation === Orientation.WHITE) {
        orientation = Orientation.BLACK
    } else {
        orientation = Orientation.WHITE
    }
}
