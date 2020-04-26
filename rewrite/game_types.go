package main

import "github.com/pciet/wichess/rules"

type (
	// Each game has a unique GameIdentifier used to find it in the database.
	GameIdentifier int

	// TODO: make Active and PreviousActive rules.Orientation

	// A GameHeader is associated with a Board to describe the current status of a game.
	// The header does not have any information about the board position (which piece is where).
	GameHeader struct {
		ID             GameIdentifier `json:"-"`
		Conceded       bool           `json:"-"`
		White          GamePlayerHeader
		Black          GamePlayerHeader
		Active         string
		PreviousActive string
		From           rules.AddressIndex
		To             rules.AddressIndex
		DrawTurns      int `json:"-"`
		Turn           int
	}

	GamePlayerHeader struct {
		Name        string
		Acknowledge bool `json:"-"`
	}

	Game struct {
		Header GameHeader
		Board
	}
)
