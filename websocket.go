package wichess

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var WebSocketUpgrader = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

func WebSocketUpgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return WebSocketUpgrader.Upgrade(w, r, nil)
}
