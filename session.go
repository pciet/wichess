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
)

func init() {
	keys = make(map[string]string)
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

func newSession(name, key, address string) {
	// invalidate previous key for name
	delete(sessions, keys[name])
	// add new session
	sessions[key] = address
	// set new key for name
	keys[name] = key
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
