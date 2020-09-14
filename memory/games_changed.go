package memory

import (
	"sync"
	"time"
)

var (
	gameFileWrites      = make([]GameIdentifier, 0, 8)
	gameFileWritesMutex sync.Mutex
)

// If longWait passes without the game being idle then a game file is written anyway.
const longWait = 2 * time.Minute

// continuouslyWriteGameFiles rewrites files with new information stored in gamesCache periodically.
// A file write only starts if no player actions are happening or longWait has passed.
func continuouslyWriteGameFiles(idle, shutdown chan signal) {
	timer := time.NewTimer(longWait)
	for {
		select {
		case <-idle:
			if timer.Stop() == false {
				<-timer.C
			}
			fallthrough
		case <-timer.C:
			timer.Reset(longWait)
			writeOneGameFile()
		case <-shutdown:
			// TODO: stop new player actions, write all files, then os.Exit
		}
	}
}

func writeOneGameFile() {
	gameFilesWriteMutex.Lock()
	if len(gameFileWrites) == 0 {
		gameFilesWriteMutex.Unlock()
		return
	}

	id := gameFileWrites[0]
	if len(gameFileWrites) == 1 {
		gameFileWrites = gameFileWrites[:0]
	} else {
		gameFileWrites = gameFileWrites[1:]
	}
	gameFilesWriteMutex.Unlock()

	writeGameFile(id)
}

func writeGameFile(id GameIdentifier) {

}
