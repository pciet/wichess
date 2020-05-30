package main

import (
	"sync"
	"time"
)

const OpponentRequestTimeoutSeconds = 120

type OpponentRequest struct {
	Opponent PlayerIdentifier
	With     ArmyRequest
	Notify   chan GameIdentifier
}

var (
	OpponentRequests     = map[PlayerIdentifier]OpponentRequest{}
	OpponentRequestsLock = sync.Mutex{}
)

func EndOpponentRequest(requester PlayerIdentifier) {
	OpponentRequestsLock.Lock()
	req, has := OpponentRequests[requester]
	if has {
		delete(OpponentRequests, requester)
	}
	OpponentRequestsLock.Unlock()
	req.Notify <- 0
}

// RequestOpponent blocks until the opponent requests this player, a timeout occurs, or
// EndOpponentRequest is called from another goroutine. If the players successfully match then
// a new game is created and the identifier for it returned, or 0 is returned otherwise.
func RequestOpponent(opponent, requester PlayerIdentifier, with ArmyRequest) GameIdentifier {
	OpponentRequestsLock.Lock()
	opp, has := OpponentRequests[requester]
	if has {
		OpponentRequestsLock.Unlock()
		DebugPrintln(requester, "already has opponent request for", opp.Opponent, "not", opponent)
		return 0
	}

	// does the opponent already have a request for this player?
	// if so then create a new game and notify that opponent of the GameIdentifier
	oppReq, has := OpponentRequests[opponent]
	if has && (oppReq.Opponent == requester) {
		delete(OpponentRequests, opponent)
		OpponentRequestsLock.Unlock()

		tx := DatabaseTransaction()
		id := NewGame(tx, with, oppReq.With,
			Player{PlayerName(tx, requester), requester},
			Player{PlayerName(tx, opponent), opponent})
		tx.Commit()

		oppReq.Notify <- id
		return id
	}

	ready := make(chan GameIdentifier)
	OpponentRequests[requester] = OpponentRequest{opponent, with, ready}
	OpponentRequestsLock.Unlock()

	var id GameIdentifier
	select {
	case <-time.After(OpponentRequestTimeoutSeconds * time.Second):
	case id = <-ready:
	}

	OpponentRequestsLock.Lock()
	delete(OpponentRequests, requester)
	OpponentRequestsLock.Unlock()

	return id
}
