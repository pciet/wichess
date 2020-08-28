package wichess

import (
	"database/sql"
	"net/http"
)

const QuitPath = "/quit"

var QuitHandler = AuthenticRequestHandler{Get: QuitGet}

func QuitGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx, requester Player) {
	EndSession(tx, requester.ID)
	tx.Commit()
	ClearBrowserSession(w, r)
}
