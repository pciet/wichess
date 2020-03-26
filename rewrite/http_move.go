package main

import (
	"database/sql"
	"net/http"

	"github.com/pciet/wichess/rules"
)

const MovePath = "/move/"

var MoveHandler = AuthenticRequestHandler{
	Post: GameIdentifierParse(PlayerNamed(MovePost), MovePath),
}

func MovePost(w http.ResponseWriter, r *http.Request, tx *sql.Tx, id GameIdentifier, player string) {
	move, promotionKind := ParseMove(r)
	if move == rules.NoMove {
		tx.Commit()
		DebugPrintln(MovePath, "failed to parse move by", player)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	changes := Move(tx, id, player, move, promotionKind)
	if (changes == nil) || (len(changes) == 0) {
		tx.Commit()
		DebugPrintln(MovePath, "bad move from", player, "for game", id, ":", move, promotionKind)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	opponent := GameOpponent(tx, id, player)

	tx.Commit()

	go Alert(opponent, id, changes)

	JSONResponse(w, changes)
}
