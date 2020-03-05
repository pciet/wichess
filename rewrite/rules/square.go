package rules

// A board is made of squares which each might have a piece.
// An empty square is indicated by the piece's Kind set to NoKind.
type Square Piece

func (a Square) FortifiedAgainst(t Square) bool { return a.Fortified && (BasicKind(t.Kind) == Pawn) }

type AddressedSquare struct {
	Address
	Square
}

func (a Square) NotEmpty() bool { return a.Kind != NoKind }
func (a Square) Empty() bool    { return a.Kind == NoKind }

func MergeReplaceAddressedSquares(base, overwrite []AddressedSquare) []AddressedSquare {
LOOP:
	for _, s := range overwrite {
		// either it needs to replace or be added
		for i, bs := range base {
			if bs.Address == s.Address {
				base[i].Square = s.Square
				continue LOOP
			}
		}
		base = append(base, s)
	}
	return base
}

func CombineAddressedSquares(a, b []AddressedSquare) []AddressedSquare {
	for _, s := range b {
		a = append(a, s)
	}
	return a
}

func (a Square) String() string {
	if a.Kind == NoKind {
		return "empty"
	}
	k := ""
	if IsBasicKind(a.Kind) == false {
		k = "+"
	}
	m := ""
	if a.Moved {
		m = "moved "
	}
	return m + a.Orientation.String() + " " + k + a.Kind.String()
}
