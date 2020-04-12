import { boardIndexToAddress } from './game/board.js'
import { Pieces, NoKind } from './pieceDefs.js'

export function Piece(kind, orientation) {
    this.kind = kind
    this.orientation = orientation
}

export const Orientation = {
    WHITE: 0,
    BLACK: 1
}

export function pieceImageName(addressIndex, kind, orientation) {
    let name = ''
    if (orientation == Orientation.WHITE) {
        name += 'w'
    } else if (orientation == Orientation.BLACK) {
        name += 'b'
    } else {
        throw new Error('unknown orientation ' + orientation)
    }
    if (kind === NoKind) {
        name = 'empty'
    } else {
        name += Pieces[kind].codeName
    }
    const addr = boardIndexToAddress(addressIndex)
    return name + '_' + addr.file.toString() + '_' + addr.rank.toString() + '.png'
}
