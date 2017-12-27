// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"fmt"
	"net/http"
)

const (
	app_index_template  = "web/html/index.html"
	game_setup_template = "web/html/setup.html"
	game_template       = "web/html/game.html"
)

type indexTemplate struct {
	Wins       int
	Losses     int
	Draws      int
	Name       string
	F0         string
	F0Matching string
	F0Moveable bool
	F1         string
	F1Matching string
	F1Moveable bool
	F2         string
	F2Matching string
	F2Moveable bool
	F3         string
	F3Matching string
	F3Moveable bool
	F4         string
	F4Matching string
	F4Moveable bool
	F5         string
	F5Matching string
	F5Moveable bool
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		if debug {
			fmt.Println("index: request not GET")
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
	id := database.playersCompetitive5HourGameID(name)
	if id != 0 {
		http.Redirect(w, r, "/competitive5", http.StatusFound)
		return
	}
	id = database.playersCompetitive15HourGameID(name)
	if id != 0 {
		http.Redirect(w, r, "/competitive15", http.StatusFound)
		return
	}
	record := database.playerRecord(name)
	friendGames, playerActive := database.playersFriendSlotOpponentsAndActive(name)
	friendMatching := database.playersFriendMatching(name)
	executeWebTemplate(w, app_index_template, indexTemplate{
		Name:       name,
		Wins:       record.wins,
		Losses:     record.losses,
		Draws:      record.draws,
		F0:         friendGames[0],
		F0Matching: friendMatching[0],
		F0Moveable: playerActive[0],
		F1:         friendGames[1],
		F1Matching: friendMatching[1],
		F1Moveable: playerActive[1],
		F2:         friendGames[2],
		F2Matching: friendMatching[2],
		F2Moveable: playerActive[2],
		F3:         friendGames[3],
		F3Matching: friendMatching[3],
		F3Moveable: playerActive[3],
		F4:         friendGames[4],
		F4Matching: friendMatching[4],
		F4Moveable: playerActive[4],
		F5:         friendGames[5],
		F5Matching: friendMatching[5],
		F5Moveable: playerActive[5],
	})
}
