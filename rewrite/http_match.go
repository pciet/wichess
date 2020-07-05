package main

import (
	"database/sql"
	"net/http"
	"strconv"
)

const (
	MatchPath         = "/match"
	MatchHTMLTemplate = "html/match.tmpl"
)

var MatchHandler = AuthenticRequestHandler{
	Get: MatchGet,
}

func init() { ParseHTMLTemplate(MatchHTMLTemplate) }

type MatchHTMLTemplateData struct {
	Name            string
	RecentOpponents [RecentOpponentCount]string
}

func MatchGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx, requester Player) {
	peopleGameID := PlayerActivePeopleGame(tx, requester.ID)
	if peopleGameID != 0 {
		// if there's an active people game then the index webpage is disallowed until the
		// game is completed
		tx.Commit()
		http.Redirect(w, r,
			PeoplePath+strconv.Itoa(int(peopleGameID)), http.StatusTemporaryRedirect)
		return
	}

	opp := PlayerRecentOpponents(tx, requester.ID)
	tx.Commit()

	WriteHTMLTemplate(w, MatchHTMLTemplate, MatchHTMLTemplateData{
		Name:            requester.Name,
		RecentOpponents: opp,
	})
}
