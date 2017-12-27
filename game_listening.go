// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
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
					tx := database.Begin()
					g := tx.gameWithIdentifier(gameid, false)
					tx.Commit()
					if (g.White != name) && (g.Black != name) {
						if debug {
							fmt.Printf("player %v not white %v or black %v\n", name, g.White, g.Black)
						}
						gameMonitorsLock.Lock()
						delete(gameMonitors, gameid)
						gameMonitorsLock.Unlock()
						return
					}
					b := wichessingBoard(g.Points, g.From, g.To)
					active := g.activeOrientation()
					activePlayer := g.Active
					tx = database.Begin()
					if (g.DrawTurns >= draw_turn_count) || b.Draw(active) || b.Checkmate(active) || g.timeLoss(active, total, tx) || g.Conceded {
						notifyFriendMoveAvailable(g.White, g.ID)
						notifyFriendMoveAvailable(g.Black, g.ID)
						tx.Commit()
						// this pattern also needed here in case a final move is made but then another move is sent before the game can be torn down
						acq := make(chan struct{})
						go func() {
							gameMonitorsLock.Lock()
							acq <- struct{}{}
						}()
					OUTER1:
						for {
							select {
							case <-channels.move:
							case <-channels.done:
							case <-acq:
								break OUTER1
							}
						}
						delete(gameMonitors, gameid)
						gameMonitorsLock.Unlock()
						return
					}
					tx.Commit()
					elapsed := g.orientationsElapsedTime(active)
					select {
					case <-channels.done:
						return
					case <-channels.move:
					case <-time.After(total - elapsed):
						// between hitting this case and reading the game in updateGameTimes the active player may have switched
						_ = g.DB.updateGameTimes(gameid, total, activePlayer)
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
						gameListeningLock.RUnlock()
						// if the lock can't be acquired we have to try a read here because the notifier could be holding it waiting to send
						// this won't work serially: lock blocks us from trying to read and trying to read first gives the notifier a chance to lock before we do
						acq := make(chan struct{})
						go func() {
							gameMonitorsLock.Lock()
							acq <- struct{}{}
						}()
					OUTER2:
						for {
							select {
							case <-channels.move:
							case <-channels.done:
							case <-acq:
								break OUTER2
							}
						}
						delete(gameMonitors, gameid)
						gameMonitorsLock.Unlock()
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
					gameListeningLock.Unlock()
					gameMonitorsLock.Lock()
					monitor, has := gameMonitors[id]
					if has {
						monitor.done <- struct{}{}
						delete(gameMonitors, id)
					}
					gameMonitorsLock.Unlock()
				} else {
					gameListeningLock.Unlock()
				}
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
					gameListeningLock.Unlock()
					gameMonitorsLock.Lock()
					monitor, has := gameMonitors[id]
					if has {
						monitor.done <- struct{}{}
						delete(gameMonitors, id)
					}
					gameMonitorsLock.Unlock()
				} else {
					gameListeningLock.Unlock()
				}
				return
			}
		}
	}(l, socket)
}
