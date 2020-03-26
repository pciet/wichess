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
	Get: GameIdentifierParse(RequesterInGame(MovesGet), MovesPath),
}

func MovesGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx, id GameIdentifier) {
	defer tx.Commit()

	turn, err := ParseURLIntQuery(TurnQuery, r.URL.Query())
	if err != nil {
		DebugPrintln(MovesPath, "failed to parse turn query:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	/*
		if GameTurnEqual(tx, id, turn) == false {
			JSONResponse(w, OutdatedMovesRequest)
			return
		}
	*/

	moves, state := MovesForGame(tx, id)

	switch state {
	case rules.Checkmate, rules.Draw, rules.Conceded, rules.TimeOver:
		JSONResponse(w, state)
	case rules.Normal, rules.Promotion, rules.Check:
		JSONResponse(w, moves)
	default:
		Panic("unknown game state", state)
	}
}
