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
	if ((s.Orientation == Black) && (m.From.Rank != 3)) ||
		((s.Orientation == White) && (m.From.Rank != 4)) || (s.Kind.Basic() != piece.Pawn) {
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

func (a *Board) enPassantMove(changes, takes []Square, m Move) ([]Square, []Square) {
	s := a[m.From.Index()]
	taking := enPassantCaptureAddress(s.Orientation, m.To)
	changes = append(changes, Square{taking, Piece{}})
	changes = append(changes, Square{m.From, Piece{}})
	return append(changes, Square{m.To, s}), append(takes, Square{taking, a[taking.Index()]})
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
