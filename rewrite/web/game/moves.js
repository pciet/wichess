import { PlayerOrientation, ActiveOrientation, Condition, replaceCondition, 
    gameDone, replaceMoves, Moves } from '../game.js'

import { boardAddressToIndex } from './board.js'
import { State } from './state.js'
import { replaceAndWriteGameCondition } from './condition.js'
import { replaceAndWriteBoardMoves } from './board_moves.js'
import { showPromotion } from './promotion.js'

export function fetchedMoves(moves) {
    if (moves.state !== Condition) {
        replaceAndWriteGameCondition(moves.state)
    }

    if ((gameDone() === true) || (PlayerOrientation !== ActiveOrientation)) {
        return
    }

    if (moves.state === State.PROMOTION) {
        showPromotion()
    } else {
        replaceAndWriteBoardMoves(moves.moveSets)
    }
}

export function addShowMovesHandler() {
    const sm = document.querySelector('#showmoves')
    sm.onclick = () => {
        if (ActiveOrientation !== PlayerOrientation) {
            return
        }
        showMoveablePieces()
        sm.onclick = () => {
            unshowMoveablePieces()
        }
    }
}

export function showMoveablePieces() {
    for (const move of Moves) {
        document.querySelector('#s'+boardAddressToIndex(move.from)).classList.add('moveable')
    }
}

export function unshowMoveablePieces() {
    for (const move of Moves) {
        document.querySelector('#s'+boardAddressToIndex(move.from)).classList.remove('moveable')
    }
    addShowMovesHandler()
}
