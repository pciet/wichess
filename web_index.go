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
	Wins   int
	Losses int
	Draws  int
	Name   string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		if debug {
			fmt.Println("index: request not GET")
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
	executeWebTemplate(w, app_index_template, indexTemplate{
		Name:   name,
		Wins:   record.wins,
		Losses: record.losses,
		Draws:  record.draws,
	})
}
