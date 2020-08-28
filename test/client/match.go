package client

import (
	"bytes"
	"encoding/json"

	"github.com/pciet/wichess"
)

func Match(a, b Instance) (wichess.GameIdentifier, error) {
	done := make(chan error)
	var gameID wichess.GameIdentifier

	match := func(an Instance, opponent string, writeID bool) {
		b, err := json.Marshal(wichess.RegularArmyRequest)
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

	go match(a, b.Name, false)
	go match(b, a.Name, true)

	err := <-done
	if err != nil {
		return 0, err
	}
	err = <-done
	return gameID, err
}
