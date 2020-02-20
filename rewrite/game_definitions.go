package main

import (
	"time"

	"github.com/pciet/wichess/rules"
)

type GameIdentifier int

type GameHeader struct {
	ID             GameIdentifier  `json:"-"`
	Competitive    bool            `json:"-"`
	PrizePiece     rules.PieceKind `json:"-"`
	Recorded       bool            `json:"-"`
	Conceded       bool            `json:"-"`
	White          PlayerGameHeader
	Black          PlayerGameHeader
	Active         rules.Orientation
	PreviousActive rules.Orientation
	From           int
	To             int
	DrawTurns      int `json:"-"`
	Turn           int
}

type PlayerGameHeader struct {
	Name           string
	Acknowledge    bool `json:"-"`
	LatestMove     time.Time
	Elapsed        time.Duration
	ElapsedUpdated time.Time
}

type Game struct {
	GameHeader
	rules.Game
	PieceIdentifiers []AddressedPieceIdentifier
}

// value set for an initialized game with no move in from/to
const no_move = 64

var basic_army = func() [16]rules.PieceKind {
	var b [16]rules.PieceKind
	for i := 0; i < 8; i++ {
		b[i] = rules.Pawn
	}

	b[8] = rules.Rook
	b[15] = rules.Rook

	b[9] = rules.Knight
	b[14] = rules.Knight

	b[10] = rules.Bishop
	b[13] = rules.Bishop

	b[11] = rules.Queen
	b[12] = rules.King

	return b
}()
