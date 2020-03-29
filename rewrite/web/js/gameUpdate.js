import { boardIndex } from './board.js'
import { writePieceImage, removePieceImage } from './gameLayoutGenerate.js'

// TODO: race condition if updateBoard is writing Board and layout is called

// A board update, either from a WebSocket alert message or the HTTP response
// to /move, is always a non-zero array of addressed pieces.
export function updateBoard(board, update) {
    for (const u of update) {
        const index = boardIndex(u.a.f, u.a.r)
        // TODO: if (u.p === undefined)
        board[index] = u.p.p
        if (u.p.p.k === 0) {
            removePieceImage(index)
            continue
        }
        writePieceImage(index, u.p.p)
    }
}
