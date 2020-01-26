package main

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

// The username and password are assumed to be valid by this function.
func Login(name, password string) string {
	tx := DatabaseTransaction()
	defer CommitTransaction(tx)

	var c string
	err := tx.QueryRow(player_crypt_query, name).Scan(&c)
	if err == sql.ErrNoRows {
		c = CreateLogin(tx, name, password)
	} else if err != nil {
		panic(err)
	}

	// if the password is correct then invalidate the existing session if present
}

func CreateLogin(tx *sql.Tx, name, password string) string {
	crypt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
}
