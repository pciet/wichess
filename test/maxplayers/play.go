package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pciet/wichess/test/client"
)

type MeasurementChans struct {
	Move  chan time.Duration
	Moves chan time.Duration
}

// TODO: use special pieces, add an extra helper HTTP path for the collection list

// MatchAndPlay does an infinite loop of matching the players and randomly playing games.
// Time measurements of making moves and getting legal moves are sent on the measurement channels.
func MatchAndPlay(players ClientPair, r MeasurementChans) {
	for {
		DebugPrintln("MATCHING", players.A.Name, players.B.Name)

		id, err := client.Match(players.A, players.B)
		if err != nil {
			fmt.Println("Match:", err)
			os.Exit(1)
		}

		DebugPrintln("MATCHED", players.A.Name, players.B.Name, id)

		done := make(chan error)

		go StartPlayGameMachine(done, id, players.A, r)
		go StartPlayGameMachine(done, id, players.B, r)

		err = <-done
		if err != nil {
			fmt.Println("PlayGameMachine:", err)
			os.Exit(1)
		}

		err = <-done
		if err != nil {
			fmt.Println("PlayGameMachine:", err)
			os.Exit(1)
		}

		close(done)
	}
}
