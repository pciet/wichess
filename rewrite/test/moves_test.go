package test

import (
	"testing"

	"github.com/pciet/wichess/rules"
)

type MovesTestCase struct {
	// name of the test case
	Case string

	// active player the moves will be calculated for
	By Orientation

	// previous move used by rules.Game
	Prev Move

	// expected game state returned by rules.Game.Moves
	State

	// position, the pieces on the board
	Pos []AddressedSquare

	// expected moves returned by rules.Game.Moves
	Moves []MoveSet
}

var MovesTestCases = make([]MovesTestCase, 0, 8)

// TestMoves tests the rules.Game.Moves method that determines
// which moves are available for the active player in a position.
// Test cases are MovesTestCase Go variables that are added to
// MovesTestCases in package test init functions.
func TestMoves(t *testing.T) {
	if len(MovesTestCases) == 0 {
		t.Fatal("no test cases")
	}

	for _, tc := range MovesTestCases {
		var b Board
		for _, s := range tc.Pos {
			s.Square = Square(Piece(s.Square).ApplyCharacteristics())
			b[s.Address.Index()] = s.Square
		}
		g := rules.MakeGame(b,
			tc.Prev.From.Index(), tc.Prev.To.Index())

		moves, state := g.Moves(tc.By)
		if state != tc.State {
			t.Fatal(tc.Case, ":",
				"expected state", tc.State, "got", state)
		}
		found := 0
	LOOP:
		for _, moveset := range tc.Moves {
			for _, calcMoveset := range moves {
				if calcMoveset.From != moveset.From {
					continue
				}
				if len(calcMoveset.Moves) != len(moveset.Moves) {
					t.Fatal(tc.Case, ":",
						calcMoveset.From, "moves mismatch,",
						"expected", len(moveset.Moves), "moves",
						moveset.Moves, "got",
						len(calcMoveset.Moves), "moves:", calcMoveset.Moves)
				}
				for _, addr := range calcMoveset.Moves {
					if HasAddress(moveset.Moves, addr) == false {
						t.Fatal(tc.Case, ":", "calculated to have", addr,
							"at", moveset.From, ":", moveset.Moves)
					}
				}
				moves = rules.RemoveMoveSet(moves, moveset.From)
				found++
				continue LOOP
			}
			t.Fatal(tc.Case, ":",
				"expected move set", moveset,
				"not in calculated moves")
		}
		if len(moves) != 0 {
			t.Fatal(tc.Case, ":",
				"rules.Game.Moves returned extra moves", moves)
		}
	}
}

func HasAddress(in []Address, has Address) bool {
	for _, a := range in {
		if a != has {
			continue
		}
		return true
	}
	return false
}
