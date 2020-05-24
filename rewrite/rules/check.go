package rules

func (a Board) InCheck(active Orientation, takes []Address) bool {
	king := a.KingLocation(active)
	for _, take := range takes {
		if take == king {
			return true
		}
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

			if ga.Board.NoKing(active) || ga.Board.InCheck(active, threats) {
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
