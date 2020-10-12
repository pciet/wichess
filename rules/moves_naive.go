package rules

import "github.com/pciet/wichess/piece"

// Moves that don't consider check aren't legal but are a first calculation, done with naiveMoves.
// All types of moves are calculated except for just castle which is never available in check.
func (a *Board) naiveMoves(active Orientation, previous Move) []MoveSet {
	moves := make([]MoveSet, 0, 16)
	for i, s := range a {
		if (s.Kind == piece.NoKind) || (s.Orientation != active) {
			continue
		}
		at := AddressIndex(i).Address()
		nm := a.naiveMovesAt(at, previous)
		if len(nm) == 0 {
			continue
		}
		moves = append(moves, MoveSet{at, nm})
	}
	return removeDuplicateMoveSetMoves(moves)
}

func (a *Board) naiveMovesAt(the Address, previous Move) []Address {
	s := a[the.Index()]

	// TODO: NoKind check done twice from Game.NaiveMoves
	if (s.Kind == piece.NoKind) || s.is.stopped {
		return []Address{}
	}

	pathvariations := appliedPaths(s.Kind, the, s.Orientation)

	moves := make([]Address, 0, 8)

	if s.Moved == false {
		moves = a.appendNaiveMoves(moves, pathvariations[piece.FirstPaths], the)
	} else {
		moves = a.appendNaiveMoves(moves, pathvariations[piece.NormalPaths], the)
	}

	if s.is.enabled {
		moves = a.appendNaiveMoves(moves, pathvariations[piece.RallyPaths], the)
	}

	moves = a.appendNaiveCaptureMoves(moves, pathvariations[piece.CapturePaths], the)
	moves = a.appendEnPassantMove(moves, the, previous)

	// TODO: only look at reveals moves when the piece has the recon characteristic
	moves = a.appendRevealMoves(moves, the)

	return moves
}

// TODO: NaiveCaptureMoves and NaiveCaptureMovesAt are mostly copied from the above two funcs,
// don't repeat

func (a *Board) naiveCaptureMoves(active Orientation, previous Move) []MoveSet {
	moves := make([]MoveSet, 0, 16)
	for i, s := range a {
		if (s.Kind == piece.NoKind) || (s.Orientation != active) {
			continue
		}
		at := AddressIndex(i).Address()
		nm := a.naiveCaptureMovesAt(at, previous)
		if len(nm) == 0 {
			continue
		}
		moves = append(moves, MoveSet{at, nm})
	}
	return moves
}

func (a *Board) naiveCaptureMovesAt(the Address, previous Move) []Address {
	s := a[the.Index()]

	if (s.Kind == piece.NoKind) || s.is.stopped {
		return []Address{}
	}

	// TODO: to expedite getting it working rewriting this to just get the take moves was skipped
	takes := appliedPaths(s.Kind, the, s.Orientation)[piece.CapturePaths]

	moves := make([]Address, 0, 8)

	moves = a.appendNaiveCaptureMoves(moves, takes, the)
	moves = a.appendEnPassantMove(moves, the, previous)

	return moves
}

func (a *Board) appendNaiveMoves(moves []Address, paths []path, from Address) []Address {
	s := a[from.Index()]
	for _, path := range paths {
		if path.Truncated && s.flags.mustEnd {
			continue
		}
		for i, move := range path.Addresses {
			p := a[move.Index()]
			if p.Kind == piece.NoKind {
				if s.flags.mustEnd && (len(path.Addresses) != i+1) {
					continue
				}
				moves = append(moves, move)
				continue
			}
			// normalized doesn't apply to quick
			if s.flags.quick {
				continue
			}
			break
		}
	}
	return moves
}

func (a *Board) appendNaiveCaptureMoves(moves []Address, paths []path, from Address) []Address {
	s := a[from.Index()]
	for _, path := range paths {
		if path.Truncated && s.flags.mustEnd {
			continue
		}
		for i, move := range path.Addresses {
			p := a[move.Index()]
			if (p.Kind == piece.NoKind) ||
				(s.flags.quick && s.flags.mustEnd && (len(path.Addresses) != i+1)) {

				continue
			}
			if (s.Orientation != p.Orientation) && ((&p).immaterialAgainst(&s) == false) &&
				((p.flags.tense == false) || p.is.normalized ||
					(s.Kind == piece.King) || (s.Kind == piece.Queen)) &&
				((s.flags.mustEnd == false) || (len(path.Addresses) == i+1)) &&
				((p.is.protected == false) || p.is.normalized) {

				moves = append(moves, move)
			}
			if s.flags.quick && (s.flags.noOverCapture == false) {
				continue
			}
			break
		}
	}
	return moves
}
