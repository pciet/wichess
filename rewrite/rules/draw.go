package rules

import (
	"log"
)

// TODO: same position has occurred three times

// Possible states for an insufficient material draw:
//   king v king
//   king v king+bishop
//   king v king+knight
//   king+bishop v king+bishop of the same bishop color
func (a Board) InsufficientMaterialDraw() bool {
	w := make([]AddressedSquare, 0, 2)
	b := make([]AddressedSquare, 0, 2)
	for i, p := range a {
		if p.Kind == NoKind {
			continue
		}
		switch BasicKind(p.Kind) {
		case Queen, Rook, Pawn:
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
		log.Panic("side has no pieces")
	}

	if (len(w) == 1) && (len(b) == 1) {
		if (w[0].Kind != King) || (b[0].Kind != King) {
			log.Panic("side missing king")
		}
		return true
	}

	if (len(w) == 2) && (len(b) == 2) {
		var bishop1, bishop2 AddressedSquare
		if w[0].Kind == King {
			bishop1 = w[1]
		} else {
			bishop1 = w[0]
		}
		if b[0].Kind == King {
			bishop2 = b[1]
		} else {
			bishop2 = b[0]
		}
		if (BasicKind(bishop1.Kind) != Bishop) || (BasicKind(bishop2.Kind) != Bishop) {
			return false
		}
		if bishop1.SquareEven() != bishop2.SquareEven() {
			return false
		}
	}

	// because pieces besides king, bishop, and knight were filtered out, here it must be a true 1v2 case

	return true
}
