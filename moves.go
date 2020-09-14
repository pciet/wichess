package wichess

import (
	"net/http"

	"github.com/pciet/wichess/game"
	"github.com/pciet/wichess/rules"
)

type MovesJSON struct {
	Moves       []rules.MoveSet `json:"m,omitempty"`
	rules.State `json:"s"`
}

func movesGet(w http.ResponseWriter, r *http.Request, g game.Instance, pid PlayerIdentifier) {
	// TODO: does the rewrite still have the race that made counting turns necessary?
	moves, state := g.Moves()

	if (g.ComputerGame() == false) &&
		((state == rules.Check) || (state == rules.Draw) || (state == rules.Checkmate)) {

		var alertState game.UpdateState
		switch state {
		case rules.Check:
			alertState = game.CheckCalculatedUpdate
		case rules.Draw:
			alertState = game.DrawCalculatedUpdate
		case rules.Checkmate:
			alertState = game.CheckmateCalculatedUpdate
		}

		oppID := g.OpponentOf(pid)

		go game.Alert(gm.GameIdentifier, g.OrientationOf(oppID), oppID, game.Update{
			State:    alertState,
			FromMove: rules.NoMove,
		})
	}

	jsonResponse(w, MovesJSON{moves, state})
}
