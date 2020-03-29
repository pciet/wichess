import { addCSSRuleProperty } from './layout.js'
import { pieceImageName } from './piece.js'
import { BoardAddress, boardAddressToIndex } from './board.js'
import { writeSquareClick } from './gameClick.js'

export function writeBoardImages(board) {
    for (let i = 0; i < board.length; i++) {
        if (board[i] === undefined) {
            continue
        }
        if (board[i].k === 0) {
            continue
        }
        writePieceImage(i, board[i])
    }
}

export function writePieceImage(boardIndex, piece) {
    let h = '<img class="pieceimg" src="/img/'
    h += pieceImageName(boardIndex, piece.k, piece.o) + '">'
    document.querySelector('#s'+boardIndex).innerHTML = h
}

export function removePieceImage(boardIndex) {
    document.querySelector('#s'+boardIndex).innerHTML = ''
}

export function writeBoardMoves(moves) {
    for (let i = 0; i < moves.length; i++) {
        const from = boardAddressToIndex(new BoardAddress(moves[i].f.f, moves[i].f.r))
        const tos = [moves[i].m.length]
        for (let j = 0; j < moves[i].m.length; j++) {
            tos[j] = boardAddressToIndex(new BoardAddress(moves[i].m[j].f, moves[i].m[j].r))
        }
        writeSquareClick(from, tos)
    }
}

export function writeBoardDimension(lowerSquareRatio, upperSquareRatio) {
    let d
    if (window.innerWidth < window.innerHeight) {
        d = window.innerWidth
    } else {
        d = window.innerHeight
    }

    const r = window.innerWidth / window.innerHeight
    if ((r <= upperSquareRatio) && (r > lowerSquareRatio)) {
        d = d * 0.75
        addCSSRuleProperty('#boardrow', 'height', ((d / window.innerHeight) * 100) + '%')
        addCSSRuleProperty('#board', 'height', '100%')
        addCSSRuleProperty('#board', 'width', ((d / window.innerWidth) * 100) + '%')
        return
    }

    // if not the square layout then this:
    addCSSRuleProperty('#board', 'width', ((d / window.innerWidth) * 100) + '%')
    addCSSRuleProperty('#board', 'height', ((d / window.innerHeight) * 100) + '%')
}
