package rules

import "github.com/pciet/wichess/piece"

// Capturing a neutralizing piece will cause adjacent neutralizers to also neutralize. This method
// changes the board.
func (a *Board) neutralizesMove(changes, captures []Square, m Move) ([]Square, []Square) {
	// before doing the neutralization capture the piece that moved and apply that change
	captures = append(captures, Square{m.From, a[m.From.Index()]})
	a[m.From.Index()] = Piece{}
	changes = append(changes, Square{m.From, Piece{}})

	for _, capture := range a.neutralizeCaptureAddresses(m.To) {
		changes = append(changes, Square{capture, Piece{}})
		captures = append(captures, Square{capture, a[capture.Index()]})
	}

	return changes, captures
}

func (a *Board) neutralizeCaptureAddresses(startingAt Address) []Address {
	var recursiveNeutralize func(Address)

	toCapture := make([]Address, 0, 4)
	neutralizing := make([]Address, 0, 2)

	recursiveNeutralize = func(at Address) {
		// A neutralize is happening at this address, so it and all surrounding squares will have
		// their piece captured. If another neutralizing piece is captured then a recursive
		// neutralizing chain happens.

		if addressSliceHas(toCapture, at) == false {
			toCapture = append(toCapture, at)
		}

		if addressSliceHas(neutralizing, at) {
			return
		}
		neutralizing = append(neutralizing, at)

		for _, surrounding := range a.surroundingSquares(at) {
			if surrounding.Kind == piece.NoKind {
				continue
			}
			if (surrounding.flags.neutralizes && (surrounding.is.normalized == false)) ||
				surrounding.is.ordered {

				recursiveNeutralize(surrounding.Address)
			}
			if addressSliceHas(toCapture, surrounding.Address) == false {
				toCapture = append(toCapture, surrounding.Address)
			}
		}
	}

	recursiveNeutralize(startingAt)

	return toCapture
}
