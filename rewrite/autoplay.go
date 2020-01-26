package main

import (
	"database/sql"
	"fmt"
)

// Uses an "easy" game algorithm to choose and do a move for a player.
// Returns a copy of the squares changed in the database.
// A nil return means the move wasn't possible.
func EasyAutoplay(tx *sql.Tx, n gameID, playerName string) []AddressedSquare {
	g := GameWithID(tx, n, true)
	if g.ID != 0 {
		DebugPrintf("game ID %d doesn't exist in database\n", n)
		return nil
	}

	o, ok := g.OrientationOf(playerName)
	if ok == false {
		DebugPrintf("name (%s) doesn't match white (%s) or black (%s) for game id %d\n", playerName, g.White, g.Black, g.ID)
		return nil
	}

	d, prom := g.DoMove(tx, g.Game.EasyMove(o), o)

	if prom {
		d = g.DoPromote(tx, d, o, rules.Queen)
	}

	return d
}
