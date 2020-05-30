import { boardIndexToAddress } from './game/board.js'
import { Pieces, NoKind } from './pieceDefs.js'

export function Piece(kind, orientation) {
    this.kind = kind
    this.orientation = orientation
}

// These constants indicate to the host where this piece can be found in a player's collection. 
// NoSlot means it's a basic piece, left and right are the random picks, and a positive integer 
// is in the collection. See docs/collection.md for more details.
export const NoSlot = 0
export const LeftPick = -1
export const RightPick = -2

export function CollectionPiece(slot, kind) {
    this.slot = slot
    this.kind = kind
}

export const Orientation = {
    WHITE: 0,
    BLACK: 1
}

export function orientationString(orientation) {
    if (orientation === Orientation.WHITE) {
        return 'White'
    } else if (orientation === Orientation.BLACK) {
        return 'Black'
    }
    throw new Error('orientation ' + orientation + ' not white or black')
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

export function pieceLookImageName(kind) {
    return 'look_' + Pieces[kind].codeName + '.png'
}

export function piecePickImageName(kind) {
    return 'pick_' + Pieces[kind].codeName + '.png'
}
