package rules

import "github.com/pciet/wichess/piece"

func (a Board) PieceLocked(at Address) bool {
	p := a[at.Index()]
	if p.Kind == piece.NoKind {
		Panic("no piece at", at, a)
	}

	for _, s := range a.SurroundingSquares(at) {
		if (s.Kind == piece.NoKind) || (s.Orientation == p.Orientation) {
			continue
		}
		if s.Locks {
			return true
		}
	}

	return false
}
