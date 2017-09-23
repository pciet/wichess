// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const (
	request_from = "From"
	request_to   = "To"
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
	gameid, err := strconv.ParseInt(r.URL.Path[6:len(r.URL.Path)], 10, 8)
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
	to, err = strconv.Atoi(r.FormValue(request_to))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	game := gameWithIdentifier(int(gameid))
	if (game.White != name) && (game.Black != name) {
		http.NotFound(w, r)
		return
	}
	diff := game.move(from, to, name)
	if (diff == nil) || (len(diff) == 0) {
		http.NotFound(w, r)
		return
	}
	json, err := json.Marshal(diff)
	if err != nil {
		panicExit(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
