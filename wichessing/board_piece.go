// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

func (b Board) PieceCount(player Orientation) uint8 {
	count := 0
	for _, point := range b {
		if point.Piece == nil {
			continue
		}
		if point.Orientation != player {
			continue
		}
		count++
	}
	return uint8(count)
}

func (b Board) PiecesInDanger(player Orientation) uint8 {
	var opponent Orientation
	if player == White {
		opponent = Black
	} else {
		opponent = White
	}
	moves := b.AllMovesFor(opponent)
	count := 0
	for _, point := range b {
		if point.Piece == nil {
			continue
		}
		if point.Orientation != player {
			continue
		}
	OUTER:
		for _, set := range moves {
			for pt, _ := range set {
				if (pt.File == point.File) && (pt.Rank == point.Rank) {
					count++
					break OUTER
				}
			}
		}
	}
	return uint8(count)
}

func (b Board) TotalPieceScore(player Orientation) int {
	total := 0
	for _, point := range b {
		if point.Piece == nil {
			continue
		}
		if point.Orientation != player {
			continue
		}
		switch point.Kind {
		case Pawn:
			total += 1
		case Knight:
			total += 2
		case Bishop:
			total += 3
		case Rook:
			total += 4
		case Queen:
			total += 5
		case Swap, Lock, Recon:
			total += 3
		case Detonate, Ghost, Steal:
			total += 4
		case Guard, Rally, Fortify:
			total += 5
		}
	}
	return total
}
