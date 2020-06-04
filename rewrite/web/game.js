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
            let k, o
            if (at === undefined) {
                k = NoKind
            } else {
                k = at.kind
                o = at.orientation
            }
            writePieceCharacteristics(k, o)
        })
    }
    writePlayersIndicator()
    writeBoardImages()
    writeBoardMoves()
    writeGameCondition()
    setMuteIcon(muted())
    scaleFont()

    document.querySelector('#mute').onclick = toggleMute 

    const back = document.querySelector('#back')
    if (window.location.pathname.includes('people')) {
        // In people mode back is disabled because the game must be completed or conceded before
        // going back to the index page. The button is changed to a concede button.
        document.querySelector('#backtext').innerHTML = '&#x2718;'
        back.onclick = () => {
            fetch('/concede/'+GameInformation.ID).then(() => { window.location = '/' })
        }
    } else {
        document.querySelector('#backtext').innerHTML = '&#8592;'
        back.onclick = () => { window.location = '/' }

    }
    back.addEventListener('mouseenter', () => { 
        document.querySelector('#back').classList.add('backhover')
    })
    back.addEventListener('mouseleave', () => {
        document.querySelector('#back').classList.remove('backhover') 
    })

    document.querySelector('#ack').onclick = () => {
        fetch('/acknowledge/'+GameInformation.ID).then(() => {
            window.location = '/'
        })
    }
}
