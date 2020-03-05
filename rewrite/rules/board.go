package rules

import (
	"log"
	"strconv"
	"strings"
)

// A Wisconsin Chess board is a regular 8x8 chess board.
type Board [8 * 8]Square

func (a Board) SurroundingSquares(at Address) []AddressedSquare {
	s := make([]AddressedSquare, 0, 8)
	for x := -1; x <= 1; x++ {
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

func (a Board) KingLocation(of Orientation) Address {
	for i, s := range a {
		if (s.Kind == King) && (s.Orientation == of) {
			return AddressIndex(i).Address()
		}
	}
	log.Panicln("no king found for", of)
	return Address{}
}

func (a Board) NoKing(of Orientation) bool {
	for _, s := range a {
		if (s.Kind == King) && (s.Orientation == of) {
			return false
		}
	}
	return true
}

func (a *Board) ApplyChanges(c []AddressedSquare) {
	for _, change := range c {
		a[change.Address.Index()] = change.Square
	}
}

func (a Board) String() string {
	var s strings.Builder
	for rank := 7; rank >= 0; rank-- {
		s.WriteString(strconv.Itoa(rank + 1))
		s.WriteRune(' ')
		for file := 0; file < 8; file++ {
			s.WriteRune('[')
			s.WriteString(a[Address{uint8(file), uint8(rank)}.Index()].String())
			s.WriteString("] ")
		}
		s.WriteString("\n")
	}
	return s.String()
}
