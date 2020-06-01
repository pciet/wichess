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

	requestedOpponent, err := ParseURLQuery(r.URL.Query(), RequestedOpponentQuery)
	if err != nil {
		DebugPrintln(PeoplePath, "couldn't parse", RequestedOpponentQuery, ":", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	opponentID := PlayerID(tx, requestedOpponent)
	tx.Commit()

	done := make(chan struct{})
	clientCanceled := r.Context().Done()
	go func() {
		select {
		case <-clientCanceled:
			EndOpponentRequest(requester.ID)
		case <-done:
		}
	}()

	gameID := RequestOpponent(opponentID, requester.ID, a)
	if gameID != 0 {
		close(done)
	}
	JSONResponse(w, PeoplePostJSON{gameID})
}
