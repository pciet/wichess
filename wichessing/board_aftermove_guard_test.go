// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

var GuardAfterMoveCases = []PositionAfterMoveCase{
	{
		Name: "Simple Guard Take",
		Initial: PointSet{
			{
				AbsPoint: AbsPoint{0, 2},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{0, 0},
				Piece: &Piece{
					Kind:        GuardRook,
					Orientation: White,
				},
			},
			WhiteKingStart,
			BlackKingStart,
		},
		PreviousFrom: AbsPoint{3, 3},
		PreviousTo:   AbsPoint{3, 4},
		From:         AbsPoint{0, 2},
		To:           AbsPoint{0, 1},
		Diff: PointSet{
			{
				AbsPoint: AbsPoint{0, 2},
			},
			{
				AbsPoint: AbsPoint{0, 0},
			},
			{
				AbsPoint: AbsPoint{0, 1},
				Piece: &Piece{
					Kind:        GuardRook,
					Orientation: White,
				},
			},
		},
	},
	{
		Name: "Guard Take Chain",
		Initial: PointSet{
			{
				AbsPoint: AbsPoint{0, 3},
				Piece: &Piece{
					Kind:        GuardPawn,
					Orientation: Black,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{1, 2},
				Piece: &Piece{
					Kind:        GuardPawn,
					Orientation: Black,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{1, 4},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{2, 3},
				Piece: &Piece{
					Kind:        GuardPawn,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{2, 4},
				Piece: &Piece{
					Kind:        GuardPawn,
					Orientation: White,
					Moved:       true,
				},
			},
			WhiteKingStart,
			BlackKingStart,
		},
		PreviousFrom: AbsPoint{3, 3},
		PreviousTo:   AbsPoint{3, 4},
		From:         AbsPoint{1, 4},
		To:           AbsPoint{1, 3},
		Diff: PointSet{
			{
				AbsPoint: AbsPoint{0, 3},
			},
			{
				AbsPoint: AbsPoint{1, 2},
			},
			{
				AbsPoint: AbsPoint{2, 3},
			},
			{
				AbsPoint: AbsPoint{2, 4},
			},
			{
				AbsPoint: AbsPoint{1, 4},
			},
			{
				AbsPoint: AbsPoint{1, 3},
				Piece: &Piece{
					Kind:        GuardPawn,
					Orientation: Black,
				},
			},
		},
	},
}
