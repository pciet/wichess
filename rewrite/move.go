package main

import (
	"database/sql"
	"fmt"

	"github.com/pciet/wichess/rules"
)

// Returns a copy of the squares changed in the database and if a promotion is needed for which orientation.
// Returns nil if a move wasn't possible.
func (a Game) DoMove(tx *sql.Tx, m rules.Move, by rules.Orientation) (SquareSet, bool, rules.Orientation) {
	if a.Active != by {
		if debug {
			fmt.Printf("%v trying to move but %v active in game ID %v\n", by, a.Active, a.ID)
		}
		return nil, false, rules.White
	}

	if a.Game.Promoting() {
		if debug {
			fmt.Printf("%v trying to move but game ID %v needs promotion\n", by, a.ID)
		}
		return nil, false, rules.White
	}

	if (a.DrawTurns >= draw_turn_count) || a.Drawn() || a.Conceded {
		if debug {
			fmt.Printf("%v trying to move but game ID %v drawn or conceded\n", by, a.ID)
		}
		return nil, false, rules.White
	}

	squares, taken := a.Game.Move(m, by)
	if len(squares) == 0 {
		if debug {
			fmt.Printf("move %v by %v for game ID %v was invalid\n", m, by, a.ID)
		}
		return nil, false, rules.White
	}

	// add the app piece identifiers to the list of changed squares when necessary
	out := make(SquareSet, len(squares))
	for i, s := range squares {
		out[i].Square = s.Square
		pid := a.PieceIdentifiers[s.Previous.BoardAddress.Index()]
		if pid == 0 {
			continue
		}
		out[i].ID = pid
	}

}
