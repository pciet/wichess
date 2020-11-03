package rules

import "log"

func (a *Board) inCheck(active Orientation, captures []Address, previous Move) bool {
	king := a.kingLocation(active)
	for _, capture := range captures {
		if capture == king {
			return true
		}
	}
	if a.threatenedNeutralizerAdjacent(nil, captures, king) ||
		a.neutralizesAssertedChainAdjacent(king, previous) {

		return true
	}
	return false
}

// removeMovesIntoCheck is called by Moves but not DoMove.
func (a *Board) removeMovesIntoCheck(moves []MoveSet, active Orientation, previous Move) []MoveSet {
	out := make([]MoveSet, 0, len(moves))
	for _, moveset := range moves {
		outset := MoveSet{moveset.From, make([]Address, 0, len(moveset.Moves))}
		for _, moveAddr := range moveset.Moves {
			move := Move{moveset.From, moveAddr}

			// stash original squares and apply move changes
			changes, _ := a.DoMove(move, previous)
			if multipleSquareInSlice(changes) {
				// TODO: consider making this a debug only check
				// this has been a symptom of multiple obscure mistakes, so added this if here
				log.Panicln("multiple versions of square returned by Board.DoMove:", move,
					changes, "\n", a.String())
			}
			orig := make([]Square, len(changes))
			for i, change := range changes {
				addr := change.Address.Index()
				orig[i] = Square{change.Address, a[addr]}
				a[addr] = change.Piece
			}
			a.applyConveyedCharacteristics()

			if (a.noKing(active) == false) &&
				(a.inCheck(active, movesAddressSlice(a.naiveCaptureMoves(active.Opponent(), move)),
					move) == false) {

				outset.Moves = append(outset.Moves, moveAddr)
			}

			// revert board back to original state
			for _, s := range orig {
				a[s.Address.Index()] = s.Piece
			}
			a.applyConveyedCharacteristics()
		}
		if len(outset.Moves) == 0 {
			continue
		}
		out = append(out, outset)
	}
	return out
}
