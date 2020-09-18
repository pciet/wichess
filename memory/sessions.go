package memory

import (
	"io/ioutil"
	"sync"
)

var (
	sessionCache = make(map[SessionKey]PlayerIdentifier)
	sessionMutex sync.RWMutex
)

func addSession(id PlayerIdentifier, k *SessionKey) {
	activeMutex.RLock()
	sessionsMutex.Lock()

	SessionsCache[*i] = id

	sessionsMutex.Unlock()
	activeMutex.RUnlock()
}

func removeSession(k *SessionKey) {
	activeMutex.RLock()
	sessionsMutex.Lock()

	_, has := SessionsCache[k]
	if has == false {
		panic("tried to remove nonexistent session key")
	}

	delete(SessionsCache, k)

	sessionsMutex.Unlock()
	activeMutex.RUnlock()
}
