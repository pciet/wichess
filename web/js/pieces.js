var k = {
    KING: 1,
    QUEEN: 2,
    ROOK: 3,
    BISHOP: 4,
    KNIGHT: 5,
    PAWN: 6,
    SWAP: 7,
    LOCK: 8,
    RECON: 9,
    DETONATE: 10,
    GHOST: 11,
    STEAL: 12,
    GUARD: 13,
    RALLY: 14,
    FORTIFY: 15
};
var o = {
    WHITE: 0,
    BLACK: 1
};
function imageNameForKind(kind, point, orientation) {
    var name;
    if (orientation == o.WHITE) {
        name = 'w';
    } else {
        name = 'b';
    }
    switch (kind) {
        case k.KING:
            name += 'king';
            break;
        case k.QUEEN:
            name += 'queen';
            break;
        case k.ROOK:
            name += 'rook';
            break;
        case k.BISHOP:
            name += 'bishop';
            break;
        case k.KNIGHT:
            name += 'knight';
            break;
        case k.PAWN:
            name += 'pawn';
            break;
        case k.SWAP:
            name += 'swap';
            break;
        case k.LOCK:
            name += 'lock';
            break;
        case k.RECON:
            name += 'recon';
            break;
        case k.DETONATE:
            name += 'detonate';
            break;
        case k.GHOST:
            name += 'ghost';
            break;
        case k.STEAL:
            name += 'steal';
            break;
        case k.GUARD:
            name += 'guard';
            break;
        case k.RALLY:
            name += 'rally';
            break;
        case k.FORTIFY:
            name += 'fortify';
            break;
    }
    return name + '_' + point + '.png';
}
