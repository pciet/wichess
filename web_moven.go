// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"fmt"
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
	gameid, err := strconv.ParseInt(r.URL.Path[7:len(r.URL.Path)], 10, 0)
	if err != nil {
		if debug {
			fmt.Println(err.Error())
		}
		http.NotFound(w, r)
		return
	}
	rLockGame(int(gameid))
	game := database.gameWithIdentifier(int(gameid))
	rUnlockGame(int(gameid))
	if (game.White != name) && (game.Black != name) {
		if debug {
			fmt.Println("moven: player not white or black")
		}
		http.NotFound(w, r)
		return
	}
	if game.White == name {
		if game.WhiteAcknowledge {
			if debug {
				fmt.Println("moven: white already acknowledged")
			}
			http.NotFound(w, r)
			return
		}
	} else {
		if game.BlackAcknowledge {
			if debug {
				fmt.Println("moven: black already acknowledged")
			}
			http.NotFound(w, r)
			return
		}
	}
	var turnTime time.Duration
	for _, id := range game.DB.playersCompetitive48Games(name) {
		if id == int(gameid) {
			turnTime = competitive48_turn_time
		}
	}
	var totalTime time.Duration
	c5 := database.playersCompetitive5HourGameID(name)
	if c5 != 0 {
		totalTime = competitive5_total_time
	} else {
		c15 := database.playersCompetitive15HourGameID(name)
		if c15 != 0 {
			totalTime = competitive15_total_time
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
		if debug {
			fmt.Println(err.Error())
		}
		http.NotFound(w, r)
		return
	}
	listeningToGame(name, game.White, game.Black, turnTime, totalTime, previousMove, game.ID, conn)
}
