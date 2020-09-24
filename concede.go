package wichess

import (
	"net/http"

	"github.com/pciet/wichess/game"
	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/rules"
)

func concedeGet(w http.ResponseWriter, r *http.Request, g game.Instance, p *memory.Player) {
	done, _ := g.Completed()
	if done {
		// TODO: conceded race, this error shouldn't necessarily stop the web interface
		debug(ConcedePath, "called by", p.Name, "after complete")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if g.HasComputerPlayer() {
		g.CanDelete()
		p.ComputerGame = 0
		p.BestComputerStreak = 0
	} else {
		// TODO: when opponent acknowledges is the game deleted?
		g.Conceded = true

		oppID := g.OpponentOf(p.PlayerIdentifier)
		go game.Alert(g.GameIdentifier, g.OrientationOf(oppID), oppID,
			game.Update{UpdateState: game.ConcededUpdate, FromMove: rules.NoMove})

		p.PeopleGame = 0
	}
}
