package main

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

// Login queries the database for the username and encrypted password.
// The password argument is the unencrypted password which is then re-encrypted for
// comparsion to the database value.
// If the credentials are correct then a session key is returned, otherwise an empty
// string is returned.
func Login(name, password string) string {
	tx := DatabaseTransaction()
	defer tx.Commit()

	var c string
	err := tx.QueryRow(PlayerCryptQuery, name).Scan(&c)
	switch err {
	case sql.ErrNoRows:
		crypt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			Panic(err)
		}
		NewPlayer(tx, name, string(crypt))
	case nil:
		err = bcrypt.CompareHashAndPassword([]byte(c), []byte(password))
		if err != nil {
			return ""
		}
	default:
		Panic(err)
	}

	return NewSession(tx, name)
}
