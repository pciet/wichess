package wichess

import (
	"net/http"

	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/piece"
)

type CollectionJSON struct {
	piece.Collection `json:"c"`
}

func collectionGet(w http.ResponseWriter, r *http.Request, p *memory.Player) {
	jsonResponse(w, CollectionJSON{p.Collection})
}
