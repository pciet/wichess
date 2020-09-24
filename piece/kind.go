package piece

// Kind is the kind of the piece, such as king, queen, pawn, or a new special piece added for
// Wisconsin Chess. The kind can be known by how the piece looks, and a kind has characteristics
// associated with it.
type Kind int

// These are all of the piece kinds of Wisconsin Chess, except for NoKind and KindCount which can
// be useful constants.
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
	Imperfect
	Derange
	KindCount
)

const basicKindCount = 6

// All special pieces are based on a regular chess piece like the knight or bishop. The kind of
// the normal piece is the basic kind of the special piece.
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
	case Pawn, War, Form, Confined, Evident, Imperfect, Derange:
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

func (a Kind) String() string { return codeNames[a] }
