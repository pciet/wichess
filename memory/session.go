package memory

import (
	"crypto/rand"
	"encoding/base64"
)

// SessionKeySize is the number of random bytes in a session key.
const SessionKeySize = 16

// A player's SessionKey is an array of random bytes. This key is encoded to base64 before being
// communicated to the player's web browser. Keys do not persist when the host is restarted.
type SessionKey [SessionKeySize]byte

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

// SessionKeyFromBase64 decodes the input base64 string into a SessionKey. Nil is returned if the
// input isn't valid.
func SessionKeyFromBase64(a string) *SessionKey {
	k, err := base64.StdEncoding.DecodeString(a)
	if (err != nil) || (len(k) != SessionKeySize) {
		// TODO: debug print this error
		return nil
	}
	var key SessionKey
	copy(key[:], k)
	return &key
}

func (a *SessionKey) Base64String() string { return base64.StdEncoding.EncodeToString(a[:]) }

func newSessionKey() *SessionKey {
	b := make([]byte, SessionKeySize)
	_, err := rand.Read(b)
	if err != nil {
		panic(err.Error())
	}
	var key SessionKey
	copy(key[:], b)
	return &key
}
