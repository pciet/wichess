import { Board } from '../game.js'
import { NoKind } from '../pieceDefs.js'
import { Piece } from '../piece.js'

import { boardIndex, BoardAddress, AddressedPiece } from './board.js'
import { writePieceImage, removePieceImage } from './board_images.js'

// TODO: race condition if updateBoard is writing Board and layout is called

// A board update, either from a WebSocket alert message or the HTTP response to /move, is 
// always a non-zero array of addressed pieces. The input update var should first be parsed 
// by parseBoardUpdate.
export function updateBoard(update) {
    for (const u of update) {
        const index = boardIndex(u.address.file, u.address.rank)
        if (u.piece.kind === NoKind) {
            Board[index] = undefined
            removePieceImage(index)
            continue
        }
        Board[index] = u.piece
        writePieceImage(index, Board[index])
    }
}

export function parseBoardUpdate(j) {
    const addressedPieces = []
    for (const update of j) {
        addressedPieces.push(new AddressedPiece(
            new BoardAddress(update.a.f, update.a.r),
            new Piece(update.p.k, update.p.o)
        ))
    }
    return addressedPieces
}
