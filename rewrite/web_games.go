package main

import (
	"net/http"
)

const GamesRelPath = "/games/"

// A GET to /games/[game identifier] responds with the chess pieces on the board.
// The player must be in the game to be able to read the page.
// A Bad Request (400) response means the request was incorrect.

func GamesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		DebugPrintln(GamesRelPath, "HTTP method", r.Method, "not", http.MethodGet)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	name := ValidSessionHandler(w, r)
	if name == "" {
		return
	}

	id := ParseURLGameIdentifier(w, r, GamesRelPath)
	if id == 0 {
		return
	}

	tx := DatabaseTransaction()
	defer tx.Commit()

	if PlayerInGame(tx, id, name) == false {
		DebugPrintln(GamesRelPath, "player", name, "requested game", id, "that they're not in or doesn't exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	JSONResponse(w, LoadGameBoard(tx, id))
}
