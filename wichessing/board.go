// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

type Board [64]Point

func (b Board) Copy() Board {
	var board Board
	for i, pt := range b {
		board[i] = Point{
			AbsPoint: AbsPoint{
				File: pt.File,
				Rank: pt.Rank,
			},
			Piece: pt.Piece.Copy(),
		}
	}
	return board
}
