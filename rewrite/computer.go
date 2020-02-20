package main

import (
	"database/sql"
)

const computer_player_name = "Computer Player"

func ComputerGameID(tx *sql.Tx, player string) GameIdentifier {
	var id GameIdentifier
	err := tx.QueryRow(computer_game_id_query, player).Scan(&id)
	if err != nil {
		DebugPrintln("tx.QueryRow failed:", err)
		return 0
	}
	return id
}
