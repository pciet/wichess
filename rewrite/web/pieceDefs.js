import { Characteristic } from './pieceCharacteristics.js'

function PieceDef(codeName, name, basicKind = NoKind, char1 = undefined, char2 = undefined) {
    this.codeName = codeName
    this.name = name
    this.basicKind = basicKind
    this.characteristics = [char1, char2]
}

export const NoKind = 0
export const King = 1
export const Pawn = 6

// These four piece kinds can be used for promotion.
export const Queen = 2
export const Rook = 3
export const Bishop = 4
export const Knight = 5

// The index of Pieces matches the rules.PieceKind enum in the host source.
export const Pieces = [
    new PieceDef('empty', 'No Kind'),
    new PieceDef('king', 'King', King),
    new PieceDef('queen', 'Queen', Queen),
    new PieceDef('rook', 'Rook', Rook),
    new PieceDef('bishop', 'Bishop', Bishop),
    new PieceDef('knight', 'Knight', Knight, Characteristic.QUICK),
    new PieceDef('pawn', 'Pawn', Pawn),
    new PieceDef('war', 'War', Pawn, Characteristic.NEUTRALIZES),
    new PieceDef('formpawn', 'Form', Pawn, Characteristic.REVEALS, Characteristic.ENABLES),
    new PieceDef('constructive', 'Constructive', Knight, 
        Characteristic.QUICK, Characteristic.ASSERTS)
]
