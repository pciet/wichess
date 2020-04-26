import { Moves, replaceMoves } from '../game.js'

import { boardAddressToIndex } from './board.js'
import { writeSquareClick } from './board_click.js'

export function replaceAndWriteBoardMoves(moves) {
    replaceMoves(moves)
    writeBoardMoves()
}

export function writeBoardMoves() {
    for (let i = 0; i < Moves.length; i++) {
        const from = boardAddressToIndex(Moves[i].from)
        const tos = []
        for (let j = 0; j < Moves[i].tos.length; j++) {
            tos[j] = boardAddressToIndex(Moves[i].tos[j])
        }
        writeSquareClick(from, tos)
    }
}
