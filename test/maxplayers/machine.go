package main

import (
	"fmt"
	"os"

	"github.com/pciet/wichess"
	"github.com/pciet/wichess/rules"
	"github.com/pciet/wichess/test/client"
)

// A state machine can be implemented as recursive functions, where a state transition is done
// by calling the next function. A function state machine has an unbounded function call stack.
// Here a state machine is implemented as goroutines communicating with channels instead.

// A Signal is a communication on a Go channel without data.
type Signal struct{}

// StartPlayGameMachine starts a state machine made of communicating goroutines that represent
// states of communication with the host, including getting moves, making moves, doing a
// promotion, and waiting for an alert on the WebSocket.
func StartPlayGameMachine(done chan error, id wichess.GameIdentifier,
	with client.Instance, records MeasurementChans) {

	DebugPrintln("DIAL", with.Name)

	err := (&with).DialWebSocket(id)
	if err != nil {
		fmt.Println("DialWebSocket:", err)
		os.Exit(1)
	}

	DebugPrintln("DIALED", with.Name)

	// TODO: race condition where alert is lost during the websocket dial?

	// state goroutines that read from a channel are responsible for closing it
	stop := make(chan Signal)
	machineDone := make(chan error)
	moves := make(chan Signal)
	listen := make(chan Signal)
	promote := make(chan Signal)
	move := make(chan []rules.MoveSet)

	// these are in the machine_[statename].go files
	go DoneState(done, stop, machineDone, with, id)
	go WebSocketState(stop, machineDone, moves, listen, with)
	go MovesState(stop, machineDone, promote, move, moves, records.Moves, with, id)
	go MoveState(stop, machineDone, listen, promote, move, records.Move, with, id)
	go PromoteState(stop, machineDone, listen, promote, with, id)

	active, err := with.ActivePlayer(id)
	if err != nil {
		fmt.Println("ActivePlayer:", err)
		os.Exit(1)
	}

	if active {
		moves <- Signal{}
	} else {
		listen <- Signal{}
	}
}

// DoneState receives an error condition or game done signal (draw, checkmate, concede) on the
// done channel. All other states are closed when this state closes the close channel. This state
// also signals the caller of StartPlayGameMachine that the machine has stopped on gameDone.
func DoneState(gameDone chan<- error, stop chan Signal, done <-chan error,
	with client.Instance, id wichess.GameIdentifier) {

	err := <-done
	close(stop)

	DebugPrintln("CLOSE WEBSOCKET", with.Name)

	(&with).CloseWebSocket()

	DebugPrintln("CLOSED WEBSOCKET", with.Name)

	if err == nil {
		DebugPrintln("ACK", with.Name)
		with.Acknowledge(id)
		DebugPrintln("ACKED", with.Name)
	}

	gameDone <- err
}
