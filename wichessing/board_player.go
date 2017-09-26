// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import (
	"fmt"
)

type PlayerMove struct {
	From AbsPoint
	To   AbsPoint
}

type MoveRating int

const MaxMoveRating MoveRating = 100

// Returns nil if the indicated player is in checkmate.
func (b Board) ComputerMove(player Orientation) *PlayerMove {
	moves, _, checkmate := b.Moves(player)
	if checkmate {
		return nil
	}
	ratings := make(map[*PlayerMove]MoveRating)
	for point, set := range moves {
		if b[point.Index()].Orientation != player {
			continue
		}
		for to, _ := range set {
			ratings[&PlayerMove{
				From: point,
				To:   *to,
			}] = b.ComputerRating(point, *to, player)
		}
	}
	var best *PlayerMove
	var bestRating MoveRating
	for move, rating := range ratings {
		if best == nil {
			best = move
			bestRating = rating
			continue
		}
		if bestRating < rating {
			best = move
			bestRating = rating
		}
	}
	return best
}

func (b Board) ComputerRating(from AbsPoint, to AbsPoint, player Orientation) MoveRating {
	fromPoint := b[from.Index()]
	if fromPoint.Piece == nil {
		panic("wichessing: no piece in specified from point")
	}
	if fromPoint.Orientation != player {
		panic(fmt.Sprintf("wichessing: from point (%v) orientation (%v) does not match player orientation (%v)", from, fromPoint.Orientation, player))
	}
	var opponent Orientation
	if player == White {
		opponent = Black
	} else {
		opponent = White
	}
	rating := MoveRating(0)
	state := b.AfterMove(from, to, player)
	if state.Checkmate(opponent) {
		rating = MaxMoveRating
		return rating
	}
	if state.PieceCount(player) < b.PieceCount(player) {
		rating--
	}
	if state.PieceCount(opponent) < b.PieceCount(opponent) {
		rating++
	}
	if state.Check(opponent) {
		rating++
	}
	if state.PiecesInDanger(player) > b.PiecesInDanger(player) {
		rating--
	}
	if state.PiecesInDanger(opponent) > b.PiecesInDanger(opponent) {
		rating++
	}
	npScore := state.TotalPieceScore(player)
	opScore := b.TotalPieceScore(player)
	if npScore < opScore {
		rating = rating - MoveRating(opScore) - MoveRating(npScore)
	}
	npScore = state.TotalPieceScore(opponent)
	opScore = b.TotalPieceScore(opponent)
	if npScore < opScore {
		rating = rating + MoveRating(opScore) - MoveRating(npScore)
	}
	return rating
}
