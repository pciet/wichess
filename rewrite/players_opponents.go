package main

import (
	"database/sql"

	"github.com/lib/pq"
)

const RecentOpponentCount = 5

func PlayerRecentOpponents(tx *sql.Tx, id PlayerIdentifier) [RecentOpponentCount]string {

	ids := PlayerRecentOpponentIDs(tx, id)

	var out [RecentOpponentCount]string
	for i, oid := range ids {
		if oid == 0 {
			out[i] = ""
			continue
		}
		out[i] = PlayerName(tx, oid)
	}

	return out
}

func PlayerRecentOpponentIDs(tx *sql.Tx,
	id PlayerIdentifier) [RecentOpponentCount]PlayerIdentifier {

	var values []sql.NullInt64
	err := tx.QueryRow(PlayersRecentOpponentsQuery, id).Scan(pq.Array(&values))
	if err != nil {
		DebugPrintln(PlayersRecentOpponentsQuery, id)
		Panic(err)
	}

	var out [RecentOpponentCount]PlayerIdentifier
	for i, oid := range values {
		if oid.Valid == false {
			continue
		}
		out[i] = PlayerIdentifier(oid.Int64)
	}

	return out
}

func UpdatePlayerRecentOpponents(tx *sql.Tx,
	id PlayerIdentifier, opponents [RecentOpponentCount]PlayerIdentifier) {

	r, err := tx.Exec(PlayersRecentOpponentsUpdate, pq.Array(opponents), id)
	if err != nil {
		Panic(err)
	}
	c, err := r.RowsAffected()
	if err != nil {
		Panic(err)
	}
	if c != 1 {
		Panic(c, "rows affected by", PlayersRecentOpponentsUpdate, opponents, id)
	}
}

// AddPlayerRecentOpponent updates the list of recent opponents by inserting this one at the top.
// The list is updated to remove duplicates, and opponents past the bottom of the list are lost.
func AddPlayerRecentOpponent(player, opponent PlayerIdentifier) {
	tx := DatabaseTransaction()

	// TODO: does this need to be a FOR UPDATE query?
	rec := PlayerRecentOpponentIDs(tx, player)

	// remove possible one duplicate of this opponent then condense the list
	for i, opp := range rec {
		if opp != opponent {
			continue
		}
		for j := i; j < RecentOpponentCount; j++ {
			if j == (RecentOpponentCount - 1) {
				rec[j] = 0
				break
			}
			rec[j] = rec[j+1]
		}
		break
	}

	// insert the opponent at the start of the list
	for i := (RecentOpponentCount - 1); i > 0; i-- {
		rec[i] = rec[i-1]
	}
	rec[0] = opponent

	UpdatePlayerRecentOpponents(tx, player, rec)

	tx.Commit()
}
