// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pciet/wichess/wichessing"
)

func (g game) randomMoveAtTime(at time.Time) game {
	board := wichessingBoard(g.Points)
	var active wichessing.Orientation
	if g.Active == g.White {
		active = wichessing.White
	} else {
		active = wichessing.Black
	}
	if board.Draw(active) {
		return g
	}
	m, _, checkmate := board.Moves(active)
	if checkmate {
		return g
	}
	for addr, _ := range m {
		if g.Points[addr.Index()].Orientation != active {
			delete(m, addr)
		}
	}
	piece := rand.Intn(len(m))
	i := 0
OUTER:
	for addr, moves := range m {
		if i != piece {
			i++
			continue
		}
		move := rand.Intn(len(moves))
		i = 0
		for point, _ := range moves {
			if i != move {
				i++
				continue
			}
			_, promoting := g.move(int(addr.Index()), int(point.Index()), g.Active, true)
			// TODO: this logic is duplicated in easyComputerMoveForGame
			if promoting {
				after := board.AfterMove(addr, *point, active)
				var from int
				if active == wichessing.White {
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
				_ = g.DB.gameWithIdentifier(g.ID).promote(from, g.Active, wichessing.Queen, true)
			}
			break OUTER
		}
		panicExit("game_moving: reached unreachable execution path")
	}
	g.updateGameTimesWithMove(at)
	return g.DB.gameWithIdentifier(g.ID)
}

func (g game) updateGameTimesWithMove(at time.Time) {
	var timeKey, elapsedKey string
	var elapsed time.Duration
	if g.Active == g.White {
		timeKey = games_white_latest_move
		elapsedKey = games_white_elapsed
		elapsed = g.WhiteElapsed + at.Sub(g.WhiteElapsedUpdated)
	} else {
		timeKey = games_black_latest_move
		elapsedKey = games_black_elapsed
		elapsed = g.BlackElapsed + at.Sub(g.BlackElapsedUpdated)
	}
	result, err := g.DB.Exec("UPDATE "+games_table+" SET "+timeKey+" = $1, "+elapsedKey+" = $2, "+games_white_elapsed_updated+" = $3, "+games_black_elapsed_updated+" = $4 WHERE "+games_identifier+" = $5;", at, elapsed, at, at, g.ID)
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
}

// TODO: lock database access to this game while this sort of method is executing its multiple reads and writes
func (db DB) updateGameTimes(id int, turn time.Duration) GameInfo {
	g := db.gameWithIdentifier(id)
	if turn == time.Duration(0) {
		return g.GameInfo
	}
	var active wichessing.Orientation
	if g.Active == g.White {
		active = wichessing.White
	} else {
		active = wichessing.Black
	}
	b := wichessingBoard(g.Points)
	if b.Draw(active) || b.Checkmate(active) {
		return g.GameInfo
	}
	sinceMove := g.sinceMove()
	// if a turn timer is set then make a random move for every turn duration that has occurred.
	for sinceMove > turn {
		if g.Active == g.Black {
			g = g.randomMoveAtTime(g.WhiteLatestMove.Add(turn))
		} else {
			g = g.randomMoveAtTime(g.BlackLatestMove.Add(turn))
		}
		sinceMove = g.sinceMove()
	}
	var elapsedKey, elapsedUpdatedKey string
	var elapsed time.Duration
	var elapsedUpdated time.Time
	if g.Active == g.White {
		elapsedKey = games_white_elapsed
		elapsedUpdatedKey = games_white_elapsed_updated
		elapsed = g.WhiteElapsed + time.Now().Sub(g.WhiteElapsedUpdated)
		elapsedUpdated = time.Now()
		g.WhiteElapsedUpdated = elapsedUpdated
		g.WhiteElapsed = elapsed
	} else {
		elapsedKey = games_black_elapsed
		elapsedUpdatedKey = games_black_elapsed_updated
		elapsed = g.BlackElapsed + time.Now().Sub(g.BlackElapsedUpdated)
		elapsedUpdated = time.Now()
		g.BlackElapsedUpdated = elapsedUpdated
		g.BlackElapsed = elapsed
	}
	result, err := g.DB.Exec("UPDATE "+games_table+" SET "+elapsedKey+" = $1, "+elapsedUpdatedKey+" = $2 WHERE "+games_identifier+" = $3;", elapsed, elapsedUpdated, g.ID)
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
	return g.GameInfo
}

func (g game) sinceMove() time.Duration {
	if g.Active == g.Black {
		return time.Now().Sub(g.WhiteLatestMove)
	} else {
		return time.Now().Sub(g.BlackLatestMove)
	}
}
