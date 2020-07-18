package test

import (
	"testing"

	"github.com/pciet/wichess/rules"
)

// TestMoves tests the rules.Game.Moves method that determines which moves are available for
// the active player in a position. Test cases are defined by wichess/test/builder and saved into
// wichess/tests/cases as JSON to be loaded by this test.
func TestMoves(t *testing.T) {
	for _, tc := range LoadMovesCases() {
		var board rules.Board
		for _, piece := range tc.Position {
			board[piece.Address.Index()] = rules.Square(rules.Piece{
				Kind:        piece.Kind,
				Orientation: piece.Orientation,
				Moved:       piece.Moved,
			}.ApplyCharacteristics())
		}
		g := rules.MakeGame(board, tc.PreviousMove.From.Index(), tc.PreviousMove.To.Index())

		moves, state := g.Moves(tc.Active)
		if state != tc.State {
			t.Fatal(tc.Name, ":", "expected state", tc.State, "got", state)
		}
		found := 0
	LOOP:
		for _, moveset := range tc.Moves {
			for _, calcMoveset := range moves {
				if calcMoveset.From != moveset.From {
					continue
				}
				if len(calcMoveset.Moves) != len(moveset.Moves) {
					t.Fatal(tc.Name, ":", calcMoveset.From, "moves mismatch,", "expected",
						len(moveset.Moves), "moves", moveset.Moves, "got", len(calcMoveset.Moves),
						"moves:", calcMoveset.Moves)
				}
				for _, addr := range calcMoveset.Moves {
					if HasAddress(moveset.Moves, addr) == false {
						t.Fatal(tc.Name, ":", "calculated to have", addr, "at", moveset.From,
							":", moveset.Moves)
					}
				}
				moves = rules.RemoveMoveSet(moves, moveset.From)
				found++
				continue LOOP
			}
			t.Fatal(tc.Name, ":", "expected move set", moveset, "not in calculated moves")
		}
		if len(moves) != 0 {
			t.Fatal(tc.Name, ":", "rules.Game.Moves returned extra moves", moves)
		}
	}
}

func HasAddress(in []rules.Address, has rules.Address) bool {
	for _, a := range in {
		if a != has {
			continue
		}
		return true
	}
	return false
}
