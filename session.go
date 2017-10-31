// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"sync"
)

const (
	key_cookie = "k"

	key_length = 64
)

// TODO: to support multiple instances of the logic server this will have to be stored in the database

var (
	// map[name]key
	keys = map[string]string{}
	// map[key]name
	names = map[string]string{}

	sessionLock = sync.RWMutex{}
)

func validSession(r *http.Request) string {
	keyCookie, err := r.Cookie(key_cookie)
	if err != nil {
		return ""
	}
	sessionLock.RLock()
	defer sessionLock.RUnlock()
	_, has := names[keyCookie.Value]
	if has {
		return keyCookie.Value
	}
	return ""
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

func newSession(name, key string) {
	sessionLock.Lock()
	defer sessionLock.Unlock()
	oldkey, has := keys[name]
	if has {
		delete(names, oldkey)
	}
	keys[name] = key
	names[key] = name
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

func nameFromSessionKey(key string) string {
	sessionLock.RLock()
	defer sessionLock.RUnlock()
	return names[key]
}
