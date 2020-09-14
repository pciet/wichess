package wichess

import (
	"net/http"
	"time"

	"github.com/pciet/wichess/auth"
)

// These constants describe every HTTP path and HTML template.
const (
	// First when a new person connects to this host process they are given a webpage to login with.
	// If the username and password provided exist together in memory, or the username is new, then
	// the player's web browser is issued a unique session key to be stored as an HttpOnly cookie.
	// All other handlers inspect this session key cookie value using the auth.Handler to be
	// sure the request is authentic.
	LoginPath         = auth.LoginPath
	LoginHTMLTemplate = "html/login.tmpl"

	// A login session is ended by a GET of /quit.
	QuitPath = "/quit"

	// The index webpage is where the army is picked from the collection. Before a game is played
	// the match webpage is navigated to to pick an opponent. If a game against a person is already
	// in progress then the browser is redirected to it.
	IndexPath         = "/"
	IndexHTMLTemplate = "html/index.tmpl"

	// After configuring the army on index the player picks a match on the match webpage.
	MatchPath         = "/match"
	MatchHTMLTemplate = "html/match.tmpl"

	// Details of the rules of individual pieces is in the details path. This information is
	// not interactive and doesn't require authentication to read.
	DetailsPath         = "/details"
	DetailsHTMLTemplate = "html/details.tmpl"

	// A simple webpage describes an overview of the rules and features of Wisconsin Chess.
	RulesPath = "/rules"
	RulesHTML = "html/rules.html"

	// GameHTMLTemplate is the template for all of the game modes.
	GameHTMLTemplate = "html/game.tmpl"

	// "Computer" is the artificial opponent mode, where the other player's moves are
	// made automatically. POST sets up the match and GET retrieves the game webpage.
	ComputerPath = "/computer"

	// Players match against each other in the People mode by typing in their opponent's username,
	// or by clicking on a recent opponent's name that's displayed. This causes a POST to
	// /people?o=PlayersName which returns the game ID when the opponent does the same or an error
	// if there's a timeout by the opponent not requesting. A GET to /people/[game identifier] loads
	// the game webpage for the match.
	PeoplePath     = "/people/"
	PeopleRootPath = "/people"

	// The /peopleid path isn't used by the regular game webpage but helps with testing by
	// providing the game identifier of the people game if one exists.
	PeopleIDPath = "/peopleid"

	// /collection and /picks are testing paths for accessing a player's collection.
	PicksPath      = "/picks"
	CollectionPath = "/collection"

	// A game can be conceded by clicking on the concede button which does a GET of
	// /concede/[game identifier].
	ConcedePath = "/concede/"

	// Which piece is where on a game's board is initially loaded into the web browser with a
	// GET of /boards/[game identifier].
	BoardsPath = "/boards/"

	// The players path isn't used by the regular game webpage but helps with testing by providing
	// which player is which orientation and which player is active.
	PlayersPath = "/players/"

	// The webpages requests a calculation of all possible moves for a turn with a GET to
	// /moves/[game identifier]?turn=[turn number].
	MovesPath = "/moves/"

	// A move is requested by POST to /move/[game identifier] with the move encoded as JSON
	// in the request body.
	MovePath = "/move/"

	// An opponent is alerted to board changes caused by a move with a WebSocket message.
	AlertPath = "/alert/"

	// When a game is complete the player navigates to the reward webpage to optionally put the
	// picked pieces (if selected before the start of the game) and a random reward piece into
	// their collection. The acknowledge button is on this webpage. POST is for any selection.
	RewardPath         = "/reward/"
	RewardHTMLTemplate = "html/reward.tmpl"

	// A player acknowledges they've reviewed the rewards and final position and don't need to
	// see them again with a GET of /acknowledge/[game identifier].
	AcknowledgePath = "/acknowledge/"
)

