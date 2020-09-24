package memory

import (
	"io/ioutil"
	"log"
	"os"
)

// The HashFile is a binary encoded file listing the password hashes of all players in order of id
// starting with 1. Each entry starts with an unsigned byte that indicates how many bytes
// represent the hash that follows.
const HashFile = "hash"

func writeHashFile() {
	f, err := os.OpenFile(filePath(HashFile), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err.Error())
	}

	for _, hash := range hashCache {
		if len(hash) > 256 {
			log.Panicln("hash longer than can fit into byte", len(hash))
		}
		_, err = f.Write([]byte{byte(len(hash))})
		if err != nil {
			panic(err.Error())
		}
		_, err = f.Write(hash)
		if err != nil {
			panic(err.Error())
		}
	}

	err = f.Close()
	if err != nil {
		panic(err.Error())
	}
}

// initializeHashCache returns the number of entries read from the HashFile.
func initializeHashCache() int {
	content, err := ioutil.ReadFile(filePath(HashFile))
	if os.IsNotExist(err) {
		return 0
	} else if err != nil {
		panic(err.Error())
	}

	i := 0
	c := 0
	for i != len(content) {
		length := int(content[i])
		hashCache = append(hashCache, content[i+1:i+length])
		i += 1 + length
		c++
	}

	return c
}
