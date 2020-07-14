package main

import (
	"database/sql"
	"net/http"
)

const AcknowledgePath = "/acknowledge/"

var AcknowledgeHandler = AuthenticRequestHandler{
	Get: GameIdentifierParsed(PlayerNamed(AcknowledgeGet), AcknowledgePath),
}

func AcknowledgeGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx,
	id GameIdentifier, requester Player) {
	defer tx.Commit()

	if GameComplete(tx, id) == false {
		DebugPrintln(AcknowledgePath, requester,
			"requested acknowledge of", id, "but not complete")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	AcknowledgeGameComplete(tx, id, requester.Name)

	if PlayerActivePeopleGame(tx, requester.ID) == id {
		UpdatePlayerActivePeopleGame(tx, requester.ID, 0)
	}
}
