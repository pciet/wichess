package wichess

import (
	"log"
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
		log.Panicln(ComputerPath, "no game", p.ComputerGame, "for", p.Name)
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
		White: GamePlayerHTMLTemplateData{
			PlayerName: whiteName,
			Captures:   g.White.Captures,
		},
		Black: GamePlayerHTMLTemplateData{
			PlayerName: blackName,
			Captures:   g.Black.Captures,
		},
		Active:       g.Active,
		PreviousMove: g.PreviousMove,
		Player:       g.OrientationAgainst(memory.ComputerPlayerIdentifier),
	}

	g.Unlock()

	writeHTMLTemplate(w, GameHTMLTemplate, t)
}

func computerPost(w http.ResponseWriter, r *http.Request, pid memory.PlayerIdentifier) {
	p := memory.RLockPlayer(pid)
	if p == nil {
		debug(ComputerPath, "POST got nil player for", pid)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if p.ComputerGame != memory.NoGame {
		debug(ComputerPath, "POST but computer game already exists for", p.Name)
		p.RUnlock()
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	army, err := piece.DecodeArmyRequest(r.Body)
	if err != nil {
		debug(ComputerPath, "POST failed to decode army request of", p.Name, ":", err)
		p.RUnlock()
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
		blackArmy = army
		black = p.PlayerIdentifier
	}
	p.RUnlock()

	// game.New takes a lock on both players so this player's can't be held here
	gameID := game.New(whiteArmy, blackArmy, white, black)
	if gameID == memory.NoGame {
		debug(ComputerPath, "NewGame failed for", p.Name)
		// likely caused by a bad army request
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p = memory.LockPlayer(pid)
	if p == nil {
		debug(ComputerPath, "POST got nil player for", pid)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p.ComputerGame = gameID
	p.Unlock()
}
