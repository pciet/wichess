// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	competitive5Waiting     = make(map[string]chan struct{})
	competitive5WaitingLock = sync.RWMutex{}
)

func competitive5CancelHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		if debug {
			fmt.Println("competitive5cancel: not POST")
		}
		http.NotFound(w, r)
		return
	}
	key, name := database.validSession(r)
	if key == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if name == "" {
		clearClientSession(w)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	competitive5WaitingLock.RLock()
	cancel, has := competitive5Waiting[name]
	competitive5WaitingLock.RUnlock()
	if has == false {
		if debug {
			fmt.Println("competitive5cancel: not waiting")
		}
		http.NotFound(w, r)
		return
	}
	select {
	case cancel <- struct{}{}:
		return
	case <-time.After(time.Second * 5):
		if debug {
			fmt.Println("competitive5cancel: failed to cancel after five seconds")
		}
		http.NotFound(w, r)
		return
	}
}

func competitive5Handler(w http.ResponseWriter, r *http.Request) {
	key, name := database.validSession(r)
	if key == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if name == "" {
		clearClientSession(w)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	id := database.playersCompetitive5HourGameID(name)
	if r.Method == "POST" {
		if id != 0 {
			http.Redirect(w, r, "/competitive5", http.StatusFound)
			return
		}
		if competitive5Matcher.Matching(name) != nil {
			if debug {
				fmt.Println("competitive5: matching")
			}
			http.NotFound(w, r)
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
		setup, err := gameSetupFromRequest(assignments.Assignments)
		if err != nil {
			if debug {
				fmt.Println(err.Error())
			}
			http.NotFound(w, r)
			return
		}
		rdy := make(chan struct{})
		competitive5Matcher.Match(name, competitive5Setup{gameSetup: setup, ready: rdy}, database.playerRating(name))
		cancel := make(chan struct{})
		competitive5WaitingLock.Lock()
		competitive5Waiting[name] = cancel
		competitive5WaitingLock.Unlock()
		defer func() {
			competitive5WaitingLock.Lock()
			delete(competitive5Waiting, name)
			competitive5WaitingLock.Unlock()
		}()
		select {
		case <-rdy:
			return // the client redirects to a GET /competitive5
		case <-cancel:
			competitive5Matcher.Cancel(name)
			http.NotFound(w, r) // the client ignores the POST response
			return
		}
	} else if r.Method == "GET" {
		if id == 0 {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		lockGame(id)
		info := database.updateGameTimes(id, competitive5_total_time, "")
		unlockGame(id)
		executeWebTemplate(w, game_template, gameTemplate{
			GameInfo:  info,
			Name:      name,
			TotalTime: competitive5_total_time,
			NowTime:   time.Now(),
		})
	} else {
		if debug {
			fmt.Println("competitive5: not GET or POST")
		}
		http.NotFound(w, r)
	}
}
