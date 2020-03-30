package main

import (
	"database/sql"

	"github.com/pciet/wichess/rules"
)

func GameComplete(tx *sql.Tx, id GameIdentifier) bool {
	// TODO: write completion state into db when first calculated instead
	_, state := LoadGame(tx, id).Moves()
	if (state == rules.Normal) ||
		(state == rules.Promotion) ||
		(state == rules.Check) {
		return false
	}
	return true
}
