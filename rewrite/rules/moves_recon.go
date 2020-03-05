package rules

func (a Board) AppendReconMoves(moves []Address, at Address) []Address {
	var possibleRecon [3]Address
	s := a[at.Index()]
	if s.Orientation == White {
		possibleRecon[0] = Address{at.File - 1, at.Rank + 1}
		possibleRecon[1] = Address{at.File, at.Rank + 1}
		possibleRecon[2] = Address{at.File + 1, at.Rank + 1}
	} else {
		possibleRecon[0] = Address{at.File - 1, at.Rank - 1}
		possibleRecon[1] = Address{at.File, at.Rank - 1}
		possibleRecon[2] = Address{at.File + 1, at.Rank - 1}
	}

	for _, r := range possibleRecon {
		p := a[r.Index()]
		if (p.Kind == NoKind) ||
			(s.Orientation != p.Orientation) ||
			(p.Recons == false) {
			continue
		}

		var reconMove Address
		if s.Orientation == White {
			reconMove = Address{r.File, r.Rank + 1}
		} else {
			reconMove = Address{r.File, r.Rank - 1}
		}
		mp := a[reconMove.Index()]
		if mp.Kind == NoKind {
			moves = append(moves, reconMove)
			continue
		}
		// recon takes aren't legal, but swaps can happen
		if mp.Orientation != s.Orientation {
			continue
		}
		if s.Swaps {
			moves = append(moves, reconMove)
		}
	}

	return moves
}
