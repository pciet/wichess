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

type indexTemplate struct {
	Wins           int
	Losses         int
	Draws          int
	BestPieceName  string
	BestPieceTakes int
	Name           string
}

type setupTemplate struct {
	Name string
}

type gameTemplate struct {
	White  string
	Black  string
	GameID int
	Name   string
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
	record := playerRecordFromDatabase(name)
	piece := bestPieceForPlayerFromDatabase(name)
	executeWebTemplate(w, app_index_template, indexTemplate{
		Name:           name,
		Wins:           record.wins,
		Losses:         record.losses,
		Draws:          record.draws,
		BestPieceName:  nameForKind(piece.Kind),
		BestPieceTakes: piece.Takes,
	})
}
