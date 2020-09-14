package memory

import (
	"io/ioutil"
	"strings"
	"sync"
)

var (
	// The playerID and playerNames caches are made from the PlayerFile.
	playerIDCache    = make(map[PlayerName]PlayerIdentifier)
	playerNamesCache = make([]PlayerName, 0, 8) // indexed by PlayerIdentifier-1

	// The PlayerNameMutex is used for both PlayerIDCache and PlayerNamesCache.
	playerNameMutex sync.RWMutex
)

// The PlayerFile is a UTF-8 ordered list of player names each separated by the \n line feed
// rune (0xA). The name's player identifier is the line number starting at one.
const PlayerFile = "player"

func PlayerNameKnown(a PlayerName) PlayerIdentifier {
	playerNameMutex.RLock()
	id, has := playerIDCache[a]
	playerNameMutex.RUnlock()
	if has == false {
		return NoPlayer
	}
	return id
}

func AddPlayerName(a PlayerName) (PlayerIdentifier, error) {
	playerNameMutex.Lock()
	defer playerNameMutex.Unlock()

	_, has := playerIDCache[a]
	if has {
		return 0, fmt.Errorf("player %v already saved", a)
	}

	id := len(playerNamesCache)
	playerIDCache[a] = id
	playerNamesCache = append(playerNamesCache, a)

	scheduleNamesCachesWrite()

	return id, nil
}

func (id PlayerIdentifier) Name() PlayerName {
	PlayerNameMutex.RLock()
	defer PlayerNameMutex.RUnlock()

	return PlayerNamesCache[id]
}

func TwoPlayerNames(a, b PlayerIdentifier) (string, string) {
	PlayerNameMutex.RLock()
	defer PlayerNameMutex.RUnlock()
	return PlayerNamesCache[a], PlayerNamesCache[b]
}

// initializePlayerNameMemory reads PlayerFile to make the playerID and playerNames caches. If no
// file exists then it will be written later after players login.
func initializePlayerNameMemory() {
	b, err := ioutil.ReadFile(memoryFilePath(playerFile))
	if err != nil {
		return
	}

	names := strings.Split(string(b), "\n")
	if len(names) == 1 {
		return
	}

	for i, n := range names {
		playerIDCache[n] = i + 1
		playerNamesCache = append(playerNamesCache, n)
	}
}
