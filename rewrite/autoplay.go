package main

import (
	"math/rand"

	"github.com/pciet/wichess/rules"
)

// Autoplay picks and does a move for a player then alerts the opponent.
func Autoplay(id GameIdentifier, player string) {
	tx := DatabaseTransaction()

	g := LoadGame(tx, id)
	if g.Header.ID == 0 {
		Panic("game", id, "not found")
	}
	if g.Header.Active != player {
		Panic("tried to autoplay for inactive player", player)
	}

	move, promotion := AutoplayMove(rules.MakeGame(
		g.Board.Board,
		rules.AddressIndex(g.Header.From),
		rules.AddressIndex(g.Header.To)),
		ActiveOrientation(g.Header.Active, g.Header.White.Name, g.Header.Black.Name),
	)

	changes := g.DoMove(tx, move, promotion)

	tx.Commit()

	if changes == nil {
		Panic("autoplay failed for", player)
	}

	go Alert(id, Opponent(player, g.Header.White.Name, g.Header.Black.Name), changes)
}

// Looking forward more than one move takes too much time.

// The autoplay algorithm in AutoplayMove inspects all moves this turn and picks the best.
// A random move is picked amongst ties. The returned PieceKind is the promotion pick if needed.
func AutoplayMove(g rules.Game, o rules.Orientation) (rules.Move, rules.PieceKind) {
	moves, state := g.Moves(o)
	if (state != rules.Normal) && (state != rules.Check) {
		Panic("tried to calculate autoplay move when game already complete")
	}

	var best rules.Move
	bestRating := -100
	for _, moveset := range moves {
		for _, moveTo := range moveset.Moves {
			move := rules.Move{moveset.From, moveTo}
			rating := AutoplayRating(g, move, o)
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
	}

	return best, rules.Queen
}

func AutoplayRating(g rules.Game, of rules.Move, by rules.Orientation) int {
	opponent := by.Opponent()
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

	if of.Forward(by) {
		rating++
	}

	rating += future.Board.PieceCount(by) - g.Board.PieceCount(by)
	rating += g.Board.PieceCount(opponent) - future.Board.PieceCount(opponent)

	ts := g.Board[of.To.Index()]
	if (ts.Kind != rules.NoKind) && (ts.Orientation != by) {
		if rules.IsBasicKind(ts.Kind) == false {
			rating++
		}
		switch rules.BasicKind(ts.Kind) {
		case rules.Queen:
			rating += 4
		case rules.Rook:
			rating += 3
		case rules.Bishop:
			rating += 2
		case rules.Knight:
			rating += 1
		}
	}

	return rating
}
