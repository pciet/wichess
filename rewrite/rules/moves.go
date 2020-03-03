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

	moves, opponentMoves := a.NaiveMoves(active)

	check := a.Check(active, opponentMoves)

	if check {
		moves = a.RemoveCastling(active, moves)
	}

	// remove all moves that would cause check
	for i, moveSet := range moves {
		for j, move := range moveSet {
			ga := a.AfterMove(moveSet.From, move)
			om, _ := ga.NaiveMoves(Opponent(active))
			if ga.Check(active, om) {
				moves = RemoveMove(moves, i, j)
			}
		}
	}

	if len(moves) == 0 {
		if check {
			return nil, Checkmate
		}
		return nil, Draw
	}

	return moves, Normal
}

// Moves that don't consider if they'll allow check aren't legal but are a first calculation.
// Returns the active player's moves and the opponent's moves in separate slices.
func (a Game) NaiveMoves(active Orientation) ([]MoveSet, []MoveSet) {
	m := make([]MoveSet, 0, 16)
	om := make([]MoveSet, 0, 16)

	for i, p := range a.Board {
		if p.Kind == NoKind {
			continue
		}
		addr := AddressIndex(i).Address()
		nm := a.NaiveMovesAt(addr)
		if len(nm) == 0 {
			continue
		}
		move := MoveSet{Addr, nm}
		if p.Orientation == active {
			m = append(m, move)
		} else {
			om = append(om, move)
		}
	}

	return m, om
}

// Both castling and en passant are included in NaiveMovesAt.
func (a Game) NaiveMovesAt(the Address) []Address {

}
