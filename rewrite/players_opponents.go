package main

import (
	"database/sql"

	"github.com/lib/pq"
)

const RecentOpponentCount = 5

func PlayerRecentOpponents(tx *sql.Tx, playerID int) [RecentOpponentCount]string {
	var values []sql.NullInt64
	err := tx.QueryRow(PlayersRecentOpponentsQuery, playerID).Scan(pq.Array(&values))
	if err != nil {
		DebugPrintln(PlayersRecentOpponentsQuery, playerID)
		Panic(err)
	}

	var out [RecentOpponentCount]string
	for i, id := range values {
		if id.Valid == false {
			continue
		}
		// TODO: one SQL query for names?
		out[i] = PlayerName(tx, int(id.Int64))
	}

	return out
}
