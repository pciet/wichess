package main

import (
	"database/sql"
	"net/http"
)

const ComputerPath = "/computer"

var ComputerHandler = AuthenticRequestHandler{
	Get:  ComputerGet,
	Post: ComputerPost,
}

func ComputerGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx, requester string) {
	defer tx.Commit()

	id := ComputerGameIdentifier(tx, requester)
	if id == 0 {
		DebugPrintln(ComputerPath, "game not found for", requester)
		http.NotFound(w, r)
		return
	}

	WriteHTMLTemplate(w, GameHTMLTemplate, GameHTMLTemplateData{requester, LoadGameHeader(tx, id)})
}

func ComputerPost(w http.ResponseWriter, r *http.Request, tx *sql.Tx, requester string) {
	defer tx.Commit()

	id := ComputerGameIdentifier(tx, requester)
	if id != 0 {
		http.Redirect(w, r, ComputerPath, http.StatusSeeOther)
		return
	}

	a, err := DecodeArmyRequest(r.Body)
	if err != nil {
		DebugPrintln(ComputerPath, "failed to decode request body for", requester, ":", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id = NewGame(tx, false, requester, a, ComputerPlayerName, ArmyRequest{})
	if id == 0 {
		DebugPrintln(ComputerPath, "bad army request for", requester)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
