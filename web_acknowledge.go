// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"net/http"
	"strconv"
)

const (
	request_player = "player"
	request_gameid = "gameid"
)

func acknowledgeGameCompletionHandler(w http.ResponseWriter, r *http.Request) {
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
	p := r.FormValue(request_player)
	if p != name {
		http.NotFound(w, r)
		return
	}
	id := r.FormValue(request_gameid)
	if id == "" {
		http.NotFound(w, r)
		return
	}
	gid, err := strconv.Atoi(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	game := gameWithIdentifier(gid)
	if (game.White != name) && (game.Black != name) {
		http.NotFound(w, r)
		return
	}
	success := game.acknowledge(name)
	if success == false {
		http.NotFound(w, r)
		return
	}
}
