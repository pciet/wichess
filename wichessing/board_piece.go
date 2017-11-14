// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import (
	"fmt"
)

func (b Board) KingLocationFor(turn Orientation) (AbsPoint, bool) {
	for _, point := range b {
		if point.Piece == nil {
			continue
		}
		if point.Orientation != turn {
			continue
		}
		if point.Base == King {
			return point.AbsPoint, true
		}
	}
	return AbsPoint{}, false
}

func (b Board) IndexPositionOfPiece(the *Piece) uint8 {
	for index, point := range b {
		if point.Piece == the {
			return uint8(index)
		}
	}
	panic(fmt.Sprintf("wichessing: piece %v not on board", *the))
	return 0
}

func (b Board) PiecesFor(player Orientation) PieceSet {
	set := make(PieceSet)
	for _, point := range b {
		if point.Piece == nil {
			continue
		}
		if point.Orientation != player {
			continue
		}
		set = set.Add(point.Piece)
	}
	return set
}

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

func (b Board) PiecesInDanger(player Orientation, previousFrom AbsPoint, previousTo AbsPoint) uint8 {
	var opponent Orientation
	if player == White {
		opponent = Black
	} else {
		opponent = White
	}
	moves := b.AllNaiveMovesFor(opponent, previousFrom, previousTo)
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
		switch point.Base {
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
		}
	}
	return total
}
