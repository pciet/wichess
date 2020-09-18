package memory

import (
	"io/ioutil"
	"strings"
	"sync"
)

var (
	playerIDCache    = make(map[PlayerName]PlayerIdentifier)
	playerNamesCache = make([]PlayerName, 0, 8) // indexed by PlayerIdentifier-1

	// The PlayerNameMutex is used for both PlayerIDCache and PlayerNamesCache.
	playerNameMutex sync.RWMutex
)

// PlayerNameKnown returns the id for the player name or NoPlayer if it's not a saved name.
func PlayerNameKnown(a PlayerName) PlayerIdentifier {
	activeMutex.RLock()
	playerNameMutex.RLock()

	id, has := playerIDCache[a]

	playerNameMutex.RUnlock()
	activeMutex.RUnlock()

	if has == false {
		return NoPlayer
	}
	return id
}

func (id PlayerIdentifier) Name() PlayerName {
	activeMutex.RLock()
	playerNameMutex.RLock()

	n := playerNamesCache[id]

	playerNameMutex.RUnlock()
	activeMutex.RUnlock()

	return n
}

// TwoPlayerNames reads the names of two players while only having to acquire the cache mutex once.
func TwoPlayerNames(a, b PlayerIdentifier) (string, string) {
	activeMutex.RLock()
	playerNameMutex.RLock()

	na, nb := playerNamesCache[a], playerNamesCache[b]

	playerNameMutex.RUnlock()
	activeMutex.RUnlock()

	return na, nb
}
