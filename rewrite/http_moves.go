package main

import (
	"database/sql"
	"net/http"

	"github.com/pciet/wichess/rules"
)

const (
	MovesPath            = "/moves/"
	TurnQuery            = "turn"
	OutdatedMovesRequest = "outdated"
)

var MovesHandler = AuthenticRequestHandler{
	Get: GameIdentifierParsed(RequesterInGame(MovesGet), MovesPath),
}

type MovesJSON struct {
	Moves       []rules.MoveSet `json:"m,omitempty"`
	rules.State `json:"s"`
}

func MovesGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx, id GameIdentifier) {
	defer tx.Commit()

	/*
		turn, err := ParseURLIntQuery(r.URL.Query(), TurnQuery)
		if err != nil {
			DebugPrintln(MovesPath, "failed to parse turn query:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if GameTurnEqual(tx, id, turn) == false {
			JSONResponse(w, OutdatedMovesRequest)
			return
		}
	*/

	moves, state := MovesForGame(tx, id)

	JSONResponse(w, MovesJSON{moves, state})
}
