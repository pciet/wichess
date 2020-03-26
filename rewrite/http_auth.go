package main

import (
	"database/sql"
	"net/http"
)

// This http.Handler type exists to avoid code repetition across most wichess HTTP paths.
// The resulting generalized functionality is sometimes extended with functions like GameIdentifierParse.

// An AuthenticRequestHandlerFunc is used in an AuthenticRequestHandler.
// The SQL transaction must be completed in the fuction.
// The string argument is the requester's username.
type AuthenticRequestHandlerFunc func(http.ResponseWriter, *http.Request, *sql.Tx, string)

// An AuthenticRequestHandler confirms the authenticity of an HTTP request
// by comparing its session key cookie to the session keys saved in the database.
// If the key doesn't match any sessions then the requester gets an HTTP error
// code response and must start a new session at /login.
// If the request is authentic (there's a matching session key) then the requester's
// username and an open database transaction are passed to the handler function
// matching the HTTP method.
// If a function isn't set for an HTTP method then that method is not allowed.
type AuthenticRequestHandler struct {
	Get  AuthenticRequestHandlerFunc
	Post AuthenticRequestHandlerFunc
}

// ServeHTTP makes AuthenticRequestHandler an http.Handler.
func (an AuthenticRequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if ((r.Method == http.MethodGet) && (an.Get == nil)) ||
		((r.Method == http.MethodPost) && (an.Post == nil)) ||
		((r.Method != http.MethodGet) && (r.Method != http.MethodPost)) {
		DebugPrintln(r.URL.Path, "bad request HTTP method", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sc, err := r.Cookie(SessionKeyCookie)
	if err == http.ErrNoCookie {
		ClearBrowserSession(w, r)
		return
	} else if err != nil {
		Panic(r.URL.Path, "failed to read session key cookie", SessionKeyCookie, "from HTTP request:", err)
	}

	if len(sc.Value) == 0 {
		DebugPrintln(r.URL.Path, "session key cookie", SessionKeyCookie, "length zero")
		ClearBrowserSession(w, r)
		return
	}

	tx := DatabaseTransaction()

	name := SessionRequester(tx, sc.Value)
	if name == "" {
		tx.Commit()
		DebugPrintln(r.URL.Path, "no username associated with session key", sc.Value)
		ClearBrowserSession(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		an.Get(w, r, tx, name)
	case http.MethodPost:
		an.Post(w, r, tx, name)
	default:
		// the method should have been rejected by the starting if
		Panic(r.URL.Path, r.Method, "HTTP method not caught by AuthenticRequestHandler")
	}

	err = tx.Rollback()
	if err != sql.ErrTxDone {
		Panic(r.URL.Path, "AuthenticRequestHandlerFunc for HTTP method", r.Method, "didn't end the SQL transaction")
	}
}
