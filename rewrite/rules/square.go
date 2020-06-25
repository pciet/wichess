package rules

import "github.com/pciet/wichess/piece"

// A board is made of squares which each might have a piece.
// An empty square is indicated by the piece's Kind set to NoKind.
type Square Piece

func (a Square) FortifiedAgainst(t Square) bool {
	return a.Fortified && (t.Kind.Basic() == piece.Pawn)
}

type AddressedSquare struct {
	Address `json:"a"`
	Square  `json:"p"`
}

func (a Square) NotEmpty() bool { return a.Kind != piece.NoKind }
func (a Square) Empty() bool    { return a.Kind == piece.NoKind }

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
	if a.Kind == piece.NoKind {
		return "empty"
	}
	k := ""
	if a.Kind.IsBasic() == false {
		k = "+"
	}
	m := ""
	if a.Moved {
		m = "moved "
	}
	return m + a.Orientation.String() + " " + k + a.Kind.String()
}
