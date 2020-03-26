package main

import (
	"database/sql"
	"net/http"
)

// GameHasPlayerHandler is used by RequesterInGame and PlayerNamed to
// verify that the requester for a game is a player in the game.
func GameHasPlayerHandler(w http.ResponseWriter, tx *sql.Tx, id GameIdentifier, requester string) bool {
	if GameHasPlayer(tx, id, requester) == false {
		tx.Commit()
		DebugPrintln(requester, "not in game", id)
		w.WriteHeader(http.StatusNotFound)
		return false
	}
	return true
}

// A RequesterInGameFunc is the input for RequesterInGame.
type RequesterInGameFunc func(http.ResponseWriter, *http.Request, *sql.Tx, GameIdentifier)

// RequesterInGame adds to an AuthenticRequestHandler that's calling GameIdentifierParsed.
// The named requester from AuthenticRequestHandler is verified to be in the
// game parsed by GameIdentifierParsed.
//
//  // added to handlers with http.Handle(BoardsPath, BoardsHandler)
//  const BoardsPath = "/boards/"
//
//  var BoardsHandler = AuthenticRequestHandler{
//    Get: GameIdentifierParsed(RequesterInGame(BoardsGet), BoardsPath),
//  }
//
//  func BoardsGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx, id GameIdentifier) {
//    // customized BoardsPath get handling here
//
// The requester's name is not an argument to the customized handler (BoardsGet in
// the above example), but if it's needed then use PlayerNamed instead.
func RequesterInGame(calls RequesterInGameFunc) GameIdentifedFunc {
	return func(w http.ResponseWriter, r *http.Request, tx *sql.Tx, id GameIdentifier, requester string) {
		if GameHasPlayerHandler(w, tx, id, requester) == false {
			return
		}
		calls(w, r, tx, id)
	}
}

// A PlayerNamedFunc is the input for PlayerNamed.
type PlayerNamedFunc func(http.ResponseWriter, *http.Request, *sql.Tx, GameIdentifier, string)

// PlayerNamed is the same as RequesterInGame except that is also passes
// along the player name.
func PlayerNamed(calls PlayerNamedFunc) GameIdentifiedFunc {
	return func(w http.ResponseWriter, r *http.Request, tx *sql.Tx, id GameIdentifier, requester string) {
		if GameHasPlayerHandler(w, tx, id, requester) == false {
			return
		}
		calls(w, r, tx, id, requester)
	}
}
