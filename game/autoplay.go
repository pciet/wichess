package game

import (
	"fmt"
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
		u.Squares = []rules.AddressedSquare{}
		u.FromMove = rules.NoMove
		go Alert(an.GameIdentifier, an.InactivePlayerIdentifier(), u)
		return
	}

	var promotionNeeded bool
	if promotion != piece.NoKind {
		u.Squares = a.Promote(promotion)
	} else {
		u.FromMove = move
		u.Squares, u.Captures, promotionNeeded = a.Move(move)
	}

	if promotionNeeded {
		// TODO: is reverse promotion properly handled?
		u.Squares = rules.MergeReplaceAddressedSquares(u.Squares, an.Promote(piece.Queen))
	}

	go Alert(an.GameIdentifier, an.ActivePlayerIdentifier(), u)
}

// This autoplay algorithm inspects the results of all moves. This already takes significant time,
// so this way looking forward more than one move isn't possible.

// autoplayMove inspects all of this turn's moves and picks the best based on a rating system.
// Ties are broken randomly. If a piece.Kind other than NoKind is returned then the move is a
// promotion. If the game is complete then rules.NoMove and piece.NoKind are returned.
func (an Instance) autoplayMove() (rules.Move, piece.Kind) {
	moves, threats, state := an.MovesWithThreats()
	if (state != rules.Normal) && (state != rules.Check) {
		return rules.NoMove, piece.NoKind
	}
	if state == rules.Promotion {
		return rules.NoMove, piece.Queen
	}

	count := rules.MoveSetMoveCount(moves)
	trials := make([]Instance, count)
	for i := 0; i < count; i++ {
		trials[i] = an.Copy()
	}

	var best rules.Move
	bestRating := -100
	var i int
	for _, moveset := range moves {
		// since this calculation takes 10-100x the time of a normal move it yields processor time
		// to people game moves where a quick response is more critical for a good player experience
		memory.QuickWaitForIdle()

		for _, to := range moveset.Moves {
			move := rules.Move{moveset.From, to}
			rating := trials[i].autoplayRating(move, threats)
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

// autoplayRating rates the move for comparison to other moves. The Game is changed as if the move
// was done, so the caller likely uses a copy of the Game for this calculation.
func (an Instance) autoplayRating(m rules.Move, threats []rules.Address) int {
	active := an.Active
	changes, captures, promotion := an.Move(m)
	opponentMoves, state := a.Moves()

	rating := 0

	switch state {
	case rules.Checkmate:
		return 100
	case rules.Draw:
		return -99
	case rules.Check:
		rating++
	}

	rating += len(captures)
	rating -= len(opponentMoves)

	if m.Forward() {
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

	for _, t := range threats {
		if t == m.To {
			rating -= 5
			break
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
