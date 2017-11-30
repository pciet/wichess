// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"fmt"

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
	tx := db.Begin()
	g := tx.gameWithIdentifier(id, true)
	var orientation wichessing.Orientation
	if easy_computer_player == g.White {
		orientation = wichessing.White
	} else {
		orientation = wichessing.Black
	}
	move := wichessingBoard(g.Points).ComputerMove(orientation, wichessing.AbsPointFromIndex(uint8(g.From)), wichessing.AbsPointFromIndex(uint8(g.To)))
	if move == nil {
		tx.Commit()
		if debug {
			fmt.Println("nil return for computer move")
		}
		return nil
	}
	// TODO: reverse promotion isn't handled (computer guard pawn moved by person playing into promotion square)
	diff, promoting, promotingOrientation := g.move(int(move.From.Index()), int(move.To.Index()), easy_computer_player, tx)
	tx.Commit()
	if promoting && (promotingOrientation == orientation) {
		after := wichessingBoard(g.Points).AfterMove(move.From, move.To, orientation, wichessing.AbsPointFromIndex(uint8(g.From)), wichessing.AbsPointFromIndex(uint8(g.To)))
		var from int
		if orientation == wichessing.White {
			for i := 56; i < 64; i++ {
				p := after[i]
				if p.Piece == nil {
					continue
				}
				if (p.Base == wichessing.Pawn) && (p.Orientation == wichessing.White) {
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
				if (p.Base == wichessing.Pawn) && (p.Orientation == wichessing.Black) {
					from = i
					break
				}
			}
		}
		tx = db.Begin()
		// TODO: g.move to update g with the database writes so this second query isn't necessary
		pdiff := tx.gameWithIdentifier(id, true).promote(from, easy_computer_player, wichessing.Queen, tx)
		tx.Commit()
		for point, piece := range pdiff {
			diff[point] = piece
		}
	}
	return diff
}

func (db DB) easyComputerGame(player string) int {
	var id int
	err := db.QueryRow("SELECT "+games_identifier+" FROM "+games_table+" WHERE ("+games_white+" = $1 AND "+games_black+" = $2) OR ("+games_white+" = $3 AND "+games_black+" = $4);", player, easy_computer_player, easy_computer_player, player).Scan(&id)
	if err != nil {
		if debug {
			fmt.Println(err.Error())
		}
		return 0
	}
	return id
}
