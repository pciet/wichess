package game

import (
	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

// Piece is a reduced representation of a piece on the board with only information needed by the
// web interface. The full representation is rules.Piece.
type Piece struct {
	rules.Orientation `json:"o"`
	piece.Kind        `json:"k"`
}

// Square is similar to rules.Square except it has the reduced package game Piece representation.
type Square struct {
	rules.Address `json:"a"`
	Piece         `json:"p"`
}

// SquaresFromRules converts a slice of rules.Square into a slice of game.Square.
func SquaresFromRules(in []rules.Square) []Square {
	out := make([]Square, len(in))
	for i, s := range in {
		out[i] = Square{
			Address: s.Address,
			Piece: Piece{
				Kind:        s.Piece.Kind,
				Orientation: s.Piece.Orientation,
			},
		}
	}
	return out
}

// PiecesFromRules converts a slice of rules.Piece into a slice of game.Piece.
func PiecesFromRules(in []rules.Piece) []Piece {
	out := make([]Piece, len(in))
	for i, p := range in {
		out[i] = Piece{
			Kind:        p.Kind,
			Orientation: p.Orientation,
		}
	}
	return out
}

// Copied from rules.MergeReplaceSquares.
func mergeReplaceSquares(base, overwrite []Square) []Square {
LOOP:
	for _, s := range overwrite {
		// either it needs to replace or be added
		for i, bs := range base {
			if bs.Address == s.Address {
				base[i].Piece = s.Piece
				continue LOOP
			}
		}
		base = append(base, s)
	}
	return base
}
