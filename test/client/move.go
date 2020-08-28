package client

import (
	"bytes"
	"encoding/json"

	"github.com/pciet/wichess"
	"github.com/pciet/wichess/rules"
)

// Move returns the wichess.Update State field string which is "" for a normal state.
func (an Instance) Move(id wichess.GameIdentifier, m rules.Move) (string, error) {
	b, err := json.Marshal(wichess.MoveJSON{From: m.From.Index().Int(), To: m.To.Index().Int()})
	if err != nil {
		return "", err
	}

	resp, err := an.JSONResponsePost(wichess.MovePath+id.String(),
		"application/json", bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}

	var u wichess.Update
	err = json.Unmarshal([]byte(resp), &u)
	if err != nil {
		return "", err
	}

	return u.State, nil
}
