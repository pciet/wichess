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

const (
	competitive_48_index_key = "index"
)

func competitive48RequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		if debug {
			fmt.Println("competitive48request: not POST")
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
	var assignments BoardAssignments
	err := json.NewDecoder(r.Body).Decode(&assignments)
	if err != nil {
		if debug {
			fmt.Println(err.Error())
		}
		http.NotFound(w, r)
		return
	}
	defer r.Body.Close()
	if (assignments.Index < 0) || (assignments.Index > 7) {
		if debug {
			fmt.Println("competitive48request: index out of range")
		}
		http.NotFound(w, r)
		return
	}
	if database.playersCompetitive48HourGameID(name, assignments.Index) != 0 {
		http.Redirect(w, r, fmt.Sprintf("/competitive48/%v", assignments.Index), http.StatusFound)
		return
	}
	if competitive48Matcher.Matching(name) != nil {
		if debug {
			fmt.Println("competitive48request: already matching")
		}
		http.NotFound(w, r)
		return
	}
	setup, err := gameSetupFromRequest(assignments.Assignments)
	if err != nil {
		if debug {
			fmt.Println(err.Error())
		}
		http.NotFound(w, r)
		return
	}
	competitive48Matcher.Match(name, competitive48Setup{gameSetup: setup, slot: uint8(assignments.Index)}, database.playerRating(name))
}

func competitive48NotificationWebsocketHandler(w http.ResponseWriter, r *http.Request) {
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
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if debug {
			fmt.Println(err.Error())
		}
		http.NotFound(w, r)
		return
	}
	listeningForCompetitive48Matches(name, conn)
}

func competitive48Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
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
	index, err := strconv.ParseInt(r.URL.Path[15:len(r.URL.Path)], 10, 0)
	if err != nil {
		if debug {
			fmt.Println(err.Error())
		}
		http.NotFound(w, r)
		return
	}
	id := database.playersCompetitive48HourGameID(name, int(index))
	if id == 0 {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	lockGame(id)
	info := database.updateGameTimes(id, competitive48_turn_time, competitive48_total_time, "")
	unlockGame(id)
	executeWebTemplate(w, game_template, gameTemplate{
		GameInfo:  info,
		Name:      name,
		TotalTime: competitive48_total_time,
		TurnTime:  competitive48_turn_time,
		NowTime:   time.Now(),
	})
}
