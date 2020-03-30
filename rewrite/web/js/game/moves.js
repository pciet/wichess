import { Condition, replaceCondition, 
        gameDone, replaceMoves } from './game.js'
import { writeGameCondition } from './condition.js'
import { writeBoardMoves } from './board_moves.js'

export function fetchedMoves(moves) {
    if (moves.state !== Condition) {
        replaceCondition(moves.state)
        writeGameCondition()
    }

    if (gameDone() === true) { return }

    replaceMoves(moves.moveSets)
    writeBoardMoves(moves.moveSets)
}
