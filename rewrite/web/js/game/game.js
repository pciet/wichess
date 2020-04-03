import { addLayout, layout, scaleFont }  from '../layout.js'
import * as layouts from './layouts.js'

import { writeBoardDimension } from './board_dimensions.js'
import { updateBoard } from './board_update.js'
import { writeBoardImages } from './board_images.js'
import { writeBoardMoves } from './board_moves.js'

import { writeGameCondition } from './condition.js'

import { fetchBoardPromise } from './fetch_board.js'
import { fetchMovesPromise } from './fetch_moves.js'
import { webSocketPromise, webSocketOnMessage } from './websocket.js'

import { fetchedMoves } from './moves.js'

// The current board is represtened by this Board var.
// The array is indexed by Wisconsin Chess address indices, 0-63.
// If a board square is empty then the matching array element
// is undefined or has kind set to Kind.NO_KIND from piece.js.
export let Board = []

// A game is in one of these six states, only continuing from
// the abnormal states of promotion and check.
export const State = {
    NORMAL: 0,
    PROMOTION: 1,
    CHECK: 2,
    CHECKMATE: 3,
    DRAW: 4,
    CONCEDED: 5,
    TIME_OVER: 6,
    REVERSE_PROMOTION: 7
}

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

// Turns are numbered to guarantee synchronization with the host.
export let Turn = GameInformation.Turn

// Available moves are held so that, along with Board, the
// information can be rewritten into the webpage if the interface
// needs to be recalculated during a window resize.
export let Moves = []
export function replaceMoves(withMoves) { Moves = withMoves }

// Host requests are started here so the webpage can do some work
// while the requests are being processed. The promised values are
// looked at in window.onload later.
const boardPromise = fetchBoardPromise(GameInformation.ID)
const movesPromise = fetchMovesPromise(GameInformation.ID, Turn)
const websocketPromise = webSocketPromise(GameInformation.ID)

const lowerSquareRatio = 0.8
const upperSquareRatio = 1.5

addLayout(lowerSquareRatio, layouts.portrait)
addLayout(upperSquareRatio, layouts.square)
addLayout(1.8, layouts.landscape)
addLayout(2, layouts.landscapeFloating)
addLayout(3, layouts.landscapeWideFloating)
addLayout(3.4, layouts.landscapeVeryWideFloating)
addLayout(1000, layouts.unsupportedWindowDimension)

window.onload = () => {
    boardPromise.then(b => {
        Board = b
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

function layoutPage() {
    writeBoardDimension(lowerSquareRatio, upperSquareRatio)
    layout()
    writeBoardImages()
    writeBoardMoves()
    writeGameCondition()
    scaleFont()

    document.querySelector('#back').onclick = () => {
        window.location = '/'
    }

    document.querySelector('#ack').onclick = () => {
        fetch('/acknowledge/'+GameInformation.ID).then(() => {
            window.location = '/'
        })
    }
}
