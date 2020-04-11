import { boardIndex } from './board.js'
import { Piece } from '../piece.js'

export function fetchBoardPromise(gameID) {
    return fetch('/boards/'+gameID).then(r => r.json()).then(j => {
        const board = []
        for (let i = 0; i < 64; i++) {
            board[i] = undefined
        }
        for (const as of j) {
            board[boardIndex(as.a.f, as.a.r)] = new Piece(as.p.k, as.p.o)
        }
        return board
    })
}
