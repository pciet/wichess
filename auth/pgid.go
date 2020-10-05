package auth

import (
	"log"
	"net/http"
	"os"
	"runtime/debug"

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

		defer func() {
			pv := recover()
			if pv == nil {
				return
			}
			log.Println(pv, "\nPlayer", pid, "\nGame", gid)
			debug.PrintStack()
			os.Exit(1)
		}()

		calls(w, r, gid, pid)
	}
}
