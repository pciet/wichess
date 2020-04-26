export const Characteristic = {
    SOVEREIGN: 0,
    GHOST: 1,
    SWAP: 2,
    LOCK: 3,
    RECON: 4,
    DETONATE: 5,
    GUARD: 6,
    RALLY: 7,
    FORTIFY: 8
}

function CharDef(name, description) {
    this.name = name
    this.description = description
}

export const Characteristics = [
    new CharDef('Sovereign', 'Putting this piece into checkmate wins the game.'),
    new CharDef('Ghost', 'This piece can move over other pieces to get to a square.'),
    new CharDef('Swap', 'This piece can swap squares with a friendly piece.'),
    new CharDef('Lock', 'Opponent pieces adjacent to this piece cannot move.'),
    new CharDef('Recon', 'Friendly pieces close behind this one can move to the square in front of it.'),
    new CharDef('Detonate', 'When taken this piece takes all adjacent pieces.'),
    new CharDef('Guard', 'When an opponent piece moves adjacent this piece takes that piece automatically.'),
    new CharDef('Rally', 'Friendly pieces adjacent to this one gain additional moves.'),
    new CharDef('Fortify', 'This piece cannot be taken by pawns.')
]
