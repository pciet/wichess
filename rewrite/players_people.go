package main

import "database/sql"

func PlayerActivePeopleGame(tx *sql.Tx, id PlayerIdentifier) GameIdentifier {
	var gameID sql.NullInt64
	err := tx.QueryRow(PlayersPeopleGameQuery, id).Scan(&gameID)
	if err != nil {
		DebugPrintln(PlayersPeopleGameQuery, id)
		Panic(err)
	}
	if gameID.Valid == false {
		return 0
	}
	return GameIdentifier(gameID.Int64)
}

// UpdatePlayerActivePeopleGame sets the game ID for the active people game or can indicate no
// active game with a gid of 0.
func UpdatePlayerActivePeopleGame(tx *sql.Tx, id PlayerIdentifier, gid GameIdentifier) {
	r, err := tx.Exec(PlayersPeopleGameUpdate, gid, id)
	if err != nil {
		DebugPrintln(PlayersPeopleGameUpdate, gid, id)
		Panic(err)
	}
	c, err := r.RowsAffected()
	if err != nil {
		Panic(err)
	}
	if c != 1 {
		Panic(c, "rows affected by", PlayersPeopleGameUpdate, gid, id)
	}
}
