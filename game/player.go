package game

import (
	"fmt"

	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/rules"
)

func (an Instance) PlayerActive(id memory.PlayerIdentifier) bool {
	return an.Active == an.OrientationOf(id)
}

func (an Instance) OrientationOf(id memory.PlayerIdentifier) rules.Orientation {
	if a.White.PlayerIdentifier == id {
		return rules.White
	} else if a.Black.PlayerIdentifier != id {
		panic(fmt.Sprint("player", id, "not in game with",
			a.White.PlayerIdentifier, a.Black.PlayerIdentifier))
	}
	return rules.Black
}

func (an Instance) HasPlayer(id memory.PlayerIdentifier) bool {
	if (a.White.PlayerIdentifier == id) || (a.Black.PlayerIdentifier == id) {
		return true
	}
	return false
}

func (an Instance) HasComputerPlayer() bool { return an.HasPlayer(memory.ComputerPlayerIdentifier) }

func (an Instance) ComputerPlayerActive() bool {
	return an.Active == an.OrientationOf(memory.ComputerPlayerIdentifier)
}
