package rules

func (a Board) GuardWillTake(target Square, guard AddressedSquare) bool {
	return guard.NotEmpty() && guard.Guards &&
		(guard.Orientation != target.Orientation) &&
		(target.FortifiedAgainst(guard.Square) == false) && (target.Locks == false) &&
		(a.PieceLocked(guard.Address) == false)
}

func (a Board) GuardTakesDetonate(changes, takes []AddressedSquare, m Move,
	guard Address) ([]AddressedSquare, []AddressedSquare) {
	// treat this like another move
	(&a).ApplyChanges(changes)

	guardDetonateChanges := make([]AddressedSquare, 0, 8)
	guardDetonateTakes := make([]AddressedSquare, 0, 2)

	guardDetonateChanges, guardDetonateTakes = a.DetonateMove(guardDetonateChanges,
		guardDetonateTakes, Move{guard, m.To})

	changes = MergeReplaceAddressedSquares(changes, guardDetonateChanges)
	takes = CombineAddressedSquares(takes, guardDetonateTakes)

	// fix take address of detonator to match original board
	for i, s := range takes {
		if s.Address == m.To {
			takes[i].Address = m.From
			break
		}
	}

	return changes, takes
}

func (a Board) GuardChain(changes, takes []AddressedSquare, m Move,
	guard Address) ([]AddressedSquare, []AddressedSquare) {
	changes = append(changes, AddressedSquare{guard, Square{}})
	g := a[guard.Index()]
	g.Moved = true
	changes = append(changes, AddressedSquare{m.To, g})
	takes = append(takes, AddressedSquare{m.From, a[m.From.Index()]})
	previousGuard := AddressedSquare{guard, g}

	(&a).ApplyChanges(changes)

	// if the newly moved guard is now adjacent to an enemy guard then more guard moves happen
	// keep applying guard moves until none are left
LOOP:
	for {
		for _, s := range a.SurroundingSquares(m.To) {
			if a.GuardWillTake(a[m.To.Index()], s) == false {
				continue
			}
			takes = append(takes, previousGuard)
			previousGuard = s
			s.Moved = true
			gchanges := make([]AddressedSquare, 0, 2)
			gchanges = append(gchanges, AddressedSquare{m.To, s.Square})
			gchanges = append(gchanges, AddressedSquare{s.Address, Square{}})
			changes = MergeReplaceAddressedSquares(changes, gchanges)
			(&a).ApplyChanges(gchanges)
			continue LOOP
		}
		break
	}

	return changes, takes
}
