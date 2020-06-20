package main

import "github.com/pciet/wichess/rules"

type (
	Piece struct {
		Slot        CollectionSlot `json:"s"`
		InUse       bool           `json:"-"`
		rules.Piece `json:"p"`
	}

	AddressedPiece struct {
		rules.Address `json:"a"`
		Piece         `json:"p"`
	}

	CapturedPiece struct {
		rules.Orientation `json:"o"`
		rules.PieceKind   `json:"k"`
		// CaptureSlot is the slot in a game's captured list for a player. The array index starts
		// at 1 instead of 0.
		CaptureSlot int `json:"-"`
	}
)
