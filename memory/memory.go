// Package memory implements a memory heirarchy for saving Wisconsin Chess information that
// persists after the host is powered off. This information includes player names and hashed
// passwords, web browser session keys, piece collections, active games, and more.
//
// Memory is fastest from volatile process memory caches (private vars in this package) and backed
// by files in the mem folder. File changes are tried to be written only when the host process is
// idle to minimize perceptible added time to gameplay actions. Player requests also never depend
// on a file write.
package memory

import (
	"os"
	"os/syscall"
)

// Folder is the name of the folder within the folder the host process is started in that holds all
// files used to save information across power off of the host computer.
const Folder = "mem"

// Initialize reads specific files in Folder to build a process memory cache of host information.
// As players interact memory is written back to files or read from files, and shutdown signals
// are listened for to guarantee files are written with the latest state for all players.
func Initialize() {
	initializePlayerNameCache()
	initializeSessionCache()

	continuouslyWriteGameFiles(continuouslySignalIdle(), listenForShutdownSignal())
}

func filePath(filename string) string { return Folder + "/" + filename }

type signal struct{}

func listenForShutdownSignal() <-chan signal {
	out := make(chan signal)
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		out <- Signal{}
	}()
	return out
}
