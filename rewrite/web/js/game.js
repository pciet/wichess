// The webpage look varies depending on window size, done with the code in layout.js.
import { addLayout, layout, scaleFont }  from './layout.js'
import * as layouts from './gameLayouts.js'

import { writeBoardImages, writeBoardMoves, writeBoardDimension } from './gameLayoutGenerate.js'
import { updateBoard } from './gameUpdate.js'
import { State } from './gameState.js'

// A Wisconsin Chess board is a regular 8x8 chess board.
export const Board = []

// Turns are numbered to guarantee synchronization with the host.
export let Turn = GameInformation.Turn

export let Moves = []

const getBoardPromise = fetch('/boards/'+GameInformation.ID).then(r => r.json())
const getMovesPromise = fetch('/moves/'+GameInformation.ID+'?turn='+Turn).then(r => r.json())
const websocketPromise = new Promise(resolve => {
    resolve(new WebSocket('ws://'+window.location.host+'/alert/'+GameInformation.ID))
})

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
    Promise.all([getBoardPromise, getMovesPromise, websocketPromise]).then(values => {
        for (const a in values[0]) {
            // TODO: host not send p JSON level?
            Board[parseInt(a)] = values[0][a].p
        }
        layoutPage()
        fetchedMoves(values[1])
        values[2].onmessage = event => {
            updateBoard(Board, JSON.parse(event.data))
            fetch('/moves/'+GameInformation.ID+'?turn='+Turn).then(r => r.json()).then(m => {
                fetchedMoves(m)
            })
        }
    })
}

function fetchedMoves(moves) {
    switch (moves) {
    case State.CHECKMATE:
        console.log('checkmate')
        return
    case State.DRAW:
        console.log('draw')
        return
    case State.CONCEDED:
        console.log('conceded')
        return
    case State.TIME_OVER:
        console.log('time over')
        return
    }
    Moves = moves
    writeBoardMoves(Moves)
}

window.onresize = layoutPage

function layoutPage() {
    writeBoardDimension(lowerSquareRatio, upperSquareRatio)
    layout()
    writeBoardImages(Board)
    writeBoardMoves(Moves)
    scaleFont()

    document.querySelector('#back').onclick = () => {
        window.location = '/'
    }
}
