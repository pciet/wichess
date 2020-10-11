package memory

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// PlayersFilePrefix is the first characters in a player filename followed by the player identifier.
const PlayersFilePrefix = "pf"

// WritePlayerFile creates or replaces the file that backs the identified player's process memory.
func WritePlayerFile(id PlayerIdentifier) {
	p := RLockPlayer(id)
	if p == nil {
		return
	}
	p.writePlayerFile()
	p.RUnlock()
}

func (a *Player) writePlayerFile() {
	if a == nil {
		return
	}
	b, err := json.Marshal(a)
	if err != nil {
		panic(err.Error())
	}

	f, err := os.OpenFile(filePath(PlayersFilePrefix+a.PlayerIdentifier.String()),
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

func writePlayerFiles() {
	for _, p := range playersCache {
		p.writePlayerFile()
	}
}

func initializePlayersCaches(ids []PlayerIdentifier) {
	c := initializeHashCache()
	if c != len(ids) {
		log.Panicln(HashFile, "entry count", c, "not equal to number of player files", len(ids))
	}
	nextPlayerID = PlayerIdentifier(c + 1)

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
