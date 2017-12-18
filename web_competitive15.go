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
	competitive15Waiting     = make(map[string]chan struct{})
	competitive15WaitingLock = sync.RWMutex{}
)

func competitive15CancelHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		if debug {
			fmt.Println("competitive15cancel: not POST")
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
	competitive15WaitingLock.RLock()
	cancel, has := competitive15Waiting[name]
	competitive15WaitingLock.RUnlock()
	if has == false {
		if debug {
			fmt.Println("competitive15cancel: not matching")
		}
		http.NotFound(w, r)
		return
	}
	select {
	case cancel <- struct{}{}:
		return
	case <-time.After(time.Second * 5):
		if debug {
			fmt.Println("competitive15cancel: failed to cancel after five seconds")
		}
		http.NotFound(w, r)
		return
	}
}

func competitive15Handler(w http.ResponseWriter, r *http.Request) {
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
	id := database.playersCompetitive15HourGameID(name)
	if r.Method == "POST" {
		if id != 0 {
			http.Redirect(w, r, "/competitive15", http.StatusFound)
			return
		}
		if competitive15Matcher.Matching(name) != nil {
			if debug {
				fmt.Println("competitive15: already matching")
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
		case <-w.(http.CloseNotifier).CloseNotify(): // https://groups.google.com/forum/#!topic/golang-nuts/ROxbuskAglc
			competitive5Matcher.Cancel(name)
			http.NotFound(w, r)
			return
		}
	} else if r.Method == "GET" {
		if id == 0 {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		info := database.updateGameTimes(id, competitive15_total_time, "")
		executeWebTemplate(w, game_template, gameTemplate{
			GameInfo:  info,
			Name:      name,
			TotalTime: competitive15_total_time,
			NowTime:   time.Now(),
		})
	} else {
		http.NotFound(w, r)
	}
}
