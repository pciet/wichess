package memory

import (
	"io/ioutil"
	"strconv"
)

// PlayersFilePrefix is the first character in a player filename followed by the player identifier.
const PlayersFilePrefix = "p"

func writePlayerFile(id PlayerIdentifier) {
	p := LockPlayer(id)
	if p == nil {
		return
	}
	defer g.Unlock()

	b, err := json.Marshal(p)
	if err != nil {
		panic(err.Error())
	}

	f, err := os.OpenFile(filePath(PlayersilePrefix+id.String()),
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

func initializePlayersCaches(ids []PlayerIdentifier) {
	c := initializeHashCache()
	if c != len(ids) {
		panic(fmt.Sprint(HashFile,
			"entry count", c, "not equal to number of player files", len(ids)))
	}

	c = initializePlayerNameCaches()
	if c != len(ids) {
		panic(fmt.Sprint(PlayerFile,
			"entry count", c, "not equal to number of player files", len(ids)))
	}

	c = initializeSessionsCache()
	if c != len(ids) {
		panic(fmt.Sprint(SessionFile,
			"entry count", c, "not equal to number of player files", len(ids)))
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
