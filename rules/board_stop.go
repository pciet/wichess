package rules

import (
	"log"

	"github.com/pciet/wichess/piece"
)

func (a *Board) pieceStopped(at Address) bool {
	p := a[at.Index()]
	if p.Kind == piece.NoKind {
		log.Panicln("no piece at", at, a)
	}

	if (p.Kind == piece.King) || (p.Kind == piece.Queen) {
		return false
	}

	for _, s := range a.SurroundingSquares(at) {
		if (s.Kind == piece.NoKind) || (s.Orientation == p.Orientation) {
			continue
		}
		if s.flags.stops {
			return true
		}
	}

	return false
}
