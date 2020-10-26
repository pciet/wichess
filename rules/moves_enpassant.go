package rules

import "github.com/pciet/wichess/piece"

// if a piece can get to the en passant take square then this will need to be updated

func (a *Board) appendEnPassantMove(moves []Address, at Address, previous Move) []Address {
	if (previous == NoMove) || (previous.From.File != previous.To.File) {
		// if the file isn't the same then an incorrect en passant capture can be done on form moves
		return moves
	}

	s := a[at.Index()]
	if s.Kind.Basic() != piece.Pawn {
		return moves
	}

	p := a[previous.To.Index()]
	if p.Kind.Basic() != piece.Pawn {
		return moves
	}

	var left, right Address
	if s.Orientation == White {
		if (previous.From.Rank != 6) || (previous.To.Rank != 4) {
			return moves
		}
		if s.Kind != piece.Evident {
			if at.Rank != 4 {
				return moves
			}
			left = Address{at.File - 1, at.Rank + 1}
			right = Address{at.File + 1, at.Rank + 1}
		} else {
			// Evident is a special case because it captures backwards, otherwise all pawns capture
			// regularly or capture the square behind so can't do the en passant capture
			if at.Rank != 6 {
				return moves
			}
			left = Address{at.File - 1, at.Rank - 1}
			right = Address{at.File + 1, at.Rank - 1}
		}
	} else {
		if (previous.From.Rank != 1) || (previous.To.Rank != 3) {
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

	switch previous.To.File {
	case left.File:
		return append(moves, left)
	case right.File:
		return append(moves, right)
	}

	return moves
}
