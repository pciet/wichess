package main

import (
	"sync"

	"github.com/gorilla/websocket"
)

type (
	Connection struct {
		Username        string
		*websocket.Conn // nil for no connection
	}

	// Each game can have a WebSocket connection for each player,
	// represented as an unordered array called GameConnections.
	GameConnections [2]Connection
)

// TODO: test that the Connections map correctly grows and shrinks
// TODO: documentation for sync suggests that channels can be a better sync mechanism

var (
	// All WebSocket connections for games are held in the Connections map.
	Connections      = map[GameIdentifier]GameConnections{}
	ConnectionsMutex = sync.RWMutex{}
)

// If the player is currently connected to the host for the specified game
// then that WebSocket is returned by Connected, otherwise false is returned.
func Connected(id GameIdentifier, player string) (*websocket.Conn, bool) {
	ConnectionsMutex.RLock()
	defer ConnectionsMutex.RUnlock()

	gcs, has := Connections[id]
	if has == false {
		return nil, false
	}

	if gcs[0].Username == player {
		if gcs[0].Conn == nil {
			return nil, false
		}
		return gcs[0].Conn, true
	}

	if gcs[1].Username != player {
		Panic(player, "not in game", id)
	}

	if gcs[1].Conn == nil {
		return nil, false
	}
	return gcs[1].Conn, true
}

// Connect adds or replaces a player's connection for a game.
// If a connection is replaced then the original is closed.
func Connect(id GameIdentifier, player string, add *websocket.Conn) {
	ConnectionsMutex.Lock()
	defer ConnectionsMutex.Unlock()

	gcs, has := Connections[id]
	if has == false {
		Connections[id] = GameConnections{Connection{player, add}}
		return
	}

	replace := func(c Connection, wc *websocket.Conn) Connection {
		if c.Conn != nil {
			DebugPrintln("replaced WebSocket for", player)
			err := c.Close()
			if err != nil {
				DebugPrintln(player, "WebSocket close error:", err)
			}
		}
		c.Conn = wc
		return c
	}

	if gcs[0].Username == player {
		gcs[0] = replace(gcs[0], add)
	} else if gcs[1].Username == player {
		gcs[1] = replace(gcs[1], add)
	} else if gcs[0].Username == "" {
		gcs[0] = Connection{player, add}
	} else if gcs[1].Username == "" {
		gcs[1] = Connection{player, add}
	} else {
		Panic(player, "can't be put into game connections", gcs)
	}

	Connections[id] = gcs

	go ConnectionCloseWait(add, id, player)
}

// ConnectionCloseWait waits for the web browser to close the WebSocket then
// updates the Connections map. Any message received from the web browser
// is assumed to be a close message.
func ConnectionCloseWait(conn *websocket.Conn, id GameIdentifier, player string) {
	// blocks until close message received
	conn.ReadMessage()

	ConnectionsMutex.Lock()
	defer ConnectionsMutex.Unlock()

	gcs, has := Connections[id]
	if has == false {
		Panic("game", id, "got WebSocket close from", player, "but no connections are tracked")
	}

	close := func(c Connection) Connection {
		err := c.Close()
		if err != nil {
			DebugPrintln("error when responding to WebSocket close for", player, ":", err)
		}
		c.Conn = nil
		return c
	}

	if gcs[0].Username == player {
		gcs[0] = close(gcs[0])
	} else if gcs[1].Username == player {
		gcs[1] = close(gcs[1])
	} else {
		Panic(player, "closed untracked WebSocket in game", id)
	}

	if (gcs[0].Conn == nil) && (gcs[1].Conn == nil) {
		delete(Connections, id)
	} else {
		Connections[id] = gcs
	}
}
