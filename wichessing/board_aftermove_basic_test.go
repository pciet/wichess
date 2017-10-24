// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

var BasicAfterMoveCases = []PositionAfterMoveCase{
	{
		Name: "Moving White Queen",
		Initial: PointSet{
			&WhiteKingStart:  {},
			&WhiteQueenStart: {},
			&BlackKingStart:  {},
			&BlackQueenStart: {},
		},
		From: AbsPoint{File: 3, Rank: 0},
		To:   AbsPoint{File: 3, Rank: 5},
		Diff: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 3, Rank: 0},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 3, Rank: 5},
				Piece: &Piece{
					Kind:        Queen,
					Orientation: White,
				},
			}: {},
		},
	},
	{
		Name: "Moving Black Bishop",
		Initial: PointSet{
			&WhiteKingStart:        {},
			&WhiteLeftBishopStart:  {},
			&WhiteRightBishopStart: {},
			&BlackKingStart:        {},
			&BlackLeftBishopStart:  {},
			&BlackRightBishopStart: {},
		},
		From: AbsPoint{File: 2, Rank: 7},
		To:   AbsPoint{File: 0, Rank: 5},
		Diff: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 7},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 5},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: Black,
				},
			}: {},
		},
	},
	{
		Name: "Moving White Knight",
		Initial: PointSet{
			&WhiteKingStart:        {},
			&WhiteLeftKnightStart:  {},
			&WhiteRightKnightStart: {},
			&BlackKingStart:        {},
			&BlackLeftKnightStart:  {},
			&BlackRightKnightStart: {},
		},
		From: AbsPoint{File: 6, Rank: 0},
		To:   AbsPoint{File: 7, Rank: 2},
		Diff: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 6, Rank: 0},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 7, Rank: 2},
				Piece: &Piece{
					Kind:        Knight,
					Orientation: White,
				},
			}: {},
		},
	},
	{
		Name: "Moving Black Rook",
		Initial: PointSet{
			&WhiteKingStart:      {},
			&WhiteLeftRookStart:  {},
			&WhiteRightRookStart: {},
			&BlackKingStart:      {},
			&BlackLeftRookStart:  {},
			&BlackRightRookStart: {},
		},
		From: AbsPoint{File: 7, Rank: 7},
		To:   AbsPoint{File: 7, Rank: 3},
		Diff: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 7, Rank: 7},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 7, Rank: 3},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: Black,
				},
			}: {},
		},
	},
	{
		Name: "Moving Black King",
		Initial: PointSet{
			&WhiteKingStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 1},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       false,
					Previous:    AbsPointToIndex(AbsPoint{File: 0, Rank: 1}),
				},
			}: {},
			&BlackKingStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 4, Rank: 6},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       false,
					Previous:    AbsPointToIndex(AbsPoint{File: 4, Rank: 6}),
				},
			}: {},
		},
		From: AbsPoint{File: 4, Rank: 7},
		To:   AbsPoint{File: 5, Rank: 7},
		Diff: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 4, Rank: 7},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 5, Rank: 7},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
				},
			}: {},
		},
	},
}
