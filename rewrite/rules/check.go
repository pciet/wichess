package rules

// TODO: does check need to use un-naive opponent moves?

func (a Game) Check(active Orientation, naiveOpponentMoves []MoveSet) bool {
	kl, _ := a.KingLocation(active)
	for _, moveSet := range naiveOpponentMoves {
		for _, move := range moveSet {
			if move == kl {
				return true
			}
		}
	}

	// TODO: isolate specific cases where indirect takes happens to reduce this computation

	// some special piece moves cause indirect takes
	for _, moveSet := range naiveOpponentMoves {
		for _, move := range moveSet {
			_, has := a.AfterMove(Opponent(active), moveSet.From, move).KingLocation(active)
			if has == false {
				return true
			}
		}
	}

	return false
}
