// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func gamesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		if debug {
			fmt.Println("games: request not GET")
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
	gameid, err := strconv.ParseInt(r.URL.Path[7:len(r.URL.Path)], 10, 0)
	if err != nil {
		if debug {
			fmt.Println(err.Error())
		}
		http.NotFound(w, r)
		return
	}
	tx := database.Begin()
	defer tx.Commit()
	game := tx.gameWithIdentifier(int(gameid), false)
	if (game.White != name) && (game.Black != name) {
		if debug {
			fmt.Println("games: player not white or black")
		}
		http.NotFound(w, r)
		return
	}
	json, err := json.Marshal(game)
	if err != nil {
		panicExit(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
