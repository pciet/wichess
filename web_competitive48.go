// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const (
	competitive_48_index_key = "index"
)

func competitive48RequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
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
	err := r.ParseForm()
	if err != nil {
		http.NotFound(w, r)
		return
	}
	istr, has := r.PostForm[competitive_48_index_key]
	if has == false {
		http.NotFound(w, r)
		return
	}
	index, err := strconv.Atoi(istr[0])
	if err != nil {
		http.NotFound(w, r)
		return
	}
	if (index < 0) || (index > 7) {
		http.NotFound(w, r)
		return
	}
	if database.playersCompetitive48HourGameID(name, index) != 0 {
		http.Redirect(w, r, fmt.Sprintf("/competitive48/%v", index), http.StatusFound)
		return
	}
	if competitive48Matcher.Matching(name) != nil {
		http.NotFound(w, r)
		return
	}
	setup, err := gameSetupFromForm(r.PostForm[request_assignments])
	if err != nil {
		http.NotFound(w, r)
		return
	}
	competitive48Matcher.Match(name, competitive48Setup{gameSetup: setup, slot: uint8(index)}, database.playerRating(name))
}

func competitive48NotificationWebsocketHandler(w http.ResponseWriter, r *http.Request) {
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
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	listeningForCompetitive48Matches(name, conn)
}

func competitive48Handler(w http.ResponseWriter, r *http.Request) {
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
	index, err := strconv.ParseInt(r.URL.Path[15:len(r.URL.Path)], 10, 8)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	id := database.playersCompetitive48HourGameID(name, int(index))
	if id == 0 {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	white, black := database.gamePlayers(id)
	executeWebTemplate(w, game_template, gameTemplate{
		White:  white,
		Black:  black,
		GameID: id,
		Name:   name,
	})
}
