import { setBoardAddPieceHandlers, board } from './board.js'

import { NoKind } from '../../wichess/pieceDefs.js'
import { boardAddressToIndex, boardIndexToAddress } from '../../wichess/game/board.js'

export function addStartHandler() {
    document.querySelector('#pickstart').onclick = () => {
        addBoardStartClickHandlers()
        addStartCancelButton()
    }
}

let starting

function addBoardStartClickHandlers() {
    for (let i = 0; i < 64; i++) {
        document.querySelector('#s'+i).onclick = () => {
            if ((board[i] === undefined) || (board[i].kind == NoKind)) {
                cancelStartSelecting()
                return
            }
            document.querySelector('#s'+i).classList.add('starting')
            starting = i
            if (board[i].start !== undefined) {
                document.querySelector('#s'+
                    boardAddressToIndex(board[i].start)).classList.add('started')
            }
            addBoardStartAtClickHandlers()
        }
    }
}

function addBoardStartAtClickHandlers() {
    for (let i = 0; i < 64; i++) {
        document.querySelector('#s'+i).onclick = () => {
            document.querySelector('#s'+starting).classList.remove('starting')
            if (board[starting].start !== undefined) {
                document.querySelector('#s'+
                    boardAddressToIndex(board[starting].start)).classList.remove('started')
            }
            board[starting].start = boardIndexToAddress(i)
            starting = undefined
            cancelStartSelecting()
        }
    }
}

function addStartCancelButton() {
    const e = document.querySelector('#pickstart')
    e.innerText = 'Cancel'
    e.onclick = () => {
        cancelStartSelecting()
    }
}

function cancelStartSelecting() {
    if (starting !== undefined) {
        document.querySelector('#s'+starting).classList.remove('starting')
        if (board[starting].start !== undefined) {
            document.querySelector('#s'+
                boardAddressToIndex(board[starting].start)).classList.remove('started')
        }
    }
    document.querySelector('#pickstart').innerText = 'Select Start'
    addStartHandler()
    setBoardAddPieceHandlers()
}
