package main

import (
	"database/sql"

	"github.com/pciet/wichess/rules"
)

func GamePlayersOrientation(tx *sql.Tx, id GameIdentifier, name string) rules.Orientation {
	white, black := GamePlayers(tx, id)
	if name == white {
		return rules.White
	} else if name != black {
		Panic("player", name, "not in game", id, "with", white, black)
	}
	return rules.Black
}

// GamesActiveAndOpponentName queries the database to show if this player is the active player
// and get the opponent name.
// An empty string is returned if the game doesn't exist.
// If the game is conceded then the player is always indicated as active.
func GameActiveAndOpponentName(tx *sql.Tx, id GameIdentifier, player string) (bool, string) {
	var conceded bool
	var active, white, black string
	err := tx.QueryRow(GamesOpponentAndActiveQuery, id).Scan(
		&active,
		&white,
		&black,
		&conceded,
	)
	if err == sql.ErrNoRows {
		DebugPrintln("no rows found for id", id, "and player", player)
		return false, ""
	} else if err != nil {
		Panic(err)
	}

	var opponent string
	if player == white {
		opponent = black
	} else if player == black {
		opponent = white
	} else {
		Panic("player", player, "doesn't match white", white, "or black", black)
	}

	if (active == player) || conceded {
		return true, opponent
	}

	return false, opponent
}

func GameHasPlayer(tx *sql.Tx, id GameIdentifier, name string) bool {
	var s sql.NullString
	err := tx.QueryRow(GamesHasPlayerQuery, id, name).Scan(&s)
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		Panic(err)
	}
	return true
}

func GameOpponent(tx *sql.Tx, id GameIdentifier, of string) string {
	var white, black string
	err := tx.QueryRow(GamesOpponentQuery, id).Scan(&white, &black)
	if err == sql.ErrNoRows {
		return ""
	} else if err != nil {
		Panic(err)
	}
	if of == white {
		return black
	} else if of == black {
		return white
	}
	return ""
}

// GamePlayers returns the white and black player names in that order.
func GamePlayers(tx *sql.Tx, id GameIdentifier) (string, string) {
	var white, black string
	err := tx.QueryRow(GamesPlayersQuery, id).Scan(&white, &black)
	if err == sql.ErrNoRows {
		return "", ""
	} else if err != nil {
		Panic(err)
	}
	return white, black
}

func GamePreviousActive(tx *sql.Tx, id GameIdentifier) string {
	var pa bool
	err := tx.QueryRow(GamesPreviousActiveQuery, id).Scan(&pa)
	if err == sql.ErrNoRows {
		return ""
	} else if err != nil {
		Panic(err)
	}
	white, black := GamePlayers(tx, id)
	if pa == false {
		return white
	}
	return black
}
