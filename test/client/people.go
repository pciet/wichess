package client

import (
	"encoding/json"

	"github.com/pciet/wichess"
)

func (an Instance) PeopleGame() (wichess.GameIdentifier, error) {
	respBody, err := an.JSONResponseGet(wichess.PeopleIDPath)
	if err != nil {
		return 0, err
	}
	var j wichess.PeopleIDJSON
	err = json.Unmarshal([]byte(respBody), &j)
	if err != nil {
		return 0, err
	}
	return j.GameIdentifier, nil
}
