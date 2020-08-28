package wichess

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

// Login queries the database for the username and encrypted password.
// The password argument is the unencrypted password which is then re-encrypted for comparsion
// to the database value.
// If the credentials are correct then the player's ID in the players table and a new session
// key are returned.
// If the username doesn't exist in the database then it and the encrypted password are inserted.
func Login(name, password string) (PlayerIdentifier, string) {
	tx := DatabaseTransaction()
	defer tx.Commit()

	var id PlayerIdentifier
	var c string
	err := tx.QueryRow(PlayersCryptQuery, name).Scan(&id, &c)
	switch err {
	case sql.ErrNoRows:
		crypt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			Panic(err)
		}
		id = NewPlayer(tx, name, string(crypt))
	case nil:
		err = bcrypt.CompareHashAndPassword([]byte(c), []byte(password))
		if err != nil {
			return 0, ""
		}
	default:
		Panic(err)
	}

	return id, NewSession(tx, id)
}
