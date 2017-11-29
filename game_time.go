// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/pciet/wichess/wichessing"
)

func (g game) orientationsElapsedTime(active wichessing.Orientation) time.Duration {
	if active == wichessing.White {
		return g.WhiteElapsed
	} else if active == wichessing.Black {
		return g.BlackElapsed
	}
	panicExit(fmt.Sprintf("unexpected orientation %v", active))
	return time.Duration(0)
}

var timeLossLock = sync.Mutex{}

func (g game) timeLoss(active wichessing.Orientation, total time.Duration) bool {
	if total == time.Duration(0) {
		return false
	}
	if active == wichessing.White {
		if g.WhiteElapsed > total {
			timeLossLock.Lock()
			if g.DB.gameRecorded(g.ID) == false {
				g.DB.updatePlayerRecords(g.Black, g.White, false)
				g.DB.setGameRecorded(g.ID)
			}
			timeLossLock.Unlock()
			return true
		}
	} else if active == wichessing.Black {
		if g.BlackElapsed > total {
			timeLossLock.Lock()
			if g.DB.gameRecorded(g.ID) == false {
				g.DB.updatePlayerRecords(g.White, g.Black, false)
				g.DB.setGameRecorded(g.ID)
			}
			timeLossLock.Unlock()
			return true
		}
	} else {
		panicExit(fmt.Sprintf("unexpected orientation %v", active))
	}
	return false
}

func (g *game) updateGameTimesWithMove(at time.Time) {
	var timeKey, elapsedKey string
	var elapsed time.Duration
	if g.Active == g.White {
		timeKey = games_white_latest_move
		elapsedKey = games_white_elapsed
		elapsed = g.WhiteElapsed + at.Sub(g.WhiteElapsedUpdated)
		g.WhiteLatestMove = at
		g.WhiteElapsed = elapsed
		g.WhiteElapsedUpdated = at
	} else {
		timeKey = games_black_latest_move
		elapsedKey = games_black_elapsed
		elapsed = g.BlackElapsed + at.Sub(g.BlackElapsedUpdated)
		g.BlackLatestMove = at
		g.BlackElapsed = elapsed
		g.BlackElapsedUpdated = at
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

func (db DB) updateGameTimes(id int, total time.Duration, activePlayer string) GameInfo {
	// there is a case where the game listening goroutine can determine a time loss but before the lock can be acquired a move is made - if the move signal cannot be sent in time to reset that routine then the game is considered a loss for the original player even though the database shows the opponent as the active player
	var active wichessing.Orientation
	g := db.gameWithIdentifier(id)
	if activePlayer == "" {
		if g.Active == g.White {
			active = wichessing.White
		} else {
			active = wichessing.Black
		}
	} else {
		if activePlayer == g.White {
			active = wichessing.White
		} else if activePlayer == g.Black {
			active = wichessing.Black
		} else {
			panicExit(fmt.Sprintf("unknown player %v", activePlayer))
		}
	}
	var opponent wichessing.Orientation
	if active == wichessing.White {
		opponent = wichessing.Black
	} else {
		opponent = wichessing.White
	}
	// if the database contains a timeLoss state then the game is already over and recorded by the listening routine or page load logic
	if g.timeLoss(active, total) || g.timeLoss(opponent, total) {
		return g.GameInfo
	}
	var elapsedKey, elapsedUpdatedKey string
	var elapsed time.Duration
	var elapsedUpdated time.Time
	if active == wichessing.White {
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
