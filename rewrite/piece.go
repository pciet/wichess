package main

import (
	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

type (
	Piece struct {
		Slot        CollectionSlot `json:"s"`
		rules.Piece `json:"p"`
	}

	AddressedPiece struct {
		rules.Address `json:"a"`
		Piece         `json:"p"`
	}

	CapturedPiece struct {
		rules.Orientation `json:"o"`
		piece.Kind        `json:"k"`
		// CaptureSlot is the slot in a game's captured list for a player. The array index starts
		// at 1 instead of 0.
		CaptureSlot int `json:"-"`
	}
)
