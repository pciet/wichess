package main

import (
	"database/sql"
	"net/http"
)

const ComputerPath = "/computer"

var ComputerHandler = AuthenticRequestHandler{
	Get:  ComputerGet,
	Post: ArmyParsed(ComputerPost),
}

func ComputerGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx, requester Player) {
	defer tx.Commit()

	id := ComputerGameIdentifier(tx, requester.Name)
	if id == 0 {
		DebugPrintln(ComputerPath, "game not found for", requester)
		http.NotFound(w, r)
		return
	}

	WriteHTMLTemplate(w, GameHTMLTemplate,
		GameHTMLTemplateData{requester.Name, LoadGameHeader(tx, id, false)})
}

func ComputerPost(w http.ResponseWriter, r *http.Request, tx *sql.Tx,
	requester Player, a ArmyRequest) {

	defer tx.Commit()

	// TODO: http handler pattern for computer game id
	id := ComputerGameIdentifier(tx, requester.Name)
	if id != 0 {
		http.Redirect(w, r, ComputerPath, http.StatusSeeOther)
		return
	}

	id = NewGame(tx, a, RegularArmyRequest, requester, Player{ComputerPlayerName, 0})
	if id == 0 {
		DebugPrintln(ComputerPath, "bad army request for", requester)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
