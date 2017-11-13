// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

var SwapAfterMoveCases = []PositionAfterMoveCase{
	{
		Name: "Black Bishop Swaps With Pawn Across Board",
		Initial: PointSet{
			&WhiteKingStart: {},
			&BlackKingStart: {},
			{
				AbsPoint: AbsPoint{1, 1},
				Piece: &Piece{
					Kind:        SwapBishop,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPoint{1, 1}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{7, 7},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPoint{7, 7}.Index(),
				},
			}: {},
		},
		From: AbsPoint{1, 1},
		To:   AbsPoint{7, 7},
		Diff: PointSet{
			{
				AbsPoint: AbsPoint{1, 1},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
				},
			}: {},
			{
				AbsPoint: AbsPoint{7, 7},
				Piece: &Piece{
					Kind:        SwapBishop,
					Orientation: Black,
				},
			}: {},
		},
	},
	{
		Name: "White Swap Pawn Swaps Regular Pawn",
		Initial: PointSet{
			&WhiteKingStart: {},
			&BlackKingStart: {},
			{
				AbsPoint: AbsPoint{0, 1},
				Piece: &Piece{
					Kind:        SwapPawn,
					Orientation: White,
				},
			}: {},
			{
				AbsPoint: AbsPoint{0, 2},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPoint{0, 2}.Index(),
				},
			}: {},
		},
		From: AbsPoint{0, 1},
		To:   AbsPoint{0, 2},
		Diff: PointSet{
			{
				AbsPoint: AbsPoint{0, 1},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
				},
			}: {},
			{
				AbsPoint: AbsPoint{0, 2},
				Piece: &Piece{
					Kind:        SwapPawn,
					Orientation: White,
				},
			}: {},
		},
	},
}