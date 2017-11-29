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

func listeningToGame(name string, white string, black string, totalTime time.Duration, previousMove time.Time, id int, socket *websocket.Conn) {
	gameListeningLock.Lock()
	defer gameListeningLock.Unlock()
	_, has := gameListening[id]
	if has == false {
		gameListening[id] = &gameListeners{}
		d := make(chan struct{})
		if totalTime > time.Duration(0) {
			gameMonitorsLock.Lock()
			gameMonitors[id] = gameMonitor{
				done: d,
				move: make(chan time.Time),
			}
			go func(channels gameMonitor, gameid int, total time.Duration, move time.Time) {
				for {
					lockGame(gameid)
					g := database.gameWithIdentifier(gameid)
					if (g.White != name) && (g.Black != name) {
						unlockGame(gameid)
						gameMonitorsLock.Lock()
						delete(gameMonitors, gameid)
						gameMonitorsLock.Unlock()
						return
					}
					b := wichessingBoard(g.Points)
					active := g.activeOrientation()
					activePlayer := g.Active
					if (g.DrawTurns >= draw_turn_count) || b.Draw(active, wichessing.AbsPointFromIndex(uint8(g.From)), wichessing.AbsPointFromIndex(uint8(g.To))) || b.Checkmate(active, wichessing.AbsPointFromIndex(uint8(g.From)), wichessing.AbsPointFromIndex(uint8(g.To))) || g.timeLoss(active, total) {
						unlockGame(gameid)
						gameMonitorsLock.Lock()
						delete(gameMonitors, gameid)
						gameMonitorsLock.Unlock()
						return
					}
					unlockGame(gameid)
					elapsed := g.orientationsElapsedTime(active)
					select {
					case <-channels.done:
						return
					case <-channels.move:
					case <-time.After(total - elapsed):
						// between hitting this case and reading the game in updateGameTimes the active player may have switched
						lockGame(gameid)
						_ = g.DB.updateGameTimes(gameid, total, activePlayer)
						unlockGame(gameid)
						// by sending an empty notification the client will request /moves, which says time has expired
						gameListeningLock.RLock()
						cs, has := gameListening[gameid]
						if has {
							if cs.white != nil {
								cs.white <- make(map[string]piece)
							}
							if cs.black != nil {
								cs.black <- make(map[string]piece)
							}
						}
						gameMonitorsLock.Lock()
						delete(gameMonitors, gameid)
						gameMonitorsLock.Unlock()
						gameListeningLock.RUnlock()
						return
					}
				}
			}(gameMonitors[id], id, totalTime, previousMove)
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
			diff, ok := <-listenTo
			if ok == false {
				_ = conn.Close()
				// TODO: logic duplicated below for a connection error/close
				gameListeningLock.Lock()
				cs, has := gameListening[id]
				if has == false {
					gameListeningLock.Unlock()
					return
				}
				if name == white {
					cs.white = nil
				} else {
					cs.black = nil
				}
				if (cs.white == nil) && (cs.black == nil) {
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
			err := conn.WriteJSON(diff)
			if err != nil {
				_ = conn.Close()
				gameListeningLock.Lock()
				cs, has := gameListening[id]
				if has == false {
					gameListeningLock.Unlock()
					return
				}
				if name == white {
					cs.white = nil
				} else {
					cs.black = nil
				}
				if (cs.white == nil) && (cs.black == nil) {
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
