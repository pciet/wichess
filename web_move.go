// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/pciet/wichess/wichessing"
)

const (
	request_from         = "From"
	request_to           = "To"
	request_promote_kind = "Kind"
)

func moveRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}
	if r.URL.Path == "/" {
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
	gameid, err := strconv.ParseInt(r.URL.Path[6:len(r.URL.Path)], 10, 0)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	var from, to int
	from, err = strconv.Atoi(r.FormValue(request_from))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	var kind int
	if r.FormValue(request_promote_kind) != "" {
		kind, err = strconv.Atoi(r.FormValue(request_promote_kind))
		if err != nil {
			http.NotFound(w, r)
			return
		}
		if (wichessing.Kind(kind) != wichessing.Knight) &&
			(wichessing.Kind(kind) != wichessing.Bishop) &&
			(wichessing.Kind(kind) != wichessing.Rook) &&
			(wichessing.Kind(kind) != wichessing.Queen) {
			http.NotFound(w, r)
			return
		}
	} else {
		to, err = strconv.Atoi(r.FormValue(request_to))
		if err != nil {
			http.NotFound(w, r)
			return
		}
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
	lockGame(int(gameid))
	defer unlockGame(int(gameid))
	game := database.gameWithIdentifier(int(gameid))
	if (game.White != name) && (game.Black != name) {
		http.NotFound(w, r)
		return
	}
	(&game).updateGameTimesWithMove(time.Now())
	var diff map[string]piece
	if game.timeLoss(game.activeOrientation(), totalTime) {
		diff = map[string]piece{}
	} else {
		var promoting bool
		var promotingOrientation wichessing.Orientation
		if kind != 0 { // promotion
			diff = game.promote(from, name, wichessing.Kind(kind), false)
			if (diff == nil) || (len(diff) == 0) {
				http.NotFound(w, r)
				return
			}
		} else {
			diff, promoting, promotingOrientation = game.move(from, to, name, false)
			if (diff == nil) || (len(diff) == 0) {
				http.NotFound(w, r)
				return
			}
		}
		var orientation wichessing.Orientation
		if game.White == name {
			orientation = wichessing.White
		} else {
			orientation = wichessing.Black
		}
		if (promoting == false) || (promotingOrientation != orientation) {
			if (game.White == easy_computer_player) || (game.Black == easy_computer_player) {
				cdiff := database.easyComputerMoveForGame(int(gameid))
				if (cdiff != nil) && (len(cdiff) != 0) {
					for addr, piece := range cdiff {
						diff[addr] = piece
					}
				}
			} else if (game.White == hard_computer_player) || (game.Black == hard_computer_player) {
				cdiff := database.hardComputerMoveForGame(int(gameid))
				if (cdiff != nil) && (len(cdiff) != 0) {
					for addr, piece := range cdiff {
						diff[addr] = piece
					}
				}
			}
		}
	}
	json, err := json.Marshal(diff)
	if err != nil {
		panicExit(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
