// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import ()

const (
	player_table = "players"

	player_name_key   = "name"
	player_crypt_key  = "crypt"
	player_wins_key   = "wins"
	player_losses_key = "losses"

	initial_piece_count = 6
)

const player_crypt_query = "SELECT " + player_crypt_key + " FROM " + player_table + " WHERE " + player_name_key + "=$1"

func playerCryptFromDatabase(name string) (bool, string) {
	row := database.QueryRow(player_crypt_query, name)
	var c string
	err := row.Scan(&c)
	if err != nil {
		return false, ""
	}
	return true, c

}

const new_player_insert = "INSERT INTO " + player_table + "(" + player_name_key + ", " + player_crypt_key + ", " + player_wins_key + ", " + player_losses_key + ") VALUES ($1, $2, $3, $4)"

func newPlayerInDatabase(name, crypt string) {
	_, err := database.Exec(new_player_insert, name, crypt, 0, 0)
	if err != nil {
		panicExit(err.Error())
	}
	newPlayerPiecesIntoDatabase(name)
}

type record struct {
	wins   int
	losses int
}

const player_record_query = "SELECT " + player_wins_key + ", " + player_losses_key + " FROM " + player_table + " WHERE " + player_name_key + "=$1"

func playerRecordFromDatabase(name string) record {
	row := database.QueryRow(player_record_query, name)
	var wins, losses int
	err := row.Scan(&wins, &losses)
	if err != nil {
		panicExit(err.Error())
	}
	return record{
		wins:   wins,
		losses: losses,
	}
}
