// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
)

const (
	key_cookie = "k"

	key_length = 64

	session_table = "sessions"
	session_name  = "name"
	session_key   = "key"
)

// key, name
func (db DB) validSession(r *http.Request) (string, string) {
	keyCookie, err := r.Cookie(key_cookie)
	if err != nil {
		return "", ""
	}
	var name string
	err = db.QueryRow("SELECT "+session_name+" FROM "+session_table+" WHERE "+session_key+"=$1;", keyCookie.Value).Scan(&name)
	if err != nil {
		if debug {
			fmt.Println(err.Error())
		}
		return "", ""
	}
	return keyCookie.Value, name
}

func clearClientSession(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     key_cookie, // from web_login.go
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // TODO: set true after TLS certification
	})
}

func (db DB) newSession(name, key string) {
	tx := db.Begin()
	defer tx.Commit()
	var playerKey []byte
	err := tx.QueryRow("SELECT "+session_key+" FROM "+session_table+" WHERE "+session_name+"=$1 FOR UPDATE;", name).Scan(&playerKey)
	if err == nil {
		if string(playerKey) != key {
			_, err = tx.Exec("UPDATE "+session_table+" SET "+session_key+" =$1 WHERE "+session_name+" =$2;", []byte(key), name)
			if err != nil {
				panicExit(err.Error())
			}
		}
		return
	} else if err != sql.ErrNoRows {
		panicExit(err.Error())
	}
	_, err = tx.Exec("INSERT INTO "+session_table+"("+session_name+", "+session_key+") VALUES ($1, $2);", name, []byte(key))
	if err != nil {
		panicExit(err.Error())
	}
}

func newSessionKey() string {
	key := make([]byte, key_length)
	count, err := rand.Read(key)
	if err != nil {
		panicExit(err.Error())
	}
	if count != key_length {
		panicExit(fmt.Sprintf("count %v does not match key length %v", count, key_length))
	}
	return base64.StdEncoding.EncodeToString(key)
}
