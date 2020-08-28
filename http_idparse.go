package wichess

import (
	"database/sql"
	"net/http"
)

// A GameIdentifiedFunc is the input for GameIdentifierParsed.
type GameIdentifiedFunc func(http.ResponseWriter, *http.Request, *sql.Tx, GameIdentifier, Player)

// TODO: id in query instead of path

// GameIdentifierParsed adds URL path game identifier parsing to an AuthenticRequestHandler.
//
//  // added to handlers with http.Handle(StatisticsPath, StatisticsHandler)
//  const StatisticsPath = "/stats/"
//
//  var StatisticsHandler = AuthenticRequestHandler{
//    Get:  GameIdentifierParsed(StatsGet, StatisticsPath),
//    Post: GameIdentifierParsed(StatsPost, StatisticsPath),
//  }
//
//  func StatsGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx,
//		id GameIdentifier, requester Player) {
//    // customized StatisticsPath get handling here
//
// The path is expected to be formed as [pathPrefix][id]. For example, /games/867 has pathPrefix
// /games/ and identifier 867.
func GameIdentifierParsed(calls GameIdentifiedFunc, pathPrefix string) AuthenticRequestHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, tx *sql.Tx, requester Player) {
		id, err := ParseURLGameIdentifier(r.URL.Path, pathPrefix)
		if err != nil {
			tx.Commit()
			DebugPrintln(pathPrefix, "bad game ID in URL", r.URL, "from", requester, ":", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		calls(w, r, tx, id, requester)
	}
}
