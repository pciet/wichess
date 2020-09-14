package wichess

import (
	"net/http"

	"github.com/pciet/wichess/game"
	"github.com/pciet/wichess/memory"
)

func alertGet(w http.ResponseWriter, r *http.Request, g game.Instance, id memory.PlayerIdentifier) {
	conn, err := webSocketUpgrade(w, r)
	if err != nil {
		debug(err)
		// the upgrade func in WebSocketUpgrade writes an error response, so nothing to add here
		return
	}
	game.Connect(g.GameIdentifier, g.OrientationOf(id), conn)
}
