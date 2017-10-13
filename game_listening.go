// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pciet/wichess/wichessing"
)

type gameListeners struct {
	white chan map[string]piece
	black chan map[string]piece
}

type gameMonitor struct {
	done chan struct{}
	move chan time.Time
}

var (
	gameListening     = make(map[int]*gameListeners)
	gameListeningLock = sync.RWMutex{}

	gameMonitors     = make(map[int]gameMonitor)
	gameMonitorsLock = sync.RWMutex{}
)

func listeningToGame(name string, white string, black string, turnTime time.Duration, previousMove time.Time, id int, socket *websocket.Conn) {
	gameListeningLock.Lock()
	defer gameListeningLock.Unlock()
	_, has := gameListening[id]
	if has == false {
		gameListening[id] = &gameListeners{}
		d := make(chan struct{})
		if turnTime > time.Duration(0) {
			gameMonitorsLock.Lock()
			gameMonitors[id] = gameMonitor{
				done: d,
				move: make(chan time.Time),
			}
			go func(channels gameMonitor, gameid int, turn time.Duration, move time.Time) {
				g := database.gameWithIdentifier(gameid)
				for {
					b := wichessingBoard(g.Points)
					var active wichessing.Orientation
					if g.Active == g.White {
						active = wichessing.White
					} else {
						active = wichessing.Black
					}
					if b.Draw(active) || b.Checkmate(active) {
						gameMonitorsLock.Lock()
						delete(gameMonitors, gameid)
						gameMonitorsLock.Unlock()
						return
					}
					timeout := move.Add(turn)
					select {
					case <-channels.done:
						return
					case move = <-channels.move:
						g = database.gameWithIdentifier(gameid)
					case <-time.After(timeout.Sub(time.Now())):
						move = time.Now()
						g = database.gameWithIdentifier(gameid).randomMoveAtTime(move)
					}
				}
			}(gameMonitors[id], id, turnTime, previousMove)
			gameMonitorsLock.Unlock()
		}
	}
	var l chan map[string]piece
	if name == white {
		gameListening[id].white = make(chan map[string]piece)
		l = gameListening[id].white
	} else if name == black {
		gameListening[id].black = make(chan map[string]piece)
		l = gameListening[id].black
	} else {
		panicExit("unexpected name " + name)
	}
	go func(listenTo chan map[string]piece, conn *websocket.Conn) {
		for {
			err := conn.WriteJSON(<-listenTo)
			if err != nil {
				gameListeningLock.Lock()
				if name == white {
					gameListening[id].white = nil
				} else {
					gameListening[id].black = nil
				}
				if (gameListening[id].white == nil) && (gameListening[id].black == nil) {
					delete(gameListening, id)
					gameMonitorsLock.Lock()
					monitor, has := gameMonitors[id]
					if has {
						monitor.done <- struct{}{}
						delete(gameMonitors, id)
					}
					gameMonitorsLock.Unlock()
				}
				gameListeningLock.Unlock()
				return
			}
		}
	}(l, socket)
}
