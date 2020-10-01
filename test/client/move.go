package client

import (
	"bytes"
	"encoding/json"

	"github.com/pciet/wichess"
	"github.com/pciet/wichess/game"
	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/rules"
)

// Move returns the game.Update UpdateState field string which is "" for a normal state.
func (an Instance) Move(id memory.GameIdentifier, m rules.Move) (string, error) {
	b, err := json.Marshal(wichess.MoveJSON{From: m.From.Index(), To: m.To.Index()})
	if err != nil {
		return "", err
	}

	resp, err := an.JSONResponsePost(wichess.MovePath+id.String(),
		"application/json", bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}

	var u game.Update
	err = json.Unmarshal([]byte(resp), &u)
	if err != nil {
		return "", err
	}

	return string(u.UpdateState), nil
}
