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

	h := LoadGameHeader(tx, id, false)
	o := OrientationOf(requester.Name, h.White.Name, h.Black.Name)
	if h.Active != o {
		Autoplay(id, ComputerPlayerName)
	}
	h.Active = o

	WriteHTMLTemplate(w, GameHTMLTemplate, GameHTMLTemplateData{requester.Name, h})
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

	var whiteArmy, blackArmy ArmyRequest
	var white, black Player
	if RandomBool() {
		whiteArmy = a
		white = requester
		blackArmy = RegularArmyRequest
		black = Player{ComputerPlayerName, 0}
	} else {
		whiteArmy = RegularArmyRequest
		white = Player{ComputerPlayerName, 0}
		blackArmy = a
		black = requester
	}
	id = NewGame(tx, whiteArmy, blackArmy, white, black)
	if id == 0 {
		DebugPrintln(ComputerPath, "bad army request for", requester)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
