package main

import "database/sql"

// GamesActiveAndOpponentName queries the database to show if
// this player is the active player and get the opponent name.
// An empty string is returned if the game doesn't exist.
// If the game is conceded then the player is always indicated as active.
func GameActiveAndOpponentName(tx *sql.Tx, id GameIdentifier,
	player string) (bool, string) {
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

func GamePreviousActive(tx *sql.Tx, id GameIdentifier) string {
	var pa string
	err := tx.QueryRow(GamesPreviousActiveQuery, id).Scan(&pa)
	if err == sql.ErrNoRows {
		return ""
	} else if err != nil {
		Panic(err)
	}
	return pa
}
