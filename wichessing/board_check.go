// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

func (b Board) Checkmate(turn Orientation) bool {
	// if the King is not in check then its not a checkmate - the no moves case is a stalemate
	if b.Check(turn) == false {
		return false
	}
	for from, set := range b.AllNaiveMovesFor(turn) {
		for _, move := range set {
			if (b.Points[from.Index()].Kind == King) && (b.Points[from.Index()].Moved == false) {
				// castle isn't allowed when king in check
				// TODO: test cases for this
				if (move.File == 2) || (move.File == 6) {
					continue
				}
			}
			if b.AfterMove(from, move, turn).Check(turn) == false {
				return false
			}
		}
	}
	return true
}

// TODO: verify fixed: if detonate bishop is in danger next to king, then king is in check
// TODO: possible optimization may involve reducing the check checks for naive moves (into check)

func (b Board) Check(turn Orientation) bool {
	var orientation Orientation
	if turn == White {
		orientation = Black
	} else {
		orientation = White
	}
	allMoves := b.AllNaiveMovesFor(orientation)
	king, _ := b.KingLocationFor(turn)
	for orig, moves := range allMoves {
		for _, pt := range moves {
			if pt == king {
				return true
			}
			// some cases cause reactions that remove the King, such as an enemy detonator move and friendly guard adjacent
			_, has := b.AfterMove(orig, pt, orientation).KingLocationFor(turn)
			if has == false {
				return true
			}
		}
	}
	return false
}
