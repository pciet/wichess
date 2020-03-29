package main

import (
	"database/sql"
	"net/http"

	"github.com/pciet/wichess/rules"
)

const BoardsPath = "/boards/"

var BoardsHandler = AuthenticRequestHandler{
	Get: GameIdentifierParsed(RequesterInGame(BoardsGet), BoardsPath),
}

type (
	BoardJSON map[rules.AddressIndex]PieceJSON

	PieceJSON struct {
		rules.Piece     `json:"p"`
		PieceIdentifier `json:"i,omitempty"`
	}
)

func BoardsGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx, id GameIdentifier) {
	b := LoadGameBoard(tx, id)
	tx.Commit()

	jr := make(BoardJSON)
	for i, s := range b.Board {
		if s.Kind == rules.NoKind {
			continue
		}
		jr[rules.AddressIndex(i)] = PieceJSON{rules.Piece(s), 0}
	}
	for _, pid := range b.PieceIdentifiers {
		a := pid.Address.Index()
		p := jr[a]
		p.PieceIdentifier = pid.ID
		jr[a] = p
	}

	JSONResponse(w, jr)
}
