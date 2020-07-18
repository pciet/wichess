import { PlayerOrientation, ActiveOrientation, Condition, replaceCondition, 
    gameDone, replaceMoves, Moves, PreviousMove } from '../game.js'

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

export let moveablePiecesShown = false

export function addShowMovesHandler() {
    const sm = document.querySelector('#showmoves')
    sm.onclick = () => {
        if ((ActiveOrientation !== PlayerOrientation) || (gameDone() === true)) {
            return
        }
        showMoveablePieces()
        sm.onclick = () => {
            unshowMoveablePieces()
        }
    }
}

function unmoveableSquareIndices() {
    const indices = []
loop:
    for (let i = 0; i < 64; i++) {
        for (const move of Moves) {
            if (boardAddressToIndex(move.from) === i) {
                continue loop
            }
        }
        indices.push(i)
    }
    return indices
}

export function showMoveablePieces() {
    unshowPreviousMove()
    moveablePiecesShown = true
    for (const i of unmoveableSquareIndices()) {
        document.querySelector('#s'+i).classList.add('unmoveable')
    }
}

export function unshowMoveablePieces() {
    moveablePiecesShown = false
    for (const i of unmoveableSquareIndices()) {
        document.querySelector('#s'+i).classList.remove('unmoveable')
    }
    addShowMovesHandler()
}

export let previousMoveShown = false

export function addShowPreviousMoveHandler() {
    const sp = document.querySelector('#showprev')
    sp.onclick = () => {
        if ((PreviousMove.from === 64) && (PreviousMove.to === 64)) {
            // first move hasn't been made yet
            return
        }
        showPreviousMove()
        sp.onclick = () => {
            unshowPreviousMove()
        }
    }
}

export function showPreviousMove() {
    previousMoveShown = true
    unshowMoveablePieces()
    document.querySelector('#s'+PreviousMove.from).classList.add('previousmove')
    document.querySelector('#s'+PreviousMove.to).classList.add('previousmove')
}

export function unshowPreviousMove() {
    if ((PreviousMove.from === 64) && (PreviousMove.to === 64)) {
        // first move hasn't been made yet
        return
    }
    previousMoveShown = false
    document.querySelector('#s'+PreviousMove.from).classList.remove('previousmove')
    document.querySelector('#s'+PreviousMove.to).classList.remove('previousmove')
    addShowPreviousMoveHandler()
}
