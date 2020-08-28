// The test/maxplayers tool is a benchmark of the average and maximum response time for varying
// numbers of simulated parallel players. The results should help determine the maximum number of
// parallel players to claim for a version of Wisconsin Chess on a platform.
//
// Players are simulated by making random moves in times evenly distributed between 0.3 and 3
// seconds to try to represent the fastest people that could be playing together.
package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var Debug = false

func main() {
	playerCount := flag.Int("count", 2, "even `number` of players to simulate")
	hostAddress := flag.String("host", "localhost:8080", "target host network `address`")
	lengthSeconds := flag.Int("length", 5, "`seconds` to run the benchmark")
	debugEnabled := flag.Bool("debug", false, "print game state and network requests")
	flag.Parse()

	if *playerCount%2 != 0 {
		fmt.Println("odd player count", *playerCount)
		os.Exit(1)
	}

	if *lengthSeconds == 0 {
		fmt.Println("zero length benchmark")
		os.Exit(1)
	}

	Debug = *debugEnabled

	// each MatchAndPlay goroutine sends duration from request to response measurements
	out := MeasurementChans{make(chan time.Duration), make(chan time.Duration)}

	pairs := ClientPairs(*hostAddress, *playerCount)

	// wait for all clients to have an account and session before starting the benchmark
	for _, pair := range pairs {
		pair.A.Login()
		pair.B.Login()
	}

	for _, pair := range pairs {
		go MatchAndPlay(pair, out)
	}

	done := time.After(time.Second * time.Duration(*lengthSeconds))

	var moveMax, movesMax, moveTotal, movesTotal time.Duration
	var moveCount, movesCount int64

LOOP:
	for {
		select {
		case <-done:
			break LOOP
		case r := <-out.Move:
			if r > moveMax {
				moveMax = r
			}
			moveTotal += r
			moveCount++
		case r := <-out.Moves:
			if r > movesMax {
				movesMax = r
			}
			movesTotal += r
			movesCount++
		}
	}

	fmt.Println("players", *playerCount)
	fmt.Println("length seconds", *lengthSeconds)
	fmt.Println("host", *hostAddress)
	fmt.Println("moves")
	fmt.Println("  max", movesMax.Truncate(time.Microsecond))
	fmt.Println("  avg", (movesTotal / time.Duration(movesCount)).Truncate(time.Microsecond))

	fmt.Println("move")
	fmt.Println("  max", moveMax.Truncate(time.Microsecond))
	fmt.Println("  avg", (moveTotal / time.Duration(moveCount)).Truncate(time.Microsecond))
}
