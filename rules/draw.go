package rules

import "github.com/pciet/wichess/piece"

// TODO: same position has occurred three times

// Possible states for an insufficient material draw:
//   king v king
//   king v king+bishop
//   king v king+knight
//   king+bishop v king+bishop of the same bishop color
func (a *Board) insufficientMaterialDraw() bool {
	w := make([]AddressedSquare, 0, 2)
	b := make([]AddressedSquare, 0, 2)
	for i, p := range a {
		if p.Kind == piece.NoKind {
			continue
		}
		switch p.Kind.Basic() {
		case piece.Queen, piece.Rook, piece.Pawn:
			return false
		}
		switch p.Orientation {
		case White:
			if len(w) == 2 {
				return false
			}
			w = append(w, AddressedSquare{AddressIndex(i).Address(), p})
		case Black:
			if len(b) == 2 {
				return false
			}
			b = append(b, AddressedSquare{AddressIndex(i).Address(), p})
		}
	}

	if (len(w) == 0) || (len(b) == 0) {
		panic("side has no pieces")
	}

	if (len(w) == 1) && (len(b) == 1) {
		if (w[0].Kind != piece.King) || (b[0].Kind != piece.King) {
			panic("side missing king")
		}
		return true
	}

	if (len(w) == 2) && (len(b) == 2) {
		var bishop1, bishop2 AddressedSquare
		if w[0].Kind == piece.King {
			bishop1 = w[1]
		} else {
			bishop1 = w[0]
		}
		if b[0].Kind == piece.King {
			bishop2 = b[1]
		} else {
			bishop2 = b[0]
		}
		if (bishop1.Kind.Basic() != piece.Bishop) || (bishop2.Kind.Basic() != piece.Bishop) {
			return false
		}
		if bishop1.SquareEven() != bishop2.SquareEven() {
			return false
		}
	}

	// because pieces besides king, bishop, and knight were filtered out, here it must be
	// a true 1v2 case

	return true
}
