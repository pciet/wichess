package main

import (
	"github.com/pciet/wichess"
	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/test/client"
)

func PromoteState(stop chan Signal, done chan<- error,
	listen chan<- Signal, promote <-chan Signal, with client.Instance, id wichess.GameIdentifier) {

	for {
		select {
		case <-stop:
			return
		case <-promote:
			RandomHumanWait()
			DebugPrintln("PROMOTE", id, with.Name)

			err := with.Promote(id, piece.Queen)
			if err != nil {
				done <- err
				break
			}

			DebugPrintln("PROMOTED", id, with.Name)

			listen <- Signal{}
		}
	}
}
