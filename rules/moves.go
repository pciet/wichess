package rules

import "github.com/pciet/wichess/piece"

// TODO: move results are all calculated here, so cache that for when a move is picked

// Moves calculates all moves that can be done by the active player and the state of the game.
// The previous move is needed because en passant calculations need to know when the pawn to
// capture was moved.
func (a *Board) Moves(active Orientation, previous Move) ([]MoveSet, State) {
	// promotion is part of the previous move
	_, needed := a.PromotionNeeded()
	if needed {
		return nil, Promotion
	}

	if a.insufficientMaterialDraw() {
		return nil, Draw
	}

	bcopy := a.Copy()

	// apply characteristic changes caused by other pieces

	// first remove characteristics due to normalizes
	for i, s := range bcopy {
		if (s.Kind == piece.NoKind) || (s.flags.normalizes == false) {
			continue
		}
		for _, ss := range bcopy.surroundingSquares(AddressIndex(i).Address()) {
			if ss.Kind == piece.NoKind {
				continue
			}
			normalize(&(bcopy[ss.Address.Index()].flags))
		}
	}

	// apply keep and orders
	for i, s := range bcopy {
		if (s.Kind == piece.NoKind) || ((s.flags.keep == false) && (s.flags.orders == false)) {
			continue
		}
		for _, ss := range bcopy.surroundingSquares(AddressIndex(i).Address()) {
			if ss.Kind == piece.NoKind {
				continue
			}
			if s.flags.keep && (ss.Orientation == s.Orientation) {
				bcopy[ss.Address.Index()].flags.immaterial = true
			}
			if s.flags.orders {
				bcopy[ss.Address.Index()].flags.neutralizes = true
			}
		}
	}

	// calculate all moves the player can make without considering check
	moves := bcopy.naiveMoves(active, previous)

	// TODO: what should previous be for the threats call?

	// check is a threat of capture, which means takes into check count
	threats := movesAddressSlice(bcopy.naiveMoves(active.Opponent(), previous))

	check := bcopy.inCheck(active, threats)

	if check == false {
		moves = bcopy.appendCastleMoves(moves, active, threats)
	} else {
		moves = bcopy.appendExtricateMoves(moves, active)
	}

	// if the king is on a threatened square or taken after a move then that move is removed
	moves = bcopy.removeMovesIntoCheck(moves, active, previous)

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
