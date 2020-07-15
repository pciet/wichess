package piece

type Kind int

const (
	NoKind Kind = iota
	King
	Queen
	Rook
	Bishop
	Knight
	Pawn
	War
	Form
	Constructive
	Confined
	Original
	Irrelevant
	KindCount
)

const BasicKindCount = 6

// All special pieces are based on a normal piece. The kind of the normal piece is the basic kind
// of the special piece.
func (a Kind) Basic() Kind {
	switch a {
	case King:
		return King
	case Queen:
		return Queen
	case Bishop, Original:
		return Bishop
	case Rook, Irrelevant:
		return Rook
	case Knight, Constructive:
		return Knight
	case Pawn, War, Form, Confined:
		return Pawn
	}
	return NoKind
}

func (a Kind) IsBasic() bool {
	switch a {
	case King, Queen, Rook, Bishop, Knight, Pawn:
		return true
	}
	return false
}

func (a Kind) String() string { return CodeNames[a] }
