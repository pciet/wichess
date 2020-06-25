package rules

import "github.com/pciet/wichess/piece"

func (a Board) IsEnPassantMove(m Move) bool {
	if m.From.File == m.To.File {
		return false
	}
	s := a[m.From.Index()]
	if s.Kind.Basic() != piece.Pawn {
		return false
	}
	to := a[m.To.Index()]
	if to.Empty() == false {
		return false
	}
	taking := a[EnPassantTakeAddress(s.Orientation, m.To).Index()]
	if taking.Empty() || (taking.Orientation == s.Orientation) {
		return false
	}
	return true
}

func (a Board) EnPassantMove(changes, takes []AddressedSquare,
	m Move) ([]AddressedSquare, []AddressedSquare) {

	s := a[m.From.Index()]
	taking := EnPassantTakeAddress(s.Orientation, m.To)

	changes = append(changes, AddressedSquare{taking, Square{}})
	changes = append(changes, AddressedSquare{m.From, Square{}})

	return append(changes, AddressedSquare{m.To, s}),
		append(takes, AddressedSquare{taking, a[taking.Index()]})
}

func EnPassantTakeAddress(taker Orientation, to Address) Address {
	if taker == White {
		return Address{to.File, to.Rank - 1}
	}
	if taker != Black {
		Panic("orientation", taker, "not white or black")
	}
	return Address{to.File, to.Rank + 1}
}
