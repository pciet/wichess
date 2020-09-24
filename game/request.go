package game

import (
	"sync"
	"time"

	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/piece"
)

const opponentRequestTimeout = 120 * time.Second

type opponentRequest struct {
	Opponent memory.PlayerName
	With     piece.ArmyRequest
	Notify   chan memory.GameIdentifier
}

var (
	opponentRequests     = map[memory.PlayerIdentifier]opponentRequest{}
	opponentRequestsLock = sync.Mutex{}
)

// RequestOpponent causes the calling goroutine to wait for the requested opponent to request back
// that player. When a match is successful then a new game is created. The opponent is requested
// by name instead of id so a match can happen even if the name hasn't been created yet.
// RequestOpponent returns early if a timeout occurs or EndOpponentRequest is called.
func RequestOpponent(opponent memory.PlayerName, requester memory.PlayerIdentifier,
	with piece.ArmyRequest) (memory.GameIdentifier, memory.PlayerIdentifier) {

	opponentRequestsLock.Lock()
	_, has := opponentRequests[requester]
	if has {
		opponentRequestsLock.Unlock()
		return memory.NoGame, memory.NoPlayer
	}

	// TODO: is a mutex taken twice for these?
	rname := requester.Name()
	oppID := memory.PlayerNameKnown(opponent)

	if oppID != memory.NoPlayer {
		oppReq, has := opponentRequests[oppID]
		if has && (oppReq.Opponent == rname) {
			delete(opponentRequests, oppID)
			opponentRequestsLock.Unlock()
			gameID := New(with, oppReq.With, requester, oppID)
			oppReq.Notify <- gameID
			return gameID, oppID
		}
		// otherwise there's no request or request for another player and this player should wait
	}

	ready := make(chan memory.GameIdentifier)
	opponentRequests[requester] = opponentRequest{opponent, with, ready}
	opponentRequestsLock.Unlock()

	var id memory.GameIdentifier
	select {
	case <-time.After(opponentRequestTimeout):
	case id = <-ready:
	}

	opponentRequestsLock.Lock()
	delete(opponentRequests, requester)
	opponentRequestsLock.Unlock()

	if id == memory.NoGame {
		return id, memory.NoPlayer
	}
	return id, memory.PlayerNameKnown(opponent)
}

// EndOpponentRequest ends the request started with RequestOpponent.
func EndOpponentRequest(requester memory.PlayerIdentifier) {
	opponentRequestsLock.Lock()
	req, has := opponentRequests[requester]
	if has {
		delete(opponentRequests, requester)
	} else {
		opponentRequestsLock.Unlock()
		return
	}
	opponentRequestsLock.Unlock()
	req.Notify <- memory.NoGame
}
