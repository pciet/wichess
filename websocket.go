package wichess

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var webSocketUpgrader = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

func webSocketUpgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return webSocketUpgrader.Upgrade(w, r, nil)
}
