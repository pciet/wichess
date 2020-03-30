import { Board } from './game.js'
import { Kind, pieceImageName } from '../piece.js'

export function writeBoardImages() {
    for (let i = 0; i < Board.length; i++) {
        if (Board[i] === undefined) {
            continue
        }
        if (Board[i].kind === Kind.NO_KIND) {
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
    document.querySelector('#s'+boardIndex).innerHTML = ''
}
