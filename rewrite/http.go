package main

import (
	"log"
	"net/http"
	"time"
)

const (
	// TODO: does this unnecessarily limit the maximum number of concurrent players on server hardware that can support more?
	// github.com/pciet/wichess/issues/9 - 100 http + 100 postgres should be under macOS 256 limit
	http_max_idle_conns    = 50
	http_idle_conn_timeout = time.Duration(time.Minute * 5)
)

func InitializeHTTP() {
	// First when a new person connects to this program they are given a webpage to login with.
	// If the username and password provided exist together in the database, or the username is new,
	// then the person's web browser is issued a unique session key to be stored as a cookie.
	http.HandleFunc(LoginRelPath, LoginHandler)

	// The index webpage is a place to choose what kind of match to play next. This is where the army is picked.
	// If a timed game is in progress then the web browser is redirected to it.
	// If the session key cookie is invalid then the web browser is redirected to the login page.
	http.HandleFunc(IndexRelPath, IndexHandler)

	// "Computer" is the artificial opponent mode, where the other player's moves are made automatically.
	// In HTTP terms, this path's POST is used to setup the match and the GET is to load the game page.
	http.HandleFunc(ComputerRelPath, ComputerHandler)

	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("web/js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("web/css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("web/img"))))
	http.Handle("/sound/", http.StripPrefix("/sound/", http.FileServer(http.Dir("web/sound"))))
	http.Handle("/font/", http.StripPrefix("/font/", http.FileServer(http.Dir("web/fonts"))))

	dt, ok := http.DefaultTransport.(*http.Transport)
	if ok == false {
		log.Panic("http.DefaultTransport.(*http.Transport) failed")
	}
	dt.MaxIdleConns = http_max_idle_conns
	dt.IdleConnTimeout = http_idle_conn_timeout
}
