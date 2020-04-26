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
	Name string
}

func IndexGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx, requester string) {
	tx.Commit()
	WriteHTMLTemplate(w, IndexHTMLTemplate, IndexHTMLTemplateData{requester})
}
