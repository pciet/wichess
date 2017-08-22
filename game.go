// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"fmt"

	"github.com/pciet/wichess/wichessing"
)

const (
	games_table = "games"

	games_white      = "white"
	games_black      = "black"
	games_identifier = "game_id"
)

const game_players_query = "SELECT " + games_white + ", " + games_black + " FROM " + games_table + " WHERE " + games_identifier + "=$1"

func opponentFor(name string, gameID int) string {
	row := database.QueryRow(game_players_query, gameID)
	var white, black string
	err := row.Scan(&white, &black)
	if err != nil {
		panicExit(err.Error())
	}
	if name == white {
		return black
	} else if name == black {
		return white
	} else {
		panicExit(fmt.Sprintf("gameID %v has no player %v", gameID, name))
	}
	return ""
}

const new_game_insert = "INSERT INTO " + games_table + " (" + games_white + ", " + games_black + ", s0, s1, s2, s3, s4, s5, s6, s7, s8, s9, s10, s11, s12, s13, s14, s15, s16, s17, s18, s19, s20, s21, s22, s23, s24, s25, s26, s27, s28, s29, s30, s31, s32, s33, s34, s35, s36, s37, s38, s39, s40, s41, s42, s43, s44, s45, s46, s47, s48, s49, s50, s51, s52, s53, s54, s55, s56, s57, s58, s59, s60, s61, s62, s63) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66) RETURNING " + games_identifier

// Returns the game identifier.
func newGameIntoDatabase(player1 string, player1setup gameSetup, player2 string, player2setup gameSetup) int {
	// https://github.com/lib/pq/issues/24
	var id int
	err := database.QueryRow(new_game_insert, player1, player2, player1setup.leftRookID, player1setup.leftKnightID, player1setup.leftBishopID, wichessing.Queen, wichessing.King, player1setup.rightBishopID, player1setup.rightKnightID, player1setup.rightRookID, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, wichessing.Pawn, player2setup.leftRookID, player2setup.leftKnightID, player2setup.leftBishopID, wichessing.King, wichessing.Queen, player2setup.rightBishopID, player2setup.rightKnightID, player2setup.rightRookID).Scan(&id)
	if err != nil {
		panicExit(err.Error())
	}
	return id
}
