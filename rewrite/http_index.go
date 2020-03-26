package main

import (
	"database/sql"
	"net/http"
)

const (
	IndexPath        = "/"
	IndexWebTemplate = "web/html/index.tmpl"
)

var IndexHandler = AuthenticRequestHandler{
	Get: IndexGet,
}

func IndexGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx, requester string) {
	defer tx.Commit()
	/*
		if PlayerHasTimedGame(tx, name) {
			// TODO: don't use redirect
			http.Redirect(w, r, TimedGameRelPath, http.StatusFound)
			return
		}
	*/

	WriteIndexWebTemplate(IndexWebTemplateData{
		requester,
		LoadPlayerRecord(tx, requester),
		LoadPlayerFriendStatus(tx, requester),
	}, w)
}

type IndexWebTemplateData struct {
	Name string
	PlayerRecord
	PlayerFriendStatus
}

func WriteIndexWebTemplate(t IndexWebTemplateData, w http.ResponseWriter) {
	WriteWebTemplate(w, IndexWebTemplate, t)
}

func init() { ParseHTMLTemplate(IndexWebTemplate) }
