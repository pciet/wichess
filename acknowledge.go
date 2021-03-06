package wichess

import (
	"net/http"

	"github.com/pciet/wichess/game"
	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/rules"
)

func acknowledgeGet(w http.ResponseWriter, r *http.Request, g game.Instance, p *memory.Player) {
	done, state := g.Completed()
	if done == false {
		debug(AcknowledgePath, "called by", p.PlayerName, "on incomplete game", g.GameIdentifier)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if g.HasComputerPlayer() {
		if (g.Active != g.OrientationOf(p.PlayerIdentifier)) && (state == rules.Checkmate) {
			p.IncrementComputerStreak()
		} else {
			p.ResetComputerStreak()
		}
		p.ComputerGame = memory.NoGame
	} else {
		p.PeopleGame = memory.NoGame
	}

	g.Acknowledge(p.PlayerIdentifier)
}
