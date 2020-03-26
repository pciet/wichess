package main

import (
	"database/sql"

	"github.com/pciet/wichess/rules"
)

func (g Game) Moves() ([]rules.MoveSet, rules.State) {

}

func MovesForGame(tx *sql.Tx, id GameIdentifier) ([]rules.MoveSet, rules.State) {
	// TODO: just load the rules.Game

	h := LoadGameHeader(tx, id)

	if h.Conceded {
		return nil, rules.Conceded
	}

	/*
		if h.TimeLoss() {
			return nil, rules.TimeOver
		}
	*/

	/*
		if h.DrawTurnsOver() {
			return nil, rules.Draw
		}
	*/

	// TODO: rules.Game.Moves needs to return moves for both players for display
	var o rules.Orientation
	if h.Active == h.White.Name {
		o = rules.White
	} else {
		o = rules.Black
	}

	return rules.Game{
		Board: LoadGameBoard(tx, id).Board,
		Previous: rules.Move{
			From: rules.AddressIndex(h.From).Address(),
			To:   rules.AddressIndex(h.To).Address(),
		},
	}.Moves(o)
}
