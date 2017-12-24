// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"math"
	"time"

	"github.com/pciet/wichess/match"
)

// A zero for an ID value means use the regular piece. Index starts at left pawn and ends at right rook.
type gameSetup [16]int

// Returns the game ID or zero if the setup is invalid.
func (db DB) requestEasyComputerMatch(player string, setup gameSetup) int {
	db.reservePieces(setup)
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

var competitive15_total_time = time.Duration(15 * time.Minute)

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

var competitive5_total_time = time.Duration(5 * time.Minute)

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
