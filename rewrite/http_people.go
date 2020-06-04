package main

import (
	"database/sql"
	"net/http"
)

const (
	PeoplePath     = "/people/"
	PeopleRootPath = "/people"

	RequestedOpponentQuery = "o"
)

var MatchPeopleHandler = AuthenticRequestHandler{
	Get:  GameIdentifierParsed(PlayerNamed(PeopleGet), PeoplePath),
	Post: ArmyParsed(PeoplePost),
}

func PeopleGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx,
	id GameIdentifier, requester Player) {

	h := LoadGameHeader(tx, id, false)
	tx.Commit()

	WriteHTMLTemplate(w, GameHTMLTemplate, GameHTMLTemplateData{requester.Name, h})
}

type PeoplePostJSON struct {
	GameIdentifier `json:"id"`
}

func PeoplePost(w http.ResponseWriter, r *http.Request, tx *sql.Tx,
	requester Player, a ArmyRequest) {

	tx.Commit()

	requestedOpponent, err := ParseURLQuery(r.URL.Query(), RequestedOpponentQuery)
	if err != nil {
		DebugPrintln(PeoplePath, "couldn't parse", RequestedOpponentQuery, ":", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	done := make(chan struct{})
	clientCanceled := r.Context().Done()
	go func() {
		select {
		case <-clientCanceled:
			EndOpponentRequest(requester.ID)
		case <-done:
		}
	}()

	gameID, opponentID := RequestOpponent(requestedOpponent, requester, a)
	if gameID != 0 {
		// game was successfully created
		close(done)
		go AddPlayerRecentOpponent(requester.ID, opponentID)
		tx = DatabaseTransaction()
		UpdatePlayerActivePeopleGame(tx, requester.ID, gameID)
		tx.Commit()
	}
	// a gameID of 0 is normal and a signal to the webpage
	JSONResponse(w, PeoplePostJSON{gameID})
}
