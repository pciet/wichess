package wichess

import (
	"net/http"

	"github.com/pciet/wichess/auth"
	"github.com/pciet/wichess/memory"
)

// FormPlayerName and FormPassword are the HTML form keys for LoginPath.
const (
	FormPlayerName = "name"
	FormPassword   = "password"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		loginAttemptHandler(w, r)
		return
	}

	if r.Method != http.MethodGet {
		debug(LoginPath, "HTTP method", r.Method, "not GET")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	writeHTMLTemplate(w, LoginHTMLTemplate, nil)
}

func loginAttemptHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		debug(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// TODO: verify README correctly describes username/password constraints
	name := r.FormValue(FormPlayerName)
	pass := r.FormValue(FormPassword)
	if (name == "") || (pass == "") {
		debug("missing username or password")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(name) > memory.PlayerNameMaxSize {
		debug("username too long")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if name == memory.ComputerPlayerName {
		debug("username matches", memory.ComputerPlayerName)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, key := auth.Login(memory.PlayerName(name), pass)
	if (*key == memory.NoSessionKey) || (id == memory.NoPlayer) {
		debug("bad password for", name)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	auth.CreateBrowserSession(w, key)
	http.Redirect(w, r, IndexPath, http.StatusSeeOther)
}
