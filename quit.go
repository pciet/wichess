package wichess

import (
	"log"
	"net/http"

	"github.com/pciet/wichess/auth"
	"github.com/pciet/wichess/memory"
)

func quitGet(w http.ResponseWriter, r *http.Request, pid memory.PlayerIdentifier) {
	defer auth.ClearBrowserSession(w, r)

	sc, err := r.Cookie(auth.SessionKeyCookie)
	if err == http.ErrNoCookie {
		debug("auth accepted", pid, "but no session key cookie found")
		return
	} else if err != nil {
		log.Panicln(r.URL.Path, "failed to read session key cookie",
			auth.SessionKeyCookie, ":", err)
	}

	key := memory.SessionKeyFromBase64(sc.Value)
	if key == nil {
		debug("bad session key cookie value for quit")
		return
	}

	memory.EndSession(key)
}
