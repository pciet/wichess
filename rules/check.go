package rules

import "github.com/pciet/wichess/piece"

func (a Board) InCheck(active Orientation, takes []Address) bool {
	king := a.KingLocation(active)
	for _, take := range takes {
		if take == king {
			return true
		}
	}
	if a.ThreatenedDetonatorAdjacent(nil, takes, king) {
		return true
	}
	return false
}

func (a Game) RemoveMovesIntoCheck(moves []MoveSet, active Orientation) []MoveSet {
	out := make([]MoveSet, 0, len(moves))

	for _, moveset := range moves {
		outset := MoveSet{moveset.From, make([]Address, 0, len(moveset.Moves))}
		for _, move := range moveset.Moves {
			ga := a.AfterMove(Move{moveset.From, move})
			threats := MovesAddressSlice(ga.NaiveTakeMoves(active.Opponent()))

			if ga.Board.NoKing(active) || ga.Board.InCheck(active, threats) ||
				ga.Board.ThreatenedDetonatorAdjacent(nil, threats, ga.Board.KingLocation(active)) {
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

// ThreatenedDetonatorAdjacent indicates if a square is adjacent to a threatened piece that
// detonates or adjacent to a chain of threatened detonators. This method is recursive, the
// initial call sets the inspected argument to nil.
func (a Board) ThreatenedDetonatorAdjacent(inspected, threats []Address, at Address) bool {
	var insp []Address
	if inspected == nil {
		insp = []Address{at}
	} else {
		insp = append(inspected, at)
	}

	// TODO: loop over threats instead?
LOOP:
	for _, as := range a.SurroundingSquares(at) {
		s := a[as.Address.Index()]
		if (s.Kind == piece.NoKind) || (s.Detonates == false) {
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

		// even if a detonator isn't threatened it counts if adjacent ones to that are
		if a.ThreatenedDetonatorAdjacent(insp, threats, as.Address) {
			return true
		}
	}
	return false
}
