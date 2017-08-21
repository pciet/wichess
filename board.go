// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"fmt"
)

const (
	database_board_table = "boards"

	database_board_table_name_key       = "player"
	database_board_table_index_key      = "slot"
	database_board_table_identifier_key = "game_id"
)

func gameIdentifierAtPlayerBoardIndexFromDatabase(name string, index int) int {
	rows, err := database.Query(fmt.Sprintf("SELECT %v FROM %v WHERE %v=$1 AND %v=$2", database_board_table_identifier_key, database_board_table, database_board_table_name_key, database_board_table_index_key), name, index)
	if err != nil {
		panicExit(err.Error())
		return 0
	}
	defer rows.Close()
	exists := rows.Next()
	if exists == false {
		return 0
	}
	var id int
	err = rows.Scan(&id)
	if err != nil {
		panicExit(err.Error())
		return 0
	}
	if rows.Next() {
		panicExit(fmt.Sprintf("multiple game database entries for %v at %v", name, index))
		return 0
	}
	return id
}

func newBoardIntoDatabase(player1 string, player1setup gameSetup, player2 string, player2setup gameSetup) {
	id := newGameIntoDatabase(player1, player1setup, player2, player2setup)
	_, err := database.Exec(fmt.Sprintf("INSERT INTO %v (%v, %v, %v) VALUES ($1, $2, $3)", database_board_table, database_board_table_name_key, database_board_table_index_key, database_board_table_identifier_key), player1, player1setup.slot, id)
	if err != nil {
		panicExit(err.Error())
		return
	}
	_, err = database.Exec(fmt.Sprintf("INSERT INTO %v (%v, %v, %v) VALUES ($1, $2, $3)", database_board_table, database_board_table_name_key, database_board_table_index_key, database_board_table_identifier_key), player2, player2setup.slot, id)
	if err != nil {
		panicExit(err.Error())
		return
	}
}

type boardInfo struct {
	Slot     int
	GameID   int
	Opponent string
}

func playerBoardInfo(name string) []boardInfo {
	// pending matches without an opponent
	boards := pendingMatchesFor(name)
	// active games
	rows, err := database.Query(fmt.Sprintf("SELECT %v, %v FROM %v WHERE %v=$1", database_board_table_index_key, database_board_table_identifier_key, database_board_table, database_board_table_name_key), name)
	if err != nil {
		panicExit(err.Error())
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		bi := boardInfo{}
		err = rows.Scan(&bi.Slot, &bi.GameID)
		if err != nil {
			panicExit(err.Error())
			return nil
		}
		bi.Opponent = opponentFor(name, bi.GameID)
		boards = append(boards, bi)
	}
	err = rows.Err()
	if err != nil {
		panicExit(err.Error())
		return nil
	}
	return boards
}
