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
	id GameIdentifier, player string) {
	defer tx.Commit()

	if GameComplete(tx, id) == false {
		DebugPrintln(AcknowledgePath, player,
			"requested acknowledge of", id, "but not complete")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	AcknowledgeGameComplete(tx, id, player)
}
