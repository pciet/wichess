package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var WebSocketUpgrader = websocket.Upgrader{1024, 1024}

func WebSocketUpgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return WebSocketUpgrader.Upgrade(w, r, nil)
}
