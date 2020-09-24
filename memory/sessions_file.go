package memory

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

// The SessionFile is a binary encoded file listing the session keys of all players in order of id
// starting with 1. Like the HashFile, each entry starts with an unsigned byte that indicates how
// many bytes represent the key that follows.
const SessionFile = "session"

// TODO: SessionFile and HashFile have a similar format, with these read/write funcs overlapping

func writeSessionFile() {
	sessionsSlice := make([]SessionKey, len(sessionCache))
	for key, id := range sessionCache {
		sessionsSlice[id-1] = key
	}

	f, err := os.OpenFile(filePath(SessionFile), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err.Error())
	}

	for _, key := range sessionsSlice {
		enc := []byte(key.String())
		if len(enc) > 256 {
			panic(fmt.Sprint("session key len longer than can fit into byte", len(enc)))
		}
		_, err = f.Write([]byte{byte(len(enc))})
		if err != nil {
			panic(err.Error())
		}
		_, err = f.Write(enc)
		if err != nil {
			panic(err.Error())
		}
	}

	err = f.Close()
	if err != nil {
		panic(err.Error())
	}
}

// initializeSessionsCache returns the number of entries.
func initializeSessionsCache() int {
	content, err := ioutil.ReadFile(filePath(SessionFile))
	if os.IsNotExist(err) {
		return 0
	} else if err != nil {
		panic(err.Error())
	}

	c := 0
	i := 0
	id := PlayerIdentifier(1)
	for i != len(content) {
		length := int(content[i])
		var key SessionKey
		copy(key[:], bytes.Runes(content[i+1:i+length]))
		sessionCache[key] = id
		i += 1 + length
		id++
		c++
	}

	return c
}
