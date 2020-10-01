package client

import (
	"encoding/json"

	"github.com/pciet/wichess"
	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/rules"
)

func (an Instance) Moves(id memory.GameIdentifier) ([]rules.MoveSet, rules.State, error) {
	resp, err := an.JSONResponseGet(wichess.MovesPath + id.String())
	if err != nil {
		return nil, rules.Normal, err
	}

	var mj wichess.MovesJSON
	err = json.Unmarshal([]byte(resp), &mj)
	if err != nil {
		return nil, rules.Normal, err
	}

	return mj.Moves, mj.State, nil
}
