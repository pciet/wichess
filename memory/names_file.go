package memory

import (
	"io/ioutil"
	"os"
	"strings"
)

// The PlayerFile is a UTF-8 ordered list of player names each separated by the \n line feed
// rune (0xA). The name's player identifier is the line number starting at one.
const PlayerFile = "player"

func writePlayerNamesFile() {
	f, err := os.OpenFile(filePath(PlayerFile), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err.Error())
	}

	for _, name := range playerNamesCache {
		_, err = f.WriteString(string(name + "\n"))
		if err != nil {
			panic(err.Error())
		}
	}

	err = f.Close()
	if err != nil {
		panic(err.Error())
	}
}

// initializePlayerNameCaches returns the number of entries in the file.
func initializePlayerNameCaches() int {
	b, err := ioutil.ReadFile(filePath(PlayerFile))
	if os.IsNotExist(err) {
		return 0
	} else if err != nil {
		panic(err.Error())
	}

	names := strings.Split(string(b), "\n")
	if len(names) == 1 {
		return 0
	}

	c := 0
	for i, n := range names {
		playerIDCache[PlayerName(n)] = PlayerIdentifier(i + 1)
		playerNamesCache = append(playerNamesCache, PlayerName(n))
		c++
	}

	return c
}
