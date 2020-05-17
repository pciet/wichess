package main

import "github.com/pciet/wichess/rules"

// A Board represents the 64 squares of an active game, along with a list of the pieces that
// are in the players' collections.
type Board struct {
	rules.Board      `json:"b"`
	CollectionPieces []AddressedCollectionSlot `json:"pi"`
}
