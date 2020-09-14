package wichess

import (
	"bytes"
	"net/http"

	"github.com/pciet/wichess/memory"
)

// handleBodyRead reads the request body up to 1024 bytes and returns it. If nil is returned then an
// error occurred, HTTP handling was done by writing http.StatusBadRequest to the response header,
// and the caller just returns.
func handleLimitedBodyRead(w http.ResponseWriter, r *http.Request) []byte {
	var body bytes.Buffer
	_, err := body.ReadFrom(http.MaxBytesReader(w, r.Body, 1024))
	if err != nil {
		debug("body read failed:", err)
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}
	return body.Bytes()
}

// handleInPeopleGame does the HTTP interaction if the player is already in a people game. If true
// is returned then handling was done and the calling handler doesn't continue.
func handleInPeopleGame(w http.ResponseWriter, r *http.Request, p *memory.Player) bool {
	if p.PeopleGame != memory.NoGame {
		http.Redirect(w, r, PeoplePath+p.PeopleGame.String(), http.StatusTemporaryRedirect)
		return true
	}
	return false
}
