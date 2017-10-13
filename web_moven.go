// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"net/http"
	"strconv"
	"time"
)

func moveNotificationWebsocketHandler(w http.ResponseWriter, r *http.Request) {
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
	gameid, err := strconv.ParseInt(r.URL.Path[7:len(r.URL.Path)], 10, 8)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	game := database.gameWithIdentifier(int(gameid))
	if (game.White != name) && (game.Black != name) {
		http.NotFound(w, r)
		return
	}
	var turnTime time.Duration
	for _, id := range game.DB.playersCompetitive48Games(name) {
		if id == int(gameid) {
			turnTime = competitive48_turn_time
		}
	}
	var previousMove time.Time
	if game.Active == game.White {
		previousMove = game.BlackLatestMove
	} else {
		previousMove = game.WhiteLatestMove
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	listeningToGame(name, game.White, game.Black, turnTime, previousMove, game.ID, conn)
}
