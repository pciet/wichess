// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"net/http"
)

const key_cookie = "k"

// A session is active if this map contains a key with the http.Request RemoteAddr field matching the http request's.
var sessions map[string]string

func init() {
	sessions = make(map[string]string)
}

func validSession(r *http.Request) bool {
	keyCookie, err := r.Cookie(key_cookie)
	if err != nil {
		return false
	}
	addr, has := sessions[keyCookie.Value]
	if has == false {
		return false
	}
	if r.RemoteAddr != addr {
		delete(sessions, keyCookie.Value)
		return false
	}
	return true
}
