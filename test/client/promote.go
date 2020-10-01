package client

import (
	"bytes"
	"encoding/json"

	"github.com/pciet/wichess"
	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/piece"
)

func (an Instance) Promote(id memory.GameIdentifier, k piece.Kind) error {
	b, err := json.Marshal(wichess.MoveJSON{Promotion: k})
	if err != nil {
		return err
	}
	_, err = an.JSONResponsePost(wichess.MovePath+id.String(),
		"application/json", bytes.NewBuffer(b))
	// TODO: does reverse promotion need to read the JSON response?
	return err
}
