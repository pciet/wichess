package rules

func (a *Board) assertsWillCapture(target Piece, asserts Square) bool {
	return asserts.notEmpty() && asserts.flags.asserts &&
		(asserts.Orientation != target.Orientation) &&
		(target.immaterialAgainst(&asserts.Piece) == false) && (target.flags.stops == false) &&
		(a.pieceStopped(asserts.Address) == false)
}

// This function changes the Board.
func (a *Board) assertsCapturesNeutralizes(changes, captures []Square, m Move,
	asserts Address) ([]Square, []Square) {

	// treat this like another move
	a.applyChanges(changes)

	assertsNeutralizeChanges := make([]Square, 0, 8)
	assertsNeutralizeCaptures := make([]Square, 0, 2)

	assertsNeutralizeChanges, assertsNeutralizeCaptures =
		a.neutralizesMove(assertsNeutralizeChanges, assertsNeutralizeCaptures, Move{asserts, m.To})

	changes = MergeReplaceSquares(changes, assertsNeutralizeChanges)
	captures = combineSquares(captures, assertsNeutralizeCaptures)

	// fix take address of neutralizes to match original board
	for i, s := range captures {
		if s.Address == m.To {
			captures[i].Address = m.From
			break
		}
	}

	return changes, captures
}

// This function changes the Board.
func (a *Board) assertsChain(changes, captures []Square, m Move,
	asserts Address) ([]Square, []Square) {

	changes = append(changes, Square{asserts, Piece{}})
	g := a[asserts.Index()]
	g.Moved = true
	changes = append(changes, Square{m.To, g})
	captures = append(captures, Square{m.From, a[m.From.Index()]})
	previousAsserts := Square{asserts, g}

	a.applyChanges(changes)

	// if the newly moved asserts is now adjacent to an enemy asserts then more assert moves happen
	// keep applying asserts moves until none are left
LOOP:
	for {
		for _, s := range a.surroundingSquares(m.To) {
			if a.assertsWillCapture(a[m.To.Index()], s) == false {
				continue
			}
			captures = append(captures, previousAsserts)
			previousAsserts = s
			s.Moved = true
			gchanges := make([]Square, 0, 2)
			gchanges = append(gchanges, Square{m.To, s.Piece})
			gchanges = append(gchanges, Square{s.Address, Piece{}})
			changes = MergeReplaceSquares(changes, gchanges)
			a.applyChanges(gchanges)
			continue LOOP
		}
		break
	}

	return changes, captures
}
