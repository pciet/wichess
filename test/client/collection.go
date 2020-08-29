package client

import (
	"encoding/json"
	"math/rand"

	"github.com/pciet/wichess"
	"github.com/pciet/wichess/piece"
)

func (an Instance) Picks() (piece.Kind, piece.Kind, error) {
	respBody, err := an.JSONResponseGet(wichess.PicksPath)
	if err != nil {
		return piece.NoKind, piece.NoKind, err
	}

	var j wichess.PicksJSON
	err = json.Unmarshal([]byte(respBody), &j)
	if err != nil {
		return piece.NoKind, piece.NoKind, err
	}

	return j.Left, j.Right, nil
}

func (an Instance) Collection() (wichess.Collection, error) {
	respBody, err := an.JSONResponseGet(wichess.CollectionPath)
	if err != nil {
		return wichess.Collection{}, err
	}

	var j wichess.CollectionJSON
	err = json.Unmarshal([]byte(respBody), &j)
	if err != nil {
		return wichess.Collection{}, err
	}

	return j.Collection, nil
}

func RandomizeCollectionOrder(of wichess.Collection) ([]piece.Kind, []wichess.CollectionSlot) {
	indices := make([]int, len(of))
	for i := 0; i < len(indices); i++ {
		indices[i] = i
	}
	rand.Shuffle(len(indices), func(i, j int) {
		indices[i], indices[j] = indices[j], indices[i]
	})

	outKinds := make([]piece.Kind, len(of))
	outSlots := make([]wichess.CollectionSlot, len(of))
	for i, s := range indices {
		outKinds[i] = of[s].Kind
		outSlots[i] = wichess.CollectionSlot(s)
	}

	return outKinds, outSlots
}
