import { Condition, replaceCondition,
        gameDone, replaceMoves } from '../game.js'

import { State } from './state.js'
import { replaceAndWriteGameCondition } from './condition.js'
import { replaceAndWriteBoardMoves } from './board_moves.js'
import { showPromotion } from './promotion.js'

export function fetchedMoves(moves) {
    if (moves.state !== Condition) {
        replaceAndWriteGameCondition(moves.state)
    }

    if (gameDone() === true) {
        return
    }

    if (moves.state === State.PROMOTION) {
        showPromotion()
    } else {
        replaceAndWriteBoardMoves(moves.moveSets)
    }
}
