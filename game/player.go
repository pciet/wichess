package game

import (
	"log"

	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/rules"
)

func (an Instance) PlayerActive(id memory.PlayerIdentifier) bool {
	return an.Active == an.OrientationOf(id)
}

func (an Instance) OrientationOf(id memory.PlayerIdentifier) rules.Orientation {
	if an.White.PlayerIdentifier == id {
		return rules.White
	} else if an.Black.PlayerIdentifier != id {
		log.Panicln("player", id, "not in game with",
			an.White.PlayerIdentifier, an.Black.PlayerIdentifier)
	}
	return rules.Black
}

func (an Instance) OrientationAgainst(id memory.PlayerIdentifier) rules.Orientation {
	if an.OrientationOf(id) == rules.White {
		return rules.Black
	}
	return rules.White
}

func (an Instance) OpponentOf(id memory.PlayerIdentifier) memory.PlayerIdentifier {
	if id == an.White.PlayerIdentifier {
		return an.Black.PlayerIdentifier
	} else if id != an.Black.PlayerIdentifier {
		log.Panicln("player", id, "not in game")
	}
	return an.White.PlayerIdentifier
}

func (an Instance) HasPlayer(id memory.PlayerIdentifier) bool {
	if (an.White.PlayerIdentifier == id) || (an.Black.PlayerIdentifier == id) {
		return true
	}
	return false
}

func (an Instance) HasComputerPlayer() bool { return an.HasPlayer(memory.ComputerPlayerIdentifier) }

func (an Instance) ComputerPlayerActive() bool {
	return an.Active == an.OrientationOf(memory.ComputerPlayerIdentifier)
}

func (an Instance) opponentOf(player rules.Orientation) rules.Orientation {
	if player == rules.White {
		return rules.Black
	} else if player != rules.Black {
		log.Panicln("bad player orientation", player)
	}
	return rules.White
}

func (an Instance) inactiveOrientation() rules.Orientation {
	if an.Active == rules.White {
		return rules.Black
	} else if an.Active != rules.Black {
		log.Panicln("bad active orientation", an.Active)
	}
	return rules.White
}

func (an Instance) inactivePlayerIdentifier() memory.PlayerIdentifier {
	if an.Active == rules.White {
		return an.Black.PlayerIdentifier
	} else if an.Active != rules.Black {
		log.Panicln("bad active orientation", an.Active)
	}
	return an.White.PlayerIdentifier
}

func (an Instance) activePlayerIdentifier() memory.PlayerIdentifier {
	if an.Active == rules.White {
		return an.White.PlayerIdentifier
	} else if an.Active != rules.Black {
		log.Panicln("bad active orientation", an.Active)
	}
	return an.Black.PlayerIdentifier
}
