package rules

import "github.com/pciet/wichess/piece"

// Square represents a square on a board.
type Square struct {
	Address `json:"a"`
	Piece   `json:"p"`
}

// MergeReplaceSquares will apply the overwrite slice to the base slice. If base doesn't include
// the Square then it's added, otherwise it's replaced.
func MergeReplaceSquares(base, overwrite []Square) []Square {
LOOP:
	for _, s := range overwrite {
		// either it needs to replace or be added
		for i, bs := range base {
			if bs.Address == s.Address {
				base[i].Piece = s.Piece
				continue LOOP
			}
		}
		base = append(base, s)
	}
	return base
}

// TODO: the piece's start square isn't included in SquaresEquivalent

// SquaresEquivalent determines if the argument unordered slices include the same squares.
func SquaresEquivalent(a, b []Square) bool {
	if len(a) != len(b) {
		return false
	}

	comparison := func(an Square) Square {
		return Square{
			Address: an.Address,
			Piece: Piece{
				Kind:        an.Piece.Kind,
				Orientation: an.Piece.Orientation,
				Moved:       an.Piece.Moved,
				//Start:       an.Piece.Start,
			},
		}
	}

	mapCount := func(slice []Square) map[Square]int {
		m := make(map[Square]int)
		for _, as := range slice {
			cas := comparison(as)
			c, has := m[cas]
			if has == false {
				m[cas] = 1
			} else {
				m[cas] = c + 1
			}
		}
		return m
	}
	am := mapCount(a)
	bm := mapCount(b)
	for k, v := range am {
		if bm[k] != v {
			return false
		}
	}
	return true
}

func (a Square) String() string {
	k := a.Address.String() + " "
	if a.Kind == piece.NoKind {
		return k + "empty"
	}
	if a.Moved {
		k += "moved "
	}
	return k + a.Orientation.String() + " " + a.Kind.String()
}
