package wichess

import (
	"log"
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

	var name memory.PlayerName
	if g.White.PlayerIdentifier == pid {
		name = whiteName
	} else if g.Black.PlayerIdentifier == pid {
		name = blackName
	} else {
		log.Panicln(PeoplePath, "GET by player", pid, "not in game", whiteName, "vs", blackName)
	}

	t := GameHTMLTemplateData{
		GameIdentifier: g.GameIdentifier,
		Conceded:       g.Conceded,
		White: GamePlayerHTMLTemplateData{
			PlayerName: whiteName,
			Captures:   g.White.Captures,
		},
		Black: GamePlayerHTMLTemplateData{
			PlayerName: blackName,
			Captures:   g.Black.Captures,
		},
		Active: g.Active,
		PreviousMove: PreviousMoveHTMLTemplateData{
			From: int(g.PreviousMove.From.Index()),
			To:   int(g.PreviousMove.To.Index()),
		},
		Player: name,
	}
	writeHTMLTemplate(w, GameHTMLTemplate, t)
}

type PeoplePostJSON struct {
	memory.GameIdentifier `json:"id"`
}

func peoplePost(w http.ResponseWriter, r *http.Request, pid memory.PlayerIdentifier) {
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

	// if the opponent hasn't requested this player yet then the handler waits here
	gameID, opponentID := game.RequestOpponent(memory.PlayerName(requestedOpponent), pid, army)

	close(done)

	if gameID != memory.NoGame {
		p = memory.LockPlayer(pid)
		p.AddRecentOpponent(opponentID)
		p.PeopleGame = gameID
		p.Unlock()
	} // else a timeout or cancel

	jsonResponse(w, PeoplePostJSON{gameID})
}
