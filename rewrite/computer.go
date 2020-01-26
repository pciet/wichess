package main

import (
	"database/sql"
)

const (
	computer_player_name = "Computer Player"

	computer_game_id_query = "SELECT " + games_identifier + " FROM " + games_table + " WHERE " + games_white + "=$1 AND " + games_black + "=" + computer_player_name + ";"
)

func ComputerGameID(tx *sql.Tx, player string) GameID {
	var id GameID
	err := tx.QueryRow(computer_game_id_query, player).Scan(&id)
	if err != nil {
		DebugPrintln("tx.QueryRow failed:", err)
		return 0
	}
	return id
}
