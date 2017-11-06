// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"sync"
)

var gameLocks = map[int]*sync.RWMutex{}
var gameLocksLock = sync.Mutex{}

func rLockGame(id int) {
	gameLocksLock.Lock()
	lock, has := gameLocks[id]
	if has == false {
		lock = &sync.RWMutex{}
		gameLocks[id] = lock
	}
	gameLocksLock.Unlock()
	lock.RLock()
}

func rUnlockGame(id int) {
	gameLocksLock.Lock()
	lock := gameLocks[id]
	gameLocksLock.Unlock()
	lock.RUnlock()
}

func lockGame(id int) {
	gameLocksLock.Lock()
	lock, has := gameLocks[id]
	if has == false {
		lock = &sync.RWMutex{}
		gameLocks[id] = lock
	}
	gameLocksLock.Unlock()
	lock.Lock()
}

func unlockGame(id int) {
	gameLocksLock.Lock()
	lock := gameLocks[id]
	gameLocksLock.Unlock()
	lock.Unlock()
}

// TODO: verify that deleteGameLock is used correctly

// the caller is responsible for being sure no locking is needed during or after this function
func deleteGameLock(id int) {
	gameLocksLock.Lock()
	_, has := gameLocks[id]
	if has == false {
		gameLocksLock.Unlock()
		return
	}
	delete(gameLocks, id)
	gameLocksLock.Unlock()
}
