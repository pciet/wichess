package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"net/http"
)

const (
	SessionKeyCookie = "k"
	SessionKeyLength = 64
)

// SessionRequester queries the database for a username that matches the key.
// If none is found then an empty string is returned.
func SessionRequester(tx *sql.Tx, sessionKey string) string {
	var name string
	err = tx.QueryRow(SessionNameQuery, sessionKey).Scan(&name)
	if err == sql.ErrNoRows {
		return ""
	} else if err != nil {
		Panic("failed to query database:", err)
	}
	return name
}

// NewSession returns a new session key that is also inserted into the database.
func NewSession(tx *sql.Tx, name string) string {
	k := make([]byte, SessionKeyLength)
	count, err := rand.Read(k)
	if err != nil {
		Panic(err)
	}
	if count != SessionKeyLength {
		Panic("count", count, "not equal to key length", SessionKeyLength)
	}
	key := base64.StdEncoding.EncodeToString(k)

	var s sql.NullString // QueryRow panics without this argument
	err = tx.QueryRow(SessionExistsQuery, name).Scan(&s)
	if err == sql.ErrNoRows {
		_, err := tx.Exec(SessionInsert, name, []byte(key))
		if err != nil {
			Panic(err)
		}
		return key
	} else if err != nil {
		Panic(err)
	}

	_, err = tx.Exec(SessionUpdate, []byte(key), name)
	if err != nil {
		Panic(err)
	}

	return key
}

func ClearBrowserSession(w http.ResponseWriter, r *http.Request) {
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
