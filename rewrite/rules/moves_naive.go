package rules

// Moves that don't consider check aren't legal but are a first calculation, done with NaiveMoves.
// All types of moves are calculated except for just castle which is never available in check.
func (a Game) NaiveMoves(active Orientation) []MoveSet {
	moves := make([]MoveSet, 0, 16)
	for i, s := range a.Board {
		if (s.Kind == NoKind) || (s.Orientation != active) {
			continue
		}
		at := AddressIndex(i).Address()
		nm := a.NaiveMovesAt(at)
		if len(nm) == 0 {
			continue
		}
		moves = append(moves, MoveSet{at, nm})
	}
	return RemoveDuplicateMoveSetMoves(moves)
}

func (a Game) NaiveMovesAt(the Address) []Address {
	s := a.Board[the.Index()]

	if (s.Kind == NoKind) || a.Board.PieceLocked(the) {
		return []Address{}
	}

	pathvariations := AppliedRelPaths(s.Kind, the, s.Orientation)

	moves := make([]Address, 0, 8)

	if s.Moved == false {
		moves = a.Board.AppendNaiveMoves(moves, pathvariations[First], the)
	} else {
		moves = a.Board.AppendNaiveMoves(moves, pathvariations[NormalMove], the)
	}

	if a.Board.PieceRallied(the) {
		moves = a.Board.AppendNaiveMoves(moves, pathvariations[RallyMove], the)
	}

	moves = a.Board.AppendNaiveTakeMoves(moves, pathvariations[Take], the)
	moves = a.AppendEnPassantMove(moves, the)

	// TODO: only look at recon moves when the piece has the recon characteristic
	moves = a.Board.AppendReconMoves(moves, the)

	return moves
}

// TODO: NaiveTakeMoves and NaiveTakeMovesAt are mostly copied from the above two funcs,
// don't repeat

func (a Game) NaiveTakeMoves(active Orientation) []MoveSet {
	moves := make([]MoveSet, 0, 16)
	for i, s := range a.Board {
		if (s.Kind == NoKind) || (s.Orientation != active) {
			continue
		}
		at := AddressIndex(i).Address()
		nm := a.NaiveTakeMovesAt(at)
		if len(nm) == 0 {
			continue
		}
		moves = append(moves, MoveSet{at, nm})
	}
	return moves
}

func (a Game) NaiveTakeMovesAt(the Address) []Address {
	s := a.Board[the.Index()]

	if (s.Kind == NoKind) || a.Board.PieceLocked(the) {
		return []Address{}
	}

	// TODO: to expedite getting it working rewriting this to just get the take moves was skipped
	takes := AppliedRelPaths(s.Kind, the, s.Orientation)[Take]

	moves := make([]Address, 0, 8)

	moves = a.Board.AppendNaiveTakeMoves(moves, takes, the)
	moves = a.AppendEnPassantMove(moves, the)

	return moves
}

func (a Board) AppendNaiveMoves(moves []Address, paths []Path, from Address) []Address {
	s := a[from.Index()]
	for _, path := range paths {
		if path.Truncated && s.MustEnd {
			continue
		}
		for i, move := range path.Addresses {
			piece := a[move.Index()]
			if piece.Kind == NoKind {
				if s.MustEnd && (len(path.Addresses) != i+1) {
					continue
				}
				moves = append(moves, move)
				continue
			}
			if (s.Orientation == piece.Orientation) && s.Swaps {
				moves = append(moves, move)
			}
			if s.Ghost {
				continue
			}
			break
		}
	}
	return moves
}

func (a Board) AppendNaiveTakeMoves(moves []Address, paths []Path, from Address) []Address {
	s := a[from.Index()]
	for _, path := range paths {
		if path.Truncated && s.MustEnd {
			continue
		}
		for i, move := range path.Addresses {
			piece := a[move.Index()]
			// TODO: these next two ifs are the only difference from AppendNaiveMoves, combine?
			if (piece.Kind == NoKind) ||
				(s.MustEnd && (len(path.Addresses) != i+1)) {
				continue
			}
			if (s.Orientation != piece.Orientation) &&
				((BasicKind(s.Kind) != Pawn) || (piece.Fortified == false)) {
				moves = append(moves, move)
			}
			if s.Ghost {
				continue
			}
			break
		}
	}
	return moves
}
