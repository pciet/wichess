// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

var CastlingAfterMoveCases = []PositionAfterMoveCase{
	{
		Name: "White Left Castling",
		Initial: PointSet{
			&WhiteLeftRookStart:  {},
			&WhiteKingStart:      {},
			&WhiteRightRookStart: {},
			&BlackLeftRookStart:  {},
			&BlackKingStart:      {},
			&BlackRightRookStart: {},
		},
		From: AbsPoint{File: 4, Rank: 0},
		To:   AbsPoint{File: 2, Rank: 0},
		Diff: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 0},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 4, Rank: 0},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 0},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 3, Rank: 0},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: White,
				},
			}: {},
		},
	},
	{
		Name: "White Right Castling",
		Initial: PointSet{
			&WhiteLeftRookStart:  {},
			&WhiteKingStart:      {},
			&WhiteRightRookStart: {},
			&BlackLeftRookStart:  {},
			&BlackKingStart:      {},
			&BlackRightRookStart: {},
		},
		From: AbsPoint{File: 4, Rank: 0},
		To:   AbsPoint{File: 6, Rank: 0},
		Diff: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 7, Rank: 0},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 4, Rank: 0},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 6, Rank: 0},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 5, Rank: 0},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: White,
				},
			}: {},
		},
	},
	{
		Name: "Black Left Castling",
		Initial: PointSet{
			&WhiteLeftRookStart:  {},
			&WhiteKingStart:      {},
			&WhiteRightRookStart: {},
			&BlackLeftRookStart:  {},
			&BlackKingStart:      {},
			&BlackRightRookStart: {},
		},
		From: AbsPoint{File: 4, Rank: 7},
		To:   AbsPoint{File: 2, Rank: 7},
		Diff: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 7},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 4, Rank: 7},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 7},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 3, Rank: 7},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: Black,
				},
			}: {},
		},
	},
	{
		Name: "Black Right Castling",
		Initial: PointSet{
			&WhiteLeftRookStart:  {},
			&WhiteKingStart:      {},
			&WhiteRightRookStart: {},
			&BlackLeftRookStart:  {},
			&BlackKingStart:      {},
			&BlackRightRookStart: {},
		},
		From: AbsPoint{File: 4, Rank: 7},
		To:   AbsPoint{File: 6, Rank: 7},
		Diff: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 7, Rank: 7},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 4, Rank: 7},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 6, Rank: 7},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 5, Rank: 7},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: Black,
				},
			}: {},
		},
	},
}