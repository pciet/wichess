package rules

import (
	"log"
	"strconv"
	"strings"

	"github.com/pciet/wichess/piece"
)

// A Wisconsin Chess board is a regular 8x8 chess board. The mapping of array index to board square
// is described for the AddressIndex type. If a Board value's Kind field is piece.NoKind then that
// square doesn't have a piece on it.
type Board [8 * 8]Piece

// NotEmpty and Empty are used when a Piece is representing a square on the board which could not
// have a chess piece on it.
func (square *Piece) NotEmpty() bool { return square.Kind != piece.NoKind }
func (square *Piece) Empty() bool    { return square.Kind == piece.NoKind }

func (a *Board) String() string {
	var s strings.Builder
	for rank := 7; rank >= 0; rank-- {
		s.WriteString(strconv.Itoa(rank + 1))
		s.WriteRune(' ')
		for file := 0; file < 8; file++ {
			s.WriteRune('[')
			s.WriteString(a[Address{file, rank}.Index()].String())
			s.WriteString("] ")
		}
		s.WriteString("\n")
	}
	return s.String()
}

func (a *Board) surroundingSquares(at Address) []AddressedSquare {
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
			addr := Address{nx, ny}
			s = append(s, AddressedSquare{
				Address: addr,
				Piece:   a[addr.Index()],
			})
		}
	}
	return s
}

func (a *Board) kingLocation(of Orientation) Address {
	for i, s := range a {
		if (s.Kind == piece.King) && (s.Orientation == of) {
			return AddressIndex(i).Address()
		}
	}
	log.Panicln("no king found for", of)
	return Address{}
}

func (a *Board) noKing(of Orientation) bool {
	for _, s := range a {
		if (s.Kind == piece.King) && (s.Orientation == of) {
			return false
		}
	}
	return true
}

func (a *Board) applyChanges(c []AddressedSquare) {
	for _, change := range c {
		a[change.Address.Index()] = change.Piece
	}
}

func (a *Board) pieceCount(of Orientation) int {
	c := 0
	for _, s := range a {
		if s.Kind == piece.NoKind {
			continue
		}
		if s.Orientation == of {
			c++
		}
	}
	return c
}
