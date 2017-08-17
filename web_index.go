// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"net/http"

	"github.com/pciet/wichess/wichessing"
)

const (
	app_index_template = "web/html/index.html"
)

type indexTemplate struct {
	Wins           int
	Losses         int
	BestPieceName  string
	BestPieceTakes int
	Name           string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}
	key := validSession(r)
	if key != "" {
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
			BestPieceName:  wichessing.NameForKind(piece.Kind),
			BestPieceTakes: piece.Takes,
		})
		return
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}
