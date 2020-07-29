package rules

import "github.com/pciet/wichess/piece"

// TODO: move results are all calculated here, so cache that for when a move is picked

func (a Game) Moves(active Orientation) ([]MoveSet, State) {
	// promotion is part of the previous move
	_, needed := a.PromotionNeeded()
	if needed {
		return nil, Promotion
	}

	if a.InsufficientMaterialDraw() {
		return nil, Draw
	}

	// apply characteristic changes caused by other pieces

	// first remove characteristics due to normalizes
	for i, s := range a.Board {
		if (s.Kind == piece.NoKind) || (s.Normalizes == false) {
			continue
		}
		for _, ss := range a.Board.SurroundingSquares(AddressIndex(i).Address()) {
			if ss.Kind == piece.NoKind {
				continue
			}
			Normalize(&(a.Board[ss.Address.Index()]))
		}
	}

	// apply keep and orders
	for i, s := range a.Board {
		if (s.Kind == piece.NoKind) || ((s.Keep == false) && (s.Orders == false)) {
			continue
		}
		for _, ss := range a.Board.SurroundingSquares(AddressIndex(i).Address()) {
			if (ss.Kind == piece.NoKind) || (ss.Orientation != s.Orientation) {
				continue
			}
			if s.Keep && (ss.Orientation == s.Orientation) {
				a.Board[ss.Address.Index()].Fortified = true
			}
			if s.Orders {
				a.Board[ss.Address.Index()].Detonates = true
			}
		}
	}

	// calculate all moves the player can make without considering check
	moves := a.NaiveMoves(active)

	// check is a threat of capture, which means takes into check count
	threats := MovesAddressSlice(a.NaiveMoves(active.Opponent()))

	check := a.Board.InCheck(active, threats)

	if check == false {
		moves = a.Board.AppendCastleMoves(moves, active, threats)
	} else {
		moves = a.Board.AppendExtricateMoves(moves, active)
	}

	// if the king is on a threatened square or taken after a move then that move is removed
	moves = a.RemoveMovesIntoCheck(moves, active)

	if len(moves) == 0 {
		if check {
			return nil, Checkmate
		}
		return nil, Draw
	}

	if check {
		return moves, Check
	}

	return moves, Normal
}
