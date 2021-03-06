package game

import (
	"log"

	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/rules"
)

func PlayerNameWithOrientation(white, black memory.PlayerName,
	o rules.Orientation) memory.PlayerName {

	if o == rules.White {
		return white
	} else if o != rules.Black {
		log.Panicln("bad orientation", o)
	}
	return black
}

func OrientationOfPlayerName(white, black, player memory.PlayerName) rules.Orientation {
	if player == white {
		return rules.White
	} else if player != black {
		log.Panicln(player, "not", white, black)
	}
	return rules.Black
}

func OpponentNameOf(white, black, player memory.PlayerName) memory.PlayerName {
	if white == player {
		return black
	} else if black != player {
		log.Panicln(player, "not white", white, "or black", black)
	}
	return white
}
