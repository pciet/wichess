// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

var FortifyAfterMoveCases = []PositionAfterMoveCase{
	{
		Name: "Guard Pawn Can't Take Fortify Pawn",
		Initial: PointSet{
			&WhiteKingStart: {},
			&BlackKingStart: {},
			{
				AbsPoint: AbsPoint{0, 1},
				Piece: &Piece{
					Kind:        FortifyPawn,
					Orientation: White,
				},
			}: {},
			{
				AbsPoint: AbsPoint{0, 3},
				Piece: &Piece{
					Kind:        GuardPawn,
					Orientation: Black,
				},
			}: {},
		},
		PreviousFrom: AbsPoint{3, 3},
		PreviousTo:   AbsPoint{3, 4},
		From:         AbsPoint{0, 1},
		To:           AbsPoint{0, 2},
		Diff: PointSet{
			{
				AbsPoint: AbsPoint{0, 1},
			}: {},
			{
				AbsPoint: AbsPoint{0, 2},
				Piece: &Piece{
					Kind:        FortifyPawn,
					Orientation: White,
				},
			}: {},
		},
	},
}
