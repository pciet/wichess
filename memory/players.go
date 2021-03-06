package memory

import (
	"sync"
	"unicode"

	"github.com/pciet/wichess/piece"
)

var (
	playersCache = make(map[PlayerIdentifier]*Player)
	playersMutex sync.RWMutex

	nextPlayerID = PlayerIdentifier(1)
)

// LockPlayer and RLockPlayer work similarly to LockGame and RLockGame.
func LockPlayer(id PlayerIdentifier) *Player {
	activeMutex.RLock()
	playersMutex.RLock()

	p := playersCache[id]

	playersMutex.RUnlock()
	activeMutex.RUnlock()

	if p != nil {
		p.Lock()
	}
	return p
}

func RLockPlayer(id PlayerIdentifier) *Player {
	activeMutex.RLock()
	playersMutex.RLock()

	p := playersCache[id]

	playersMutex.RUnlock()
	activeMutex.RUnlock()

	if p != nil {
		p.RLock()
	}
	return p
}

// NewPlayer is the only place where multiple mutex are locked (ignoring activeMutex), so another
// mutex is added on top of those to guarantee there's no deadlock from concurrent NewPlayer calls.
var newPlayerMutex = sync.Mutex{}

// NewPlayer saves a new player into the package memory caches and later to the files. After this
// function returns LockPlayer and RLockPlayer can be used with the returned id.
//
// The provided name is verified to match the PlayerName requirement of unicode.IsGraphic returning
// true for all runes.
func NewPlayer(name PlayerName, passwordHash []byte) PlayerIdentifier {
	for _, r := range name {
		if unicode.IsGraphic(r) == false {
			return NoPlayer
		}
	}

	p := Player{
		PlayerIdentifier: nextPlayerID,
		PlayerName:       name,
		Left:             piece.RandomSpecialKind(),
		Right:            piece.RandomSpecialKind(),
	}

	activeMutex.RLock()
	newPlayerMutex.Lock()
	playersMutex.Lock()
	hashMutex.Lock()
	playerNameMutex.Lock()

	id := nextPlayerID
	nextPlayerID++

	playersCache[id] = &p
	hashCache = append(hashCache, passwordHash)
	playerIDCache[name] = id

	if len(playerNamesCache) < int(id) {
		// not quite append because the index has to be specifically the id
		if cap(playerNamesCache) < int(id) {
			newCache := make([]PlayerName, len(playerNamesCache), len(playerNamesCache)*2)
			copy(newCache, playerNamesCache)
			playerNamesCache = newCache
		}
		playerNamesCache = playerNamesCache[:id]
	}
	playerNamesCache[id-1] = name

	playerNameMutex.Unlock()
	hashMutex.Unlock()
	playersMutex.Unlock()
	newPlayerMutex.Unlock()
	activeMutex.RUnlock()

	// write backing files here to reduce impact of a bad shutdown
	go func() {
		// TODO: if concurrent NewPlayer the player and hash file should be one write
		activeMutex.RLock()
		playerNameMutex.RLock()
		hashMutex.RLock()

		writePlayerNamesFile()
		writeHashFile()
		WritePlayerFile(id)

		hashMutex.RUnlock()
		playerNameMutex.RUnlock()
		activeMutex.RUnlock()
	}()

	return id
}
