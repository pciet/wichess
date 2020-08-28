package wichess

import (
	"database/sql"
	"net/http"

	"github.com/pciet/wichess/rules"
)

const PlayersPath = "/players/"

var PlayersHandler = AuthenticRequestHandler{
	Get: GameIdentifierParsed(RequesterInGame(PlayersGet), PlayersPath),
}

type PlayersJSON struct {
	White  string            `json:"w"`
	Black  string            `json:"b"`
	Active rules.Orientation `json:"a"`
}

func PlayersGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx, id GameIdentifier) {
	var rj PlayersJSON
	rj.White, rj.Black = GamePlayers(tx, id)
	rj.Active = GameActive(tx, id)
	tx.Commit()
	JSONResponse(w, rj)
}
