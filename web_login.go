// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"fmt"
	"net/http"
)

const (
	login_template = "web/html/login.html"

	form_player_name = "name"
	form_password    = "password"

	name_max_length = 64
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		loginAttempt(w, r)
		return
	}
	if r.Method != "GET" {
		if debug {
			fmt.Println("login: request not GET")
		}
		http.NotFound(w, r)
		return
	}
	key, _ := database.validSession(r)
	if key != "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	executeWebTemplate(w, login_template, nil)
}

func loginAttempt(w http.ResponseWriter, r *http.Request) {
	key, _ := database.validSession(r)
	if key != "" {
		http.Redirect(w, r, "/", http.StatusFound)
	}
	err := r.ParseForm()
	if err != nil {
		if debug {
			fmt.Println(err.Error())
		}
		http.NotFound(w, r)
		return
	}
	playerName := r.FormValue(form_player_name)
	password := r.FormValue(form_password)
	if (playerName == "") || (password == "") {
		if debug {
			fmt.Println("login: form missing player or password")
		}
		http.NotFound(w, r)
		return
	}
	if len(playerName) > name_max_length {
		if debug {
			fmt.Println("login: player name longer than max length")
		}
		http.NotFound(w, r)
		return
	}
	key = database.loginOrCreate(playerName, password)
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
