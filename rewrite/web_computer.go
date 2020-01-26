package main

import (
	"net/http"
)

const ComputerRelPath = "/computer"

// A GET to /computer responds with the game page for play against the AI computer player.
// A Not Found (404) response means the game doesn't exist yet and should be requested for with a POST first.

// The web browser page sends a POST to /computer with the army configuration to start a new game.
// A Created (201) response means the game was created with the requested pieces.
// A Found (302) response means the game already exists and the requested army configuration is discarded.

func ComputerHandler(w http.ResponseWriter, r *http.Request) {
	name := ValidSessionHandler(w, r)
	if name == "" {
		return
	}

	switch r.Method {
	case http.MethodGet:
		GetComputerHandler(w, r, name)
	case http.MethodPost:
		PostComputerHandler(w, r, name)
	default:
		DebugPrintln(ComputerRelPath, "HTTP method", r.Method, "not GET or POST for player", name)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func GetComputerHandler(w http.ResponseWriter, r *http.Request, player string) {
	tx := DatabaseTransaction()
	defer CommitTransaction(tx)

	id := ComputerGameID(tx, player)
	if id == 0 {
		DebugPrintln(ComputerRelPath, "GET: no game found for player", player)
		http.NotFound(w, r)
		return
	}

	h := GameHeader(tx, id)
	if h.ID == 0 {
		panic(ComputerRelPath, "GET: game ID found for player", player, "but then couldn't get game header")
	}

	WriteGameWebTemplate(w, GameWebTemplateData{player, h})
}

func PostComputerHandler(w http.ResponseWriter, r *http.Request, player string) {
	defer r.Body.Close()
	a, err := DecodeArmyRequest(r.Body)
	if err != nil {
		DebugPrintln(ComputerRelPath, "POST: failed to decode request body for player", player, "-", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tx := DatabaseTransaction()
	defer CommitTransaction(tx)

	id := ComputerGameID(tx, player)
	if id != 0 {
		DebugPrintln(ComputerRelPath, "POST: game already found for player", player)
		http.Redirect(w, r, ComputerRelPath, http.StatusFound)
		return
	}

	err = NewGame(tx, player, a, computer_player, ArmyRequest{}, false)
	if err != nil {
		if err == ErrNoPiece {
			DebugPrintln(ComputerRelPath, "POST: one or more pieces not in database for player", player, a)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			DebugPrintln(ComputerRelPath, "POST: failed to request computer match for player", player, "-", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}
