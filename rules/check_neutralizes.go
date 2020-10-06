package rules

import "github.com/pciet/wichess/piece"

// neutralizesAssertedChainAdjacent looks for situations where a neutralizing piece could move into
// range of an asserting piece and cause a neutralizing chain that captures the address.
func (a *Board) neutralizesAssertedChainAdjacent(to Address) bool {
	p := a[to.Index()]
	if p.Kind == piece.NoKind {
		return false
	}
	o := p.Orientation

	opponentNeutralizesMoves := make([]MoveSet, 0, 4)
	for i, s := range a {
		if (s.Kind == piece.NoKind) || (s.Orientation == o) ||
			(((s.flags.neutralizes == false) || (s.is.normalized)) && (s.is.ordered == false)) {

			continue
		}
		addr := AddressIndex(i).Address()
		opponentNeutralizesMoves = append(opponentNeutralizesMoves,
			MoveSet{addr, a.naiveMovesAt(addr, NoMove)})
	}

	for _, set := range opponentNeutralizesMoves {
		for _, addr := range set.Moves {
			// if a neutralizer moves into an assert then the argument address must be shown to
			// not be in neutralize chain range of the assert square
			for _, assertsSquare := range a.surroundingSquares(addr) {
				asserts := assertsSquare.Piece
				if (asserts.Kind == piece.NoKind) || (asserts.Orientation != o) ||
					(asserts.flags.asserts == false) || asserts.is.normalized {

					continue
				}
				// the neutralizer would move into an assert at addr
				// temporarily adjust board to have the neutralizer moved
				ni := set.From.Index()
				k := a[ni].Kind
				a[ni].Kind = piece.NoKind
				captured := a.neutralizeChainCaptures(to, addr)
				a[ni].Kind = k
				if captured {
					return true
				}
			}
		}
	}

	return false
}

func (a *Board) neutralizeChainCaptures(addr, neutStart Address) bool {
	if addr == neutStart {
		return true
	}
	for _, capturedAddress := range a.neutralizeCaptureAddresses(neutStart) {
		if capturedAddress == addr {
			return true
		}
	}
	return false
}

// threatenedNeutralizerAdjacent indicates if a square is adjacent to a threatened piece that
// neutralizes or adjacent to a chain of threatened neutralizers. This method is recursive, the
// initial call sets the inspected argument to nil.
func (a *Board) threatenedNeutralizerAdjacent(inspected, threats []Address, at Address) bool {
	var insp []Address
	if inspected == nil {
		insp = []Address{at}
	} else {
		insp = append(inspected, at)
	}

LOOP:
	for _, as := range a.surroundingSquares(at) {
		s := a[as.Address.Index()]
		if (s.Kind == piece.NoKind) ||
			(((s.flags.neutralizes == false) || s.is.normalized) && (s.is.ordered == false)) {

			continue
		}
		for _, addr := range threats {
			if addr != as.Address {
				continue
			}
			return true
		}
		for _, addr := range insp {
			if addr == as.Address {
				continue LOOP
			}
		}

		// even if a neutralizer isn't threatened it counts if adjacent ones to it are
		if a.threatenedNeutralizerAdjacent(insp, threats, as.Address) {
			return true
		}
	}
	return false
}
