package wichess

import (
	"net/http"

	"github.com/pciet/wichess/game"
	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/rules"
)

type PlayersJSON struct {
	White  memory.PlayerName `json:"w"`
	Black  memory.PlayerName `json:"b"`
	Active rules.Orientation `json:"a"`
}

func playersGet(w http.ResponseWriter, r *http.Request, g game.Instance) {
	wn, bn := memory.TwoPlayerNames(g.White.PlayerIdentifier, g.Black.PlayerIdentifier)
	jsonResponse(w, PlayersJSON{
		White:  wn,
		Black:  bn,
		Active: g.Active,
	})
}
