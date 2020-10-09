package client

import (
	"bytes"
	"encoding/json"
	"math/rand"

	"github.com/pciet/wichess"
	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/piece"
)

func (an Instance) AcceptAllRewardsRandomly(id memory.GameIdentifier) error {
	randSlot := func(a, b piece.CollectionSlot) piece.CollectionSlot {
		n := piece.CollectionSlot(rand.Intn(piece.CollectionSize) + 1)
		for (n == a) || (n == b) {
			n++
			if n > piece.CollectionSize {
				n = 1
			}
		}
		return n
	}
	var j wichess.RewardJSON
	j.Left = randSlot(0, 0)
	j.Right = randSlot(j.Left, 0)
	j.Reward = randSlot(j.Left, j.Right)

	b, err := json.Marshal(j)
	if err != nil {
		return err
	}

	_, err = an.JSONResponsePost(wichess.RewardPath+id.String(),
		"application/json", bytes.NewBuffer(b))

	return err
}
