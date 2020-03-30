package main

import "database/sql"

func DeleteGame(tx *sql.Tx, id GameIdentifier) {
	r, err := tx.Exec(GamesDelete, id)
	if err != nil {
		Panic(err)
	}

	c, err := r.RowsAffected()
	if err != nil {
		Panic(err)
	}
	if c != 1 {
		Panic(c, "rows affected instead of 1 by game delete of", id)
	}
}
