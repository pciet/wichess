package memory

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// GameFilePrefix is the first characters in a game filename followed by the game identifier.
const GameFilePrefix = "gf"

func deleteGameFile(id GameIdentifier) {
	err := os.Remove(filePath(GameFilePrefix + id.String()))
	if os.IsNotExist(err) {
		return // if this game wasn't saved between reboots then there might not be a file
	} else if err != nil {
		panic(err.Error())
	}
}

func writeGamesFiles() {
	for id, g := range gamesCache {
		// handlers could be writing to a game even after getting through the cache lock
		g.RLock()

		b, err := json.Marshal(g)
		if err != nil {
			panic(err.Error())
		}
		f, err := os.OpenFile(filePath(GameFilePrefix+id.String()),
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)

		if err != nil {
			panic(err.Error())
		}

		_, err = f.Write(b)
		if err != nil {
			panic(err.Error())
		}

		err = f.Close()
		if err != nil {
			panic(err.Error())
		}

		g.RUnlock()
	}
}

func initializeGamesCache(ids []GameIdentifier) {
	for _, id := range ids {
		content, err := ioutil.ReadFile(filePath(GameFilePrefix + id.String()))
		if err != nil {
			panic(err.Error())
		}
		var g Game
		err = json.Unmarshal(content, &g)
		if err != nil {
			panic(err.Error())
		}
		(&g.Board).InitializePieces()
		gamesCache[id] = &g
		if id >= nextGameID {
			nextGameID = id + 1
		}
	}
}
