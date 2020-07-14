package main

import "database/sql"

func MarkGameConceded(tx *sql.Tx, id GameIdentifier) {
	SQLExecRow(tx, GamesConcededUpdate, true, id)
}
