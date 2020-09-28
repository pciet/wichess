package memory

import "sync"

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

	n := playerNamesCache[id-1]

	playerNameMutex.RUnlock()
	activeMutex.RUnlock()

	return n
}

// TwoPlayerNames reads the names of two players while only having to acquire the cache mutex once.
func TwoPlayerNames(a, b PlayerIdentifier) (PlayerName, PlayerName) {
	activeMutex.RLock()
	playerNameMutex.RLock()

	var na, nb PlayerName
	if a == ComputerPlayerIdentifier {
		na = ComputerPlayerName
		nb = playerNamesCache[b-1]
	} else if b == ComputerPlayerIdentifier {
		nb = ComputerPlayerName
		na = playerNamesCache[a-1]
	} else {
		na, nb = playerNamesCache[a-1], playerNamesCache[b-1]
	}

	playerNameMutex.RUnlock()
	activeMutex.RUnlock()

	return na, nb
}
