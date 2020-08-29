package wichess

import (
	"database/sql"
	"net/http"

	"github.com/pciet/wichess/piece"
)

const PicksPath = "/picks"

var PicksHandler = AuthenticRequestHandler{
	Get: PicksGet,
}

type PicksJSON struct {
	Left  piece.Kind `json:"l"`
	Right piece.Kind `json:"r"`
}

func PicksGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx, requester Player) {
	left, right := PlayerPiecePicks(tx, requester.Name)
	tx.Commit()
	JSONResponse(w, PicksJSON{left, right})
}
