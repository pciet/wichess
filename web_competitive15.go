// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

var (
	competitive15Waiting     = make(map[string]chan struct{})
	competitive15WaitingLock = sync.RWMutex{}
)

func competitive15CancelHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
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
	competitive15WaitingLock.RLock()
	cancel, has := competitive15Waiting[name]
	competitive15WaitingLock.RUnlock()
	if has == false {
		http.NotFound(w, r)
		return
	}
	select {
	case cancel <- struct{}{}:
		return
	case <-time.After(time.Second * 5):
		http.NotFound(w, r)
		return
	}
}

func competitive15Handler(w http.ResponseWriter, r *http.Request) {
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
	id := database.playersCompetitive15HourGameID(name)
	if r.Method == "POST" {
		if id != 0 {
			http.Redirect(w, r, "/competitive15", http.StatusFound)
			return
		}
		if competitive15Matcher.Matching(name) != nil {
			http.NotFound(w, r)
			return
		}
		var assignments BoardAssignments
		err := json.NewDecoder(r.Body).Decode(&assignments)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		defer r.Body.Close()
		setup, err := gameSetupFromRequest(assignments.Assignments)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		rdy := make(chan struct{})
		competitive15Matcher.Match(name, competitive15Setup{gameSetup: setup, ready: rdy}, database.playerRating(name))
		cancel := make(chan struct{})
		competitive15WaitingLock.Lock()
		competitive15Waiting[name] = cancel
		competitive15WaitingLock.Unlock()
		defer func() {
			competitive15WaitingLock.Lock()
			delete(competitive15Waiting, name)
			competitive15WaitingLock.Unlock()
		}()
		select {
		case <-rdy:
			return // the client redirects to a GET /competitive15
		case <-cancel:
			competitive15Matcher.Cancel(name)
			http.NotFound(w, r) // the client ignores the POST response
			return
		}
	} else if r.Method == "GET" {
		if id == 0 {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		executeWebTemplate(w, game_template, gameTemplate{
			GameInfo:  database.updateGameTimes(id, competitive15_turn_time, competitive15_total_time),
			Name:      name,
			TotalTime: competitive15_total_time,
			TurnTime:  competitive15_turn_time,
			NowTime:   time.Now(),
		})
	} else {
		http.NotFound(w, r)
	}
}
