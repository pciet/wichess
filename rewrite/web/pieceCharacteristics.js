export const Characteristic = {
    NEUTRALIZES: 0,
    ASSERTS: 1,
    ENABLES: 2,
    QUICK: 3,
    REVEALS: 4,
}

function CharDef(name, description) {
    this.name = name
    this.description = description
}

export const Characteristics = [
    new CharDef('Neutralizes', 'When taken, surrounding pieces from both sides are also taken.'),
    new CharDef('Asserts', 'This piece moves itself to take opponent pieces that move adjacent.'),
    new CharDef('Enables', 'Adjacent pieces from the same side gain more moves.'),
    new CharDef('Quick', 'This piece moves over other pieces.'),
    new CharDef('Reveals', 'Ally pieces behind this piece can move to the square ahead.'),
]
