package client

import (
	"bytes"
	"encoding/json"

	"github.com/pciet/wichess"
	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/piece"
)

func Match(a, b Instance, aArmy, bArmy piece.ArmyRequest) (memory.GameIdentifier, error) {
	done := make(chan error)
	var gameID memory.GameIdentifier

	match := func(an Instance, army piece.ArmyRequest, opponent string, writeID bool) {
		b, err := json.Marshal(army)
		if err != nil {
			done <- err
			return
		}

		respBody, err := an.JSONResponsePost(
			wichess.PeoplePath+"?"+wichess.RequestedOpponentQuery+"="+opponent,
			"application/json", bytes.NewBuffer(b))

		if err != nil {
			done <- err
			return
		}

		var j wichess.PeoplePostJSON
		err = json.Unmarshal([]byte(respBody), &j)
		if err != nil {
			done <- err
			return
		}

		if writeID {
			gameID = j.GameIdentifier
		}

		done <- nil
	}

	go match(a, aArmy, b.Name, false)
	go match(b, bArmy, a.Name, true)

	err := <-done
	if err != nil {
		return 0, err
	}
	err = <-done
	return gameID, err
}
