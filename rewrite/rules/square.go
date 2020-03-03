package rules

// A board is made of squares which each might have a piece.
// An empty square is indicated by the piece's Kind set to NoKind.
type Square Piece

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
				bs[i].Square = s.Square
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
