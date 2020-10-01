package main

import (
	"fmt"

	"github.com/pciet/wichess/game"
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
		LISTEN:
			DebugPrintln("WEBSOCKET READ", with.Name)

			s, err := with.WebSocketReadState()
			if err != nil {
				done <- err
				break
			}

			DebugPrintln("WEBSOCKET READ DONE", with.Name)

			switch s {
			case game.WaitUpdate:
				DebugPrintln("  WAIT", with.Name)
				goto LISTEN
			case game.CheckCalculatedUpdate:
				DebugPrintln("  CHECK", with.Name)
				goto LISTEN
			case game.DrawCalculatedUpdate, game.CheckmateCalculatedUpdate, game.ConcededUpdate:
				DebugPrintln("  DONE", with.Name)
				done <- nil
			case "", game.ContinueUpdate:
				DebugPrintln("  CONTINUE", with.Name)
				moves <- Signal{}
			default:
				DebugPrintln("  WEBSOCKET BAD UPDATE STATE", with.Name)
				done <- fmt.Errorf("unknown WebSocket update state %v", s)
			}
		}
	}
}
