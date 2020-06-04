package main

import (
	"database/sql"
	"net/http"
	"strconv"

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
	Name            string
	LeftPiece       int
	RightPiece      int
	Collection      [CollectionCount]rules.PieceKind
	RecentOpponents [RecentOpponentCount]string
}

func IndexGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx, requester Player) {
	peopleGameID := PlayerActivePeopleGame(tx, requester.ID)
	if peopleGameID != 0 {
		// if there's an active people game then the index webpage is disallowed until the
		// game is completed
		tx.Commit()
		http.Redirect(w, r,
			PeoplePath+strconv.Itoa(int(peopleGameID)), http.StatusTemporaryRedirect)
		return
	}

	left, right := PlayerPiecePicks(tx, requester.Name)
	collection := PlayerCollection(tx, requester.ID)
	opp := PlayerRecentOpponents(tx, requester.ID)
	tx.Commit()

	WriteHTMLTemplate(w, IndexHTMLTemplate, IndexHTMLTemplateData{
		Name:            requester.Name,
		LeftPiece:       int(left),
		RightPiece:      int(right),
		Collection:      collection.Kinds(),
		RecentOpponents: opp,
	})
}
