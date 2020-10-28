package rules

import (
	"log"

	"github.com/pciet/wichess/piece"
)

func (a *Board) isEnPassantMove(m Move) bool {
	if m.From.File == m.To.File {
		return false
	}
	s := a[m.From.Index()]

	// derange and imperfect are the two pawns that only captures the square behind, and the
	// evident is the only pawn that captures backwards
	if (s.Kind.Basic() != piece.Pawn) || (s.Kind == piece.Derange) || (s.Kind == piece.Imperfect) {
		return false
	} else if s.Kind == piece.Evident {
		if ((s.Orientation == Black) && (m.From.Rank != 1)) ||
			((s.Orientation == White) && (m.From.Rank != 6)) {

			return false
		}
	} else if ((s.Orientation == Black) && (m.From.Rank != 3)) ||
		((s.Orientation == White) && (m.From.Rank != 4)) {

		return false
	}

	to := a[m.To.Index()]
	if to.Empty() == false {
		return false
	}
	taking := a[enPassantCaptureAddress(s.Orientation, m.To).Index()]
	if taking.Empty() || (taking.Orientation == s.Orientation) || (taking.Basic() != piece.Pawn) {
		return false
	}
	return true
}

func (a *Board) enPassantMove(changes, captures []Square, m Move) ([]Square, []Square) {
	s := a[m.From.Index()]
	s.Moved = true // TODO: not sure why this is necessary
	taking := enPassantCaptureAddress(s.Orientation, m.To)
	target := a[taking.Index()]
	if (target.flags.neutralizes && (target.is.normalized == false)) || target.is.ordered {
		return a.neutralizesMove(changes, captures, Move{m.From, taking})
	}
	changes = append(changes, Square{taking, Piece{}})
	changes = append(changes, Square{m.From, Piece{}})
	return append(changes, Square{m.To, s}), append(captures, Square{taking, a[taking.Index()]})
}

func enPassantCaptureAddress(taker Orientation, to Address) Address {
	if taker == White {
		return Address{to.File, to.Rank - 1}
	}
	if taker != Black {
		log.Panicln("orientation", taker, "not white or black")
	}
	return Address{to.File, to.Rank + 1}
}
