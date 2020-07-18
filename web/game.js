import { addLayout, layout, layoutSelector, removeAllLayouts, scaleFont }  from './layout.js'
import * as layouts from './game/layouts.js'

import { initializeHandedness, handedness } from './game/layouts_handedness.js'
import { initializeWhitespace } from './game/layouts_whitespace.js'
import { initializeOrientation } from './game/layouts_orientation.js'
import { optionControlsShown,
    showOptions, hideOptions, controlsLayout } from './game/layouts_controls.js'

import { writeBoardDimension } from './game/board_dimensions.js'
import { updateBoard } from './game/board_update.js'
import { writeBoardImages } from './game/board_images.js'
import { writeBoardMoves } from './game/board_moves.js'
import { writePlayersIndicator, writeActivePlayerIndicator,
    hasComputerPlayer } from './game/players.js'
import { writeGameCondition } from './game/condition.js'
import { spaceCaptureImages } from './game/captures.js'
import { squareElement } from './game/board.js'

import { fetchBoardPromise } from './game/fetch_board.js'
import { fetchMovesPromise } from './game/fetch_moves.js'
import { webSocketPromise, webSocketOnMessage } from './game/websocket.js'
import { fetchedMoves, showMoveablePieces, unshowMoveablePieces } from './game/moves.js'

import { fetchNextMoveSound, muted, setMuteIcon, toggleMute } from './game/audio.js'

import { Orientation } from './piece.js'
import { NoKind } from './pieceDefs.js'
import { State } from './game/state.js'

export let PlayerOrientation

if (GameInformation.Player === GameInformation.White.Name) {
    PlayerOrientation = Orientation.WHITE
} else if (GameInformation.Player === GameInformation.Black.Name) {
    PlayerOrientation = Orientation.BLACK
} else {
    throw new Error(GameInformation.Player + ' not white or black player')
}

export let ActiveOrientation

if (GameInformation.Active === 'white') {
    ActiveOrientation = Orientation.WHITE
} else if (GameInformation.Active === 'black') {
    ActiveOrientation = Orientation.BLACK
} else {
    throw new Error('unknown orientation value ' + GameInformation.Active)
}

export function switchActive() {
    if (ActiveOrientation === Orientation.WHITE) {
        ActiveOrientation = Orientation.BLACK
    } else if (ActiveOrientation === Orientation.BLACK) {
        ActiveOrientation = Orientation.WHITE
    } else {
        throw new Error('active orientation ' + ActiveOrientation + ' not white or black')
    }
    writeActivePlayerIndicator()
}

// The current board is represtened by this Board var. The array is indexed by Wisconsin Chess 
// address indices, 0-63. If a board square is empty then the matching array element is undefined 
// or has kind set to NoKind from piece.js.
export let Board

// The condition is the current game state.
export let Condition = State.NORMAL
export function replaceCondition(c) { Condition = c }
export function gameDone() {
    if ((Condition === State.CHECKMATE) ||
        (Condition === State.DRAW) ||
        (Condition === State.CONCEDED) ||
        (Condition === State.TIME_OVER)) {
        return true
    }
    return false
}

if (GameInformation.Conceded === true) {
    Condition = State.CONCEDED
}

// Turns are numbered to guarantee synchronization with the host.
export let Turn = GameInformation.Turn

// Available moves are held so that, along with Board, the information can be rewritten into 
// the webpage if the interface needs to be recalculated during a window resize.
export let Moves = []
export function replaceMoves(withMoves) { Moves = withMoves }

export let PreviousMove = window.GameInformation.Previous
export function replacePreviousMove(from, to) {
    PreviousMove = {
        from: from,
        to: to
    }
}

// Host requests are started here so the webpage can do some work while the requests are being 
// processed. The promised values are looked at in window.onload later.
const boardPromise = fetchBoardPromise(GameInformation.ID)
const movesPromise = fetchMovesPromise(GameInformation.ID, Turn)
const websocketPromise = webSocketPromise(GameInformation.ID)

fetchNextMoveSound()

const lowerSquareRatio = 0.8
const upperSquareRatio = 1.2

export const landscapeRatio = 1.8
export const floatingLandscapeRatio = 2
export const wideFloatingLandscapeRatio = 3
export const veryWideFloatingLandscapeRatio = 3.4

function addLayouts() {
    // TODO: portrait and square
    addLayout(lowerSquareRatio, layouts.portrait)
    addLayout(upperSquareRatio, layouts.square)
    addLayout(landscapeRatio, layouts.landscape())
    addLayout(floatingLandscapeRatio, layouts.landscapeFloating())
    addLayout(wideFloatingLandscapeRatio, layouts.landscapeWideFloating())
    addLayout(veryWideFloatingLandscapeRatio, layouts.landscapeVeryWideFloating())
    addLayout(1000, layouts.unsupportedWindowDimension)
}

// Some layouts are constructed based on user settings, so the layout list has to be reset when
// settings change.
export function resetLayouts() {
    removeAllLayouts()
    addLayouts()
}

window.onload = () => {
    boardPromise.then(b => {
        Board = b
        initializeHandedness()
        initializeWhitespace()
        initializeOrientation(PlayerOrientation)
        addLayouts()
        layoutPage()
        return movesPromise
    }).then(m => {
        fetchedMoves(m)
        return websocketPromise
    }).then(w => {
        w.onmessage = webSocketOnMessage
    })
}

let resizeWait

window.onresize = () => {
    clearTimeout(resizeWait)
    resizeWait = setTimeout(layoutPage, 150)
}

export let selectedPiece

export function layoutPage() {
    writeBoardDimension(lowerSquareRatio, upperSquareRatio)
    layout()

    for (let i = 0; i < 64; i++) {
        const s = squareElement(i)
        s.onclick = () => {
            const p = Board[i]
            if ((p === undefined) || (p.kind === NoKind)) {
                selectedPiece = undefined
            } else {
                selectedPiece = p.kind
            }

            if (s.moveClickFunc !== undefined) {
                s.moveClickFunc()
            }
        }
    }

    writePlayersIndicator()
    writeBoardImages()

    if (optionControlsShown === true) {
        showOptions()
    } else {
        hideOptions()
    }

    document.querySelector('#toptakes').innerHTML = layouts.topTakes()
    document.querySelector('#bottomtakes').innerHTML = layouts.bottomTakes()
    spaceCaptureImages()

    writeBoardMoves()
    writeGameCondition()
    scaleFont()

    document.body.classList.add('visible')
}
