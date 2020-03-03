package main

import (
	"database/sql"

	"github.com/pciet/wichess/rules"
)

func MovesForGame(tx *sql.Tx, id GameIdentifier) ([]rules.MoveSet, rules.State) {
	// TODO: just load the rules.Game

	h := LoadGameHeader(tx, id)

	if h.Conceded {
		return nil, rules.Conceded
	}

	if h.TimeLoss() {
		return nil, rules.TimeOver
	}

	if h.DrawTurnsOver() {
		return nil, rules.Draw
	}

	return rules.Game{
		Board: LoadGameBoard(tx, id),
		Move: rules.Move{
			From: rules.AddressIndex(h.From).Address(),
			To:   rules.AddressIndex(h.To).Address(),
		},
	}.Moves()
}
