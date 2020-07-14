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
	SQLExecRow(tx, PlayersPeopleGameUpdate, gid, id)
}
