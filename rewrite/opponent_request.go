package main

import (
	"sync"
	"time"
)

const OpponentRequestTimeoutSeconds = 120

type OpponentRequest struct {
	Opponent string
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
// a new game is created and the identifier for it returned, or 0 is returned otherwise. The
// opponent is named by string so that the match is successful even if the username hasn't been
// created yet. The opponent's player ID is also returned.
func RequestOpponent(opponent string, requester Player,
	with ArmyRequest) (GameIdentifier, PlayerIdentifier) {
	OpponentRequestsLock.Lock()
	opp, has := OpponentRequests[requester.ID]
	if has {
		OpponentRequestsLock.Unlock()
		DebugPrintln(requester, "already has opponent request for", opp.Opponent, "not", opponent)
		return 0, 0
	}

	// TODO: does this lock cause bad delays because of this database communication?
	tx := DatabaseTransaction()
	oppID := PlayerID(tx, opponent)
	tx.Commit()

	// does the opponent already have a request for this player?
	// if so then create a new game and notify that opponent of the GameIdentifier
	oppReq, has := OpponentRequests[oppID]
	if has && (oppID == -1) {
		Panic("PlayerID -1 value meaning no player in DB in OpponentRequests map")
	}
	if has && (oppReq.Opponent == requester.Name) {
		delete(OpponentRequests, oppID)
		OpponentRequestsLock.Unlock()

		// TODO: if player can be deleted then there's a race here from above PlayerID read

		tx = DatabaseTransaction()
		id := NewGame(tx, with, oppReq.With, requester, Player{opponent, oppID})
		tx.Commit()

		oppReq.Notify <- id
		return id, oppID
	}

	ready := make(chan GameIdentifier)
	OpponentRequests[requester.ID] = OpponentRequest{opponent, with, ready}
	OpponentRequestsLock.Unlock()

	var id GameIdentifier
	select {
	case <-time.After(OpponentRequestTimeoutSeconds * time.Second):
	case id = <-ready:
	}

	OpponentRequestsLock.Lock()
	delete(OpponentRequests, requester.ID)
	OpponentRequestsLock.Unlock()

	tx = DatabaseTransaction()
	oppID = PlayerID(tx, opponent)
	tx.Commit()

	if oppID == -1 {
		Panic("unexpected missing player ID for", opponent)
	}

	return id, oppID
}
