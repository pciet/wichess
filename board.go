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
