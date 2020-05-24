package main

import (
	"database/sql"

	"github.com/pciet/wichess/rules"
)

func PicksInGameForPlayer(tx *sql.Tx,
	gameID GameIdentifier, name string) (rules.PieceKind, rules.PieceKind) {

	white, black := GamePlayers(tx, gameID)

	var query string
	if white == name {
		query = GamesWhitePicksQuery
	} else if black == name {
		query = GamesBlackPicksQuery
	} else {
		Panic(name, "not in game", gameID)
	}
	var left, right rules.PieceKind
	err := tx.QueryRow(query, gameID).Scan(&left, &right)
	if err != nil {
		DebugPrintln(query)
		DebugPrintln(gameID)
		Panic(err)
	}
	return left, right
}
