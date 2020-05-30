package main

import (
	"database/sql"

	"github.com/lib/pq"
)

const RecentOpponentCount = 5

func PlayerRecentOpponents(tx *sql.Tx, id PlayerIdentifier) [RecentOpponentCount]string {
	var values []sql.NullInt64
	err := tx.QueryRow(PlayersRecentOpponentsQuery, id).Scan(pq.Array(&values))
	if err != nil {
		DebugPrintln(PlayersRecentOpponentsQuery, id)
		Panic(err)
	}

	var out [RecentOpponentCount]string
	for i, oid := range values {
		if oid.Valid == false {
			continue
		}
		// TODO: one SQL query for names?
		out[i] = PlayerName(tx, PlayerIdentifier(oid.Int64))
	}

	return out
}
