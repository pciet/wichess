package rules

func (a *Board) assertsWillCapture(target Piece, asserts Square) bool {
	return asserts.notEmpty() && asserts.flags.asserts && (asserts.is.normalized == false) &&
		(asserts.Orientation != target.Orientation) && (asserts.is.stopped == false) &&
		(target.immaterialAgainst(&asserts.Piece) == false) && (target.is.protected == false)
}

// This function changes the Board.
func (a *Board) assertsCapturesNeutralizes(changes, captures []Square, m Move,
	asserts Address) ([]Square, []Square) {

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

// TODO: assertsChain could probably be clearer

// assertsChain does the asserts capture for DoMove and chains any additional asserting pieces that
// become in range. The Board is expected to have already been changed to put the piece that will
// be initially captured on the square it moved to.
func (a *Board) assertsChain(changes, captures []Square, m Move,
	asserts Address) ([]Square, []Square) {

	asserter := a[asserts.Index()]
	asserter.Moved = true
	changes = MergeReplaceSquares(changes, []Square{{asserts, Piece{}}, {m.To, asserter}})
	captures = append(captures, Square{m.From, a[m.To.Index()]})
	previousAsserts := Square{asserts, asserter}

	a.ApplyChanges(changes)
	a.applyConveyedCharacteristics()

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
			a.ApplyChanges(gchanges)
			a.applyConveyedCharacteristics()
			continue LOOP
		}
		break
	}

	return changes, captures
}
