package auth

import (
	"net/http"

	"github.com/pciet/wichess/game"
	"github.com/pciet/wichess/memory"
)

type PlayerAndGameReadableFunc func(http.ResponseWriter, *http.Request,
	game.Instance, *memory.Player)

// PlayerAndGameReadable verifies the requesting player is in the game and provides instances
// that should only be read from.
func PlayerAndGameReadable(
	calls PlayerAndGameReadableFunc, pathPrefix string) HandlerFunc {

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

		p := memory.RLockPlayer(pid)

		defer authRecover("\nPlayer\n", p, "\nGame\n", g)

		calls(w, r, g, p)

		p.RUnlock()
	}
}
