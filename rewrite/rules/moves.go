package rules

// TODO: move results are all calculated here, so cache that for when a move is picked

func (a Game) Moves(active Orientation) ([]MoveSet, State) {
	// promotion is part of the previous move
	if a.PromotionNeeded() {
		return nil, Promotion
	}

	if a.InsufficientMaterialDraw() {
		return nil, Draw
	}

	// calculate all moves the player can make without considering check
	moves := a.NaiveMoves(active)

	// check is a threat of capture, which means takes into check count
	threats := MovesAddressSlice(a.NaiveTakeMoves(active.Opponent()))

	check := a.Board.InCheck(active, threats)

	if check == false {
		moves = a.Board.AppendCastleMoves(moves, active, threats)
	}

	// if the king is on a threatened square or taken after a move then that move is removed
	moves = a.RemoveMovesIntoCheck(moves, active)

	if len(moves) == 0 {
		if check {
			return nil, Checkmate
		}
		return nil, Draw
	}

	return moves, Normal
}
