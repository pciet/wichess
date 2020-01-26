package main

import (
	"net/http"
	"time"
)

const (
	http_max_idle_conns    = 50
	http_idle_conn_timeout = time.Duration(time.Minute * 5)
)

func main() {
	InitializeDatabaseConnection()

	// All interactions with wichess require a session which is started by logging in with a username and password.
	// The login page is gotten with this handler, and a post is used to submit the credentials.
	http.HandleFunc(LoginRelPath, LoginHandler)

	// The index page lets the player pick their mode and army and shows existing friend games.
	// If a timed match is in progress then this handler redirects the browser to that page.
	http.HandleFunc(IndexRelPath, IndexHandler)

	// Computer is the AI player mode where the opponent's moves are made automatically by wichess.
	// The Computer path is used to setup the game and to get the computer game page.
	http.HandleFunc(ComputerRelPath, ComputerHandler)

	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("web/js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("web/css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("web/img"))))
	http.Handle("/sound/", http.StripPrefix("/sound/", http.FileServer(http.Dir("web/sound"))))

	// github.com/pciet/wichess/issues/9 - 100 http + 100 postgres should be under macOS 256 limit
	dt, ok := http.DefaultTransport.(*http.Transport)
	if ok == false {
		panic("http.DefaultTransport.(*http.Transport) failed")
	}
	dt.MaxIdleConns = http_max_idle_conns
	dt.IdleConnTimeout = http_idle_conn_timeout

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
