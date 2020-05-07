package main

import (
	"database/sql"
	"net/http"
)

const QuitPath = "/quit"

var QuitHandler = AuthenticRequestHandler{Get: QuitGet}

func QuitGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx, user string) {
	// TODO: AuthenticRequestHandler that gives the player ID instead so
	// that the PlayerID query here isn't needed
	EndSession(tx, PlayerID(tx, user))
	tx.Commit()
	ClearBrowserSession(w, r)
}
