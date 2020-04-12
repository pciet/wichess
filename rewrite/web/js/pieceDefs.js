import { Characteristic } from './pieceCharacteristics.js'

function PieceDef(codeName, name, char1 = undefined, char2 = undefined) {
    this.codeName = codeName
    this.name = name
    this.characteristics = [char1, char2]
}

export const NoKind = 0

// These four piece kinds can be used for promotion.
export const Queen = 2
export const Rook = 3
export const Bishop = 4
export const Knight = 5

// The index of Pieces matches the rules.PieceKind enum in the host source.
export const Pieces = [
    new PieceDef('empty', 'No Kind'),
    new PieceDef('king', 'King', Characteristic.SOVEREIGN),
    new PieceDef('queen', 'Queen'),
    new PieceDef('rook', 'Rook'),
    new PieceDef('bishop', 'Bishop'),
    new PieceDef('knight', 'Knight', Characteristic.GHOST),
    new PieceDef('pawn', 'Pawn'),
    new PieceDef('swappawn', 'Swap Pawn', Characteristic.SWAP),
    new PieceDef('lockpawn', 'Lock Pawn', Characteristic.LOCK),
    new PieceDef('reconpawn', 'Recon Pawn', Characteristic.RECON),
    new PieceDef('detonatepawn', 'Detonate Pawn', Characteristic.DETONATE),
    new PieceDef('guardpawn', 'Guard Pawn', Characteristic.GUARD),
    new PieceDef('rallypawn', 'Rally Pawn', Characteristic.RALLY),
    new PieceDef('fortifypawn', 'Fortify Pawn', Characteristic.FORTIFY),
    new PieceDef('extendedpawn', 'Extended Pawn'),
    new PieceDef('swapknight', 'Swap Knight', Characteristic.GHOST, Characteristic.SWAP),
    new PieceDef('lockknight', 'Lock Knight', Characteristic.GHOST, Characteristic.LOCK),
    new PieceDef('reconknight', 'Recon Knight', Characteristic.GHOST, Characteristic.RECON),
    new PieceDef('detonateknight', 'Detonate Knight', Characteristic.GHOST, Characteristic.DETONATE),
    new PieceDef('guardknight', 'Guard Knight', Characteristic.GHOST, Characteristic.GUARD),
    new PieceDef('rallyknight', 'Rally Knight', Characteristic.GHOST, Characteristic.RALLY),
    new PieceDef('fortifyknight', 'Fortify Knight', Characteristic.GHOST, Characteristic.FORTIFY),
    new PieceDef('extendedknight', 'Extended Knight', Characteristic.GHOST),
    new PieceDef('swapbishop', 'Swap Bishop', Characteristic.SWAP),
    new PieceDef('lockbishop', 'Lock Bishop', Characteristic.LOCK),
    new PieceDef('reconbishop', 'Recon Bishop', Characteristic.RECON),
    new PieceDef('detonatebishop', 'Detonate Bishop', Characteristic.DETONATE),
    new PieceDef('ghostbishop', 'Ghost Bishop', Characteristic.GHOST),
    new PieceDef('guardbishop', 'Guard Bishop', Characteristic.GUARD),
    new PieceDef('rallybishop', 'Rally Bishop', Characteristic.RALLY),
    new PieceDef('fortifybishop', 'Fortify Bishop', Characteristic.FORTIFY),
    new PieceDef('extendedbishop', 'Extended Bishop'),
    new PieceDef('swaprook', 'Swap Rook', Characteristic.SWAP),
    new PieceDef('lockrook', 'Lock Rook', Characteristic.LOCK),
    new PieceDef('reconrook', 'Recon Rook', Characteristic.RECON),
    new PieceDef('detonaterook', 'Detonate Rook', Characteristic.DETONATE),
    new PieceDef('ghostrook', 'Ghost Rook', Characteristic.GHOST),
    new PieceDef('guardrook', 'Guard Rook', Characteristic.GUARD),
    new PieceDef('rallyrook', 'Rally Rook', Characteristic.RALLY),
    new PieceDef('fortifyrook', 'Fortify Rook', Characteristic.FORTIFY),
    new PieceDef('extendedrook', 'Extended Rook')
]
