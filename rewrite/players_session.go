package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
)

const SessionKeyLength = 64

// NewSession generates a unique session key, updates the player's database row with it, and
// returns it.
func NewSession(tx *sql.Tx, id PlayerIdentifier) string {
	k := make([]byte, SessionKeyLength)
	count, err := rand.Read(k)
	if err != nil {
		Panic(err)
	}
	if count != SessionKeyLength {
		Panic("count", count, "not equal to key length", SessionKeyLength)
	}
	key := base64.StdEncoding.EncodeToString(k)

	_, err = tx.Exec(PlayersSessionUpdate, []byte(key), id)
	if err != nil {
		Panic(err)
	}

	return key
}

// PlayersSessionKey queries the database for the player's current session key. If the column
// is null (no session) or the player doesn't exist then an empty string is returned.
func PlayersSessionKey(tx *sql.Tx, id PlayerIdentifier) string {
	var key sql.NullString
	err := tx.QueryRow(PlayersSessionQuery, id).Scan(&key)
	if err == sql.ErrNoRows {
		return ""
	} else if err != nil {
		DebugPrintln(id)
		Panic(err)
	}
	if key.Valid == false {
		return ""
	}
	return key.String
}

// EndSession sets the player's row's session column to null.
func EndSession(tx *sql.Tx, id PlayerIdentifier) {
	_, err := tx.Exec(PlayersSessionUpdate, nil, id)
	if err != nil {
		Panic(err)
	}
}
