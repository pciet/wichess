package rules

import "github.com/pciet/wichess/piece"

func (a Board) AppendReconMoves(moves []Address, at Address) []Address {
	possibleRecon := make([]Address, 0, 3)
	s := a[at.Index()]
	// TODO: simplify this if
	if s.Orientation == White {
		if at.File != 0 {
			possibleRecon = append(possibleRecon, Address{at.File - 1, at.Rank + 1})
		}
		possibleRecon = append(possibleRecon, Address{at.File, at.Rank + 1})
		if at.File != 7 {
			possibleRecon = append(possibleRecon, Address{at.File + 1, at.Rank + 1})
		}
	} else {
		if at.File != 0 {
			possibleRecon = append(possibleRecon, Address{at.File - 1, at.Rank - 1})
		}
		possibleRecon = append(possibleRecon, Address{at.File, at.Rank - 1})
		if at.File != 7 {
			possibleRecon = append(possibleRecon, Address{at.File + 1, at.Rank - 1})
		}
	}

	for _, r := range possibleRecon {
		ra := r.Index()
		if ra >= 64 {
			continue
		}
		p := a[ra]
		if (p.Kind == piece.NoKind) ||
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
		ra = reconMove.Index()
		if ra >= 64 {
			continue
		}
		mp := a[ra]
		if mp.Kind == piece.NoKind {
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
