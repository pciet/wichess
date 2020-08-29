package wichess

import (
	"database/sql"
	"net/http"
)

const PeopleIDPath = "/peopleid"

var PeopleIDHandler = AuthenticRequestHandler{
	Get: PeopleIDGet,
}

type PeopleIDJSON struct {
	GameIdentifier `json:"id"`
}

func PeopleIDGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx, requester Player) {
	id := PlayerActivePeopleGame(tx, requester.ID)
	tx.Commit()
	JSONResponse(w, PeopleIDJSON{GameIdentifier: id})
}
