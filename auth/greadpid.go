package auth

import (
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/pciet/wichess/game"
	"github.com/pciet/wichess/memory"
)

type GameReadablePlayerIdentifiedFunc func(http.ResponseWriter, *http.Request,
	game.Instance, memory.PlayerIdentifier)

// GameReadablePlayerIdentified is the same as GameReadable except the player id is an added arg.
func GameReadablePlayerIdentified(
	calls GameReadablePlayerIdentifiedFunc, pathPrefix string) HandlerFunc {

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

		defer func() {
			pv := recover()
			if pv == nil {
				return
			}
			log.Println(pv, "\nPlayer", pid, "\nGame\n", g)
			debug.PrintStack()
			os.Exit(1)
		}()

		calls(w, r, g, pid)
	}
}
