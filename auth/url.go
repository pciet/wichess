package auth

import (
	"strconv"

	"github.com/pciet/wichess/memory"
)

// parseURLGameIdentifier expects the URL path to look like [pathPrefix][ID integer to be parsed].
// For the example URL http://localhost:8080/games/513 the pathPrefix argument should be /games/
// and urlPath /games/513 which will return 513 as a GameIdentifier.
func parseURLGameIdentifier(urlPath, pathPrefix string) (memory.GameIdentifier, error) {
	id, err := strconv.ParseInt(urlPath[len(pathPrefix):len(urlPath)], 10, 0)
	if err != nil {
		return 0, err
	}
	return memory.GameIdentifier(id), nil
}
