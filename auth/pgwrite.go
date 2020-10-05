package auth

import (
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/pciet/wichess/game"
	"github.com/pciet/wichess/memory"
)

type PlayerAndGameWriteableFunc func(http.ResponseWriter, *http.Request,
	game.Instance, *memory.Player)

// PlayerAndGameWriteable verifies the requesting player is in the game and provides instances
// that can be changed.
func PlayerAndGameWriteable(
	calls PlayerAndGameWriteableFunc, pathPrefix string) HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request, pid memory.PlayerIdentifier) {
		gid, err := parseURLGameIdentifier(r.URL.Path, pathPrefix)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		g := game.Lock(gid)
		if g.Nil() {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer g.Unlock()

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
			log.Println(pv, "\nPlayer\n", p, "\nGame\n", g)
			debug.PrintStack()
			os.Exit(1)
		}()

		calls(w, r, g, p)
		p.Unlock()
	}
}
