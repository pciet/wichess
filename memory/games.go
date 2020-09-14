package memory

import (
	"sync"
)

// TODO: assumptions for GamesCache's type are documented at github.com/pciet/wichess/issues/113

var (
	gamesCache = make(map[GameIdentifier]*Game)
	gamesMutex sync.RWMutex

	nextGameID GameIdentifier
)

// AddGame causes this game to be saved in the memory system.
func AddGame(g *Game) GameIdentifier {
	gamesMutex.Lock()
	g.GameIdentifier = nextGameID
	gamesCache[nextGameID] = g
	nextGameID++
	gamesMutex.Unlock()
	g.Changed()
	return id
}

// LockGame locks the indicated game from concurrent access and returns its memory. If the memory
// is changed then GameMemoryChanged is called to schedule a backing file update. *Game.Unlock,
// from the sync.RWMutex embedded in the struct, is called when access is done. If the game
// doesn't exist then nil is returned.
func LockGame(id GameIdentifier) *Game {
	gamesMutex.RLock()
	g := gamesCache[id]
	gamesMutex.RUnlock()
	if g != nil {
		g.Lock()
	}
	return g
}

// RLockGame is the same as LockGame except it does a read lock on the game's sync.RWMutex. When
// the caller is done reading from *Game then RUnlock is called.
func RLockGame(id GameIdentifier) *Game {
	gamesMutex.RLock()
	g := gamesCache[id]
	gamesMutex.RUnlock()
	if g != nil {
		g.RLock()
	}
	return g
}

// initializeGamesMemory reads all existing games from disk and puts them into the gamesCache.
func InitializeGamesMemory() {
	// TODO: set NextGameID based on the maximum filename
}
