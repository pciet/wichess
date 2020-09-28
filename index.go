package wichess

import (
	"net/http"

	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/piece"
)

type IndexHTMLTemplateData struct {
	Name        string
	Left, Right int // piece.Kind is templated as a string, so using int instead
	piece.Collection
}

func indexGet(w http.ResponseWriter, r *http.Request, p *memory.Player) {
	if handleInPeopleGame(w, r, p) {
		return
	}
	writeHTMLTemplate(w, IndexHTMLTemplate, IndexHTMLTemplateData{
		Name:       p.PlayerName.String(),
		Left:       int(p.Left),
		Right:      int(p.Right),
		Collection: p.Collection,
	})
}
