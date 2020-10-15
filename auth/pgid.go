package auth

import (
	"net/http"

	"github.com/pciet/wichess/game"
	"github.com/pciet/wichess/memory"
)

type GameAndPlayerIdentifiedFunc func(http.ResponseWriter, *http.Request,
	memory.GameIdentifier, memory.PlayerIdentifier)

// GameAndPlayerIdentified verifies the requesting player is in the game then only provides the
// identifiers of the game and player.
func GameAndPlayerIdentified(
	calls GameAndPlayerIdentifiedFunc, pathPrefix string) HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request, pid memory.PlayerIdentifier) {
		gid, err := parseURLGameIdentifier(r.URL.Path, pathPrefix)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		g := game.RLock(gid)
		if g.Nil() {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if g.HasPlayer(pid) == false {
			w.WriteHeader(http.StatusBadRequest)
			g.RUnlock()
			return
		}
		g.RUnlock()

		defer authRecover("\nPlayer", pid, "\nGame", gid)

		calls(w, r, gid, pid)
	}
}
