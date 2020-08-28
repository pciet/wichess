// The test/client package simulates one or more HTTP clients of wichess.
package client

import (
	"net/http"
	"net/http/cookiejar"

	"github.com/gorilla/websocket"
)

type Instance struct {
	Name, Password, Host string
	*http.Client
	*cookiejar.Jar
	*websocket.Dialer
	*websocket.Conn
}

// New allocates memory for an instance of a wichess client. The host string is the network name
// without the protocol, like "localhost:8080".
func New(host, username, password string) Instance {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err.Error())
	}
	inst := Instance{
		Name:     username,
		Password: password,
		Host:     host,
		Jar:      jar,
	}
	inst.Client = &http.Client{Jar: jar}
	inst.Dialer = &websocket.Dialer{Jar: jar}
	return inst
}

func (an Instance) WebSocketHost() string { return "ws://" + an.Host }
func (an Instance) HTTPHost() string      { return "http://" + an.Host }
