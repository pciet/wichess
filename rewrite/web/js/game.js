// The webpage look varies depending on window size, done with the code in layout.js.
import { addLayout, layout, scaleFont }  from './layout.js'
import * as layouts from './gameLayouts.js'
import { writeBoardImages, writeBoardMoves, writeBoardDimension } from './gameLayoutGenerate.js'

// A Wisconsin Chess board is a regular 8x8 chess board.
const Board = [64]

// Turns are numbered to guarantee synchronization with the host.
let Turn = GameInformation.Turn

let Moves = []

const getBoardPromise = fetch('/boards/'+GameInformation.ID).then(r => r.json())
const getMovesPromise = fetch('/moves/'+GameInformation.ID+'?turn='+Turn).then(r => r.json())

const websocketPromise = new Promise(resolve => {
    resolve(new WebSocket('ws://'+window.location.host+'/alert/'+GameInformation.Identifier))
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
    Promise.all([getBoardPromise, getMovesPromise]).then(values => {
        for (const a in values[0]) {
            // TODO: host not send p JSON level?
            Board[parseInt(a)] = values[0][a].p
        }
        Moves = values[1]
        layoutPage()
    })
    websocketPromise.then(websocket => {
        websocket.onmessage = event => {
            console.log(JSON.parse(event.data))
        }
    }
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


