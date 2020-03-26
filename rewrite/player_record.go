package main

import "database/sql"

type PlayerRecord struct {
	Wins   int
	Losses int
	Draws  int
}

func LoadPlayerRecord(tx *sql.Tx, name string) PlayerRecord {
	var r PlayerRecord
	err := tx.QueryRow(PlayerRecordQuery, name).Scan(
		&r.Wins,
		&r.Losses,
		&r.Draws,
	)
	if err != nil {
		Panic(err)
	}
	return r
}
