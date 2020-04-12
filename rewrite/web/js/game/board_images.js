import { Board } from './game.js'
import { Piece, Orientation, pieceImageName } from '../piece.js'
import { NoKind } from '../pieceDefs.js'

export function writeBoardImages() {
    for (let i = 0; i < Board.length; i++) {
        if ((Board[i] === undefined) || (Board[i].kind === NoKind)) {
            writePieceImage(i, new Piece(NoKind, Orientation.WHITE))
            continue
        }
        writePieceImage(i, Board[i])
    }
}

export function writePieceImage(boardIndex, piece) {
    let h = '<img class="pieceimg" src="/img/'
    h += pieceImageName(boardIndex, piece.kind, piece.orientation) + '">'
    document.querySelector('#s'+boardIndex).innerHTML = h
}

export function removePieceImage(boardIndex) {
    writePieceImage(boardIndex, new Piece(NoKind, Orientation.WHITE))
}
