package main

import (
	"database/sql"
	"net/http"
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
}

func IndexGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx, requester Player) {
	left, right := PlayerPiecePicks(tx, requester.Name)
	tx.Commit()

	WriteHTMLTemplate(w, IndexHTMLTemplate, IndexHTMLTemplateData{
		Name:       requester.Name,
		LeftPiece:  int(left),
		RightPiece: int(right),
	})
}
