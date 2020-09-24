package wichess

import (
	"net/http"

	"github.com/pciet/wichess/memory"
)

type MatchHTMLTemplateData struct {
	Name               string
	RecentOpponents    [memory.RecentOpponentCount]string
	ComputerStreak     int
	BestComputerStreak int
}

func matchGet(w http.ResponseWriter, r *http.Request, p *memory.Player) {
	if handleInPeopleGame(w, r, p) {
		return
	}
	d := MatchHTMLTemplateData{
		Name:               p.PlayerName.String(),
		ComputerStreak:     p.ComputerStreak,
		BestComputerStreak: p.BestComputerStreak,
	}
	for i, o := range p.RecentOpponents {
		if o == memory.NoPlayer {
			break
		}
		d.RecentOpponents[i] = o.Name().String()
	}
	writeHTMLTemplate(w, MatchHTMLTemplate, d)
}
