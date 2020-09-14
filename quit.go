package wichess

import (
	"net/http"

	"github.com/pciet/wichess/auth"
	"github.com/pciet/wichess/memory"
)

func quitGet(w http.ResponseWriter, r *http.Request, pid memory.PlayerIdentifier) {
	memory.EndSession(pid)
	auth.ClearBrowserSession(w, r)
}
