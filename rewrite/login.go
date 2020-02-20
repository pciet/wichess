package main

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// The session key is returned normally, or an empty string is returned if the password is incorrect.
func Login(name, password string) string {
	tx := DatabaseTransaction()
	defer CommitTransaction(tx)

	var c string
	err := tx.QueryRow(player_crypt_query, name).Scan(&c)
	switch err {
	case sql.ErrNoRows:
		crypt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Panic(err)
		}
		NewPlayer(tx, name, string(crypt))
	case nil:
		err = bcrypt.CompareHashAndPassword([]byte(c), []byte(password))
		if err != nil {
			return ""
		}
	default:
		log.Panic(err)
	}

	return NewSession(tx, name)
}
