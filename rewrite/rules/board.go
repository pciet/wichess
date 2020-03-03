package rules

import (
	"log"
)

// A Wisconsin Chess board is a regular 8x8 chess board.
type Board [8 * 8]Square

func (a Board) SurroundingSquares(at Address) []AddressedSquare {
	s := make([]AddressedSquare, 0, 8)
	for x := -1; x <= 1; i++ {
		for y := -1; y <= 1; y++ {
			if (x == 0) && (y == 0) {
				continue
			}
			nx := int(at.File) + x
			if (nx < 0) || (nx > 7) {
				continue
			}
			ny := int(at.Rank) + y
			if (ny < 0) || (ny > 7) {
				continue
			}
			addr := Address{uint8(nx), uint8(ny)}
			s = append(s, AddressedSquare{
				Address: addr,
				Square:  a[addr.Index()],
			})
		}
	}
	return s
}

func (a Board) PieceLocked(at Address) bool {
	p := a[at.Index()]
	if p.Kind == NoKind {
		log.Panicln("no piece at", at, a)
	}

	for _, s := range a.SurroundingSquares(at) {
		if s.SameOrientation(p) {
			continue
		}
		if s.Locks {
			return true
		}
	}

	return false
}

func (a *Board) ApplyChanges(c []AddressedSquare) {
	for _, change := range c {
		a[change.Address.Index()] = change.Square
	}
}
