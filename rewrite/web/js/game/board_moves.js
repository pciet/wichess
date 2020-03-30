import { boardAddressToIndex } from './board.js'
import { writeSquareClick } from './board_click.js'

export function writeBoardMoves(moves) {
    for (let i = 0; i < moves.length; i++) {
        const from = boardAddressToIndex(moves[i].from)
        const tos = []
        for (let j = 0; j < moves[i].tos.length; j++) {
            tos[j] = boardAddressToIndex(moves[i].tos[j])
        }
        writeSquareClick(from, tos)
    }
}
