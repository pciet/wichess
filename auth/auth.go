// Package auth implements players' login and request authentication on the Wisconsin Chess host.
// Passwords are hashed with bcrypt and correct logins result in a session key string used as a
// web browser cookie. This key is then inspected in most future HTTP requests to authenticate
// the player's computer is the one the login was done on.
//
// A primary authentication handler is provided, and additional functionality can be added using
// functions that do things like verify a player is in a game or setup the game memory for writes.
package auth

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/pciet/wichess/memory"
)

// LoginPath is the HTTP path for where a player enters their credentials to get a session key.
// Package auth will redirect here if a key is invalid.
const LoginPath = "/login"

// A HandlerFunc is an application of the auth package for an HTTP path and method called after
// the player has been authenticated. Other func types in the package extend this type with
// additional handling before finally calling a handling function provided to the package.
type HandlerFunc func(http.ResponseWriter, *http.Request, memory.PlayerIdentifier)

// Handler is a http.Handler that uses the requester's session key to get their player identifier.
// The struct fields are functions that do more handler work for an HTTP method. If a function is
// not set then that HTTP method is not allowed.
type Handler struct {
	Get, Post HandlerFunc
}

func (a Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if ((r.Method == http.MethodGet) && (an.Get == nil)) ||
		((r.Method == http.MethodPost) && (an.Post == nil)) ||
		((r.Method != http.MethodGet) && (r.Method != http.MethodPost)) {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sc, err := r.Cookie(SessionKeyCookie)
	if err == http.ErrNoCookie {
		ClearBrowserSession(w, r)
		return
	} else if err != nil {
		panic(fmt.Sprint(r.URL.Path,
			"failed to read session key cookie", SessionKeyCookie, ":", err))
	}

	key := memory.SessionKeyFromString(sc.Value)
	if key == memory.NoSessionKey {
		ClearBrowserSession(w, r)
		return
	}

	pid := memory.SessionPlayerIdentifier(key)
	if pid == memory.NoPlayer {
		ClearBrowserSession(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		an.Get(w, r, pid)
	case http.MethodPost:
		an.Post(w, r, pid)
	default:
		panic(fmt.Sprint(r.URL.Path, r.Method, "HTTP method not caught by Handler"))
	}
}
