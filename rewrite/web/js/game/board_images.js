import { Board } from './game.js'
import { Kind, Piece, Orientation, pieceImageName } from '../piece.js'

export function writeBoardImages() {
    console.log(Board.length)
    for (let i = 0; i < Board.length; i++) {
        if ((Board[i] === undefined) || (Board[i].kind === Kind.NO_KIND)) {
            writePieceImage(i, new Piece(Kind.NO_KIND, Orientation.WHITE))
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
    writePieceImage(boardIndex, new Piece(Kind.NO_KIND, Orientation.WHITE))
}
