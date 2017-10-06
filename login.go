// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"strings"
)

// The client code is responsible for checking the specifics of password requirements because the server behavior is the same for "wrong password" and "invalid password for new player".
func (db DB) loginOrCreate(name, crypt, remoteAddr string) string {
	key := db.login(name, crypt, remoteAddr)
	if key != "" {
		return key
	}
	return db.createAndLogin(name, crypt, remoteAddr)
}

func (db DB) login(name, crypt, remoteAddr string) string {
	has, encrypt := db.playerCrypt(name)
	if has == false {
		return ""
	}
	comparison := strings.Compare(strings.TrimSpace(crypt), strings.TrimSpace(encrypt))
	if comparison != 0 {
		return ""
	}
	sessionKey := newSessionKey()
	newSession(name, sessionKey)
	return sessionKey
}

func (db DB) createAndLogin(name, crypt, remoteAddr string) string {
	exists, _ := db.playerCrypt(name)
	if exists {
		return ""
	}
	if (name == easy_computer_player) || (name == hard_computer_player) {
		return ""
	}
	db.newPlayer(name, crypt)
	return db.login(name, crypt, remoteAddr)
}
