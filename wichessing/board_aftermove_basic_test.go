// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

var BasicAfterMoveCases = []PositionAfterMoveCase{
	{
		Name: "Chess Initial Position White Knight Move",
		Initial: PointSet{
			&WhiteLeftRookStart:    {},
			&WhiteLeftKnightStart:  {},
			&WhiteLeftBishopStart:  {},
			&WhiteQueenStart:       {},
			&WhiteKingStart:        {},
			&WhiteRightBishopStart: {},
			&WhiteRightKnightStart: {},
			&WhiteRightRookStart:   {},
			&WhitePawn0Start:       {},
			&WhitePawn1Start:       {},
			&WhitePawn2Start:       {},
			&WhitePawn3Start:       {},
			&WhitePawn4Start:       {},
			&WhitePawn5Start:       {},
			&WhitePawn6Start:       {},
			&WhitePawn7Start:       {},
			&BlackLeftRookStart:    {},
			&BlackLeftKnightStart:  {},
			&BlackLeftBishopStart:  {},
			&BlackQueenStart:       {},
			&BlackKingStart:        {},
			&BlackRightBishopStart: {},
			&BlackRightKnightStart: {},
			&BlackRightRookStart:   {},
			&BlackPawn0Start:       {},
			&BlackPawn1Start:       {},
			&BlackPawn2Start:       {},
			&BlackPawn3Start:       {},
			&BlackPawn4Start:       {},
			&BlackPawn5Start:       {},
			&BlackPawn6Start:       {},
			&BlackPawn7Start:       {},
		},
		From: AbsPoint{File: 1, Rank: 0},
		To:   AbsPoint{File: 0, Rank: 2},
		Diff: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 1, Rank: 0},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 2},
				Piece: &Piece{
					Kind:        Knight,
					Orientation: White,
				},
			}: {},
		},
	},
	{
		Name: "Chess Initial Position No Pawns White Queen Take",
		Initial: PointSet{
			&WhiteLeftRookStart:    {},
			&WhiteLeftKnightStart:  {},
			&WhiteLeftBishopStart:  {},
			&WhiteQueenStart:       {},
			&WhiteKingStart:        {},
			&WhiteRightBishopStart: {},
			&WhiteRightKnightStart: {},
			&WhiteRightRookStart:   {},
			&BlackLeftRookStart:    {},
			&BlackLeftKnightStart:  {},
			&BlackLeftBishopStart:  {},
			&BlackQueenStart:       {},
			&BlackKingStart:        {},
			&BlackRightBishopStart: {},
			&BlackRightKnightStart: {},
			&BlackRightRookStart:   {},
		},
		From: AbsPoint{File: 3, Rank: 0},
		To:   AbsPoint{File: 3, Rank: 7},
		Diff: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 3, Rank: 0},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 3, Rank: 7},
				Piece: &Piece{
					Kind:        Queen,
					Orientation: White,
				},
			}: {},
		},
	},
	{
		Name: "Moving Two White Pawn",
		Initial: PointSet{
			&WhiteKingStart:  {},
			&WhitePawn2Start: {},
			&BlackKingStart:  {},
			&BlackPawn7Start: {},
		},
		From: AbsPoint{File: 2, Rank: 1},
		To:   AbsPoint{File: 2, Rank: 3},
		Diff: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 1},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 3},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
				},
			}: {},
		},
	},
	{
		Name: "Moving Black Pawn",
		Initial: PointSet{
			&WhiteKingStart:  {},
			&WhitePawn0Start: {},
			&WhitePawn1Start: {},
			&WhitePawn2Start: {},
			&WhitePawn3Start: {},
			&WhitePawn4Start: {},
			&WhitePawn5Start: {},
			&WhitePawn6Start: {},
			&WhitePawn7Start: {},
			&BlackKingStart:  {},
			&BlackPawn0Start: {},
			&BlackPawn1Start: {},
			&BlackPawn2Start: {},
			&BlackPawn3Start: {},
			&BlackPawn4Start: {},
			&BlackPawn5Start: {},
			&BlackPawn6Start: {},
			&BlackPawn7Start: {},
		},
		From: AbsPoint{File: 5, Rank: 6},
		To:   AbsPoint{File: 5, Rank: 5},
		Diff: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 5, Rank: 6},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 5, Rank: 5},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
				},
			}: {},
		},
	},
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
