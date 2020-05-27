package main

import (
	"database/sql"
	"net/http"
	"strconv"
)

// This http.Handler type exists to avoid code repetition across most wichess HTTP paths.
// The resulting functionality is sometimes extended with functions like GameIdentifierParse.

// An AuthenticRequestHandler confirms the authenticity of an HTTP request by comparing its
// session key cookie to the session keys saved in the database.
// If the key doesn't match any sessions then the requester gets an HTTP error code response
// and must start a new session at /login.
// If the request is authentic (there's a matching session key) then the requester's username, a
// database ID for the players table, and an open database transaction are passed to the handler
// function matching the HTTP method.
// If a function isn't set for an HTTP method then that method is not allowed.
type AuthenticRequestHandler struct {
	Get  AuthenticRequestHandlerFunc
	Post AuthenticRequestHandlerFunc
}

// An AuthenticRequestHandlerFunc is used in an AuthenticRequestHandler.
// The SQL transaction must be completed in the fuction.
type AuthenticRequestHandlerFunc func(http.ResponseWriter, *http.Request, *sql.Tx, Player)

// ServeHTTP makes AuthenticRequestHandler an http.Handler.
func (an AuthenticRequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if ((r.Method == http.MethodGet) && (an.Get == nil)) ||
		((r.Method == http.MethodPost) && (an.Post == nil)) ||
		((r.Method != http.MethodGet) && (r.Method != http.MethodPost)) {
		DebugPrintln(r.URL.Path, "bad request HTTP method", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pid, err := r.Cookie(PlayerIDCookie)
	if err == http.ErrNoCookie {
		ClearBrowserSession(w, r)
		return
	} else if err != nil {
		Panic(r.URL.Path, "failed to read player ID cookie", PlayerIDCookie, ":", err)
	}

	sc, err := r.Cookie(SessionKeyCookie)
	if err == http.ErrNoCookie {
		ClearBrowserSession(w, r)
		return
	} else if err != nil {
		Panic(r.URL.Path, "failed to read session key cookie", SessionKeyCookie, ":", err)
	}

	playerID, err := strconv.Atoi(pid.Value)
	if err != nil {
		DebugPrintln(r.URL.Path, "couldn't parse player ID cookie", PlayerIDCookie, ":", err)
		ClearBrowserSession(w, r)
		return
	}

	tx := DatabaseTransaction()

	key := PlayersSessionKey(tx, playerID)
	if (key == "") || (key != sc.Value) {
		tx.Commit()
		ClearBrowserSession(w, r)
		return
	}

	name := PlayerName(tx, playerID)
	if name == "" {
		tx.Commit()
		ClearBrowserSession(w, r)
		return
	}

	p := Player{name, playerID}

	switch r.Method {
	case http.MethodGet:
		an.Get(w, r, tx, p)
	case http.MethodPost:
		an.Post(w, r, tx, p)
	default:
		// the method should have been rejected by the starting if
		Panic(r.URL.Path, r.Method, "HTTP method not caught by AuthenticRequestHandler")
	}

	err = tx.Rollback()
	if err != sql.ErrTxDone {
		Panic(r.URL.Path, "HTTP method", r.Method, "handler didn't complete the SQL transaction")
	}
}