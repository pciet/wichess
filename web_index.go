// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"net/http"
)

const (
	app_index_template = "web/html/index.html"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}
	if validSession(r) {
		executeWebTemplate(w, app_index_template, nil)
		return
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}
