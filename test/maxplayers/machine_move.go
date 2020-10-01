package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pciet/wichess/game"
	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/rules"
	"github.com/pciet/wichess/test/client"
)

func MoveState(stop chan Signal, done chan<- error,
	listen chan<- Signal, promote chan<- Signal, move <-chan []rules.MoveSet,
	record chan<- time.Duration, with client.Instance, id memory.GameIdentifier) {

	for {
		select {
		case <-stop:
			return
		case moveSets := <-move:
			RandomHumanWait()
			m := PickRandomMove(moveSets)

			DebugPrintln("MOVE", id, with.Name, m)

			before := time.Now()
			state, err := with.Move(id, m)
			if err != nil {
				done <- err
				break
			}
			record <- time.Since(before)

			DebugPrintln("MOVED", id, with.Name, m)

			switch state {
			case game.PromotionNeededUpdate:
				promote <- Signal{}
			case "", game.ContinueUpdate:
				listen <- Signal{}
			default:
				done <- fmt.Errorf("unknown move game state %v", state)
			}
		}
	}
}

func PickRandomMove(from []rules.MoveSet) rules.Move {
	moves := make([]rules.Move, 0, 16)
	for _, set := range from {
		for _, to := range set.Moves {
			moves = append(moves, rules.Move{set.From, to})
		}
	}
	return moves[rand.Intn(len(moves))]
}
