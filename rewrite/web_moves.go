package main

import (
	"log"
	"net/http"

	"github.com/pciet/wichess/rules"
)

const (
	MovesRelPath = "/moves/"

	TurnQuery = "turn"

	OutdatedMovesRequest = "outdated"
)

// A GET to /moves/[game identifier]?turn=[turn number] responds with all available moves.
// The turn must be included to avoid race conditions between the host and browser's game state.
// The player must be one of the people matched in the game and have a valid session.
// A Bad Request (400) response means the request was incorrect.

func MovesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		DebugPrintln(MovesRelPath, "HTTP method", r.Method, "not", http.MethodGet)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	name := ValidSessionHandler(w, r)
	if name == "" {
		return
	}

	id := ParseURLGameIdentifier(w, r, MovesRelPath)
	if id == 0 {
		return
	}

	turn := ParseURLIntQuery(w, r, TurnQuery)
	if turn == 0 {
		return
	}

	tx := DatabaseTransaction()
	defer tx.Commit()

	if PlayerInGame(tx, id, name) == false {
		DebugPrintln(MovesRelPath, "player", name, "requested game", id, "that they're not in or doesn't exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if GameTurnEqual(tx, id, turn) == false {
		JSONResponse(w, OutdatedMovesRequest)
		return
	}

	moves, state := MovesForGame(tx, id)
	switch state {
	case rules.Checkmate, rules.Draw, rules.Conceded, rules.TimeOver:
		JSONResponse(w, state)
	case rules.Normal, rules.Promotion, rules.Check:
		JSONResponse(w, moves)
	default:
		log.Panicln("unknown game state", state)
	}
}
