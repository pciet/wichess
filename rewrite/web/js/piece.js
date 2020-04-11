import { boardIndexToAddress } from './game/board.js'

export function Piece(kind, orientation) {
    this.kind = kind
    this.orientation = orientation
}

export const Orientation = {
    WHITE: 0,
    BLACK: 1
}

export const Kind = {
    NO_KIND: 0,
    KING: 1,
    QUEEN: 2,
    ROOK: 3,
    BISHOP: 4,
    KNIGHT: 5,
    PAWN: 6,
    SWAP_PAWN: 7,
    LOCK_PAWN: 8,
    RECON_PAWN: 9,
    DETONATE_PAWN: 10,
    GUARD_PAWN: 11,
    RALLY_PAWN: 12,
    FORTIFY_PAWN: 13,
    EXTENDED_PAWN: 14,
    SWAP_KNIGHT: 15,
    LOCK_KNIGHT: 16,
    RECON_KNIGHT: 17,
    DETONATE_KNIGHT: 18,
    GUARD_KNIGHT: 19,
    RALLY_KNIGHT: 20,
    FORTIFY_KNIGHT: 21,
    EXTENDED_KNIGHT: 22,
    SWAP_BISHOP: 23,
    LOCK_BISHOP: 24,
    RECON_BISHOP: 25,
    DETONATE_BISHOP: 26,
    GHOST_BISHOP: 27,
    GUARD_BISHOP: 28,
    RALLY_BISHOP: 29,
    FORTIFY_BISHOP: 30,
    EXTENDED_BISHOP: 31,
    SWAP_ROOK: 32,
    LOCK_ROOK: 33,
    RECON_ROOK: 34,
    DETONATE_ROOK: 35,
    GHOST_ROOK: 36,
    GUARD_ROOK: 37,
    RALLY_ROOK: 38,
    FORTIFY_ROOK: 39,
    EXTENDED_ROOK: 40
}

// TODO: convert these images names to just respond with what the host sends
// TODO: like 0_1_0.png for a white king at 0,0

// image names are like bbishop_1_5.png

export function pieceImageName(addressIndex, kind, orientation) {
    //return orientation.toString()+'-'+kind.toString()+'-'+addressIndex.toString()+'.png'
    let name = ''
    if (orientation == Orientation.WHITE) {
        name += 'w'
    } else if (orientation == Orientation.BLACK) {
        name += 'b'
    } else {
        throw new Error('unknown orientation ' + orientation)
    }
    switch (kind) {
        case Kind.NO_KIND:
            name = 'empty'
            break
        case Kind.KING:
            name += 'king'
            break
        case Kind.QUEEN:
            name += 'queen'
            break
        case Kind.ROOK:
            name += 'rook'
            break
        case Kind.BISHOP:
            name += 'bishop'
            break
        case Kind.KNIGHT:
            name += 'knight'
            break
        case Kind.PAWN:
            name += 'pawn'
            break
        case Kind.SWAP_PAWN:
            name += 'swappawn'
            break
        case Kind.LOCK_PAWN:
            name += 'lockpawn'
            break
        case Kind.RECON_PAWN:
            name += 'reconpawn'
            break
        case Kind.DETONATE_PAWN:
            name += 'detonatepawn'
            break
        case Kind.GUARD_PAWN:
            name += 'guardpawn'
            break
        case Kind.RALLY_PAWN:
            name += 'rallypawn'
            break
        case Kind.FORTIFY_PAWN:
            name += 'fortifypawn'
            break
        case Kind.EXTENDED_PAWN:
            name += 'extendedpawn'
            break
        case Kind.SWAP_KNIGHT:
            name += 'swapknight'
            break
        case Kind.LOCK_KNIGHT:
            name += 'lockknight'
            break
        case Kind.RECON_KNIGHT:
            name += 'reconknight'
            break
        case Kind.DETONATE_KNIGHT:
            name += 'detonateknight'
            break
        case Kind.GUARD_KNIGHT:
            name += 'guardknight'
            break
        case Kind.RALLY_KNIGHT:
            name += 'rallyknight'
            break
        case Kind.FORTIFY_KNIGHT:
            name += 'fortifyknight'
            break
        case Kind.EXTENDED_KNIGHT:
            name += 'extendedknight'
            break
        case Kind.SWAP_BISHOP:
            name += 'swapbishop'
            break
        case Kind.LOCK_BISHOP:
            name += 'lockbishop'
            break
        case Kind.RECON_BISHOP:
            name += 'reconbishop'
            break
        case Kind.DETONATE_BISHOP:
            name += 'detonatebishop'
            break
        case Kind.GHOST_BISHOP:
            name += 'ghostbishop'
            break
        case Kind.GUARD_BISHOP:
            name += 'guardbishop'
            break
        case Kind.RALLY_BISHOP:
            name += 'rallybishop'
            break
        case Kind.FORTIFY_BISHOP:
            name += 'fortifybishop'
            break
        case Kind.EXTENDED_BISHOP:
            name += 'extendedbishop'
            break
        case Kind.SWAP_ROOK:
            name += 'swaprook'
            break
        case Kind.LOCK_ROOK:
            name += 'lockrook'
            break
        case Kind.RECON_ROOK:
            name += 'reconrook'
            break
        case Kind.DETONATE_ROOK:
            name += 'detonaterook'
            break
        case Kind.GHOST_ROOK:
            name += 'ghostrook'
            break
        case Kind.GUARD_ROOK:
            name += 'guardrook'
            break
        case Kind.RALLY_ROOK:
            name += 'rallyrook'
            break
        case Kind.FORTIFY_ROOK:
            name += 'fortifyrook'
            break
        case Kind.EXTENDED_ROOK:
            name += 'extendedrook'
            break
    }
    const addr = boardIndexToAddress(addressIndex)
    return name + '_' + addr.file.toString() + '_' + addr.rank.toString() + '.png'
}

// TODO: reduce repetition

// kindCharacteristics returns an array of one or two
// characteristic name strings.
export function kindCharacteristics(kind) {
    switch (kind) {
    case Kind.KNIGHT:
        return ["Ghost", undefined]
    case Kind.KING:
        return ["Sovereign", undefined]
    }
    return [undefined, undefined]
}

const kindNames = [
    "No Kind", "King", "Queen", "Rook",
    "Bishop", "Knight", "Pawn"]

export function kindName(kind) {
    return kindNames[kind]
}
