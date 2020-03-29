package main

import "database/sql"

func LoadGameTurn(tx *sql.Tx, id GameIdentifier) int {
	var t int
	err := tx.QueryRow(GamesTurnQuery, id).Scan(&t)
	if err != nil {
		Panic(err)
	}
	return t
}

func GameTurnEqual(tx *sql.Tx, id GameIdentifier, turn int) bool {
	t := LoadGameTurn(tx, id)
	if t == turn {
		return true
	}
	return false
}
