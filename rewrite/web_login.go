package main

import (
	"net/http"
)

// TODO: make key cookie secure

const (
	LoginRelPath = "/login"

	login_web_template = "web/html/login.html"

	form_player_name = "name"
	form_password    = "password"

	name_max_length = 64
)

func init() { ParseHTMLTemplate(login_web_template) }

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		LoginAttemptHandler(w, r)
		return
	}

	if r.Method != http.MethodGet {
		DebugPrintln(LoginRelPath, "HTTP method", r.Method, "not GET")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	WriteWebTemplate(w, login_web_template, nil)
}

func LoginAttemptHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		DebugPrintln(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// TODO: verify README correctly describes username/password constraints
	name := r.FormValue(form_player_name)
	pass := r.FormValue(form_password)
	if (name == "") || (pass == "") {
		DebugPrintln("missing username or password")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(name) > name_max_length {
		DebugPrintln("username too long")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if name == computer_player_name {
		DebugPrintln("username matches", computer_player_name)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	key := Login(name, pass)
	if key == "" {
		panic("Login returned empty key")
	}
	http.SetCookie(w, &http.Cookie{
		Name:     key_cookie,
		Value:    key,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // TODO: set true for TLS
	})
	// TODO: is http.StatusFound the right code?
	http.Redirect(w, r, IndexRelPath, http.StatusFound)
}