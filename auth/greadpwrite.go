package auth

import (
	"net/http"

	"github.com/pciet/wichess/game"
	"github.com/pciet/wichess/memory"
)

type PlayerWriteableGameReadableFunc func(http.ResponseWriter, *http.Request,
	game.Instance, *memory.Player)

// PlayerWriteableGameReadable verifies the requesting player is in the game and provides an
// a writeable Player and readable game instance.
func PlayerWriteableGameReadable(
	calls PlayerWriteableGameReadableFunc, pathPrefix string) HandlerFunc {

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

		p := memory.LockPlayer(pid)

		defer authRecover("\nPlayer", p, "\nGame\n", g)

		calls(w, r, g, p)

		p.Unlock()
	}
}
