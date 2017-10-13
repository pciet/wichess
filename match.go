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

func (db DB) requestEasyComputerMatch(player string, setup gameSetup) {
	db.newGame(player, setup, easy_computer_player, gameSetup{}, time.Duration(0), time.Duration(0))
}

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
		id := database.newGame(a, ameta.gameSetup, b, bmeta.gameSetup, competitive48_total_time, competitive48_turn_time)
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
