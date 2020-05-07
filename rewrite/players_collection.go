package main

import (
	"database/sql"

	"github.com/lib/pq"
)

func PlayerCollection(tx *sql.Tx, playerID int) Collection {
	var values []sql.NullInt64
	err := tx.QueryRow(PlayersCollectionQuery,
		playerID).Scan(pq.Array(&values))
	if err != nil {
		DebugPrintln(PlayersCollectionQuery, playerID)
		Panic(err)
	}

	if len(values) != CollectionCount {
		Panic(playerID, "bad collection length", len(values))
	}

	var c Collection
	for i, v := range values {
		if v.Valid == false {
			Panic(playerID, "sql null at", i)
		}
		c[i] = EncodedPiece(v.Int64).Decode()
	}
	return c
}
