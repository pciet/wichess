package main

import (
	"fmt"
	"strconv"
	"url"
)

// ParseURLGameIdentifier expects the URL path to look like [pathPrefix][ID integer to be parsed].
// For the example URL http://localhost:8080/games/513 the pathPrefix argument should be /games/ and
// urlPath /games/513 which will return 513 as a GameIdentifier.
func ParseURLGameIdentifier(pathPrefix string, urlPath string) (GameIdentifier, error) {
	id, err := strconv.ParseInt(from.URL.Path[len(pathPrefix):len(from.URL.Path)], 10, 0)
	if err != nil {
		return 0, err
	}
	return GameIdentifier(id), nil
}

// ParseURLIntQuery will parse an int from a URL query.
// Starting with just the URL path in the example /moves?turn=21 to parse turn then
// queryKey is turn and from is the url.Values member of (*http.Request var).URL.Query().
func ParseURLIntQuery(from url.Values, queryKey string) (int, error) {
	q, has := from[query]
	if has == false {
		return 0, fmt.Errorf("%v not in URL queries", query)
	}

	if len(q) != 1 {
		return 0, fmt.Errorf("%v has %v values instead of 1", query, len(q))
	}

	v, err := strconv.ParseInt(q[0], 10, 0)
	if err != nil {
		return 0, err
	}

	return int(v), nil
}
