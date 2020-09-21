package rules

import "github.com/pciet/wichess/piece"

func (a *Board) inCheck(active Orientation, captures []Address) bool {
	king := a.kingLocation(active)
	for _, capture := range captures {
		if capture == king {
			return true
		}
	}
	if a.threatenedNeutralizerAdjacent(nil, captures, king) {
		return true
	}
	return false
}

func (a *Board) removeMovesIntoCheck(moves []MoveSet, active Orientation) []MoveSet {
	out := make([]MoveSet, 0, len(moves))

	for _, moveset := range moves {
		outset := MoveSet{moveset.From, make([]Address, 0, len(moveset.Moves))}
		for _, move := range moveset.Moves {
			// TODO: possible to not copy the entire board here?
			ga := a.afterMove(Move{moveset.From, move})
			threats := movesAddressSlice(ga.naiveCaptureMoves(active.Opponent()))

			if ga.noKing(active) || ga.inCheck(active, captures) ||
				ga.threatenedNeutralizerAdjacent(nil, threats, ga.kingLocation(active)) {
				continue
			}

			outset.Moves = append(outset.Moves, move)
		}
		if len(outset.Moves) == 0 {
			continue
		}
		out = append(out, outset)
	}
	return out
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

	// TODO: loop over threats instead?
LOOP:
	for _, as := range a.surroundingSquares(at) {
		s := a[as.Address.Index()]
		if (s.Kind == piece.NoKind) || (s.flags.neutralizes == false) {
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

		// even if a neutralizer isn't threatened it counts if adjacent ones to that are
		if a.threatenedNeutralizerAdjacent(insp, threats, as.Address) {
			return true
		}
	}
	return false
}
