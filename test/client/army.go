package client

import (
	"math/rand"

	"github.com/pciet/wichess/piece"
)

// RandomFullArmyRequest makes an army request with the most pieces possible picked randomly.
func (an Instance) RandomFullArmyRequest() (piece.ArmyRequest, error) {
	var out piece.ArmyRequest

	// always use picks
	left, right, err := an.Picks()
	if err != nil {
		return piece.ArmyRequest{}, err
	}

	assignRandomArmySlot := func(k piece.Kind, slot piece.CollectionSlot) {
		slots := out.AvailableSlotsForKind(k)
		if len(slots) == 0 {
			return
		}
		out[slots[rand.Intn(len(slots))]] = slot
	}

	assignRandomArmySlot(left, piece.LeftPick)
	assignRandomArmySlot(right, piece.RightPick)

	collection, err := an.Collection()
	if err != nil {
		return piece.ArmyRequest{}, err
	}

	kinds, collectionSlots := RandomizeCollectionOrder(collection)

	for i, k := range kinds {
		assignRandomArmySlot(k, collectionSlots[i])
	}

	return out, nil
}
