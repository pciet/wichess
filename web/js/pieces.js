var k = {
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
};
var o = {
    WHITE: 0,
    BLACK: 1
};
function nameForKind(kind) {
    switch (kind) {
        case k.KING:
            return 'King';
        case k.QUEEN:
            return 'Queen';
        case k.ROOK:
            return 'Rook';
        case k.BISHOP:
            return 'Bishop';
        case k.KNIGHT:
            return 'Knight';
        case k.PAWN:
            return 'Pawn';
        case k.SWAP_PAWN:
            return 'Swap Pawn';
        case k.LOCK_PAWN:
            return 'Lock Pawn';
        case k.RECON_PAWN:
            return 'Recon Pawn';
        case k.DETONATE_PAWN:
            return 'Detonate Pawn';
        case k.GUARD_PAWN:
            return 'Guard Pawn';
        case k.RALLY_PAWN:
            return 'Rally Pawn';
        case k.FORTIFY_PAWN:
            return 'Fortify Pawn';
        case k.EXTENDED_PAWN:
            return 'Extended Pawn';
        case k.SWAP_KNIGHT:
            return 'Swap Knight';
        case k.LOCK_KNIGHT:
            return 'Lock Knight';
        case k.RECON_KNIGHT:
            return 'Recon Knight';
        case k.DETONATE_KNIGHT:
            return 'Detonate Knight';
        case k.GUARD_KNIGHT:
            return 'Guard Knight';
        case k.RALLY_KNIGHT:
            return 'Rally Knight';
        case k.FORTIFY_KNIGHT:
            return 'Fortify Knight';
        case k.EXTENDED_KNIGHT:
            return 'Extended Knight';
        case k.SWAP_BISHOP:
            return 'Swap Bishop';
        case k.LOCK_BISHOP:
            return 'Lock Bishop';
        case k.RECON_BISHOP:
            return 'Recon Bishop';
        case k.DETONATE_BISHOP:
            return 'Detonate Bishop';
        case k.GHOST_BISHOP:
            return 'Ghost Bishop';
        case k.GUARD_BISHOP:
            return 'Guard Bishop';
        case k.RALLY_BISHOP:
            return 'Rally Bishop';
        case k.FORTIFY_BISHOP:
            return 'Fortify Bishop';
        case k.EXTENDED_BISHOP:
            return 'Extended Bishop';
        case k.SWAP_ROOK:
            return 'Swap Rook';
        case k.LOCK_ROOK:
            return 'Lock Rook';
        case k.RECON_ROOK:
            return 'Recon Rook';
        case k.DETONATE_ROOK:
            return 'Detonate Rook';
        case k.GHOST_ROOK:
            return 'Ghost Rook';
        case k.GUARD_ROOK:
            return 'Guard Rook';
        case k.RALLY_ROOK:
            return 'Rally Rook';
        case k.FORTIFY_ROOK:
            return 'Fortify Rook';
        case EXTENDED_ROOK:
            return 'Extended Rook';
    }
}
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
        case k.SWAP_PAWN:
            name += 'swappawn';
            break;
        case k.LOCK_PAWN:
            name += 'lockpawn';
            break;
        case k.RECON_PAWN:
            name += 'reconpawn';
            break;
        case k.DETONATE_PAWN:
            name += 'detonatepawn';
            break;
        case k.GUARD_PAWN:
            name += 'guardpawn';
            break;
        case k.RALLY_PAWN:
            name += 'rallypawn';
            break;
        case k.FORTIFY_PAWN:
            name += 'fortifypawn';
            break;
        case k.EXTENDED_PAWN:
            name += 'extendedpawn';
            break;
        case k.SWAP_KNIGHT:
            name += 'swapknight';
            break;
        case k.LOCK_KNIGHT:
            name += 'lockknight';
            break;
        case k.RECON_KNIGHT:
            name += 'reconknight';
            break;
        case k.DETONATE_KNIGHT:
            name += 'detonateknight';
            break;
        case k.GUARD_KNIGHT:
            name += 'guardknight';
            break;
        case k.RALLY_KNIGHT:
            name += 'rallyknight';
            break;
        case k.FORTIFY_KNIGHT:
            name += 'fortifyknight';
            break;
        case k.EXTENDED_KNIGHT:
            name += 'extendedknight';
            break;
        case k.SWAP_BISHOP:
            name += 'swapbishop';
            break;
        case k.LOCK_BISHOP:
            name += 'lockbishop';
            break;
        case k.RECON_BISHOP:
            name += 'reconbishop';
            break;
        case k.DETONATE_BISHOP:
            name += 'detonatebishop';
            break;
        case k.GHOST_BISHOP:
            name += 'ghostbishop';
            break;
        case k.GUARD_BISHOP:
            name += 'guardbishop';
            break;
        case k.RALLY_BISHOP:
            name += 'rallybishop';
            break;
        case k.FORTIFY_BISHOP:
            name += 'fortifybishop';
            break;
        case k.EXTENDED_BISHOP:
            name += 'extendedbishop';
            break;
        case k.SWAP_ROOK:
            name += 'swaprook';
            break;
        case k.LOCK_ROOK:
            name += 'lockrook';
            break;
        case k.RECON_ROOK:
            name += 'reconrook';
            break;
        case k.DETONATE_ROOK:
            name += 'detonaterook';
            break;
        case k.GHOST_ROOK:
            name += 'ghostrook';
            break;
        case k.GUARD_ROOK:
            name += 'guardrook';
            break;
        case k.RALLY_ROOK:
            name += 'rallyrook';
            break;
        case k.FORTIFY_ROOK:
            name += 'fortifyrook';
            break;
        case k.EXTENDED_ROOK:
            name += 'extendedrook';
            break;
    }
    return name + '_' + point + '.png';
}
