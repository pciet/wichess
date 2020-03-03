package main

import (
	"net/http"
	"strconv"
)

// Expects a pathPrefix like '/games/' in /games/[ID to be parsed].
// Writes a Bad Request HTTP response if the URL can't be parsed and returns 0.
func ParseURLGameIdentifier(w http.ResponseWriter, r *http.Request, pathPrefix string) GameIdentifier {
	id, err := strconv.ParseInt(r.URL.Path[len(pathPrefix):len(r.URL.Path)], 10, 0)
	if err != nil {
		DebugPrintln(pathPrefix, "bad game identifier string in URL", r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return 0
	}
	return GameIdentifier(id)
}

func ParseURLIntQuery(w http.ResponseWriter, r *http.Request, query string) int {
	q, has := r.URL.Query()[query]
	if has == false {
		DebugPrintln(query, "not in URL", r.URL)
		w.WriteHeader(http.StatusBadRequest)
		return 0
	}

	if len(q) != 1 {
		DebugPrintln(query, "has", len(q), "values instead of 1")
		w.WriteHeader(http.StatusBadRequest)
		return 0
	}

	v, err := strconv.ParseInt(q[0], 10, 0)
	if err != nil {
		DebugPrintln(query, "not parseable as int:", err)
		w.WriteHeader(http.StatusBadRequest)
		return 0
	}

	return v
}
