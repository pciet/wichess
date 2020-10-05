package auth

import (
	"log"
	"net/http"
	"os"
	"runtime/debug"

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

		defer func() {
			pv := recover()
			if pv == nil {
				return
			}
			log.Println(pv, "\nPlayer", p, "\nGame\n", g)
			debug.PrintStack()
			os.Exit(1)
		}()

		calls(w, r, g, p)
		p.Unlock()
	}
}
