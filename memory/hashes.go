package memory

import "sync"

var (
	hashCache = make([][]byte, 0, 8)
	hashMutex sync.RWMutex
)

// PlayerHash returns the password hash for a player, or nil if the player doesn't exist.
func PlayerHash(id PlayerIdentifier) []byte {
	activeMutex.RLock()
	hashMutex.RLock()

	if len(hashCache) < id {
		return nil
	}
	b := hashCache[id-1]

	hashMutex.RUnlock()
	activeMutex.RUnlock()

	return b
}
