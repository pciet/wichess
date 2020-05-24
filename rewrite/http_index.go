package main

import (
	"database/sql"
	"net/http"

	"github.com/pciet/wichess/rules"
)

const (
	IndexPath         = "/"
	IndexHTMLTemplate = "html/index.tmpl"
)

var IndexHandler = AuthenticRequestHandler{
	Get: IndexGet,
}

func init() { ParseHTMLTemplate(IndexHTMLTemplate) }

type IndexHTMLTemplateData struct {
	Name       string
	LeftPiece  int
	RightPiece int
	Collection [CollectionCount]rules.PieceKind
}

func IndexGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx, requester Player) {
	left, right := PlayerPiecePicks(tx, requester.Name)
	collection := PlayerCollection(tx, requester.ID)
	tx.Commit()

	WriteHTMLTemplate(w, IndexHTMLTemplate, IndexHTMLTemplateData{
		Name:       requester.Name,
		LeftPiece:  int(left),
		RightPiece: int(right),
		Collection: collection.Kinds(),
	})
}
