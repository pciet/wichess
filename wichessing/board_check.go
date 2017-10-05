// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

func (b Board) Checkmate(turn Orientation) bool {
	for from, set := range b.AllMovesFor(turn) {
		for move, _ := range set {
			if b.AfterMove(from, *move, turn).Check(turn) == false {
				return false
			}
		}
	}
	return true
}

// TODO: if detonate bishop is in danger next to king, then king is in check

func (b Board) Check(turn Orientation) bool {
	var kingLocation AbsPoint
	moves := make(AbsPointSet)
	for _, pt := range b {
		if pt.Piece == nil {
			continue
		}
		if pt.Piece.Orientation == turn {
			if pt.Piece.Kind == King {
				kingLocation = pt.AbsPoint
			}
			continue
		}
		moves = moves.Add(b.MovesFromPoint(pt)).Reduce()
	}
	for pt, _ := range moves {
		if (pt.File == kingLocation.File) && (pt.Rank == kingLocation.Rank) {
			return true
		}
	}
	return false
}
