package auth

import (
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/pciet/wichess/memory"
)

// Login determines if the named player already has an account and if the provided password's hash
// matches the saved hash. If the name is new then a new account is created. If there's a password
// mismatch then memory.NoPlayer is returned, and the second return is the base64 encoded
// session key.
func Login(name memory.PlayerName, password string) (memory.PlayerIdentifier, string) {
	id := memory.PlayerNameKnown(name)
	if id == memory.NoPlayer {
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("bcrypt.GenerateFromPassword failed:", err)
			return memory.NoPlayer, ""
		}
		id = memory.NewPlayer(name, hash)
	} else {
		err := bcrypt.CompareHashAndPassword(memory.PlayerHash(id), []byte(password))
		if err != nil {
			return memory.NoPlayer, ""
		}
	}
	return id, memory.NewSession(id)
}
