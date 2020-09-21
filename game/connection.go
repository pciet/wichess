package game

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"

	"github.com/pciet/wichess/memory"
)

// TODO: documentation for sync suggests that channels can be a better sync mechanism
// TODO: delete map entry on game done

var (
	// The connections map holds the WebSocket pointer for each player indexed by rules.Orientation.
	// If a player isn't connected then the pointer is nil or an error will be returned when the
	// Conn is used. If both players aren't connected then the game might not be in the map.
	connections      = map[memory.GameIdentifier][2]*websocket.Conn{}
	connectionsMutex = sync.RWMutex{}
)

// Connected returns the WebSocket for the player of an orientation or nil if the player isn't
// connected. It's possible the connection is closed even with a non-nil return.
func Connected(id memory.GameIdentifier, o rules.Orientation) *websocket.Conn {
	connectionsMutex.RLock()
	defer connectionsMutex.RUnlock()
	gameConns, has := connections[id]
	if has == false {
		return nil
	}
	return gameConns[o]
}

// Connect adds or replaces a player's WebSocket connection for a game. If replaced then the
// original WebSocket is closed.
func Connect(id memory.GameIdentifier, o rules.Orientation, with *websocket.Conn) {
	if with == nil {
		log.Panicln("no WebSocket for %v in %v", o, id)
	}

	connectionsMutex.Lock()
	defer connectionsMutex.Unlock()

	gameConns := connections[id]
	if gameConns[o] != nil {
		conn.Close()
	}
	gameConns[o] = with
	Connections[id] = gameConns

	// TODO: does this goroutine return when the WebSocket is replaced?

	go func() {
		for {
			// the webpages don't send anything on the WebSocket
			_, _, err := with.NextReader()
			if err != nil {
				with.Close()
				break
			}
		}
	}()
}

// RemoveConnection removes a player's WebSocket connection for a game.
func RemoveConnection(id memory.GameIdentifier, o rules.Orientation) {
	connectionsMutex.Lock()
	defer connectionsMutex.Unlock()

	gameConns := connections[id]
	if gameConns[o] == nil {
		return
	}
	gameConns[o].Close()
	gameConns[o] = nil
	connections[id] = gameConns
}
