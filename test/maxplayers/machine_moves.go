package main

import (
	"fmt"
	"time"

	"github.com/pciet/wichess"
	"github.com/pciet/wichess/rules"
	"github.com/pciet/wichess/test/client"
)

func MovesState(stop chan Signal, done chan<- error,
	promote chan<- Signal, move chan<- []rules.MoveSet, moves <-chan Signal,
	record chan<- time.Duration, with client.Instance, id wichess.GameIdentifier) {

	for {
		select {
		case <-stop:
			return
		case <-moves:
			RandomHumanWait()

			DebugPrintln("MOVES", id, with.Name)

			before := time.Now()
			movs, state, err := with.Moves(id)
			if err != nil {
				done <- err
				break
			}
			record <- time.Since(before)

			DebugPrintln("GOT MOVES", id, with.Name)

			switch state {
			case rules.Checkmate, rules.Draw:
				done <- nil
			case rules.Promotion:
				promote <- Signal{}
			case rules.Normal, rules.Check:
				move <- movs
			default:
				done <- fmt.Errorf("unknown moves game state %v", state)
			}
		}
	}
}
