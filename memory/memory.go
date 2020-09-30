// Package memory implements a memory heirarchy for saving Wisconsin Chess information that
// persists after the host is powered off. This information includes player names and hashed
// passwords, piece collections, active games, and more.
//
// Memory is fastest from volatile process memory caches (private vars in this package) and backed
// by files in the mem folder. Player requests never depend on file writes. All files are written
// when the process is ended, and game files are deleted when games end.
//
// This package also defines a variety of types used by the Wisconsin Chess host, including the
// Game and Player structs and some types for their fields like GameIdentifier and PlayerIdentifier.
package memory

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

// Folder is the name of the folder within the folder the host process is started in that holds all
// memory files.
const Folder = "mem"

// TODO: the use of mutex in this package has gotten complicated, consider a channels solution

// activeMutex is normally taken as a read so there's no waiting, but when a shutdown occurs the
// write lock is taken to be sure acquiring the other mutex doesn't cause a deadlock.
var activeMutex = sync.RWMutex{}

// Initialize reads specific files in Folder to build a process memory cache of host information.
//
// An operating system signal handling goroutine is started here that listens for os.Interrupt and
// syscall.SIGTERM then writes all pending file changes and calls os.Exit(0).
func Initialize() {
	files, err := ioutil.ReadDir(Folder)
	if err != nil {
		panic(err.Error())
	}

	gameIDs := make([]GameIdentifier, 0, 8)
	playerIDs := make([]PlayerIdentifier, 0, 8)
	for _, f := range files {
		name := f.Name()
		if strings.HasPrefix(name, GameFilePrefix) {
			var id GameIdentifier
			c, err := fmt.Sscanf(name, GameFilePrefix+"%d", &id)
			if err != nil {
				panic(err.Error())
			} else if c != 1 {
				log.Panicln("parsed", c)
			}
			gameIDs = append(gameIDs, id)
		} else if strings.HasPrefix(name, PlayersFilePrefix) {
			var id PlayerIdentifier
			c, err := fmt.Sscanf(name, PlayersFilePrefix+"%d", &id)
			if err != nil {
				panic(err.Error())
			} else if c != 1 {
				log.Panicln("parsed", c)
			}
			playerIDs = append(playerIDs, id)
		}
	}

	// initializePlayersCache also does the hash, names, and sessions caches
	initializePlayersCaches(playerIDs)
	initializeGamesCache(gameIDs)

	go func() {
		shutdown := make(chan os.Signal, 2)
		signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
		<-shutdown

		// stop new mutex acquisition before writing files
		activeMutex.Lock()

		writeGamesFiles()
		// TODO: could a game ack be lost by player gid being written, thus leaving orphan files
		writePlayerFiles()
		writePlayerNamesFile()
		writeHashFile()

		os.Exit(0)
	}()
}
