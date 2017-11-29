// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"golang.org/x/crypto/bcrypt"
)

// The client code is responsible for checking the specifics of password requirements because the server behavior is the same for "wrong password" and "invalid password for new player".
func (db DB) loginOrCreate(name, password string) string {
	key := db.login(name, password)
	if key != "" {
		return key
	}
	return db.createAndLogin(name, password)
}

func (db DB) login(name, password string) string {
	has, encrypt := db.playerCrypt(name)
	if has == false {
		return ""
	}
	err := bcrypt.CompareHashAndPassword([]byte(encrypt), []byte(password))
	if err != nil {
		return ""
	}
	sessionKey := newSessionKey()
	db.newSession(name, sessionKey)
	return sessionKey
}

func (db DB) createAndLogin(name, password string) string {
	exists, _ := db.playerCrypt(name)
	if exists {
		return ""
	}
	if (name == easy_computer_player) || (name == hard_computer_player) {
		return ""
	}
	crypt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panicExit(err.Error())
	}
	db.newPlayer(name, string(crypt))
	return db.login(name, password)
}
