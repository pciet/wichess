package wichess

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/pciet/wichess/piece"
)

const (
	IndexPath         = "/"
	IndexHTMLTemplate = "html/index.tmpl"
)

var IndexHandler = AuthenticRequestHandler{
	Get: IndexGet,
}

type IndexHTMLTemplateData struct {
	Name       string
	LeftPiece  int
	RightPiece int
	Collection [CollectionCount]piece.Kind
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
	tx.Commit()

	WriteHTMLTemplate(w, IndexHTMLTemplate, IndexHTMLTemplateData{
		Name:       requester.Name,
		LeftPiece:  int(left),
		RightPiece: int(right),
		Collection: collection.Kinds(),
	})
}
