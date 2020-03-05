package rules

import (
	"log"
)

func (a Board) PieceLocked(at Address) bool {
	p := a[at.Index()]
	if p.Kind == NoKind {
		log.Panicln("no piece at", at, a)
	}

	for _, s := range a.SurroundingSquares(at) {
		if (s.Kind == NoKind) || (s.Orientation == p.Orientation) {
			continue
		}
		if s.Locks {
			return true
		}
	}

	return false
}
