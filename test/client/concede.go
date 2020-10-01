package client

import (
	"github.com/pciet/wichess"
	"github.com/pciet/wichess/memory"
)

// ConcedeIfPeopleGame concedes the people game if one exists.
func (an Instance) ConcedeIfPeopleGame() error {
	id, err := an.PeopleGame()
	if err != nil {
		return err
	}
	if id == 0 {
		return nil
	}
	return an.ConcedeGame(id)
}

func (an Instance) ConcedeGame(id memory.GameIdentifier) error {
	return an.Get(wichess.ConcedePath + id.String())
}
