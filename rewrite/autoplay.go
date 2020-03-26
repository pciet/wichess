package main

import (
	"database/sql"
	"math/rand"

	"github.com/pciet/wichess/rules"
)

// Autoplay picks and does a move for a player then alerts the opponent.
func Autoplay(id GameIdentifier, player string) {
	tx := DatabaseTransaction()

	g := LoadGame(tx, id)
	if g.ID == 0 {
		Panic("game", id, "not found")
	}
	if g.Active != player {
		Panic("tried to autoplay for inactive player", player)
	}

	move, promotion := AutoplayMove(g.RulesGame(), g.OrientationOf(g.Active))
	changes := g.DoMove(tx, move, promotion)
	opponent := GameOpponent(tx, id, player)

	tx.Commit()

	if changes == nil {
		Panic("autoplay failed for", player)
	}

	go Alert(id, opponent, changes)
}

// Looking forward more than one move takes too much time.

// The autoplay algorithm in AutoplayMove inspects all moves this turn and picks the best.
// A random move is picked amongst ties.
func AutoplayMove(g rules.Game, o rules.Orientation) (rules.Move, rules.PieceKind) {
	moves, state := g.Moves(o)
	if (state != rules.Normal) && (state != rules.Check) {
		Panic("tried to calculate autoplay move when game already complete")
	}

	var best rules.Move
	bestRating := -100
	for _, move := range moves {
		rating := AutoplayRating(g, move)
		if rating > bestRating {
			bestRating = rating
			best = move
			continue
		}
		if rating == bestRating {
			if rand.Intn(2) == 0 {
				best = move
			}
		}
	}

	return move, rules.Queen
}

func AutoplayRating(g rules.Game, of rules.Move) int {
	opponent := g.InactivePlayer()
	future := g.AfterMove(of)
	_, state := future.Moves(opponent)

	rating := 0

	switch state {
	case rules.Checkmate:
		return 100
	case rules.Draw:
		return -99
	case rules.Check:
		rating++
	}

	rating += future.PlayerPieceCount(g.Active) - g.PlayerPieceCount(g.Active)
	rating += g.PlayerPieceCount(opponent) - future.PlayerPieceCount(opponent)

	ts := g.Board[of.To.Index()]
	if (ts.Kind != NoKind) && (ts.Orientation != g.ActiveOrientation()) {
		if rules.IsBasicKind(ts.Kind) == false {
			rating++
		}
		switch rules.BasicKind(ts.Kind) {
		case Queen:
			rating += 4
		case Rook:
			rating += 3
		case Bishop:
			rating += 2
		case Knight:
			rating += 1
		}
	}

	return rating
}
