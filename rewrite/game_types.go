package main

import (
	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

type (
	// Each game has a unique GameIdentifier used to find it in the database.
	GameIdentifier int

	// A GameHeader is associated with a Board to describe the current status of a game.
	// The header does not have any information about the board position (which piece is where).
	GameHeader struct {
		ID             GameIdentifier `json:"-"`
		Conceded       bool           `json:"-"`
		White          GamePlayerHeader
		Black          GamePlayerHeader
		Active         rules.Orientation
		PreviousActive rules.Orientation
		From           rules.AddressIndex
		To             rules.AddressIndex
		DrawTurns      int `json:"-"`
		Turn           int
	}

	// Captured pieces are ordered by when they were taken, so newly captured pieces are added
	// to the first index that's still the initial value of rules.NoKind.
	Captures [15]piece.Kind

	GamePlayerHeader struct {
		Name string
		Captures
		Acknowledge bool `json:"-"`
		// if a player chooses their left and/or right random pick pieces then they will be added
		// to the player's collection if they complete the game
		Left   piece.Kind `json:"-"`
		Right  piece.Kind `json:"-"`
		Reward piece.Kind `json:"-"`
	}

	Game struct {
		Header GameHeader
		Board
	}
)
