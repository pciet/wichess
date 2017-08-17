// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"strings"
)

// The client code is responsible for checking the specifics of password requirements because the server behavior is the same for "wrong password" and "invalid password for new player".
func loginOrCreate(name, crypt, remoteAddr string) string {
	key := login(name, crypt, remoteAddr)
	if key != "" {
		return key
	}
	return createAndLogin(name, crypt, remoteAddr)
}

func login(name, crypt, remoteAddr string) string {
	has, encrypt := playerCryptFromDatabase(name)
	if has == false {
		return ""
	}
	comparison := strings.Compare(strings.TrimSpace(crypt), strings.TrimSpace(encrypt))
	if comparison != 0 {
		return ""
	}
	sessionKey := newSessionKey()
	newSession(name, sessionKey, remoteAddr)
	return sessionKey
}

func createAndLogin(name, crypt, remoteAddr string) string {
	exists, _ := playerCryptFromDatabase(name)
	if exists {
		return ""
	}
	newPlayerInDatabase(name, crypt)
	return login(name, crypt, remoteAddr)
}
