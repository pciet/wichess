package wichess

import (
	"net/http"

	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/piece"
)

type PicksJSON struct {
	Left  piece.Kind `json:"l"`
	Right piece.Kind `json:"r"`
}

func picksGet(w http.ResponseWriter, r *http.Request, p *memory.Player) {
	jsonResponse(w, PicksJSON{p.Left, p.Right})
}
