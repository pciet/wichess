export const NoKind = 0
export const King = 1
export const Pawn = 6

// These four piece kinds can be used for promotion.
export const Queen = 2
export const Rook = 3
export const Bishop = 4
export const Knight = 5

// The index of Pieces matches the rules.PieceKind enum in the host source and is the code name
// of the piece used for the image names and encoded information.
export const Pieces = [
    'empty',
    'king', 'queen', 'rook', 'bishop', 'knight', 'pawn',
    'war',
    'formpawn',
    'constructive',
    'confined',
    'original',
    'irrelevant',
    'evident',
    'line',
    'impossible',
    'convenient',
    'appropriate',
    'warprook',
    'brilliant',
    'simple',
    'exit',
    'imperfect',
    'derange'
]

export const BasicKinds = [
    NoKind,
    King, Queen, Rook, Bishop, Knight, Pawn,
    Pawn,
    Pawn,
    Knight,
    Pawn,
    Bishop,
    Rook,
    Pawn,
    Knight,
    Rook,
    Bishop,
    Knight,
    Rook,
    Knight,
    Rook,
    Bishop,
    Pawn,
    Pawn
]
