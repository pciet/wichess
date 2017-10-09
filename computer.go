// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"github.com/pciet/wichess/wichessing"
)

const (
	easy_computer_player = "Easy Computer Player"
	hard_computer_player = "Hard Computer Player"
)

func (db DB) hardComputerMoveForGame(id int) map[string]piece {
	panicExit("game: hard computer move not implemented")
	return nil
}

func (db DB) easyComputerMoveForGame(id int) map[string]piece {
	g := db.gameWithIdentifier(id)
	var orientation wichessing.Orientation
	if easy_computer_player == g.White {
		orientation = wichessing.White
	} else {
		orientation = wichessing.Black
	}
	move := wichessingBoard(g.Points).ComputerMove(orientation)
	if move == nil {
		return nil
	}
	diff, promoting := g.move(int(move.From.Index()), int(move.To.Index()), easy_computer_player)
	if promoting {
		after := wichessingBoard(g.Points).AfterMove(move.From, move.To, orientation)
		var from int
		if orientation == wichessing.White {
			for i := 56; i < 64; i++ {
				p := after[i]
				if p.Piece == nil {
					continue
				}
				if (p.Kind == wichessing.Pawn) && (p.Orientation == wichessing.White) {
					from = i
					break
				}
			}
		} else {
			for i := 0; i < 8; i++ {
				p := after[i]
				if p.Piece == nil {
					continue
				}
				if (p.Kind == wichessing.Pawn) && (p.Orientation == wichessing.Black) {
					from = i
					break
				}
			}
		}
		// TODO: g.move to update g with the database writes so this second query isn't necessary
		pdiff := db.gameWithIdentifier(id).promote(from, easy_computer_player, wichessing.Queen)
		for point, piece := range pdiff {
			diff[point] = piece
		}
	}
	return diff
}

func (db DB) easyComputerGame(player string) int {
	row := db.QueryRow("SELECT "+games_identifier+" FROM "+games_table+" WHERE ("+games_white+" = $1 AND "+games_black+" = $2) OR ("+games_white+" = $3 AND "+games_black+" = $4);", player, easy_computer_player, easy_computer_player, player)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0
	}
	return id
}
