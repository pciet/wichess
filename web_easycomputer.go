// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func easyComputerRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		if debug {
			fmt.Println("easycomputerrequest: not POST")
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
	if database.easyComputerGame(name) != 0 {
		http.Redirect(w, r, "/easycomputer", http.StatusFound)
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
	if database.requestEasyComputerMatch(name, setup) == 0 {
		if debug {
			fmt.Println("easycomputerrequest: failed to request match")
		}
		http.NotFound(w, r)
		return
	}
	// the client is responsible for triggering a GET at /easycomputer after this POST is successful
}

func easyComputerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		if debug {
			fmt.Println("easycomputer: request not GET")
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
	id := database.easyComputerGame(name)
	if id == 0 {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	white, black := database.gamePlayers(id)
	executeWebTemplate(w, game_template, gameTemplate{
		GameInfo: GameInfo{
			White: white,
			Black: black,
			ID:    id,
		},
		Name: name,
	})
}
