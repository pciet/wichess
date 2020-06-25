package main

import (
	"database/sql"
	"net/http"

	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

const BoardsPath = "/boards/"

var BoardsHandler = AuthenticRequestHandler{
	Get: GameIdentifierParsed(RequesterInGame(BoardsGet), BoardsPath),
}

func BoardsGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx, id GameIdentifier) {
	b := LoadGameBoard(tx, id, false)
	tx.Commit()

	jr := make([]rules.AddressedSquare, 0, 32)
	for i, s := range b.Board {
		if s.Kind == piece.NoKind {
			continue
		}
		jr = append(jr, rules.AddressedSquare{
			rules.AddressIndex(i).Address(),
			s,
		})
	}

	JSONResponse(w, jr)
}
