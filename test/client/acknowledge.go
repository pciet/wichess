package client

import "github.com/pciet/wichess"

func (an Instance) Acknowledge(id wichess.GameIdentifier) error {
	return an.Get(wichess.AcknowledgePath + id.String())
}
