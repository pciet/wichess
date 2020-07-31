package main

import (
	"net/http"
	"time"
)

const (
	// TODO: does this unnecessarily limit the maximum number of concurrent players on server
	// hardware that can support more?
	// github.com/pciet/wichess/issues/9 - 100 http + 100 postgres should be under macOS 256 limit
	HTTPMaxIdleConns    = 50
	HTTPIdleConnTimeout = time.Duration(time.Minute * 5)
)

func InitializeHTTP() {
	// First when a new person connects to this program they are given a webpage to login with.
	// If the username and password provided exist together in the database, or the username
	// is new, then the person's web browser is issued a unique session key to be stored as
	// a cookie.
	// All other handlers inspect the session key cookie value to be sure the request is authentic.
	http.HandleFunc(LoginPath, LoginHandler)

	// A login is ended by a GET to /quit with the session key cookie.
	http.Handle(QuitPath, QuitHandler)

	// The index webpage is where the army is picked from the collection. Before a game is played
	// the match webpage is navigated to to pick an opponent. If a game against a person is already
	// in progress then the browser is redirected to it.
	http.Handle(IndexPath, IndexHandler)

	// After configuring the army on index the player picks a match on the match page.
	http.Handle(MatchPath, MatchHandler)

	// Details of the rules of individual pieces is in the details path. This information is
	// not interactive and doesn't require authentication to read.
	http.HandleFunc(DetailsPath, DetailsGet)

	// A simple webpage describes an overview of the rules and features of Wisconsin Chess.
	http.HandleFunc(RulesPath, RulesGet)

	// "Computer" is the artificial opponent mode, where the other player's moves are
	// made automatically.
	// In HTTP terms, this path's POST is used to setup the match and the GET is to load
	// the game webpage.
	http.Handle(ComputerPath, ComputerHandler)

	// Players match against each other in the People mode by typing in their opponent's username,
	// or by clicking on a recent opponent's name that's displayed. This causes a POST to
	// /people?o=PlayersName which returns the game ID when the opponent does the same or an error
	// if there's a timeout by the opponent not requesting. A GET to /people/[game identifier] loads
	// the game webpage for the match.
	http.Handle(PeoplePath, MatchPeopleHandler)
	http.Handle(PeopleRootPath, MatchPeopleHandler)

	// A people game can be conceded by clicking on the concede button, which does a GET to
	// /concede/[game identifier].
	http.Handle(ConcedePath, ConcedeHandler)

	// Which piece is where on a game's board is initially loaded into the web browser with a
	// GET to /boards/[game identifier].
	http.Handle(BoardsPath, BoardsHandler)

	// The webpages requests a calculation of all possible moves for a turn with a GET to
	// /moves/[game identifier]?turn=[turn number].
	// The turn number is included to guarantee the web browser and host are synchronized.
	http.Handle(MovesPath, MovesHandler)

	// A move is requested by POST to /move/[game identifier] with the move encoded as JSON
	// in the request body.
	http.Handle(MovePath, MoveHandler)

	// An opponent is alerted to board changes caused by a move with a WebSocket message.
	http.Handle(AlertPath, AlertWebSocketUpgradeHandler)

	// When a game is complete the player navigates to the reward webpage to optionally put the
	// picked pieces (if selected before the start of the game) and a random reward piece into
	// their collection. The acknowledge button is on this webpage.
	// A POST is sent with any selection.
	http.Handle(RewardPath, RewardHandler)

	// A player acknowledges they've reviewed the rewards and final position and don't need to
	// see them again with a GET of /acknowledge/[game identifier].
	http.Handle(AcknowledgePath, AcknowledgeHandler)

	// All files used by the webpages, including JavaScript and CSS, is in the web folder.
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/favicon.ico")
	})

	dt, ok := http.DefaultTransport.(*http.Transport)
	if ok == false {
		Panic("http.DefaultTransport.(*http.Transport) failed")
	}
	dt.MaxIdleConns = HTTPMaxIdleConns
	dt.IdleConnTimeout = HTTPIdleConnTimeout
}
