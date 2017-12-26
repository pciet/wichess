// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/lib/pq"
)

const (
	friend_table = "friends"

	friend_requester = "requester"
	friend_setup     = "setup"
	friend_friend    = "friend"
	friend_slot      = "slot"
)

var friendMatchListeners = map[string]*websocket.Conn{}
var friendMatchListenersLock = sync.RWMutex{}

type friendMatchMessage struct {
	Slot int `json:"slot"`
}

func notifyFriendMatch(name string, slot uint8) {
	friendMatchListenersLock.RLock()
	conn, has := friendMatchListeners[name]
	friendMatchListenersLock.RUnlock()
	if has == false {
		return
	}
	err := conn.WriteJSON(friendMatchMessage{int(slot)})
	if err != nil {
		if debug {
			fmt.Println(err.Error())
		}
		err = conn.Close()
		if debug && (err != nil) {
			fmt.Println(err.Error())
		}
	}
}

func listeningForFriendMatches(name string, conn *websocket.Conn) {
	ch := conn.CloseHandler()
	conn.SetCloseHandler(func(code int, text string) error {
		friendMatchListenersLock.Lock()
		delete(friendMatchListeners, name)
		friendMatchListenersLock.Unlock()
		return ch(code, text)
	})
	friendMatchListenersLock.Lock()
	friendMatchListeners[name] = conn
	friendMatchListenersLock.Unlock()
}

// The name should be validated before calling this function.
func (db DB) playersFriendMatching(name string) [6]string {
	rows, err := db.Query("SELECT "+friend_friend+", "+friend_slot+" FROM "+friend_table+" WHERE "+friend_requester+"=$1;", name)
	if err != nil {
		panic(err.Error())
	}
	var friends [6]string
	for rows.Next() {
		var friend string
		var slot uint8
		err = rows.Scan(&friend, &slot)
		if err != nil {
			panic(err.Error())
		}
		friends[slot] = friend
	}
	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
	return friends
}

func (db DB) concedeFriendGame(name string, slot uint8) {
	id := db.playersGameFromFriendSlot(name, slot)
	// acknowledge sets the slot to zero
	tx := db.Begin()
	tx.setGameConceded(id)
	g := tx.gameWithIdentifier(id, true)
	(&g).acknowledgeGameComplete(name, tx)
	tx.Commit()
	// notifying with an empty diff will cause the client to ask for available moves which will show the conceded state
	gameListeningLock.RLock()
	cs, has := gameListening[id]
	if has {
		if g.White == name {
			if cs.black != nil {
				cs.black <- make(map[string]piece)
			}
		} else if g.Black == name {
			if cs.white != nil {
				cs.white <- make(map[string]piece)
			}
		} else {
			panic(fmt.Sprint(name, " player not white or black ", g.White, g.Black))
		}
	}
	gameListeningLock.RUnlock()
}

func (db DB) cancelFriendRequest(requester string, slot uint8) bool {
	tx := db.Begin()
	defer tx.Commit()
	var setSql []sql.NullInt64
	err := tx.QueryRow("SELECT "+friend_setup+" FROM "+friend_table+" WHERE "+friend_requester+" = $1 AND "+friend_slot+" = $2 FOR UPDATE;", requester, slot).Scan(pq.Array(&setSql))
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		panic(err.Error())
	}
	var setup gameSetup
	for i, item := range setSql {
		setup[i] = int(item.Int64)
	}
	// these piece identifiers should have already been vetted before setting up the initial friend match request
	db.unreservePieces(setup)
	result, err := tx.Exec("DELETE FROM "+friend_table+" WHERE "+friend_requester+" = $1 AND "+friend_slot+" = $2;", requester, slot)
	if err != nil {
		panic(err.Error())
	}
	count, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	if count != 1 {
		panic(fmt.Sprint(count, " row affected by delete"))
	}
	return true
}

// Returns -1 if the request is now pending, otherwise returns the opponent's slot for the matched game.
// The calling names should be validated before calling this function.
// Any specified pieces in the provided setup are reserved so that loss during a competitive game doesn't delete the piece until after the friend game completes.
func (db DB) friendRequest(requester string, setup gameSetup, friend string, slot uint8) int {
	defer db.reservePieces(setup)
	tx := db.Begin()
	_, err := tx.Exec("LOCK TABLE " + friend_table + ";")
	if err != nil {
		panic(err.Error())
	}
	var s uint8
	var setSql []sql.NullInt64
	err = tx.QueryRow("SELECT "+friend_slot+", "+friend_setup+" FROM "+friend_table+" WHERE "+friend_requester+" = $1 AND "+friend_friend+" = $2;", friend, requester).Scan(&s, pq.Array(&setSql))
	var set gameSetup
	for i, item := range setSql {
		set[i] = int(item.Int64)
	}
	if err == sql.ErrNoRows {
		result, err := tx.Exec("INSERT INTO "+friend_table+" ("+friend_requester+", "+friend_setup+", "+friend_friend+", "+friend_slot+") VALUES ($1, $2, $3, $4);", requester, pq.Array(setup), friend, slot)
		if err != nil {
			panic(err.Error())
		}
		count, err := result.RowsAffected()
		if err != nil {
			panic(err.Error())
		}
		if count != 1 {
			panic(fmt.Sprint(count, " rows affected by insert to ", friend_table))
		}
		tx.Commit()
		return -1
	} else if err != nil {
		panic(err.Error())
	}
	result, err := tx.Exec("DELETE FROM "+friend_table+" WHERE "+friend_requester+" = $1 AND "+friend_friend+" = $2 AND "+friend_slot+" = $3;", friend, requester, s)
	if err != nil {
		panic(err.Error())
	}
	count, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	if count != 1 {
		panic(fmt.Sprint(count, " rows affected by delete from ", friend_table))
	}
	tx.Commit()
	// set friend slot 's' to game ID, set requester slot 'slot' to game ID
	id := database.newGame(requester, setup, friend, set, false)
	if id == 0 {
		panic(fmt.Sprint("invalid setup ", requester, setup, friend, set))
	}
	database.setPlayerFriendSlot(requester, slot, id)
	database.setPlayerFriendSlot(friend, s, id)
	return int(s)
}
