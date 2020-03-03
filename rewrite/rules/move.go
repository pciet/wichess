package rules

import (
	"log"
)

func (a Game) AfterMove(m Move) Game {
	d, _ := a.DoMove(m)
	for _, s := range d {
		a[s.Address.Index()] = s.Square
	}
	a.Previous = m
	return a
}

// At least these bad moves can be made with DoMove:
//   putting the king in check
//   skipping a promotion
//   moving a locked piece
//   moves that aren't in the piece's move set
//   pawn takes a fortified piece
//   en passant turns later
//   castling through threatened squares, during check, or without a rook
//   swapping with a friendly piece without having the swap ability

// Returns the squares that changed and the squares with each piece that was taken.
// No move legality is determined, bad moves either cause a panic or happen.
func (a Board) DoMove(m Move) ([]AddressedSquare, []AddressedSquare) {
	from := a[m.From.Index()]
	if from.Empty() {
		log.Panicln("no piece for move", m, a)
	}

	changes := make([]AddressedSquare, 0, 3)
	takes := make([]AddressedSquare, 0, 1)

	to := a[m.To.Index()]
	if to.NotEmpty() {
		if to.SameOrientation(from) {
			changes = a.SwapMove(changes, m)
		} else {
			if to.Detonates {
				return a.DetonateMove(changes, takes, m)
			}
			changes, takes = a.TakeMove(changes, takes, m)
		}
	} else {
		if a.IsCastleMove(m) {
			return a.CastleMove(changes, m), nil
		}
		if a.IsEnPassantMove(m) {
			changes, takes = a.EnPassantMove(changes, takes, m)
		} else {
			changes = a.NoTakeMove(changes, m)
		}
	}

	for _, s := range a.SurroundingSquares(m.To) {
		if a.GuardWillTake(from, s) == false {
			continue
		}
		if from.Detonates {
			return a.GuardTakesDetonate(changes, takes, m, s.Address)
		}
		return a.GuardChain(changes, takes, m, s.Address)
	}

	return changes, takes
}

func (a Board) NoTakeMove(changes []AddressedSquare, m Move) []AddressedSquare {
	s := a[m.From.Index()]
	s.Moved = true
	changes = append(changes, AddressedSquare{m.From, EmptySquare()})
	return append(changes, AddressedSquare{m.To, s})
}

func (a Board) TakeMove(changes, takes []AddressedSquare, m Move) ([]AddressedSquare, []AddressedSquare) {
	s := a[m.From.Index()]
	s.Moved = true
	changes = append(changes, AddressedSquare{m.From, EmptySquare()})
	return append(changes, AddressedSquare{m.To, s}), append(takes, AddressedSquare{m.To, a[m.To.Index()]})
}
