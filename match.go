// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/pciet/wichess/wichessing"
)

type gameSetup struct {
	slot          int
	leftRookID    int
	leftKnightID  int
	leftBishopID  int
	rightBishopID int
	rightKnightID int
	rightRookID   int
}

var pendingMatches = make(map[string]gameSetup)
var pendingMatchesLock = &sync.Mutex{}

type listeningPlayer struct {
	*websocket.Conn
}

var listeningPlayers = make(map[string]listeningPlayer)
var listeningPlayersLock = &sync.Mutex{}

func setWhite(s *gameSetup) {
	if s.leftRookID == 0 {
		s.leftRookID = int(wichessing.WhiteRook)
	}
	if s.leftKnightID == 0 {
		s.leftKnightID = int(wichessing.WhiteKnight)
	}
	if s.leftBishopID == 0 {
		s.leftBishopID = int(wichessing.WhiteBishop)
	}
	if s.rightBishopID == 0 {
		s.rightBishopID = int(wichessing.WhiteBishop)
	}
	if s.rightKnightID == 0 {
		s.rightKnightID = int(wichessing.WhiteKnight)
	}
	if s.rightRookID == 0 {
		s.rightRookID = int(wichessing.WhiteRook)
	}
}

func setBlack(s *gameSetup) {
	if s.leftRookID == 0 {
		s.leftRookID = int(wichessing.BlackRook)
	}
	if s.leftKnightID == 0 {
		s.leftKnightID = int(wichessing.BlackKnight)
	}
	if s.leftBishopID == 0 {
		s.leftBishopID = int(wichessing.BlackBishop)
	}
	if s.rightBishopID == 0 {
		s.rightBishopID = int(wichessing.BlackBishop)
	}
	if s.rightKnightID == 0 {
		s.rightKnightID = int(wichessing.BlackKnight)
	}
	if s.rightRookID == 0 {
		s.rightRookID = int(wichessing.BlackRook)
	}
}

// An identifier of 0 means request the normal basic piece instead of a hero piece.
func requestMatch(name string, config gameSetup) {
	pendingMatchesLock.Lock()
	defer pendingMatchesLock.Unlock()
	if len(pendingMatches) == 0 {
		pendingMatches[name] = config
		return
	}
	if len(pendingMatches) == 1 {
		_, has := pendingMatches[name]
		if has {
			return
		}
	}
	for key, value := range pendingMatches {
		if key == name {
			continue
		}
		pendingMatches[name] = config
		setWhite(&config)
		setBlack(&value)
		newBoardIntoDatabase(name, config, key, value)
		notifyMatchMadeToListeners(name, config.slot, key, pendingMatches[key].slot)
		delete(pendingMatches, name)
		delete(pendingMatches, key)
		return
	}
}

type matchNotification struct {
	Opponent string
	Slot     int
}

func notifyMatchMadeToListeners(p1 string, slot1 int, p2 string, slot2 int) {
	listeningPlayersLock.Lock()
	defer listeningPlayersLock.Unlock()
	conn, has := listeningPlayers[p1]
	if has {
		_ = conn.WriteJSON(matchNotification{
			Opponent: p2,
			Slot:     slot1,
		})
	}
	conn, has = listeningPlayers[p2]
	if has {
		_ = conn.WriteJSON(matchNotification{
			Opponent: p1,
			Slot:     slot2,
		})
	}
}

type matchListRequest struct {
	states interface{}
}

// Assumes the only kind of websocket message is a request for the match list.
func listeningForMatchChanges(name string, conn *websocket.Conn) {
	listeningPlayersLock.Lock()
	defer listeningPlayersLock.Unlock()
	listeningPlayers[name] = listeningPlayer{
		Conn: conn,
	}
	go func() {
		defer doneListeningForMatchChanges(name)
		for {
			var m matchListRequest
			err := conn.ReadJSON(&m)
			if err != nil {
				return
			}
			err = conn.WriteJSON(playerBoardInfo(name))
			if err != nil {
				return
			}
		}
	}()
}

func doneListeningForMatchChanges(name string) {
	listeningPlayersLock.Lock()
	defer listeningPlayersLock.Unlock()
	delete(listeningPlayers, name)
}

func pendingMatchesFor(name string) []boardInfo {
	pendingMatchesLock.Lock()
	defer pendingMatchesLock.Unlock()
	boards := make([]boardInfo, 0, 1)
	for key, value := range pendingMatches {
		if key != name {
			continue
		}
		boards = append(boards, boardInfo{
			Slot: value.slot,
		})
	}
	return boards
}
