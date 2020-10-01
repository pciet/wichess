package client

import (
	"encoding/json"

	"github.com/pciet/wichess"
	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/rules"
)

func (an Instance) ActivePlayer(id memory.GameIdentifier) (bool, error) {
	respBody, err := an.JSONResponseGet(wichess.PlayersPath + id.String())
	if err != nil {
		return false, err
	}
	var j wichess.PlayersJSON
	err = json.Unmarshal([]byte(respBody), &j)
	if err != nil {
		return false, err
	}

	if j.Active == rules.White {
		if memory.PlayerName(an.Name) == j.White {
			return true, nil
		}
		return false, nil
	}

	if memory.PlayerName(an.Name) == j.Black {
		return true, nil
	}
	return false, nil
}
