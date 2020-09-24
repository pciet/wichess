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

	key := memory.SessionKeyFromString(sc.Value)
	if *key == memory.NoSessionKey {
		debug("quit called but no session found for", pid)
		return
	}

	memory.EndSession(key)
}
