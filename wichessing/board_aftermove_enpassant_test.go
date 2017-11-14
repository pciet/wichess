// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

var EnPassantAfterMoveCases = []PositionAfterMoveCase{
	{
		Name: "White En Passant Take",
		Initial: PointSet{
			&WhiteKingStart: {},
			&BlackKingStart: {},
			&Point{
				AbsPoint: AbsPoint{0, 4},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{1, 4},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
				},
			}: {},
		},
		PreviousFrom: AbsPoint{0, 6},
		PreviousTo:   AbsPoint{0, 4},
		From:         AbsPoint{1, 4},
		To:           AbsPoint{0, 5},
		Diff: PointSet{
			&Point{
				AbsPoint: AbsPoint{0, 4},
			}: {},
			&Point{
				AbsPoint: AbsPoint{1, 4},
			}: {},
			&Point{
				AbsPoint: AbsPoint{0, 5},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
				},
			}: {},
		},
	},
}
