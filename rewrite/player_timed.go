package main

import "database/sql"

func PlayerHasTimedGame(tx *sql.Tx, name string) bool {
	t5, t15 := PlayerTimedGameIdentifiers(tx, name)
	if (t5 != 0) || (t15 != 0) {
		return true
	}
	return false
}

func PlayerTimedGameIdentifiers(tx *sql.Tx, name string) (t5, t15 GameIdentifier) {
	err := tx.QueryRow(PlayerTimedGameQuery, name).Scan(&t5, &t15)
	if err != nil {
		Panic(err)
	}
	return
}
