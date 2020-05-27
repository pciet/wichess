import { addLayout, layout, scaleFont }  from './layout.js'
import * as layouts from './game/layouts.js'

import { writeBoardDimension } from './game/board_dimensions.js'
import { updateBoard } from './game/board_update.js'
import { writeBoardImages } from './game/board_images.js'
import { writeBoardMoves } from './game/board_moves.js'
import { writePlayersIndicator } from './game/players.js'
import { writePieceCharacteristics} from './game/characteristics.js'
import { writeGameCondition } from './game/condition.js'

import { fetchBoardPromise } from './game/fetch_board.js'
import { fetchMovesPromise } from './game/fetch_moves.js'
import { webSocketPromise, webSocketOnMessage } from './game/websocket.js'
import { fetchedMoves } from './game/moves.js'

import { fetchNextMoveSound, muted, 
    setMuteIcon, toggleMute } from './game/audio.js'

import { NoKind } from './pieceDefs.js'
import { State } from './game/state.js'

// The current board is represtened by this Board var.
// The array is indexed by Wisconsin Chess address indices, 0-63.
// If a board square is empty then the matching array element
// is undefined or has kind set to NoKind from piece.js.
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

// Turns are numbered to guarantee synchronization with the host.
export let Turn = GameInformation.Turn

// Available moves are held so that, along with Board, the information can be rewritten into 
// the webpage if the interface needs to be recalculated during a window resize.
export let Moves = []
export function replaceMoves(withMoves) { Moves = withMoves }

// Host requests are started here so the webpage can do some work while the requests are being 
// processed. The promised values are looked at in window.onload later.
const boardPromise = fetchBoardPromise(GameInformation.ID)
const movesPromise = fetchMovesPromise(GameInformation.ID, Turn)
const websocketPromise = webSocketPromise(GameInformation.ID)

fetchNextMoveSound()

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
    for (let i = 0; i < 64; i++) {
        document.querySelector('#s'+i).addEventListener('mouseover', () => {
            const at = Board[i]
            let k
            if (at === undefined) {
                k = NoKind
            } else {
                k = at.kind
            }
            writePieceCharacteristics(k)
        })
    }
    writePlayersIndicator()
    writeBoardImages()
    writeBoardMoves()
    writeGameCondition()
    setMuteIcon(muted())
    scaleFont()

    document.querySelector('#mute').onclick = toggleMute 

    document.querySelector('#back').onclick = () => {
        window.location = '/'
    }

    document.querySelector('#ack').onclick = () => {
        fetch('/acknowledge/'+GameInformation.ID).then(() => {
            window.location = '/'
        })
    }
}