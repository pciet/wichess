package main

import (
	"fmt"

	"github.com/pciet/wichess"
	"github.com/pciet/wichess/test/client"
)

func WebSocketState(stop chan Signal, done chan<- error,
	moves chan<- Signal, listen <-chan Signal,
	with client.Instance) {

	for {
		select {
		case <-stop:
			return
		case <-listen:
			DebugPrintln("WEBSOCKET WAIT", with.Name)

			s, err := with.WebSocketReadState()
			if err != nil {
				done <- err
				break
			}

			DebugPrintln("WEBSOCKET READ", with.Name)

			switch s {
			case wichess.WaitUpdate:
				break
			case wichess.DrawCalculatedUpdate, wichess.CheckmateCalculatedUpdate,
				wichess.ConcededUpdate:
				done <- nil
			case "", wichess.ContinueUpdate, wichess.CheckCalculatedUpdate:
				moves <- Signal{}
			default:
				done <- fmt.Errorf("unknown WebSocket update state %v", s)
			}
		}
	}
}
