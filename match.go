// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"math"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pciet/wichess/match"
)

// A zero for an ID value means use the regular piece. Index starts at left pawn and ends at right rook.
type gameSetup [16]int

// Returns the game ID or zero if the setup is invalid.
func (db DB) requestEasyComputerMatch(player string, setup gameSetup) int {
	return db.newGame(player, setup, easy_computer_player, gameSetup{}, false)
}

type competitive15Setup struct {
	gameSetup
	ready chan struct{}
}

const (
	competitive15_match_period   = 5
	competitive15_threshold      = 10
	competitive15_bad_difference = 500
)

var (
	competitive15_turn_time  = time.Duration(0)
	competitive15_total_time = time.Duration(15 * time.Minute)
)

var competitive15Matcher = match.NewMatcher(competitive15_match_period, competitive15_threshold,
	func(rating, opprating int) bool {
		if math.Abs(float64(rating)-float64(opprating)) > competitive15_bad_difference {
			return false
		} else {
			return true
		}
	},
	func(a string, am interface{}, b string, bm interface{}) {
		ameta := am.(competitive15Setup)
		bmeta := bm.(competitive15Setup)
		id := database.newGame(a, ameta.gameSetup, b, bmeta.gameSetup, true)
		database.setPlayerCompetitive15Game(a, id)
		database.setPlayerCompetitive15Game(b, id)
		ameta.ready <- struct{}{}
		bmeta.ready <- struct{}{}
	})

type competitive5Setup struct {
	gameSetup
	ready chan struct{}
}

const (
	competitive5_match_period   = 5
	competitive5_threshold      = 10
	competitive5_bad_difference = 500
)

var (
	competitive5_turn_time  = time.Duration(0)
	competitive5_total_time = time.Duration(5 * time.Minute)
)

var competitive5Matcher = match.NewMatcher(competitive5_match_period, competitive5_threshold,
	func(rating, opprating int) bool {
		if math.Abs(float64(rating)-float64(opprating)) > competitive5_bad_difference {
			return false
		} else {
			return true
		}
	},
	func(a string, am interface{}, b string, bm interface{}) {
		ameta := am.(competitive5Setup)
		bmeta := bm.(competitive5Setup)
		id := database.newGame(a, ameta.gameSetup, b, bmeta.gameSetup, true)
		database.setPlayerCompetitive5Game(a, id)
		database.setPlayerCompetitive5Game(b, id)
		ameta.ready <- struct{}{}
		bmeta.ready <- struct{}{}
	})

type competitive48Setup struct {
	gameSetup
	slot uint8
}

var competitive48Listeners = make(map[string]*websocket.Conn)
var competitive48ListenersLock = sync.Mutex{}

const (
	// in seconds
	competitive48_match_period = 5
	// how many periods to trigger before making bad matches
	competitive48_threshold = 10
	// maximum difference in rating to be considered a good match
	competitive48_bad_difference = 500
)

var (
	competitive48_turn_time  = time.Duration(48 * time.Hour)
	competitive48_total_time = time.Duration(0)
)

var competitive48Matcher = match.NewMatcher(competitive48_match_period, competitive48_threshold,
	func(rating, opprating int) bool {
		if math.Abs(float64(rating)-float64(opprating)) > competitive48_bad_difference {
			return false
		} else {
			return true
		}
	},
	func(a string, am interface{}, b string, bm interface{}) {
		ameta := am.(competitive48Setup)
		bmeta := bm.(competitive48Setup)
		id := database.newGame(a, ameta.gameSetup, b, bmeta.gameSetup, true)
		database.setPlayerCompetitive48Slot(a, int(ameta.slot), id)
		database.setPlayerCompetitive48Slot(b, int(bmeta.slot), id)
		competitive48ListenersLock.Lock()
		defer competitive48ListenersLock.Unlock()
		conn, has := competitive48Listeners[a]
		if has {
			err := conn.WriteJSON(struct {
				Slot int
			}{int(ameta.slot)})
			if err != nil {
				delete(competitive48Listeners, a)
			}
		}
		conn, has = competitive48Listeners[b]
		if has {
			err := conn.WriteJSON(struct {
				Slot int
			}{int(bmeta.slot)})
			if err != nil {
				delete(competitive48Listeners, b)
			}
		}
	})

func listeningForCompetitive48Matches(player string, conn *websocket.Conn) {
	competitive48ListenersLock.Lock()
	defer competitive48ListenersLock.Unlock()
	competitive48Listeners[player] = conn
}
