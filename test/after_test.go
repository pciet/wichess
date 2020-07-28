package test

import (
	"testing"

	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

// TestAfterMove tests the rules.Board.DoMove method by comparing the result of a move to the
// expected position. Cases are made with test/builder which saves them into test/cases as JSON.
func TestAfterMove(t *testing.T) {
	for _, tc := range LoadAllAfterMoveCases() {
		var board rules.Board
		for _, piece := range tc.Position {
			board[piece.Address.Index()] = rules.Square(rules.Piece{
				Kind:        piece.Kind,
				Orientation: piece.Orientation,
				Moved:       piece.Moved,
				Start:       piece.Start,
			}.ApplyCharacteristics())
		}

		changes, _ := board.DoMove(tc.Move)

		// assuming any changes are now moved pieces
		for i := 0; i < len(tc.Changes); i++ {
			if tc.Changes[i].Kind == piece.NoKind {
				continue
			}
			tc.Changes[i].Moved = true
		}

		if rules.AddressedSquaresEquivalent(changes, tc.Changes) == false {
			t.Fatal(tc.Name, ":", "expected changes", tc.Changes, "got", changes)
		}
	}
}
