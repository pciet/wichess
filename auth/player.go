package auth

import (
	"net/http"

	"github.com/pciet/wichess/memory"
)

type PlayerReadableFunc func(http.ResponseWriter, *http.Request, *memory.Player)

// PlayerReadable provides a pointer to player memory that should only be read.
func PlayerReadable(calls PlayerReadableFunc, pathPrefix string) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, pid memory.PlayerIdentifier) {
		p := memory.RLockPlayer(pid)
		calls(w, r, p)
		p.RUnlock()
	}
}

type PlayerWriteableFunc func(http.ResponseWriter, *http.Request, player.Instance)

// PlayerWriteable provides a pointer to player memory that can be written to.
func PlayerWriteable(calls PlayerWriteableFunc, pathPrefix string) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, pid memory.PlayerIdentifier) {
		p := memory.LockPlayer(pid)
		calls(w, r, p)
		p.Unlock()
	}
}
