package wichess

import (
	"net/http"

	"github.com/pciet/wichess/memory"
)

type PeopleIDJSON struct {
	memory.GameIdentifier `json:"id"`
}

func peopleIDGet(w http.ResponseWriter, r *http.Request, p *memory.Player) {
	jsonResponse(w, PeopleIDJSON{GameIdentifier: p.PeopleGame})
}
