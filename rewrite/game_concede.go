package main

import "database/sql"

// TODO: tx.Exec func that has all of the panic handling: SQLExecRow(tx, GamesConcededUpdate

func MarkGameConceded(tx *sql.Tx, id GameIdentifier) {
	r, err := tx.Exec(GamesConcededUpdate, true, id)
	if err != nil {
		Panic(err)
	}
	c, err := r.RowsAffected()
	if err != nil {
		Panic(err)
	}
	if c != 1 {
		Panic(c, "rows affected by", GamesConcededUpdate)
	}
}
