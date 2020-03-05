package rules

// The king should not be in check when AppendCastleMoves is called.
func (a Board) AppendCastleMoves(moves []MoveSet, by Orientation, opponentThreats []Address) []MoveSet {
	var king Address
	if by == White {
		king = WhiteKingStart
	} else {
		king = BlackKingStart
	}

	s := a[king.Index()]
	if (s.Kind != King) || s.Moved {
		return moves
	}

	// The intermediates need to be ordered as a path from the king.
	appendMove := func(intermediates []Address, rook Address, castleMove Address) {
		r := a[rook.Index()]
		if (BasicKind(r.Kind) != Rook) || r.Moved {
			return
		}

		// if intermediate squares are not empty then castle not available
		for i, inter := range intermediates {
			if a[inter.Index()].Kind != NoKind {
				return
			}
			// the first two squares also need to be unthreatened
			if i > 1 {
				continue
			}
			for _, move := range opponentThreats {
				if inter == move {
					return
				}
			}
		}
		for i, moveset := range moves {
			if moveset.From != king {
				continue
			}
			moves[i].Moves = append(moves[i].Moves, castleMove)
			return
		}
		moves = append(moves, MoveSet{king, []Address{castleMove}})
	}

	var leftRook, rightRook, leftCastle, rightCastle Address
	var left, right []Address

	if by == White {
		leftRook = WhiteLeftRookStart
		rightRook = WhiteRightRookStart
		left = []Address{{3, 0}, {2, 0}, {1, 0}}
		right = []Address{{5, 0}, {6, 0}}
		leftCastle = Address{2, 0}
		rightCastle = Address{6, 0}
	} else {
		// reversed because naming in this function is from the white perspective
		leftRook = BlackRightRookStart
		rightRook = BlackLeftRookStart
		left = []Address{{3, 7}, {2, 7}, {1, 7}}
		right = []Address{{5, 7}, {6, 7}}
		leftCastle = Address{2, 7}
		rightCastle = Address{6, 7}
	}

	appendMove(left, leftRook, leftCastle)
	appendMove(right, rightRook, rightCastle)

	return moves
}
