package memory

import "sync"

// TODO: assumptions for GamesCache's type are documented at github.com/pciet/wichess/issues/113

var (
	gamesCache = make(map[GameIdentifier]*Game)
	gamesMutex sync.RWMutex

	nextGameID GameIdentifier
)

// AddGame causes this game to be saved in the memory system.
func AddGame(g *Game) GameIdentifier {
	activeMutex.RLock()
	gamesMutex.Lock()

	g.GameIdentifier = nextGameID
	gamesCache[nextGameID] = g
	id := nextGameID
	nextGameID++

	gamesMutex.Unlock()
	activeMutex.RUnlock()

	return id
}

// LockGame locks the indicated game from concurrent access and returns its memory. The Unlock
// method of the embedded Game sync.RWMutex should be called when access is done.
//
// If the game doesn't exist then nil is returned.
func LockGame(id GameIdentifier) *Game {
	activeMutex.RLock()
	gamesMutex.RLock()

	g := gamesCache[id]

	gamesMutex.RUnlock()
	activeMutex.RUnlock()

	if g != nil {
		g.Lock()
	}
	return g
}

// RLockGame is the same as LockGame except it does a read lock on the game's sync.RWMutex. When
// the caller is done reading then the RUnlock method should be called.
func RLockGame(id GameIdentifier) *Game {
	activeMutex.RLock()
	gamesMutex.RLock()

	g := gamesCache[id]

	gamesMutex.RUnlock()
	activeMutex.RUnlock()

	if g != nil {
		g.RLock()
	}
	return g
}
