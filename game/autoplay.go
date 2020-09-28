package game

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

// Autoplay picks and does a move for the active player. If the move causes a promotion to be
// needed then it's done with a queen.
func (an Instance) Autoplay() {
	move, promotion := an.autoplayMove()

	var u Update
	if (move == rules.NoMove) && (promotion == piece.NoKind) {
		// game is completed
		u.Squares = []Square{}
		u.FromMove = rules.NoMove
		go Alert(an.GameIdentifier, an.inactiveOrientation(), an.inactivePlayerIdentifier(), u)
		return
	}

	var promotionNeeded bool
	if promotion != piece.NoKind {
		u.Squares = SquaresFromRules(an.Promote(promotion))
	} else {
		u.FromMove = move
		var squares []rules.Square
		var captures []rules.Piece
		squares, captures, promotionNeeded = an.Move(move)
		u.Squares = SquaresFromRules(squares)
		u.Captures = PiecesFromRules(captures)
	}

	if promotionNeeded {
		// TODO: is reverse promotion properly handled?
		u.Squares = mergeReplaceSquares(u.Squares, SquaresFromRules(an.Promote(piece.Queen)))
	}

	go Alert(an.GameIdentifier, an.Active, an.activePlayerIdentifier(), u)
}

// This autoplay algorithm inspects the results of all moves. This already takes significant time,
// so this way looking forward more than one move isn't possible.

// autoplayMove inspects all of this turn's moves and picks the best based on a rating system.
// Ties are broken randomly. If a piece.Kind other than NoKind is returned then the move is a
// promotion. If the game is complete then rules.NoMove and piece.NoKind are returned.
func (an Instance) autoplayMove() (rules.Move, piece.Kind) {
	moves, state := an.Moves()
	if (state != rules.Normal) && (state != rules.Check) {
		return rules.NoMove, piece.NoKind
	}
	if state == rules.Promotion {
		return rules.NoMove, piece.Queen
	}

	count := 0
	for _, set := range moves {
		count += len(set.Moves)
	}

	var best rules.Move
	bestRating := -100
	var i int
	for _, moveset := range moves {
		for _, to := range moveset.Moves {
			move := rules.Move{moveset.From, to}
			rating := an.autoplayRating(move)
			i++
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

func (an Instance) autoplayRating(m rules.Move) int {
	icopy := an.Copy()

	active := icopy.Active
	changes, captures, _ := icopy.Move(m)
	if changes == nil {
		log.Panicln("move", m, "failed on board\n", icopy.Game.Board.String())
	}

	opponentMoves, state := icopy.Moves()

	rating := len(captures) - len(opponentMoves)

	switch state {
	case rules.Checkmate:
		return 100
	case rules.Draw:
		return -99
	case rules.Check:
		rating++
	}

	if m.Forward(active) {
		rating++
	}

	for _, p := range captures {
		if p.Orientation == active {
			switch p.Kind.Basic() {
			case piece.Queen:
				rating -= 8
			case piece.Rook:
				rating -= 6
			case piece.Bishop:
				rating -= 4
			case piece.Knight:
				rating -= 3
			case piece.Pawn:
				rating -= 2
			}
			continue
		}
		switch p.Kind.Basic() {
		case piece.Queen:
			rating += 8
		case piece.Rook:
			rating += 6
		case piece.Bishop:
			rating += 4
		case piece.Knight:
			rating += 3
		case piece.Pawn:
			rating += 2
		}
	}

	return rating
}

// autoplay gets the Instance and calls the Autoplay method on it.
func autoplay(in memory.GameIdentifier, by memory.PlayerIdentifier) error {
	g := Lock(in)
	if g.Nil() {
		return fmt.Errorf("bad game id %v", in)
	}
	g.Autoplay()
	g.Unlock()
	return nil
}
