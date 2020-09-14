package wichess

import (
	"net/http"

	"github.com/pciet/wichess/memory"
)

type MatchHTMLTemplateData struct {
	Name               string
	RecentOpponents    [RecentOpponentCount]string
	ComputerStreak     int
	BestComputerStreak int
}

func matchGet(w http.ResponseWriter, r *http.Request, p *memory.Player) {
	if handleInPeopleGame(w, r, p) {
		return
	}
	WriteHTMLTemplate(w, MatchHTMLTemplate, MatchHTMLTemplateData{
		Name:               p.Name,
		RecentOpponents:    p.RecentOpponents,
		ComputerStreak:     p.ComputerStreak,
		BestComputerStreak: p.BestComputerStreak,
	})
}
