package main

import (
	"math/rand"

	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

// Autoplay picks and does a move for a player then alerts the opponent.
// If a promotion needs to be done then a queen is picked.
func Autoplay(id GameIdentifier, player string) {
	tx := DatabaseTransaction()

	g := LoadGame(tx, id, true)
	if g.Header.ID == 0 {
		Panic("game", id, "not found")
	}
	if g.Header.Active != OrientationOf(player, g.Header.White.Name, g.Header.Black.Name) {
		Panic("tried to autoplay for inactive player", player)
	}

	move, promotion := AutoplayMove(rules.MakeGame(g.Board.Board,
		rules.AddressIndex(g.Header.From),
		rules.AddressIndex(g.Header.To)),
		g.Header.Active,
	)

	var u Update
	promotionNeeded := false
	if (move == rules.NoMove) && (promotion == piece.NoKind) {
		// alert player to completed game with empty diff
		u.Squares = []rules.AddressedSquare{}
		u.FromMove = rules.NoMove
	} else {
		u.FromMove = move
		u.Squares, u.Captures, promotionNeeded = g.DoMove(tx, move, promotion)
		if promotionNeeded {
			// if this is the promoting player then do it now
			(&g.Board).ApplyChanges(u.Squares)
			promoter, _ := g.Board.PromotionNeeded()
			if PlayerWithOrientation(promoter, g.Header.White.Name,
				g.Header.Black.Name) == player {
				promUpdates, _, _ := g.DoMove(tx, rules.NoMove, piece.Queen)
				u.Squares = rules.MergeReplaceAddressedSquares(u.Squares, promUpdates)
			}
		}
	}

	tx.Commit()

	go Alert(id, Opponent(g.Header.Active, g.Header.White.Name, g.Header.Black.Name), u)
}

// Looking forward more than one move takes too much time.

// The autoplay algorithm in AutoplayMove inspects all moves this turn and picks the best.
// A random move is picked amongst ties. The returned PieceKind is the promotion pick if needed.
// If the game is determined to be complete then rules.NoMove is returned.
func AutoplayMove(g rules.Game, o rules.Orientation) (rules.Move, piece.Kind) {
	moves, state := g.Moves(o)
	if (state != rules.Normal) && (state != rules.Check) {
		return rules.NoMove, piece.NoKind
	}
	if state == rules.Promotion {
		return rules.NoMove, piece.Queen
	}

	takes := rules.MovesAddressSlice(g.NaiveTakeMoves(o.Opponent()))

	var best rules.Move
	bestRating := -100
	for _, moveset := range moves {
		for _, moveTo := range moveset.Moves {
			move := rules.Move{moveset.From, moveTo}
			rating := AutoplayRating(g, move, o, takes)
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

	return best, piece.NoKind
}

func AutoplayRating(g rules.Game, of rules.Move,
	by rules.Orientation, threats []rules.Address) int {
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
	if (ts.Kind != piece.NoKind) && (ts.Orientation != by) {
		if ts.Kind.IsBasic() == false {
			rating++
		}
		switch ts.Kind.Basic() {
		case piece.Queen:
			rating += 4
		case piece.Rook:
			rating += 3
		case piece.Bishop:
			rating += 2
		case piece.Knight:
			rating += 1
		}
	}

	for _, t := range threats {
		if t == of.To {
			rating -= 3
			break
		}
	}

	return rating
}
