package client

import (
	"math/rand"

	"github.com/pciet/wichess"
	"github.com/pciet/wichess/piece"
)

// RandomFullArmyRequest makes an army request with the most pieces possible picked randomly.
func (an Instance) RandomFullArmyRequest() (wichess.ArmyRequest, error) {
	var out wichess.ArmyRequest

	// always use picks
	left, right, err := an.Picks()
	if err != nil {
		return wichess.ArmyRequest{}, err
	}

	assignRandomArmySlot := func(k piece.Kind, slot wichess.CollectionSlot) {
		slots := out.AvailableSlotsForKind(k)
		if len(slots) == 0 {
			return
		}
		out[slots[rand.Intn(len(slots))]] = slot
	}

	assignRandomArmySlot(left, wichess.LeftPick)
	assignRandomArmySlot(right, wichess.RightPick)

	collection, err := an.Collection()
	if err != nil {
		return wichess.ArmyRequest{}, err
	}

	kinds, collectionSlots := RandomizeCollectionOrder(collection)

	for i, k := range kinds {
		assignRandomArmySlot(k, collectionSlots[i])
	}

	return out, nil
}
