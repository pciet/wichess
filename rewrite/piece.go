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
)
