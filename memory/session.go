package memory

import (
	"crypto/rand"
	"math/big"
	"unicode"
)

const SessionKeySize = 16

// A player's SessionKey is an array of random runes. This key is encoded as a string before being
// communicated to the player's web browser.
type SessionKey [SessionKeySize]rune

// NoSessionKey is the zero value of a SessionKey and represents a bad key.
var NoSessionKey SessionKey

var maxSessionRune = big.Int(unicode.MaxRune)

func NewSessionKey() SessionKey {
	var key SessionKey
	for i := 0; i < SessionKeySize; i++ {
		bigIntP, err := rand.Int(rand.Reader, &maxSessionRune)
		if err != nil {
			panic(err.Error())
		}
		key[i] = rune(bigIntP.Int64())
	}
	return key
}

func NewSession() SessionKey {

}

func (a SessionKey) String() string { return string(a) }
