package main

import "github.com/pciet/wichess/rules"

// A Board represents the 64 squares of an active game.
// Pieces that have an identifier (special pieces collected by a player) are also listed.
type Board struct {
	rules.Board      `json:"b"`
	PieceIdentifiers []AddressedPieceIdentifier `json:"pi"`
}
