// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
)

const (
	key_cookie = "k"

	key_length = 64
)

var (
	// map[key]remoteAddr
	sessions map[string]string
	// map[name]key
	keys map[string]string
	// map[key]name
	names map[string]string
)

func init() {
	keys = make(map[string]string)
	sessions = make(map[string]string)
	names = make(map[string]string)
}

func validSession(r *http.Request) string {
	keyCookie, err := r.Cookie(key_cookie)
	if err != nil {
		return ""
	}
	key := keyCookie.Value
	addr, has := sessions[key]
	if has == false {
		return ""
	}
	if r.RemoteAddr != addr {
		delete(sessions, keyCookie.Value)
		return ""
	}
	return key
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

func newSession(name, key, address string) {
	// invalidate previous key for name
	delete(sessions, keys[name])
	// add new session
	sessions[key] = address
	// set new key for name
	keys[name] = key
	names[key] = name
}

func newSessionKey() string {
	key := make([]byte, key_length)
	count, err := rand.Read(key)
	if err != nil {
		panicExit(err.Error())
		return ""
	}
	if count != key_length {
		panicExit(fmt.Sprintf("count %v does not match key length %v", count, key_length))
		return ""
	}
	return base64.StdEncoding.EncodeToString(key)
}

func nameFromSessionKey(key string) string {
	return names[key]
}
