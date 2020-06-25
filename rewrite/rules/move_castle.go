package rules

import "github.com/pciet/wichess/piece"

func (a Board) IsCastleMove(m Move) bool {
	s := a[m.From.Index()]
	if (s.Kind != piece.King) || s.Moved {
		return false
	}
	if (m.To == Address{2, 0}) || (m.To == Address{6, 0}) ||
		(m.To == Address{2, 7}) || (m.To == Address{6, 7}) {
		return true
	}
	return false
}

func (a Board) CastleMove(changes []AddressedSquare, m Move) []AddressedSquare {
	var rookMove Move
	switch m.To {
	case Address{2, 0}:
		rookMove = Move{Address{0, 0}, Address{3, 0}}
	case Address{6, 0}:
		rookMove = Move{Address{7, 0}, Address{5, 0}}
	case Address{2, 7}:
		rookMove = Move{Address{0, 7}, Address{3, 7}}
	case Address{6, 7}:
		rookMove = Move{Address{7, 7}, Address{5, 7}}
	default:
		Panic("not a castle move", m, a)
	}

	rook := a[rookMove.From.Index()]
	king := a[m.From.Index()]

	king.Moved = true
	rook.Moved = true

	changes = append(changes, AddressedSquare{m.From, Square{}})
	changes = append(changes, AddressedSquare{rookMove.From, Square{}})
	changes = append(changes, AddressedSquare{rookMove.To, rook})
	return append(changes, AddressedSquare{m.To, king})
}
