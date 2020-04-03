package main

import (
	"database/sql"

	"github.com/pciet/wichess/rules"
)

// TODO: is there less work that can be done to just get State?

func (g Game) Moves() ([]rules.MoveSet, rules.State) {
	if g.Header.Conceded {
		return nil, rules.Conceded
	}

	/*
		if g.Header.TimeLoss() {
			return nil, rules.TimeOver
		}

		if g.Header.DrawTurnsOver() {
			return nil, rules.Draw
		}
	*/

	game := rules.MakeGame(g.Board.Board,
		rules.AddressIndex(g.Header.From), rules.AddressIndex(g.Header.To))

	active := ActiveOrientation(g.Header.Active, g.Header.White.Name, g.Header.Black.Name)

	moves, state := game.Moves(active)
	if state == rules.Promotion {
		promoter, _ := game.Board.PromotionNeeded()
		if promoter != active {
			return []rules.MoveSet{}, rules.ReversePromotion
		}
		return []rules.MoveSet{}, state
	}

	return moves, state
}

func MovesForGame(tx *sql.Tx, id GameIdentifier) ([]rules.MoveSet, rules.State) {
	return LoadGame(tx, id).Moves()
}
