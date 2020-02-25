package main

import (
	"database/sql"
	"log"
)

const computer_player_name = "Computer Player"

func ComputerGameID(tx *sql.Tx, player string) GameIdentifier {
	var id GameIdentifier
	err := tx.QueryRow(computer_game_id_query, player).Scan(&id)
	if err == sql.ErrNoRows {
		return 0
	} else if err != nil {
		log.Panicln("tx.QueryRow failed:", err)
	}
	return id
}
