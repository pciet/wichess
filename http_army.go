package main

import (
	"database/sql"
	"net/http"
)

// See http_auth.go and http_idparse.go for examples of this handler pattern.

func ArmyParsed(calls ArmyParsedFunc) AuthenticRequestHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, tx *sql.Tx, requester Player) {
		a, err := DecodeArmyRequest(r.Body)
		if err != nil {
			tx.Commit()
			DebugPrintln("failed to decode army request for", requester, ":", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		calls(w, r, tx, requester, a)
	}
}

type ArmyParsedFunc func(http.ResponseWriter, *http.Request, *sql.Tx, Player, ArmyRequest)
