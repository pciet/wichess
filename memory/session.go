package memory

import (
	"crypto/rand"
	"math/big"
	"unicode"
	"unicode/utf8"
)

// SessionKeySize is the number of random runes in a session key.
const SessionKeySize = 16

// A player's SessionKey is an array of random runes. This key is encoded as a string before being
// communicated to the player's web browser encoded again into base64. Keys do not persist when
// the host is restarted.
type SessionKey [SessionKeySize]rune

// NoSessionKey is the zero value of a SessionKey and represents a bad key.
var NoSessionKey SessionKey

// NewSession randomly generates a key and saves it into volatile memory for the player.
func NewSession(of PlayerIdentifier) *SessionKey {
	key := newSessionKey()
	addSession(of, key)
	return key
}

// EndSession removes the supplied key from memory.
func EndSession(with *SessionKey) { removeSession(with) }

// SessionKeyFromString decodes the input string into a SessionKey. If the input string isn't a
// valid key then nil is returned.
func SessionKeyFromString(a string) *SessionKey {
	if utf8.RuneCountInString(a) != SessionKeySize {
		return nil
	}
	out := SessionKey{}
	// range over a string returns the byte index, not the rune index, so a separate i is used here
	i := 0
	for _, r := range a {
		out[i] = r
		i++
	}
	return &out
}

func (a *SessionKey) String() string { return string(a[:]) }

var maxSessionRune = big.NewInt(unicode.MaxRune)

func newSessionKey() *SessionKey {
	var key SessionKey
	for i := 0; i < SessionKeySize; i++ {
		bigIntP, err := rand.Int(rand.Reader, maxSessionRune)
		if err != nil {
			panic(err.Error())
		}
		key[i] = rune(bigIntP.Int64())
	}
	return &key
}
