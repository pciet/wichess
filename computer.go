package wichess

import (
	"fmt"
	"net/http"

	"github.com/pciet/wichess/game"
	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/piece"
)

func computerGet(w http.ResponseWriter, r *http.Request, p *memory.Player) {
	if handleInPeopleGame(w, r, p) {
		return
	}

	if p.ComputerGame == memory.NoGame {
		debug(ComputerPath, "no game for", p.Name)
		http.NotFound(w, r)
		return
	}

	g := game.Lock(p.ComputerGame)
	if g.Nil() {
		panic(fmt.Sprint(ComputerPath, "no game", p.ComputerGame, "for", p.Name))
	}

	// TODO: is it even possible that the computer player's move hasn't been done?
	if g.ComputerPlayerActive() {
		// TODO: be sure this sets the player as active
		g.Autoplay()
	}

	whiteName, blackName :=
		memory.TwoPlayerNames(g.White.PlayerIdentifier, g.Black.PlayerIdentifier)

	t := GameHTMLTemplateData{
		GameIdentifier: p.ComputerGame,
		Conceded:       g.Conceded,
		White: {
			Name:     whiteName,
			Captures: g.White.Captures,
		},
		Black: {
			Name:     blackName,
			Captures: g.Black.Captures,
		},
		Active:   g.Active,
		Previous: g.PreviousMove,
		Player:   g.OrientationAgainst(ComputerPlayerIdentifier),
	}

	g.Unlock()

	writeHTMLTemplate(w, GameHTMLTemplate, t)
}

func computerPost(w http.ResponseWriter, r *http.Request, p *memory.Player) {
	if p.ComputerGame != memory.NoGame {
		debug(ComputerPath, "POST but computer game already exists for", p.Name)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	army, err := piece.DecodeArmyRequest(r.Body)
	if err != nil {
		debug(ComputerPath, "POST failed to decode army request of", p.Name, ":", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var whiteArmy, blackArmy piece.ArmyRequest
	var white, black memory.PlayerIdentifier
	if randomBool() {
		whiteArmy = army
		white = p.PlayerIdentifier
		blackArmy = piece.RegularArmyRequest
		black = memory.ComputerPlayerIdentifier
	} else {
		whiteArmy = piece.RegularArmyRequest
		white = memory.ComputerPlayerIdentifier
		blackArmy = a
		black = p.PlayerIdentifier
	}

	p.ComputerGame = game.New(whiteArmy, blackArmy, white, black)
	if p.ComputerGame == memory.NoGame {
		debug(ComputerPath, "NewGame failed for", p.Name)
		// likely caused by a bad army request
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p.Changed()
}
