package main

import "database/sql"

func NewPlayer(tx *sql.Tx, name, crypt string) {
	result, err := tx.Exec(PlayerNewInsert,
		name, crypt)
	if err != nil {
		Panic(err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		Panic(err)
	}
	if count != 1 {
		Panic(count, "rows affected by new player insert for", name)
	}
}
