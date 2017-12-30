// Copyright 2017 Matthew Juran
// All Rights Reserved

var MaxAttributeCount = 2;

var a = {
    GHOST: 1,
    SWAPS: 2,
    LOCKS: 3,
    RECON: 4,
    DETONATES: 5,
    GUARDS: 6,
    RALLIES: 7,
    FORTIFIED: 9,
    SOVEREIGN: 10
};

function stringForAttribute(attribute) {
    switch (attribute) {
        case a.GHOST:
            return "Ghost";
        case a.SWAPS:
            return "Swaps";
        case a.LOCKS:
            return "Locks";
        case a.RECON:
            return "Recon";
        case a.DETONATES:
            return "Detonates";
        case a.GUARDS:
            return "Guard";
        case a.RALLIES:
            return "Rallies";
        case a.FORTIFIED:
            return "Fortified";
        case a.SOVEREIGN:
            return "Sovereign";
        default:
            throw "unknown attribute " + attribute;
    }
}

function descriptionForAttribute(attribute) {
    switch (attribute) {
        case a.GHOST:
            return "This piece may move over other pieces to move into an empty square or take an opponent's piece.";
        case a.SWAPS:
            return "Moving onto a friendly piece causes it to travel back to this piece's original square.";
        case a.LOCKS:
            return "Any opponent piece adjacent to this piece cannot move.";
        case a.RECON:
            return "A friendly piece in one of the three squares behind this piece may move to the one square ahead of it.";
        case a.DETONATES:
            return "When this piece is taken all surrounding pieces and the taker are also taken.";
        case a.GUARDS:
            return "If an opponent piece moves into a square adjacent to this piece then this piece takes that piece by moving to that square."
        case a.RALLIES:
            return "Any rallyable friendly pieces adjacent to this piece gain their rally moves.";
        case a.FORTIFIED:
            return "Pieces with a pawn base cannot take this piece.";
        case a.SOVEREIGN:
            return "If this piece is threatened then you are in check and must move out of check; if you cannot move out of check then your opponent has won.";
        default:
            throw "unknown attribute " + attribute;
    }
}

// kind from pieces.js
// returns an array of attributes
function attributesForKind(kind) {
    switch (kind) {
        case k.KING:
            return [a.SOVEREIGN];
        case k.QUEEN:
        case k.PAWN:
        case k.EXTENDED_PAWN:
        case k.ROOK:
        case k.EXTENDED_ROOK:
        case k.BISHOP:
        case k.EXTENDED_BISHOP:
            return [];
        case k.KNIGHT:
        case k.EXTENDED_KNIGHT:
        case k.GHOST_BISHOP:
        case k.GHOST_ROOK:
            return [a.GHOST];
        case k.SWAP_PAWN:
        case k.SWAP_BISHOP:
        case k.SWAP_ROOK:
            return [a.SWAPS];
        case k.LOCK_PAWN:
        case k.LOCK_BISHOP:
        case k.LOCK_ROOK:
            return [a.LOCKS];
        case k.RECON_PAWN:
        case k.RECON_BISHOP:
        case k.RECON_ROOK:
            return [a.RECON];
        case k.DETONATE_PAWN:
        case k.DETONATE_BISHOP:
        case k.DETONATE_ROOK:
            return [a.DETONATES];
        case k.GUARD_PAWN:
        case k.GUARD_BISHOP:
        case k.GUARD_ROOK:
            return [a.GUARDS];
        case k.RALLY_PAWN:
        case k.RALLY_BISHOP:
        case k.RALLY_ROOK:
            return [a.RALLIES];
        case k.FORTIFY_PAWN:
        case k.FORTIFY_BISHOP:
        case k.FORTIFY_ROOK:
            return [a.FORTIFIED];
        case k.SWAP_KNIGHT:
            return [a.SWAPS, a.GHOST];
        case k.LOCK_KNIGHT:
            return [a.LOCKS, a.GHOST];
        case k.RECON_KNIGHT:
            return [a.RECON, a.GHOST];
        case k.DETONATE_KNIGHT:
            return [a.DETONATES, a.GHOST];
        case k.GUARD_KNIGHT:
            return [a.GUARDS, a.GHOST];
        case k.RALLY_KNIGHT:
            return [a.RALLIES, a.GHOST];
        case k.FORTIFY_KNIGHT:
            return [a.FORTIFIED, a.GHOST];
        default:
            throw "unknown piece " + kind;
    }
}
