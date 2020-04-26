import { chessBoard } from '../wichess/game/board.js'
import { Pieces } from '../wichess/pieceDefs.js'
import { Orientation } from '../wichess/piece.js'
import { States } from '../wichess/game/state.js'

import { addBoardPiecePickHandlers } from './piecePick.js'
import { addMovesButtons } from './moves.js'
import { prevFromClick, prevToClick } from './prev.js'

// For building a test case, the first thing done is placing pieces
// onto the board to define the position.

export let board = []

for (let i = 0; i < 64; i++) {
    board[i] = undefined
}

export function addPiece(index, kind) {
    board[index] = kind
}

document.querySelector('#board').innerHTML = chessBoard()

let pieceList = '<select id="piecechoice">'
for (let i = 0; i < Pieces.length; i++) {
    pieceList += '<option value="'+i+'">'+Pieces[i].name+'</option>'
}
pieceList += '</select>'
document.querySelector('#pieceselect').innerHTML = pieceList

addBoardPiecePickHandlers()

// For a position, building a test case mostly is defining the set
// of available moves. A "from" square is picked (which must have
// a piece), then a set of squares that can be moved to is picked.

addMovesButtons()

export function addMoveSet(from, tos) {
    console.log(from)
    console.log(tos)
}

// A test case also has other information, like what state the game
// engine should say when calculating the position's moves, or what
// the previous move was (used for en passant).

let stateList = '<select id="statechoice">'
for (let i = 0; i < States.length; i++) {
    stateList += '<option value="'+i+'">'+States[i].name+'</option>'
}
stateList += '</select>'
document.querySelector('#stateselect').innerHTML = stateList

document.querySelector('#prevfrom').onclick = prevFromClick
document.querySelector('#prevto').onclick = prevToClick
