package rules

// Each piece has a kind which indicates its possible move sets and abilities.
type PieceKind int

type Piece struct {
	Kind        PieceKind `json:"k"`
	Orientation `json:"o"`
	Moved       bool `json:"m"`

	Swaps     bool `json:"-"`
	Detonates bool `json:"-"`
	Guards    bool `json:"-"`
	Fortified bool `json:"-"`
	Locks     bool `json:"-"`
	Rallies   bool `json:"-"`
	MustEnd   bool `json:"-"`
	Ghost     bool `json:"-"`
	Recons    bool `json:"-"`
}

var (
	WhiteKingStart      = Address{4, 0}
	BlackKingStart      = Address{4, 7}
	WhiteLeftRookStart  = Address{0, 0}
	WhiteRightRookStart = Address{7, 0}
	BlackLeftRookStart  = Address{7, 7}
	BlackRightRookStart = Address{0, 7}
)

func (a Piece) ApplyCharacteristics() Piece {
	if BasicKind(a.Kind) == Knight {
		a.MustEnd = true
		a.Ghost = true
	}
	// TODO: map this?
	switch a.Kind {
	case King, Queen, Rook, Knight, Bishop, Pawn, NoKind:
		break
	case SwapPawn, SwapKnight, SwapBishop, SwapRook:
		a.Swaps = true
	case DetonatePawn, DetonateKnight, DetonateBishop, DetonateRook:
		a.Detonates = true
	case GuardPawn, GuardKnight, GuardBishop, GuardRook:
		a.Guards = true
	case FortifyPawn, FortifyKnight, FortifyBishop, FortifyRook:
		a.Fortified = true
	case LockPawn, LockKnight, LockBishop, LockRook:
		a.Locks = true
	case RallyPawn, RallyKnight, RallyBishop, RallyRook:
		a.Rallies = true
	case GhostBishop, GhostRook:
		a.Ghost = true
	case ReconPawn, ReconKnight, ReconBishop, ReconRook:
		a.Recons = true
	default:
		Panic("unknown piece kind", a.Kind, a)
	}
	return a
}

// All special pieces are based on a normal piece, called the basic kind of the piece.
func BasicKind(p PieceKind) PieceKind {
	switch p {
	case NoKind:
		return NoKind
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
		Panic("unexpected piece kind", p)
		return NoKind
	}
}

// A special piece is one that's not from the normal chess set (king, queen, rook, bishop, knight, pawn).
func RandomSpecialPieceKind() PieceKind {
	return PieceKind(randomSource.Int63n(int64(PieceKindCount-BasicPieceKindCount)) + 1 + BasicPieceKindCount)
}

func IsBasicKind(p PieceKind) bool {
	switch p {
	case King, Queen, Rook, Bishop, Knight, Pawn:
		return true
	}
	return false
}

func (a PieceKind) String() string {
	switch a {
	case King:
		return "king"
	case Queen:
		return "queen"
	case Rook:
		return "rook"
	case Bishop:
		return "bishop"
	case Knight:
		return "knight"
	case Pawn:
		return "pawn"
	default:
		return "undefined"
	}
}
