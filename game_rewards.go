package wichess

import (
	"database/sql"

	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

// GamePlayersRewards returns the left, right, and reward piece kinds in that order.
func GamePlayersRewards(tx *sql.Tx,
	id GameIdentifier, player rules.Orientation) (piece.Kind, piece.Kind, piece.Kind) {

	var query string
	if player == rules.White {
		query = GamesWhiteRewardsQuery
	} else if player == rules.Black {
		query = GamesBlackRewardsQuery
	} else {
		Panic("orientation", player, "not white or black")
	}

	var left, right, reward piece.Kind
	err := tx.QueryRow(query, id).Scan(&left, &right, &reward)
	if err != nil {
		DebugPrintln(query)
		Panic(err)
	}

	return left, right, reward
}

func PicksInGameForPlayer(tx *sql.Tx,
	gameID GameIdentifier, name string) (piece.Kind, piece.Kind) {

	white, black := GamePlayers(tx, gameID)

	var query string
	if white == name {
		query = GamesWhitePicksQuery
	} else if black == name {
		query = GamesBlackPicksQuery
	} else {
		Panic(name, "not in game", gameID)
	}
	var left, right piece.Kind
	err := tx.QueryRow(query, gameID).Scan(&left, &right)
	if err != nil {
		DebugPrintln(query)
		DebugPrintln(gameID)
		Panic(err)
	}
	return left, right
}
