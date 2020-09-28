package memory

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// PlayersFilePrefix is the first characters in a player filename followed by the player identifier.
const PlayersFilePrefix = "pf"

func writePlayerFiles() {
	for id, p := range playersCache {
		b, err := json.Marshal(p)
		if err != nil {
			panic(err.Error())
		}

		f, err := os.OpenFile(filePath(PlayersFilePrefix+id.String()),
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
	}
}

func initializePlayersCaches(ids []PlayerIdentifier) {
	c := initializeHashCache()
	if c != len(ids) {
		log.Panicln(HashFile, "entry count", c, "not equal to number of player files", len(ids))
	}

	c = initializePlayerNameCaches()
	if c != len(ids) {
		log.Panicln(PlayerFile, "entry count", c, "not equal to number of player files", len(ids))
	}

	for _, id := range ids {
		content, err := ioutil.ReadFile(filePath(PlayersFilePrefix + id.String()))
		if err != nil {
			panic(err.Error())
		}
		var p Player
		err = json.Unmarshal(content, &p)
		if err != nil {
			panic(err.Error())
		}
		playersCache[id] = &p
	}
}
