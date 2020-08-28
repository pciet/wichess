package client

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pciet/wichess"
)

func (an *Instance) DialWebSocket(id wichess.GameIdentifier) error {
	var resp *http.Response
	var err error
	an.Conn, resp, err = an.Dialer.Dial(an.WebSocketHost()+wichess.AlertPath+id.String(), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusSwitchingProtocols {
		return fmt.Errorf("bad WebSocket upgrade status %v on %v",
			resp.StatusCode, wichess.AlertPath)
	}

	return nil
}

// WebSocketReadState returns the wichess.Update State string which is "" for a normal state.
func (an Instance) WebSocketReadState() (string, error) {
	var u wichess.Update
	err := an.Conn.ReadJSON(&u)
	return u.State, err
}

func (an *Instance) CloseWebSocket() error {
	err := an.Conn.Close()
	an.Conn = nil
	return err
}
