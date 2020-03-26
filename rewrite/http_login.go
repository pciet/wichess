package main

import "net/http"

// TODO: make key cookie secure

const (
	LoginPath        = "/login"
	LoginWebTemplate = "web/html/login.tmpl"

	FormPlayerName = "name"
	FormPassword   = "password"

	NameMaxLength = 64
)

func init() { ParseHTMLTemplate(LoginWebTemplate) }

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		LoginAttemptHandler(w, r)
		return
	}

	if r.Method != http.MethodGet {
		DebugPrintln(LoginPath, "HTTP method", r.Method, "not GET")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	WriteWebTemplate(w, LoginWebTemplate, nil)
}

func LoginAttemptHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		DebugPrintln(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// TODO: verify README correctly describes username/password constraints
	name := r.FormValue(FormPlayerName)
	pass := r.FormValue(FormPassword)
	if (name == "") || (pass == "") {
		DebugPrintln("missing username or password")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(name) > NameMaxLength {
		DebugPrintln("username too long")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if name == ComputerPlayerName {
		DebugPrintln("username matches", ComputerPlayerName)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	key := Login(name, pass)
	if key == "" {
		DebugPrintln("bad password for", name)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     SessionKeyCookie,
		Value:    key,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // TODO: set true for TLS
	})
	http.Redirect(w, r, IndexRelPath, http.StatusSeeOther)
}
