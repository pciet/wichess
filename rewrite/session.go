package main

import (
	"net/http"
)

// TODO: are there login race conditions caused by having ValidSession not be part of the database transaction?
// TODO: validation per-connection instead of per-request to reduce ValidSession calls, maybe WebSocket rewrite replacing HTTP, or a non-browser client that uses UDP
// TODO: can the response be http.StatusNotAuthorized instead of http.StatusFound?

// Returns the player name associated with the session key in the request.
// If the session key cookie isn't present, is obviously wrong, or isn't in the database then the browser is redirected to /login and has the key cookie cleared.
func ValidSessionHandler(w http.ResponseWriter, r *http.Request) string {
	k, err := r.Cookie(key_cookie)
	if err == http.ErrNoCookie {
		DebugPrintln("no key cookie", key_cookie, "in HTTP request")
		ClearBrowserSession(w, r)
		return ""
	} else if err != nil {
		panic("failed to read key cookie", key_cookie, "from HTTP request:", err)
	}

	if len(k.Value) != key_length {
		DebugPrintln("key cookie", key_cookie, "length", len(k.Value), "not equal to required length", key_length)
		ClearBrowserSession(w, r)
		return ""
	}

	var name string
	err = Database.QueryRow(session_player_name_query, k.Value).Scan(&name)
	if err == sql.ErrNoRows {
		DebugPrintln("no row found for requested session key:", err)
		ClearBrowserSession(w, r)
		return ""
	} else if err != nil {
		panic("failed to query database:", err)
	} else if name == "" {
		panic("database row found for session key", k.Value, "but no name set")
	}

	return name
}

func ClearBrowserSession(w, r) {
	http.SetCookie(w, &http.Cookie{
		Name:     key_cookie,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // TODO: set true for TLS
	})
	http.Redirect(w, r, LoginRelPath, http.StatusFound)
}
