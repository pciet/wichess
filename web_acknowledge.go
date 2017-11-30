// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const (
	request_player = "player"
	request_gameid = "gameid"
)

func acknowledgeGameCompletionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		if debug {
			fmt.Println("acknowledge: request not POST")
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
	p := r.FormValue(request_player)
	if p != name {
		if debug {
			fmt.Println("acknowledge: request_player not valid")
		}
		http.NotFound(w, r)
		return
	}
	id := r.FormValue(request_gameid)
	if id == "" {
		if debug {
			fmt.Println("acknowledge: request_gameid is empty string")
		}
		http.NotFound(w, r)
		return
	}
	gid, err := strconv.Atoi(id)
	if err != nil {
		if debug {
			fmt.Println("acknowledge: failed to parse gameid")
		}
		http.NotFound(w, r)
		return
	}
	tx := database.Begin()
	defer tx.Commit()
	game := tx.gameWithIdentifier(gid, true)
	if (game.White != name) && (game.Black != name) {
		if debug {
			fmt.Println("acknowledge: player doesn't match white or black")
		}
		http.NotFound(w, r)
		return
	}
	if (&game).acknowledgeGameComplete(name, tx) == false {
		if debug {
			fmt.Println("acknowledge: game.acknowledgeGameComplete returned false")
		}
		http.NotFound(w, r)
		return
	}
	if (game.White == easy_computer_player) || (game.Black == easy_computer_player) {
		if (&game).acknowledgeGameComplete(easy_computer_player, tx) == false {
			panicExit("web_acknowledge: failed to acknowledge easy computer player")
		}
	} else if (game.White == hard_computer_player) || (game.Black == hard_computer_player) {
		if (&game).acknowledgeGameComplete(hard_computer_player, tx) == false {
			panicExit("web_acknowledge: failed to acknowledge hard computer  player")
		}
	}
	r.Body.Close()
}
