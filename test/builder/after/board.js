import { selectValue } from './builder.js'

import { chessBoard, boardIndex } from '../../wichess/game/board.js'
import { initializeOrientation } from '../../wichess/game/layouts_orientation.js'
import { pieceImageName, Orientation } from '../../wichess/piece.js'
import { NoKind } from '../../wichess/pieceDefs.js'

export const board = []
export const changeBoard = []

export function initBoard() {
    for (let i = 0; i < 64; i++) {
        board[i] = undefined
        changeBoard[i] = undefined
    }
    initializeOrientation(Orientation.WHITE)
    document.querySelector('#board').innerHTML = chessBoard()
    document.querySelector('#changesboard').innerHTML = chessBoard('c')
    setBoardAddPieceHandlers()
}

export function setBoardAddPieceHandlers() {
    for (let i = 0; i < 64; i++) {
        document.querySelector('#s'+i).onclick = addPieceHandler(i)
        document.querySelector('#c'+i).onclick = addChangePieceHandler(i)
    }
}

export function addChangeBoardPieceHandlers() {
    for (let i = 0; i < 64; i++) {
        document.querySelector('#c'+i).onclick = addChangePieceHandler(i)
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
        addBoardPiece(index, selectValue('#pieces'), selectValue('#pieceorientation'), 
            moved, undefined)
    }
}

function addChangePieceHandler(index) {
    return () => {
        addChangeBoardPiece(index, selectValue('#changepieces'), 
            selectValue('#changepieceorientation'))
    }
}

export function addBoardPieces(position) {
    if (position === undefined) {
        return
    }
    for (const p of position) {
        addBoardPiece(boardIndex(p.addr.f, p.addr.r), p.k, p.o, p.m, p.s)
    }
}

export function addChangeBoardPieces(position) {
    if (position === undefined) {
        return
    }
    for (const p of position) {
        addChangeBoardPiece(boardIndex(p.a.f, p.a.r), p.p.k, p.p.o)
    }
}

function addBoardPiece(index, kind, orientation, moved, start) {
    if (kind === NoKind) {
        board[index] = undefined
        document.querySelector('#s'+index).innerHTML = ''
        return
    }

    board[index] = {
        kind: kind,
        orientation: orientation,
        moved: moved,
    }
    if (start !== undefined) {
        board[index].start = {
            file: start.f,
            rank: start.r
        }
    }

    const img = '<img src="img/'+pieceImageName(index, kind, orientation)+'">'
    document.querySelector('#s'+index).innerHTML = img
}

function addChangeBoardPiece(index, kind, orientation) {
    changeBoard[index] = {
        kind: kind,
        orientation: orientation
    }
    document.querySelector('#c'+index).innerHTML = '<img src="img/'+
        pieceImageName(index, kind, orientation)+'">'
}
