package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"log"
	"net/http"
)

// TODO: are there login race conditions caused by having ValidSession not be part of the database transaction?
// TODO: validation per-connection instead of per-request to reduce ValidSession calls, maybe WebSocket rewrite replacing HTTP, or a non-browser client that uses UDP
// TODO: can the response be http.StatusNotAuthorized instead of http.StatusFound?
// TODO: what is the cost of comparing the entire session cookie string to the database table in every request?

// Returns the player name associated with the session key in the request.
// If the session key cookie isn't present or isn't in the database then the browser is redirected to /login and has the key cookie cleared.
func ValidSessionHandler(w http.ResponseWriter, r *http.Request) string {
	k, err := r.Cookie(session_key_cookie)
	if err == http.ErrNoCookie {
		ClearBrowserSession(w, r)
		return ""
	} else if err != nil {
		log.Panicln("failed to read key cookie", session_key_cookie, "from HTTP request:", err)
	}

	if len(k.Value) == 0 {
		DebugPrintln("key cookie", session_key_cookie, "length zero")
		ClearBrowserSession(w, r)
		return ""
	}

	var name string
	err = Database.QueryRow(session_name_query, k.Value).Scan(&name)
	if err == sql.ErrNoRows {
		DebugPrintln("no row found for requested session key:", err)
		ClearBrowserSession(w, r)
		return ""
	} else if err != nil {
		log.Panicln("failed to query database:", err)
	} else if name == "" {
		log.Panicln("database row found for session key", k.Value, "but no name set")
	}

	return name
}

// Returns the session key.
func NewSession(tx *sql.Tx, name string) string {
	k := make([]byte, session_key_length)
	count, err := rand.Read(k)
	if err != nil {
		log.Panic(err)
	}
	if count != session_key_length {
		log.Panicln("count", count, "not equal to key length", session_key_length)
	}
	key := base64.StdEncoding.EncodeToString(k)

	var s sql.NullString // no rows are selected but the SQL library panics without this argument
	err = tx.QueryRow(session_exists_query, name).Scan(&s)
	if err == sql.ErrNoRows {
		_, err := tx.Exec(session_insert, name, []byte(key))
		if err != nil {
			log.Panic(err)
		}
		return key
	} else if err != nil {
		log.Panic(err)
	}

	_, err = tx.Exec(session_update, []byte(key), name)
	if err != nil {
		log.Panic(err)
	}

	return key
}

func ClearBrowserSession(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     session_key_cookie,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false, // TODO: set true for TLS
	})
	http.Redirect(w, r, LoginRelPath, http.StatusFound)
}
