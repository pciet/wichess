package memory

import "sync"

var (
	sessionCache = make(map[SessionKey]PlayerIdentifier)
	sessionMutex sync.RWMutex
)

// SessionPlayerIdentifier returns the id associated with the key or NoPlayer.
func SessionPlayerIdentifier(with *SessionKey) PlayerIdentifier {
	sessionMutex.RLock()
	defer sessionMutex.RUnlock()

	p, has := sessionCache[*with]
	if has == false {
		return NoPlayer
	}
	return p
}

func addSession(id PlayerIdentifier, k *SessionKey) {
	activeMutex.RLock()
	sessionMutex.Lock()

	sessionCache[*k] = id

	sessionMutex.Unlock()
	activeMutex.RUnlock()
}

func removeSession(k *SessionKey) {
	activeMutex.RLock()
	sessionMutex.Lock()

	defer func() {
		sessionMutex.Unlock()
		activeMutex.RUnlock()
	}()

	_, has := sessionCache[*k]
	if has == false {
		return
	}
	delete(sessionCache, *k)
}
