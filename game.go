// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"bytes"
	"fmt"

	"github.com/pciet/wichess/wichessing"
)

const (
	database_games_table = "games"

	database_games_table_white      = "white"
	database_games_table_black      = "black"
	database_games_table_identifier = "game_id"
)

func opponentFor(name string, gameID int) string {
	rows, err := database.Query(fmt.Sprintf("SELECT %v, %v FROM %v WHERE %v=$1", database_games_table_white, database_games_table_black, database_games_table, database_games_table_identifier), gameID)
	if err != nil {
		panicExit(err.Error())
		return ""
	}
	defer rows.Close()
	exists := rows.Next()
	if exists == false {
		panicExit(fmt.Sprintf("no game with id %v", gameID))
		return ""
	}
	var white, black string
	err = rows.Scan(&white, &black)
	if err != nil {
		panicExit(err.Error())
		return ""
	}
	if rows.Next() {
		panicExit(fmt.Sprintf("duplicate database entries for gameID %v", gameID))
		return ""
	}
	if name == white {
		return black
	} else if name == black {
		return white
	} else {
		panicExit(fmt.Sprintf("gameID %v has no player %v", gameID, name))
		return ""
	}
}

// Returns the game identifier.
func newGameIntoDatabase(player1 string, player1setup gameSetup, player2 string, player2setup gameSetup) int {
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf("INSERT INTO %v (%v, %v", database_games_table, database_games_table_white, database_games_table_black))
	for i := 0; i < 64; i++ {
		query.WriteString(fmt.Sprintf(", s%v", i))
	}
	query.WriteString(") VALUES ($1, $2")
	for i := 0; i < 64; i++ {
		query.WriteString(fmt.Sprintf(", $%v", i+3))
	}
	query.WriteString(fmt.Sprintf(") RETURNING %v", database_games_table_identifier))
	// https://github.com/lib/pq/issues/24
	var id int
	err := database.QueryRow(query.String(), player1, player2, player1setup.leftRookID, player1setup.leftKnightID, player1setup.leftBishopID, wichessing.Queen, wichessing.King, player1setup.rightBishopID, player1setup.rightKnightID, player1setup.rightRookID, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, player2setup.leftRookID, player2setup.leftKnightID, player2setup.leftBishopID, wichessing.King, wichessing.Queen, player2setup.rightBishopID, player2setup.rightKnightID, player2setup.rightRookID).Scan(&id)
	if err != nil {
		panicExit(err.Error())
		return 0
	}
	return id
}
