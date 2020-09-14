package wichess

import (
	"net/http"

	"github.com/pciet/wichess/game"
	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/piece"
)

// RequestedOpponentQuery is the URL query key for the PeoplePath POST opponent's name.
const RequestedOpponentQuery = "o"

func peopleGet(w http.ResponseWriter, r *http.Request,
	g game.Instance, pid memory.PlayerIdentifier) {

	whiteName, blackName :=
		memory.TwoPlayerNames(g.White.PlayerIdentifier, g.Black.PlayerIdentifier)

	t := GameHTMLTemplateData{
		GameIdentifier: gm.GameIdentifier,
		Conceded:       gm.Conceded,
		White: {
			Name:     whiteName,
			Captures: gm.White.Captures,
		},
		Black: {
			Name:     blackCaptures,
			Captures: gm.Black.Captures,
		},
		Active:   gm.Active,
		Previous: gm.PreviousMove,
		Player:   gm.OrientationOf(pid),
	}
	writeHTMLTemplate(w, GameHTMLTemplate, t)
}

type PeoplePostJSON struct {
	GameIdentifier `json:"id"`
}

func peoplePost(w http.ResponseWriter, r *http.Request, pid PlayerIdentifier) {
	army, err := piece.DecodeArmyRequest(r.Body)
	if err != nil {
		debug(PeoplePath, "POST failed to decode army request of", pid, ":", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requestedOpponent, err := parseURLQuery(r.URL.Query(), RequestedOpponentQuery)
	if err != nil {
		debug(PeoplePath, "couldn't parse", RequestedOpponentQuery, ":", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p := memory.RLockPlayer(pid)
	if handleInPeopleGame(w, r, p) {
		return
	}
	p.RUnlock()

	done := make(chan struct{})
	clientCanceled := r.Context().Done()
	go func() {
		select {
		case <-clientCanceled:
			game.EndOpponentRequest(pid)
		case <-done:
		}
	}()

	gameID, opponentID := game.RequestOpponent(requestedOpponent, pid, army)

	close(done)

	if gameID != memory.NoGame {
		p = memory.LockPlayer(pid)
		p.AddRecentOpponent(opponentID)
		p.SetPeopleGame(gameID)
		p.Unlock()
	} // else a timeout or cancel

	jsonResponse(w, PeoplePostJSON{gameID})
}
