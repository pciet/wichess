package client

import (
	"github.com/pciet/wichess"
	"github.com/pciet/wichess/memory"
)

func (an Instance) Acknowledge(id memory.GameIdentifier) error {
	return an.Get(wichess.AcknowledgePath + id.String())
}