var (
	acknowledgeHandler = auth.Handler{
		Get: auth.PlayerAndGameWriteable(acknowledgeGet, AcknowledgePath),
	}

	quitHandler = auth.Handler{Get: auth.PlayerIdentified(quitGet)}

	indexHandler = auth.Handler{
		Get: auth.PlayerReadable(indexGet),
	}

	matchHandler = auth.Handler{
		Get: auth.PlayerReadable(matchGet),
	}

	computerHandler = auth.Handler{
		Get:  auth.PlayerReadable(computerGet, ComputerPath),
		Post: auth.PlayerWriteable(computerPost, ComputerPath),
	}

	matchPeopleHandler = auth.Handler{
		Get:  auth.GameReadablePlayerIdentified(peopleGet, PeoplePath),
		Post: peoplePost,
	}

	peopleIDHandler = auth.Handler{
		Get: auth.PlayerReadable(peopleIDGet),
	}

	picksHandler = auth.Handler{
		Get: auth.PlayerReadable(picksGet),
	}

	collectionHandler = auth.Handler{Get: auth.PlayerReadable(collectionGet)}

	concedeHandler = auth.Handler{
		Get: auth.PlayerAndGameWriteable(concedeGet, ConcedePath),
	}

	boardsHandler = auth.Handler{
		Get: auth.GameReadable(boardsGet, BoardsPath),
	}

	playersHandler = auth.Handler{
		Get: auth.GameReadable(playersGet, PlayersPath),
	}

	movesHandler = auth.Handler{
		Get: auth.GameReadablePlayerIdentified(movesGet, MovesPath),
	}

	moveHandler = auth.Handler{
		Post: auth.GameAndPlayerIdentified(movePost, MovePath),
	}

	alertWebSocketUpgradeHandler = auth.Handler{
		Get: auth.GameReadablePlayerIdentified(alertGet, AlertPath),
	}

	rewardHandler = auth.Handler{
		Get:  auth.GameAndPlayerReadable(rewardGet, RewardPath),
		Post: auth.GameReadablePlayerWriteable(rewardPost, RewardPath),
	}
)

// TODO: does HTTPMaxIdleConns unnecessarily limit the max number of concurrent players on server
// hardware that can support more?
// github.com/pciet/wichess/issues/9 - 100 http + 100 postgres should be under macOS 256 limit

const (
	httpMaxIdleConns    = 50
	httpIdleConnTimeout = time.Duration(time.Minute * 5)
)

// InitializeHTTP adds net/http handlers for all paths and loads all HTML templates. The
// http.ListenAndServe function is called to serve those HTTP paths by the caller.
func InitializeHTTP() {
	loadHTMLTemplates()

	http.HandleFunc(LoginPath, loginHandler)
	http.Handle(QuitPath, quitHandler)
	http.Handle(IndexPath, indexHandler)
	http.Handle(MatchPath, matchHandler)
	http.HandleFunc(DetailsPath, detailsGet)
	http.HandleFunc(RulesPath, rulesGet)
	http.Handle(ComputerPath, computerHandler)
	http.Handle(PeoplePath, matchPeopleHandler)
	http.Handle(PeopleRootPath, matchPeopleHandler)
	http.Handle(PeopleIDPath, peopleIDHandler)
	http.Handle(PicksPath, picksHandler)
	http.Handle(CollectionPath, collectionHandler)
	http.Handle(ConcedePath, concedeHandler)
	http.Handle(BoardsPath, boardsHandler)
	http.Handle(PlayersPath, playersHandler)
	http.Handle(MovesPath, movesHandler)
	http.Handle(MovePath, moveHandler)
	http.Handle(AlertPath, alertWebSocketUpgradeHandler)
	http.Handle(RewardPath, rewardHandler)
	http.Handle(AcknowledgePath, acknowledgeHandler)

	// All files used by the webpages, including JavaScript and CSS, is in the web folder.
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/favicon.ico")
	})

	dt, ok := http.DefaultTransport.(*http.Transport)
	if ok == false {
		Panic("http.DefaultTransport.(*http.Transport) failed")
	}
	dt.MaxIdleConns = httpMaxIdleConns
	dt.IdleConnTimeout = httpIdleConnTimeout
}
