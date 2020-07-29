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
	Evident
	Line
	Impossible
	Convenient
	Appropriate
	Warp
	Brilliant
	Simple
	Exit
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
	case Bishop, Original, Convenient, Exit:
		return Bishop
	case Rook, Irrelevant, Impossible, Warp, Simple:
		return Rook
	case Knight, Constructive, Line, Appropriate, Brilliant:
		return Knight
	case Pawn, War, Form, Confined, Evident:
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
