package wichess

import (
	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

const (
	CollectionCount = 21

	NotInCollection CollectionSlot = 0

	// The two random picks are considered part of the player's collection.
	LeftPick  CollectionSlot = -1
	RightPick CollectionSlot = -2
)

type (
	Collection [CollectionCount]Piece

	// The CollectionSlot is the array index into Collection plus one to be a human readable value.
	CollectionSlot int8

	AddressedCollectionSlot struct {
		Slot          CollectionSlot `json:"s"`
		rules.Address `json:"a"`
	}

	RandomPicks struct {
		Left, Right piece.Kind
	}
)

func (a Collection) Kinds() [CollectionCount]piece.Kind {
	var out [CollectionCount]piece.Kind
	for i, p := range a {
		out[i] = p.Kind
	}
	return out
}

func (a CollectionSlot) Int() int { return int(a) }
