package auth

import (
	"net/http"

	"github.com/pciet/wichess/game"
	"github.com/pciet/wichess/memory"
)

type GameReadableFunc func(http.ResponseWriter, *http.Request, game.Instance)

// GameReadable verifies the requesting player is in the game and provides a game instance that
// should only be read from.
func GameReadable(calls GameReadableFunc, pathPrefix string) HandlerFunc {
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
		defer g.RUnlock()

		if g.HasPlayer(pid) == false {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		calls(w, r, g)
	}
}
