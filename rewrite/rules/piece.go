package rules

import (
	"log"
)

// Each piece has a kind which indicates its possible move sets and abilities.
type PieceKind int

type Piece struct {
	Kind        PieceKind `json:"kind"`
	Orientation `json:"orientation"`
	Moved       bool `json:"moved"`

	Swaps     bool `json:"-"`
	Detonates bool `json:"-"`
	Guards    bool `json:"-"`
	Fortified bool `json:"-"`
	Locks     bool `json:"-"`
}

func (a Piece) SameOrientation(as Piece) bool { return a.Orientation == as.Orientation }
func (a Piece) FortifiedAgainst(t Piece) bool { return a.Fortified && (BasicKind(t.Kind) == Pawn) }

// All special pieces are based on a normal piece, called the basic kind of the piece.
func BasicKind(p PieceKind) PieceKind {
	switch p {
	case King:
		return King
	case Queen:
		return Queen
	case Rook, SwapRook, LockRook, ReconRook, DetonateRook, GhostRook, GuardRook, RallyRook, FortifyRook, ExtendedRook:
		return Rook
	case Bishop, SwapBishop, LockBishop, ReconBishop, DetonateBishop, GhostBishop, GuardBishop, RallyBishop, FortifyBishop, ExtendedBishop:
		return Bishop
	case Knight, SwapKnight, LockKnight, ReconKnight, DetonateKnight, GuardKnight, RallyKnight, FortifyKnight, ExtendedKnight:
		return Knight
	case Pawn, SwapPawn, LockPawn, ReconPawn, DetonatePawn, GuardPawn, RallyPawn, FortifyPawn, ExtendedPawn:
		return Pawn
	default:
		log.Panicln("unexpected piece kind", p)
		return NoKind
	}
}

// A special piece is one that's not from the normal chess set (king, queen, rook, bishop, knight, pawn).
func RandomSpecialPieceKind() PieceKind {
	return PieceKind(randomSource.Int63n(int64(PieceKindCount-BasicPieceKindCount)) + 1 + BasicPieceKindCount)
}
