package test

import (
	"testing"

	"github.com/pciet/wichess/rules"
)

// TODO: pieces start address should be included in this test

// TestMoves tests the rules.Game.Moves method that determines which moves are available for
// the active player in a position. Test cases are defined by test/builder and saved into
// test/cases as JSON to be loaded by this test.
func TestMoves(t *testing.T) {
	for _, tc := range loadAllMovesCases() {
		var board rules.Board
		for _, p := range tc.Position {
			board[p.Address.Index()] = rules.NewPiece(p.Kind,
				p.Orientation, p.Moved, rules.NoAddress)
		}

		moves, state := board.Moves(tc.Active, tc.PreviousMove)
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
					if hasAddress(moveset.Moves, addr) == false {
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

func hasAddress(in []rules.Address, has rules.Address) bool {
	for _, a := range in {
		if a != has {
			continue
		}
		return true
	}
	return false
}
