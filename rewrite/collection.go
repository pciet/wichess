package main

import "github.com/pciet/wichess/rules"

const (
	NotInCollection = 0
	CollectionCount = 24
)

type (
	Collection [CollectionCount]Piece

	// The CollectionSlot is the array index
	// into Collection plus one to be a human
	// readable value. Zero is used to
	// indicate the piece is not in the collection.
	CollectionSlot uint8

	AddressedCollectionSlot struct {
		Slot          CollectionSlot `json:"s"`
		rules.Address `json:"a"`
	}
)
