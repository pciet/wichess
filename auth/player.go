package auth

import (
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/pciet/wichess/memory"
)

type PlayerReadableFunc func(http.ResponseWriter, *http.Request, *memory.Player)

// PlayerReadable provides a pointer to player memory that should only be read.
func PlayerReadable(calls PlayerReadableFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, pid memory.PlayerIdentifier) {
		p := memory.RLockPlayer(pid)

		defer func() {
			pv := recover()
			if pv == nil {
				return
			}
			log.Println(pv, "\nPlayer", pid, "\n", p)
			debug.PrintStack()
			os.Exit(1)
		}()

		calls(w, r, p)
		p.RUnlock()
	}
}

type PlayerWriteableFunc func(http.ResponseWriter, *http.Request, *memory.Player)

// PlayerWriteable provides a pointer to player memory that can be written to.
func PlayerWriteable(calls PlayerWriteableFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, pid memory.PlayerIdentifier) {
		p := memory.LockPlayer(pid)

		defer func() {
			pv := recover()
			if pv == nil {
				return
			}
			log.Println(pv, "\nPlayer", pid, "\n", p)
			debug.PrintStack()
			os.Exit(1)
		}()

		calls(w, r, p)
		p.Unlock()
	}
}
