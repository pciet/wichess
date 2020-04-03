import { promotion } from './layouts.js'
import { Kind } from '../piece.js'
import { doMove } from './move.js'
import { writePieceCharacteristics } from './characteristics.js'
import { layoutElement } from '../layoutElement.js'
import { removeNewlines } from '../layout.js'

export function showPromotion(reverse = false) {
    const d = document.querySelector('#description')
    d.innerHTML = removeNewlines(promotion)
    layoutElement(d)
    
    document.querySelector('#queenprom').addEventListener('click', promClick(Kind.QUEEN, reverse))
    document.querySelector('#rookprom').addEventListener('click', promClick(Kind.ROOK, reverse))
    document.querySelector('#knightprom').addEventListener('click', promClick(Kind.KNIGHT, reverse))
    document.querySelector('#bishopprom').addEventListener('click', promClick(Kind.BISHOP, reverse))
}

function promClick(kind, reverse) {
    return () => { 
        // a reverse promotion is followed by another move
        doMove(0, 0, kind, reverse)
        // this will remove the promotion html
        writePieceCharacteristics(Kind.NO_KIND)
    }
}
