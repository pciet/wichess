import { Board } from '../game.js'
import { Piece, Orientation, pieceImageName } from '../piece.js'
import { NoKind } from '../pieceDefs.js'

import { orientation } from './layouts_orientation.js'

export function writeBoardImages() {
    for (let i = 0; i < Board.length; i++) {
        if ((Board[i] === undefined) || (Board[i].kind === NoKind)) {
            writePieceImage(i, new Piece(NoKind, Orientation.WHITE), true)
            continue
        }
        writePieceImage(i, Board[i], true)
    }
}

export function writePieceImage(boardIndex, piece, noanimation = false) {
    let pIndex = boardIndex
    if (orientation === Orientation.BLACK) {
        // image numbering for images isn't reversed even when the board is.
        // 63 = 0, 62 = 1
        pIndex = 63 - boardIndex
    }

    let h = '<img class="pieceimg pieceback" src="/web/img/'
    h += pieceImageName(pIndex, NoKind, Orientation.WHITE) + '">'

    if (piece.kind !== NoKind) {
        h += '<img class="pieceimg piecefront '
        if (noanimation === false) {
            h += 'animatepiece'
        }
        h += '" id="animp' + boardIndex + '" src="/web/img/'
        h += pieceImageName(pIndex, piece.kind, piece.orientation) + '">'
    }

    document.querySelector('#s'+boardIndex).innerHTML = h
}

export function removePieceImage(boardIndex) {
    const p = document.querySelector('#animp'+boardIndex)
    p.classList.remove('animatepiece')
    p.classList.add('removeanimatepiece')
    //writePieceImage(boardIndex, new Piece(NoKind, Orientation.WHITE))
}
