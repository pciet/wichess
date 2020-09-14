package wichess

import (
	"net/http"

	"github.com/pciet/wichess/game"
	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

func boardsGet(w http.ResponseWriter, r *http.Request, g game.Instance) {
	jr := make([]rules.AddressedSquare, 0, 32)
	for i, s := range g.Board {
		if s.Kind == piece.NoKind {
			continue
		}
		jr = append(jr, rules.AddressedSquare{rules.AddressIndex(i).Address(), s})
	}
	jsonResponse(w, jr)
}
