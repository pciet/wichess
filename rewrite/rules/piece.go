package rules

import (
	"log"
)

type PieceKind int

type Piece struct {
	Kind PieceKind
	Orientation
	Moved bool
}

func RandomSpecialPieceKind() PieceKind {
	return PieceKind(randomSource.Int63n(int64(PieceKindCount-BasicPieceKindCount)) + 1 + BasicPieceKindCount)
}

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
