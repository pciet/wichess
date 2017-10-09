// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"net/http"
)

const (
	app_index_template  = "web/html/index.html"
	game_setup_template = "web/html/setup.html"
	game_template       = "web/html/game.html"
)

const (
	available = iota
	matching
	matched
)

type indexTemplate struct {
	Wins   int
	Losses int
	Draws  int
	Name   string
	C48S0  int
	C48S1  int
	C48S2  int
	C48S3  int
	C48S4  int
	C48S5  int
	C48S6  int
	C48S7  int
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
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
	record := database.playerRecord(name)
	c48games := database.playersCompetitive48Games(name)
	var states [8]int
	for i := 0; i < 8; i++ {
		if c48games[i] != 0 {
			states[i] = matched
		} else {
			states[i] = available
		}
	}
	meta := competitive48Matcher.Matching(name)
	if meta != nil {
		states[meta.(competitive48Setup).slot] = matching
	}
	executeWebTemplate(w, app_index_template, indexTemplate{
		Name:   name,
		Wins:   record.wins,
		Losses: record.losses,
		Draws:  record.draws,
		C48S0:  states[0],
		C48S1:  states[1],
		C48S2:  states[2],
		C48S3:  states[3],
		C48S4:  states[4],
		C48S5:  states[5],
		C48S6:  states[6],
		C48S7:  states[7],
	})
}
