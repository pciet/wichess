// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const request_turn = "Turn"

func movesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		if debug {
			fmt.Println("moves: request not GET")
		}
		http.NotFound(w, r)
		return
	}
	// TODO: this shouldn't be possible
	if r.URL.Path == "/" {
		if debug {
			fmt.Println("moves: request.URL.Path == /")
		}
		http.NotFound(w, r)
		return
	}
	key := validSession(r)
	if key == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	name := nameFromSessionKey(key)
	if name == "" {
		clearClientSession(w)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	gameid, err := strconv.ParseInt(r.URL.Path[7:len(r.URL.Path)], 10, 0)
	if err != nil {
		if debug {
			fmt.Println(err.Error())
		}
		http.NotFound(w, r)
		return
	}
	var totalTime time.Duration
	c5 := database.playersCompetitive5HourGameID(name)
	if c5 != 0 {
		totalTime = competitive5_total_time
	} else {
		c15 := database.playersCompetitive15HourGameID(name)
		if c15 != 0 {
			totalTime = competitive15_total_time
		}
	}
	s, has := r.URL.Query()[request_turn]
	if has == false {
		if debug {
			fmt.Println("moves: request missing request_turn value")
		}
		http.NotFound(w, r)
		return
	}
	if (len(s) == 0) || (len(s) > 1) {
		if debug {
			fmt.Println("moves: request has invalid request_turn value")
		}
		http.NotFound(w, r)
		return
	}
	turn, err := strconv.ParseInt(s[0], 10, 0)
	if err != nil {
		if debug {
			fmt.Println(err.Error())
		}
		http.NotFound(w, r)
		return
	}
	lockGame(int(gameid))
	defer unlockGame(int(gameid))
	game := database.gameWithIdentifier(int(gameid))
	if (game.White != name) && (game.Black != name) {
		if debug {
			fmt.Println("moves: player not white or black")
		}
		http.NotFound(w, r)
		return
	}
	var moves map[string]map[string]struct{}
	moves = game.moves(totalTime)
	done := false
	for mv, _ := range moves {
		if (mv == "checkmate") || (mv == "time") || (mv == "draw") {
			done = true
			break
		}
	}
	if (done == false) && (int(turn) != game.Turn) {
		moves = map[string]map[string]struct{}{
			"outdated": nil,
		}
	}
	json, err := json.Marshal(moves)
	if err != nil {
		panicExit(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
