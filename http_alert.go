package wichess

import (
	"database/sql"
	"net/http"
)

const AlertPath = "/alert/"

var AlertWebSocketUpgradeHandler = AuthenticRequestHandler{
	Get: GameIdentifierParsed(PlayerNamed(AlertGet), AlertPath),
}

func AlertGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx,
	id GameIdentifier, requester Player) {
	tx.Commit()

	conn, err := WebSocketUpgrade(w, r)
	if err != nil {
		DebugPrintln(err)
		// the upgrade func in WebSocketUpgrade writes an error response, so nothing to add here
		return
	}

	Connect(id, requester.Name, conn)
}
