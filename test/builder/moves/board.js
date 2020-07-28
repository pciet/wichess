import { selectValue } from './builder.js'

import { chessBoard, boardIndex } from '../../wichess/game/board.js'
import { initializeOrientation } from '../../wichess/game/layouts_orientation.js'
import { pieceImageName, Orientation } from '../../wichess/piece.js'
import { NoKind } from '../../wichess/pieceDefs.js'

export const board = []

export function initBoard() {
    for (let i = 0; i < 64; i++) {
        board[i] = undefined
    }
    initializeOrientation(Orientation.WHITE)
    document.querySelector('#board').innerHTML = chessBoard()
    setBoardAddPieceHandlers()
}

export function setBoardAddPieceHandlers() {
    for (let i = 0; i < 64; i++) {
        document.querySelector('#s'+i).onclick = addPieceHandler(i)
    }
}

function addPieceHandler(index) {
    return () => {
        let moved
        const movedNum = selectValue('#moved')
        if (movedNum === 0) {
            moved = false
        } else if (movedNum === 1) {
            moved = true
        } else {
            throw new Error('#moved value ' + movedNum + ' not 0 or 1')
        }
        addBoardPiece(index, selectValue('#pieces'), selectValue('#pieceorientation'), moved)
    }
}

export function addBoardPieces(position) {
    for (const p of position) {
        addBoardPiece(boardIndex(p.addr.f, p.addr.r), p.k, p.o, p.m)
    }
}

function addBoardPiece(index, kind, orientation, moved) {
    if (kind === NoKind) {
        board[index] = undefined
        document.querySelector('#s'+index).innerHTML = ''
        return
    }

    board[index] = {
        kind: kind,
        orientation: orientation,
        moved: moved
    }
    const img = '<img src="img/'+pieceImageName(index, kind, orientation)+'">'
    document.querySelector('#s'+index).innerHTML = img
}
