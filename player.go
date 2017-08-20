// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"fmt"

	"github.com/pciet/wichess/wichessing"
)

const (
	database_player_table = "players"

	database_player_table_name_key   = "name"
	database_player_table_crypt_key  = "crypt"
	database_player_table_wins_key   = "wins"
	database_player_table_losses_key = "losses"

	initial_piece_count = 6
)

func playerCryptFromDatabase(name string) (bool, string) {
	rows, err := database.Query(fmt.Sprintf("SELECT %v FROM %v WHERE %v=$1", database_player_table_crypt_key, database_player_table, database_player_table_name_key), name)
	if err != nil {
		panicExit(err.Error())
		return false, ""
	}
	defer rows.Close()
	exists := rows.Next()
	if exists == false {
		return false, ""
	}
	var c string
	err = rows.Scan(&c)
	if err != nil {
		panicExit(err.Error())
		return false, ""
	}
	if rows.Next() {
		panicExit(fmt.Sprintf("duplicate database entries for %v", name))
		return false, ""
	}
	return true, c

}

func newPlayerInDatabase(name, crypt string) {
	_, err := database.Exec(fmt.Sprintf("INSERT INTO %v (%v, %v, %v, %v) VALUES ($1, $2, $3, $4)", database_player_table, database_player_table_name_key, database_player_table_crypt_key, database_player_table_wins_key, database_player_table_losses_key), name, crypt, 0, 0)
	if err != nil {
		panicExit(err.Error())
		return
	}
	newPlayerPiecesIntoDatabase(name)
	return
}

type record struct {
	wins   int
	losses int
}

func playerRecordFromDatabase(name string) record {
	rows, err := database.Query(fmt.Sprintf("SELECT %v, %v FROM %v WHERE %v=$1", database_player_table_wins_key, database_player_table_losses_key, database_player_table, database_player_table_name_key), name)
	if err != nil {
		panicExit(err.Error())
		return record{}
	}
	defer rows.Close()
	exists := rows.Next()
	if exists == false {
		panicExit(fmt.Sprintf("player %v not in database", name))
		return record{}
	}
	var wins, losses int
	err = rows.Scan(&wins, &losses)
	if err != nil {
		panicExit(err.Error())
		return record{}
	}
	if rows.Next() {
		panicExit(fmt.Sprintf("duplicate database entries for %v", name))
		return record{}
	}
	return record{
		wins:   wins,
		losses: losses,
	}
}

func newPlayerPiecesIntoDatabase(name string) {
	for i := 0; i < initial_piece_count; i++ {
		piece := wichessing.RandomPiece()
		_, err := database.Exec(fmt.Sprintf("INSERT INTO %v (%v, %v, %v, %v) VALUES ($1, $2, $3, $4)", database_piece_table, database_piece_table_kind_key, database_piece_table_owner_key, database_piece_table_takes_key, database_piece_table_ingame_key), piece.Kind, name, 0, false)
		if err != nil {
			panicExit(err.Error())
			return
		}
	}
}
