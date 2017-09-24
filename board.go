// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"fmt"
)

const (
	board_table = "boards"

	board_name_key       = "player"
	board_slot_key       = "slot"
	board_identifier_key = "game_id"
)

const game_id_query = "SELECT " + board_identifier_key + " FROM " + board_table + " WHERE " + board_name_key + "=$1 AND " + board_slot_key + "=$2"

func gameIdentifierAtPlayerBoardIndexFromDatabase(name string, index int) int {
	row := database.QueryRow(game_id_query, name, index)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0
	}
	return id
}

const new_game_board_insert = "INSERT INTO " + board_table + "(" + board_name_key + ", " + board_slot_key + ", " + board_identifier_key + ") VALUES ($1, $2, $3)"

func newBoardIntoDatabase(player1 string, player1setup gameSetup, player2 string, player2setup gameSetup) {
	id := newGameIntoDatabase(player1, player1setup, player2, player2setup)
	_, err := database.Exec(new_game_board_insert, player1, player1setup.slot, id)
	if err != nil {
		panicExit(err.Error())
	}
	_, err = database.Exec(new_game_board_insert, player2, player2setup.slot, id)
	if err != nil {
		panicExit(err.Error())
	}
}

type boardInfo struct {
	Slot     int
	GameID   int
	Opponent string
}

const select_board_slot_and_id_query = "SELECT " + board_slot_key + ", " + board_identifier_key + " FROM " + board_table + " WHERE " + board_name_key + "=$1"

// Can return metadata on up to 64 boards.
func playerBoardInfo(name string) []boardInfo {
	// pending matches without an opponent
	boards := pendingMatchesFor(name)
	// active games
	rows, err := database.Query(select_board_slot_and_id_query, name)
	if err != nil {
		panicExit(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		bi := boardInfo{}
		err = rows.Scan(&bi.Slot, &bi.GameID)
		if err != nil {
			panicExit(err.Error())
		}
		bi.Opponent = opponentFor(name, bi.GameID)
		boards = append(boards, bi)
	}
	err = rows.Err()
	if err != nil {
		panicExit(err.Error())
	}
	return boards
}

func deleteBoardFromDatabase(player string, gameID int) {
	result, err := database.Exec("DELETE FROM "+board_table+" WHERE "+board_name_key+" = $1 AND "+board_identifier_key+" = $2;", player, gameID)
	if err != nil {
		panicExit(err.Error())
	}
	count, err := result.RowsAffected()
	if err != nil {
		panicExit(err.Error())
	}
	if count != 1 {
		panicExit(fmt.Sprintf("%v rows affected by delete board exec for %v %v", count, player, gameID))
	}
}
