// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"net/http"
)

const request_assignments = "assignments[]"

func easyComputerRequestHandler(w http.ResponseWriter, r *http.Request) {
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
	if database.easyComputerGame(name) != 0 {
		http.Redirect(w, r, "/easycomputer", http.StatusFound)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.NotFound(w, r)
		return
	}
	setup, err := gameSetupFromForm(r.PostForm[request_assignments])
	if err != nil {
		http.NotFound(w, r)
		return
	}
	database.requestEasyComputerMatch(name, setup)
	// the client is responsible for triggering a GET at /easycomputer after this POST is successful
}

func easyComputerHandler(w http.ResponseWriter, r *http.Request) {
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
	id := database.easyComputerGame(name)
	if id == 0 {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	white, black := database.gamePlayers(id)
	executeWebTemplate(w, game_template, gameTemplate{
		White:  white,
		Black:  black,
		GameID: id,
		Name:   name,
	})
}
