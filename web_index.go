// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"net/http"
	"strconv"

	"github.com/pciet/wichess/wichessing"
)

const (
	app_index_template  = "web/html/index.html"
	game_setup_template = "web/html/setup.html"
	game_template       = "web/html/game.html"
)

type indexTemplate struct {
	Wins           int
	Losses         int
	BestPieceName  string
	BestPieceTakes int
	Name           string
}

type setupTemplate struct {
	Name string
}

type gameTemplate struct {
	Name   string
	GameID int
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
	if r.URL.Path == "/" {
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
	u, err := strconv.ParseInt(r.URL.Path[1:len(r.URL.Path)], 10, 8)
	if err != nil {
		http.NotFound(w, r)
		return
	} else if u > 63 {
		http.NotFound(w, r)
		return
	}
	gameID := gameIdentifierAtPlayerBoardIndexFromDatabase(name, int(u))
	if gameID == 0 {
		executeWebTemplate(w, game_setup_template, setupTemplate{
			Name: name,
		})
		return
	}
	executeWebTemplate(w, game_template, gameTemplate{
		Name:   name,
		GameID: gameID,
	})
}
