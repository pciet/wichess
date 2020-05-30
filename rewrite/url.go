package main

import (
	"fmt"
	"net/url"
	"strconv"
)

// ParseURLGameIdentifier expects the URL path to look like [pathPrefix][ID integer to be parsed].
// For the example URL http://localhost:8080/games/513 the pathPrefix argument should be /games/
// and urlPath /games/513 which will return 513 as a GameIdentifier.
func ParseURLGameIdentifier(urlPath, pathPrefix string) (GameIdentifier, error) {
	id, err := strconv.ParseInt(urlPath[len(pathPrefix):len(urlPath)], 10, 0)
	if err != nil {
		return 0, err
	}
	return GameIdentifier(id), nil
}

// ParseURLIntQuery will parse an int from a URL query.
// Starting with just the URL path in the example /moves?turn=21 to parse turn then
// queryKey is turn and from is the url.Values member of (*http.Request var).URL.Query().
func ParseURLIntQuery(from url.Values, queryKey string) (int, error) {
	intstring, err := ParseURLQuery(from, queryKey)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseInt(intstring, 10, 0)
	if err != nil {
		return 0, err
	}

	return int(v), nil
}

func ParseURLQuery(from url.Values, queryKey string) (string, error) {
	q, has := from[queryKey]
	if has == false {
		return "", fmt.Errorf("%v not in URL queries", queryKey)
	}

	if len(q) != 1 {
		return "", fmt.Errorf("%v has %v values instead of 1", queryKey, len(q))
	}

	return q[0], nil
}
