package main

import (
	"database/sql"
	"net/http"
)

const ConcedePath = "/concede/"

var ConcedeHandler = AuthenticRequestHandler{
	Get: GameIdentifierParsed(PlayerNamed(ConcedeGet), ConcedePath),
}

func ConcedeGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx,
	id GameIdentifier, requester Player) {

	// the conceding player doesn't reload the game, so if it's marked conceded then you know
	// it was your opponent that conceded
	MarkGameConceded(tx, id)
	go Alert(id, GameOpponent(tx, id, requester.Name), Update{nil, ConcededUpdate})

	// conceding only happens for people games
	UpdatePlayerActivePeopleGame(tx, requester.ID, 0)

	tx.Commit()

	// on a success response the web browser redirects to the index
}
