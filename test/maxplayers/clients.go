package main

import (
	"fmt"

	"github.com/pciet/wichess/test/client"
)

// ClientPair represents two automated players that play against each other.
type ClientPair struct {
	A, B client.Instance
}

func ClientPairs(host string, playerCount int) []ClientPair {
	count := playerCount / 2
	players := make([]ClientPair, count)
	for i := 0; i < count; i++ {
		players[i] = ClientPair{
			A: client.New(host, fmt.Sprintf("pa%d", i), fmt.Sprintf("pwa%d", i)),
			B: client.New(host, fmt.Sprintf("pb%d", i), fmt.Sprintf("pwb%d", i)),
		}
	}
	return players
}
