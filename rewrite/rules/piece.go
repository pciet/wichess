package rules

// PieceKind is a positive integer that
// indiciates the piece's moves and characteristics.
type PieceKind int

type Piece struct {
	Kind        PieceKind `json:"k"`
	Orientation `json:"o"`
	Moved       bool `json:"-"`

	Swaps bool `json:"-"`

	// Neutralizes
	Detonates bool `json:"-"`

	// Asserts
	Guards bool `json:"-"`

	Fortified bool `json:"-"`
	Locks     bool `json:"-"`

	// Enables
	Rallies bool `json:"-"`

	MustEnd bool `json:"-"`

	// Quick
	Ghost bool `json:"-"`

	// Reveals
	Recons bool `json:"-"`
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
	case War:
		a.Detonates = true
	case Constructive:
		a.Guards = true
	case Form:
		a.Recons = true
		a.Rallies = true
	default:
		Panic("unknown piece kind", a.Kind, a)
	}
	return a
}

// All special pieces are based on a normal piece, called
// the basic kind of the piece.
func BasicKind(p PieceKind) PieceKind {
	switch p {
	case NoKind:
		return NoKind
	case King:
		return King
	case Queen:
		return Queen
	case Bishop:
		return Bishop
	case Rook:
		return Rook
	case Knight, Constructive:
		return Knight
	case Pawn, War, Form:
		return Pawn
	default:
		Panic("unexpected piece kind", p)
		return NoKind
	}
}

func RandomSpecialPieceKind() PieceKind {
	return PieceKind(randomSource.Int63n(
		int64(PieceKindCount-BasicPieceKindCount-1)) + 1 + BasicPieceKindCount)
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
	case War:
		return "war"
	case Constructive:
		return "constructive"
	case Form:
		return "form"
	default:
		return "undefined"
	}
}
