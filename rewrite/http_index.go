package main

import (
	"database/sql"
	"net/http"
)

const (
	IndexPath         = "/"
	IndexHTMLTemplate = "web/html/index.tmpl"
)

var IndexHandler = AuthenticRequestHandler{
	Get: IndexGet,
}

func init() { ParseHTMLTemplate(IndexHTMLTemplate) }

type IndexHTMLTemplateData struct {
	Name string
	PlayerRecord
	PlayerFriendStatus
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

	WriteHTMLTemplate(w, IndexHTMLTemplate, IndexHTMLTemplateData{
		requester,
		LoadPlayerRecord(tx, requester),
		LoadPlayerFriendStatus(tx, requester),
	})
}
