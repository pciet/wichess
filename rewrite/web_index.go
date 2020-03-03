package main

import (
	"net/http"
)

const (
	IndexRelPath = "/"

	index_web_template = "web/html/index.tmpl"
)

func init() { ParseHTMLTemplate(index_web_template) }

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		DebugPrintln("attempt to access unsupported path", r.URL.Path)
		// TODO: what should this status be?
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if r.Method != http.MethodGet {
		DebugPrintln(IndexRelPath, "HTTP method", r.Method, "not GET")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	name := ValidSessionHandler(w, r)
	if name == "" {
		return
	}

	tx := DatabaseTransaction()
	defer CommitTransaction(tx)

	/*
		if PlayerHasTimedGame(tx, name) {
			http.Redirect(w, r, TimedGameRelPath, http.StatusFound)
			return
		}
	*/

	WriteIndexWebTemplate(w, IndexWebTemplateData{name, LoadPlayerRecord(tx, name), LoadPlayerFriendStatus(tx, name)})
}

type IndexWebTemplateData struct {
	Name string
	PlayerRecord
	PlayerFriendStatus
}

func WriteIndexWebTemplate(w http.ResponseWriter, t IndexWebTemplateData) {
	WriteWebTemplate(w, index_web_template, t)
}
