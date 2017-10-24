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
	var orientation Orientation
	if turn == White {
		orientation = Black
	} else {
		orientation = White
	}
	allMoves := b.AllMovesFor(orientation)
	king, _ := b.KingLocationFor(turn)
	for orig, moves := range allMoves {
		for pt, _ := range moves {
			if (pt.File == king.File) && (pt.Rank == king.Rank) {
				return true
			}
			// some cases cause reactions that remove the King, such as an enemy detonator move and friendly guard adjacent
			_, has := b.AfterMove(orig, *pt, orientation).KingLocationFor(turn)
			if has == false {
				return true
			}
		}
	}
	return false
}
