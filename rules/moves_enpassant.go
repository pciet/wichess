package rules

import "github.com/pciet/wichess/piece"

// if a piece can get to the en passant take square then this will need to be updated

func (a Game) AppendEnPassantMove(moves []Address, at Address) []Address {
	if a.Previous == NoPreviousMove {
		return moves
	}

	s := a.Board[at.Index()]
	if s.Kind.Basic() != piece.Pawn {
		return moves
	}

	p := a.Board[a.Previous.To.Index()]
	if p.Kind.Basic() != piece.Pawn {
		return moves
	}

	var left, right Address
	if s.Orientation == White {
		if (a.Previous.From.Rank != 6) || (a.Previous.To.Rank != 4) {
			return moves
		}
		if s.Kind != piece.Evident {
			if at.Rank != 4 {
				return moves
			}
			left = Address{at.File - 1, at.Rank + 1}
			right = Address{at.File + 1, at.Rank + 1}
		} else {
			// Evident is a special case because it captures backwards
			if at.Rank != 6 {
				return moves
			}
			left = Address{at.File - 1, at.Rank - 1}
			right = Address{at.File + 1, at.Rank - 1}
		}
	} else {
		if (a.Previous.From.Rank != 1) || (a.Previous.To.Rank != 3) {
			return moves
		}
		if s.Kind != piece.Evident {
			if at.Rank != 3 {
				return moves
			}
			left = Address{at.File - 1, at.Rank - 1}
			right = Address{at.File + 1, at.Rank - 1}
		} else {
			if at.Rank != 1 {
				return moves
			}
			left = Address{at.File - 1, at.Rank + 1}
			right = Address{at.File + 1, at.Rank + 1}
		}
	}

	switch a.Previous.To.File {
	case left.File:
		return append(moves, left)
	case right.File:
		return append(moves, right)
	}

	return moves
}
