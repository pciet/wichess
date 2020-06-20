package main

import (
	"database/sql"

	"github.com/lib/pq"
	"github.com/pciet/wichess/rules"
)

// NewGame inserts a row in the games table with the requested armies. If an army request is
// invalid then no effects occur and a 0 is returned. See ReserveArmies for the criteria of a
// valid army request.
func NewGame(tx *sql.Tx, wa, ba ArmyRequest, white, black Player) GameIdentifier {
	wp, wpicks, bp, bpicks, err := ReserveArmies(tx, wa, ba, white.ID, black.ID)
	if err != nil {
		DebugPrintln("NewGame called ReserveArmies:", white, wa, "vs", black, ba, ":", err)
		return 0
	}

	emptyCaptures := pq.Array(make([]rules.PieceKind, 15))

	// QueryRow instead of Exec: https://github.com/lib/pq/issues/24
	var id GameIdentifier
	err = tx.QueryRow(GamesNewInsert,
		false,
		white.Name, false, wpicks.Left, wpicks.Right, emptyCaptures,
		black.Name, false, bpicks.Left, bpicks.Right, emptyCaptures,
		rules.White, rules.Black,
		NoMove, NoMove,
		0, 1,
		pq.Array([]EncodedPiece{
			wp[8], wp[9], wp[10], wp[11], wp[12], wp[13], wp[14], wp[15],
			wp[0], wp[1], wp[2], wp[3], wp[4], wp[5], wp[6], wp[7],
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			bp[0], bp[1], bp[2], bp[3], bp[4], bp[5], bp[6], bp[7],
			bp[8], bp[9], bp[10], bp[11], bp[12], bp[13], bp[14], bp[15],
		}),
	).Scan(&id)
	if err != nil {
		DebugPrintln(GamesNewInsert)
		Panic("failed to insert new game:", err)
	}

	return id
}
