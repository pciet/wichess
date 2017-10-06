// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"fmt"

	"github.com/pciet/wichess/wichessing"
)

const (
	check_key     = "check"
	checkmate_key = "checkmate"
	promote_key   = "promote"
	draw_key      = "draw"
)

// Returns address that have changed, Kind 0 piece for a now empty point, but does not update the board points. Returns true if promoting.
func (g game) move(from, to int, mover string) (map[string]piece, bool) {
	var nextMover string
	var orientation, nextOrientation wichessing.Orientation
	if g.White == g.Active {
		if g.White != mover {
			return nil, false
		}
		orientation = wichessing.White
		nextOrientation = wichessing.Black
		nextMover = g.Black
	} else {
		if g.Black != mover {
			return nil, false
		}
		orientation = wichessing.Black
		nextOrientation = wichessing.White
		nextMover = g.White
	}
	b := wichessingBoard(g.Points)
	if b.HasPawnToPromote() {
		return nil, false
	}
	if b.Draw(orientation) {
		return nil, false
	}
	diff := make(map[string]piece)
	for point, _ := range b.Move(absPoint(from), absPoint(to), orientation) {
		if point.Piece == nil {
			diff[point.AbsPoint.String()] = piece{
				Piece: wichessing.Piece{
					Kind: 0,
				},
			}
		} else {
			diff[point.AbsPoint.String()] = piece{
				Piece:      *point.Piece,
				Identifier: g.Points[from].Identifier, // TODO: this could be incorrect
			}
		}
	}
	if len(diff) == 0 {
		return diff, false
	}
	after := b.AfterMove(absPoint(from), absPoint(to), orientation)
	var promoting bool
	if after.HasPawnToPromote() {
		g.DB.updateGame(g.ID, diff, mover)
		promoting = true
	} else {
		g.DB.updateGame(g.ID, diff, nextMover)
	}
	if (mover != easy_computer_player) && (mover != hard_computer_player) {
		if orientation == wichessing.White {
			if gameListening[g.ID].black != nil {
				gameListening[g.ID].black <- diff
			}
		} else {
			if gameListening[g.ID].white != nil {
				gameListening[g.ID].white <- diff
			}
		}
	}
	if after.Checkmate(nextOrientation) {
		if nextOrientation == wichessing.White {
			g.DB.updatePlayerRecords(g.Black, g.White, false)
		} else {
			g.DB.updatePlayerRecords(g.White, g.Black, false)
		}
	} else if after.Draw(nextOrientation) {
		g.DB.updatePlayerRecords(g.White, g.Black, true)
	}
	return diff, promoting
}

func (g game) promote(from int, player string, kind wichessing.Kind) map[string]piece {
	var nextMover string
	var orientation wichessing.Orientation
	if g.White == player {
		nextMover = g.Black
		orientation = wichessing.White
		if (from < 56) || (from > 63) {
			return nil
		}
	} else {
		nextMover = g.White
		orientation = wichessing.Black
		if (from < 0) || (from > 7) {
			return nil
		}
	}
	point := g.Points[from]
	if point.Kind == 0 {
		return nil
	}
	if (point.Orientation != orientation) || (point.Kind != wichessing.Pawn) {
		return nil
	}
	b := wichessingBoard(g.Points)
	if b.Draw(orientation) {
		return nil
	}
	diff := make(map[string]piece)
	for point, _ := range b.PromotePawn(wichessing.AbsPointFromIndex(uint8(from)), wichessing.Kind(kind)) {
		if point.Piece == nil {
			diff[point.AbsPoint.String()] = piece{
				Piece: wichessing.Piece{
					Kind: 0,
				},
			}
		} else {
			diff[point.AbsPoint.String()] = piece{
				Piece: *point.Piece,
			}
		}
	}
	if (diff == nil) || (len(diff) == 0) {
		return diff
	}
	g.DB.updateGame(g.ID, diff, nextMover)
	if (player != easy_computer_player) && (player != hard_computer_player) {
		if orientation == wichessing.White {
			if gameListening[g.ID].black != nil {
				gameListening[g.ID].black <- diff
			}
		} else {
			if gameListening[g.ID].white != nil {
				gameListening[g.ID].white <- diff
			}
		}
	}
	var checkOrientation wichessing.Orientation
	if orientation == wichessing.White {
		checkOrientation = wichessing.Black
	} else {
		checkOrientation = wichessing.White
	}
	after := b.AfterPromote(absPoint(from), wichessing.Kind(kind))
	if after.Checkmate(checkOrientation) {
		if checkOrientation == wichessing.White {
			g.DB.updatePlayerRecords(g.Black, g.White, false)
		} else {
			g.DB.updatePlayerRecords(g.White, g.Black, false)
		}
	} else if after.Draw(checkOrientation) {
		g.DB.updatePlayerRecords(g.White, g.Black, true)
	}
	return diff
}

// The map keys are wichessing.AbsPoint converted to "x/file-y/rank" formatted string.
// If the game is in a check or checkmate state, or a piece is to be promoted, then a corresponding key with a nil value will be set.
func (g game) moves() map[string]map[string]struct{} {
	var board wichessing.Board
	for i := 0; i < 64; i++ {
		var p *wichessing.Piece
		if g.Points[i].Piece.Kind == 0 {
			p = nil
		} else {
			p = &g.Points[i].Piece
		}
		board[i] = wichessing.Point{
			Piece: p,
			AbsPoint: wichessing.AbsPoint{
				File: wichessing.FileFromIndex(uint8(i)),
				Rank: wichessing.RankFromIndex(uint8(i)),
			},
		}
	}
	moves := make(map[string]map[string]struct{})
	if board.HasPawnToPromote() {
		moves[promote_key] = nil
		return moves
	}
	var active wichessing.Orientation
	if g.Active == g.White {
		active = wichessing.White
	} else {
		active = wichessing.Black
	}
	if board.Draw(active) {
		moves[draw_key] = nil
		return moves
	}
	m, check, checkmate := board.Moves(active)
	for point, set := range m {
		moves[point.String()] = set.String()
	}
	if checkmate {
		moves[checkmate_key] = nil
	} else if check {
		moves[check_key] = nil
	}
	return moves
}

func (g *game) acknowledge(player string) bool {
	var active wichessing.Orientation
	if g.Active == g.Black {
		active = wichessing.Black
	} else {
		active = wichessing.White
	}
	b := wichessingBoard(g.Points)
	if (b.Checkmate(active) == false) && (b.Draw(active) == false) {
		return false
	}
	var ackKey string
	if player == g.Black {
		ackKey = games_black_acknowledge
		g.BlackAcknowledge = true
	} else if player == g.White {
		ackKey = games_white_acknowledge
		g.WhiteAcknowledge = true
	} else {
		panicExit("player " + player + " is not " + g.Black + " (black) or " + g.White + " (white)")
	}
	if g.BlackAcknowledge && g.WhiteAcknowledge {
		g.DB.deleteGame(g.ID)
		return true
	}
	result, err := g.DB.Exec("UPDATE "+games_table+" SET "+ackKey+" = $1 WHERE "+games_identifier+" = $2;", true, g.ID)
	if err != nil {
		panicExit(err.Error())
	}
	count, err := result.RowsAffected()
	if err != nil {
		panicExit(err.Error())
	}
	if count != 1 {
		panicExit(fmt.Sprintf("%v rows affected by ack update exec", count))
	}
	return true
}
