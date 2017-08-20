// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"net/http"
	"strconv"
)

const (
	request_left_rook    = "lrook"
	request_left_knight  = "lknight"
	request_left_bishop  = "lbishop"
	request_right_bishop = "rbishop"
	request_right_knight = "rknight"
	request_right_rook   = "rrook"

	request_point = "point"
)

func requestMatchHandler(w http.ResponseWriter, r *http.Request) {
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
	var lr, lk, lb, rb, rk, rr int
	var err error
	if r.FormValue(request_left_rook) != "" {
		lr, err = strconv.Atoi(r.FormValue(request_left_rook))
		if err != nil {
			http.NotFound(w, r)
			return
		}
	}
	if r.FormValue(request_left_knight) != "" {
		lk, err = strconv.Atoi(r.FormValue(request_left_knight))
		if err != nil {
			http.NotFound(w, r)
			return
		}
	}
	if r.FormValue(request_left_bishop) != "" {
		lb, err = strconv.Atoi(r.FormValue(request_left_bishop))
		if err != nil {
			http.NotFound(w, r)
			return
		}
	}
	if r.FormValue(request_right_bishop) != "" {
		rb, err = strconv.Atoi(r.FormValue(request_right_bishop))
		if err != nil {
			http.NotFound(w, r)
			return
		}
	}
	if r.FormValue(request_right_knight) != "" {
		rk, err = strconv.Atoi(r.FormValue(request_right_knight))
		if err != nil {
			http.NotFound(w, r)
			return
		}
	}
	if r.FormValue(request_right_rook) != "" {
		rr, err = strconv.Atoi(r.FormValue(request_right_rook))
		if err != nil {
			http.NotFound(w, r)
			return
		}
	}
	requestMatch(name, lr, lk, lb, rb, rk, rr)
}
