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
	Get: GameIdentifierParsed(PlayerNamed(MovesGet), MovesPath),
}

type MovesJSON struct {
	Moves       []rules.MoveSet `json:"m,omitempty"`
	rules.State `json:"s"`
}

func MovesGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx,
	id GameIdentifier, requester Player) {

	// TODO: turn number isn't currently used, does the rewrite still have the race condition that
	// made that counting necessary?
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

	if (GameHasPlayer(tx, id, ComputerPlayerName) == false) &&
		((state == rules.Check) || (state == rules.Draw) || (state == rules.Checkmate)) {
		var alertState string
		switch state {
		case rules.Check:
			alertState = CheckCalculatedUpdate
		case rules.Draw:
			alertState = DrawCalculatedUpdate
		case rules.Checkmate:
			alertState = CheckmateCalculatedUpdate
		}
		go Alert(id, GameOpponent(tx, id, requester.Name), Update{
			State:    alertState,
			FromMove: rules.NoMove,
		})
	}

	tx.Commit()

	JSONResponse(w, MovesJSON{moves, state})
}
