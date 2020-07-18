import { appendTakePieceImage } from './layouts.js'
import { handedness } from './layouts_handedness.js'

import { Orientation } from '../piece.js'
import { NoKind } from '../pieceDefs.js'
import { interiorDimensions } from '../layout.js'

const whiteCaptures = window.GameInformation.White.Captures
const blackCaptures = window.GameInformation.Black.Captures

export function updateCapturedPieces(orientation, kind) {
    // orientation is of the captured piece
    let capArray
    if (orientation === Orientation.WHITE) {
        capArray = blackCaptures
    } else {
        capArray = whiteCaptures
    }
    let i = 0
    for (; i < 15; i++) {
        if (capArray[i] === NoKind) {
            break
        }
    }
    capArray[i] = kind
    appendTakePieceImage(orientation, kind, i)
}

export function captureList(orientation) {
    if (orientation === Orientation.WHITE) {
        return whiteCaptures
    } if (orientation === Orientation.BLACK) {
        return blackCaptures
    }
    throw new Error('orientation ' + orientation + ' not white or black')
}

export function spaceCaptureImages() {
    const dims = interiorDimensions(document.querySelector('#toptakes'))
    let i = 0
    for (; i < whiteCaptures.length; i++) {
        if (whiteCaptures[i] === NoKind) {
            break
        }
    }
    const whiteOffset = (dims.width - dims.height) / i
    i = 0
    for (; i < blackCaptures.length; i++) {
        if (blackCaptures[i] === NoKind) {
            break
        }
    }
    const blackOffset = (dims.width - dims.height) / i
    const width = dims.height + 'px'
    for (let i = 0; i < whiteCaptures.length; i++) {
        if (whiteCaptures[i] === NoKind) {
            break
        }
        const e = document.querySelector('#tw'+i)
        const o = i*whiteOffset + 'px'
        if (handedness === true) {
            e.style.right = o
        } else {
            e.style.left = o
        }
        e.style.zIndex = i
        e.style.width = width
        e.style.height = width
    }
    for (let i = 0; i < blackCaptures.length; i++) {
        if (blackCaptures[i] === NoKind) {
            break
        }
        const e = document.querySelector('#tb'+i)
        const o = i*blackOffset + 'px'
        if (handedness === true) {
            e.style.right = o
        } else {
            e.style.left = o
        }
        e.style.zIndex = i
        e.style.width = width
        e.style.height = width
    }
}

