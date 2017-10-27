// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

var GhostAfterMoveCases = []PositionAfterMoveCase{
	{
		Name: "White Rook Moves Through White King",
		Initial: PointSet{
			&WhiteKingStart: {},
			&BlackKingStart: {},
			{
				AbsPoint: AbsPoint{0, 0},
				Piece: &Piece{
					Kind:        GhostRook,
					Orientation: White,
				},
			}: {},
		},
		From: AbsPoint{0, 0},
		To:   AbsPoint{7, 0},
		Diff: PointSet{
			{
				AbsPoint: AbsPoint{0, 0},
			}: {},
			{
				AbsPoint: AbsPoint{7, 0},
				Piece: &Piece{
					Kind:        GhostRook,
					Orientation: White,
				},
			}: {},
		},
	},
}
