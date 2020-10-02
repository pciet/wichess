package client

import (
	"errors"

	"github.com/pciet/wichess"
	"github.com/pciet/wichess/memory"
)

func (an Instance) AcknowledgePeopleGame() error {
	id, err := an.PeopleGame()
	if err != nil {
		return err
	}
	if id == memory.NoGame {
		return errors.New("no people game")
	}
	return an.Acknowledge(id)
}

func (an Instance) Acknowledge(id memory.GameIdentifier) error {
	return an.Get(wichess.AcknowledgePath + id.String())
}
