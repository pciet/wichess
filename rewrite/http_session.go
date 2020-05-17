package main

import (
	"net/http"
	"strconv"
)

const (
	PlayerIDCookie   = "p"
	SessionKeyCookie = "k"
)

func CreateBrowserSession(w http.ResponseWriter, playerID int, sessionKey string) {
	http.SetCookie(w, &http.Cookie{
		Name:     PlayerIDCookie,
		Value:    strconv.Itoa(playerID),
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     SessionKeyCookie,
		Value:    sessionKey,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // TODO: set true for TLS
	})
}

func ClearBrowserSession(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     PlayerIDCookie,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     SessionKeyCookie,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false, // TODO: set true for TLS
	})
	http.Redirect(w, r, LoginPath, http.StatusSeeOther)
}
