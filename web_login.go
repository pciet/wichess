// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"net/http"
)

const (
	login_template = "web/html/login.html"

	form_player_name = "name"
	form_password    = "password"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		loginAttempt(w, r)
		return
	}
	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}
	key := validSession(r)
	if key != "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	executeWebTemplate(w, login_template, nil)
}

func loginAttempt(w http.ResponseWriter, r *http.Request) {
	if validSession(r) != "" {
		http.Redirect(w, r, "/", http.StatusFound)
	}
	err := r.ParseForm()
	if err != nil {
		webError(w, r, "failed to parse form", err)
		return
	}
	playerName := r.FormValue(form_player_name)
	password := r.FormValue(form_password)
	if (playerName == "") || (password == "") {
		webError(w, r, "missing form field", nil)
		return
	}
	key := database.loginOrCreate(playerName, password)
	if key == "" {
		executeWebTemplate(w, login_template, nil)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     key_cookie,
		Value:    key,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // TODO: set true after TLS certification
	})
	http.Redirect(w, r, "/", http.StatusFound)
}
