// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

var DetonateAfterMoveCases = []PositionAfterMoveCase{
	{
		Name: "Pawn Takes Detonator Adjacent To Guard",
		Initial: PointSet{
			WhiteKingStart,
			BlackKingStart,
			WhitePawn1Start,
			{
				AbsPoint: AbsPoint{0, 2},
				Piece: &Piece{
					Kind:        DetonatePawn,
					Orientation: Black,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{1, 3},
				Piece: &Piece{
					Kind:        GuardPawn,
					Orientation: Black,
					Moved:       true,
				},
			},
		},
		PreviousFrom: AbsPoint{3, 3},
		PreviousTo:   AbsPoint{3, 4},
		From:         AbsPoint{1, 1},
		To:           AbsPoint{0, 2},
		Diff: PointSet{
			{
				AbsPoint: AbsPoint{0, 2},
			},
			{
				AbsPoint: AbsPoint{1, 1},
			},
			{
				AbsPoint: AbsPoint{1, 3},
			},
		},
	},
	{
		Name: "False Detonator Chain",
		Initial: PointSet{
			{
				AbsPoint: AbsPoint{0, 0},
				Piece: &Piece{
					Kind:        DetonateKnight,
					Orientation: White,
				},
			},
			{
				AbsPoint: AbsPoint{1, 1},
				Piece: &Piece{
					Kind:        DetonatePawn,
					Orientation: Black,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{3, 2},
				Piece: &Piece{
					Kind:        DetonateRook,
					Orientation: White,
				},
			},
			WhiteKingStart,
			BlackKingStart,
			BlackLeftRookStart,
		},
		PreviousFrom: AbsPoint{3, 3},
		PreviousTo:   AbsPoint{3, 4},
		From:         AbsPoint{0, 7},
		To:           AbsPoint{0, 0},
		Diff: PointSet{
			{
				AbsPoint: AbsPoint{0, 7},
			},
			{
				AbsPoint: AbsPoint{0, 0},
			},
			{
				AbsPoint: AbsPoint{1, 1},
			},
		},
	},
	{
		Name: "Detonator Chain",
		Initial: PointSet{
			{
				AbsPoint: AbsPoint{0, 0},
				Piece: &Piece{
					Kind:        DetonateRook,
					Orientation: White,
					Moved:       false,
				},
			},
			{
				AbsPoint: AbsPoint{1, 1},
				Piece: &Piece{
					Kind:        DetonateRook,
					Orientation: Black,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{1, 2},
				Piece: &Piece{
					Kind:        DetonateKnight,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{1, 3},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{2, 1},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: Black,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{2, 2},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{0, 4},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: Black,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{2, 4},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{3, 2},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
				},
			},
			WhiteKingStart,
			BlackKingStart,
		},
		PreviousFrom: AbsPoint{3, 3},
		PreviousTo:   AbsPoint{3, 4},
		From:         AbsPoint{0, 4},
		To:           AbsPoint{0, 0},
		Diff: PointSet{
			{
				AbsPoint: AbsPoint{0, 4},
			},
			{
				AbsPoint: AbsPoint{0, 0},
			},
			{
				AbsPoint: AbsPoint{1, 1},
			},
			{
				AbsPoint: AbsPoint{2, 1},
			},
			{
				AbsPoint: AbsPoint{1, 2},
			},
			{
				AbsPoint: AbsPoint{2, 2},
			},
			{
				AbsPoint: AbsPoint{1, 3},
			},
		},
	},
	{
		Name: "White Pawn Takes Black Detonate Pawn Adjacent Black Guard Bishop",
		Initial: PointSet{
			WhitePawn1Start,
			WhiteKingStart,
			{
				AbsPoint: AbsPoint{0, 2},
				Piece: &Piece{
					Kind:        DetonatePawn,
					Orientation: Black,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{1, 3},
				Piece: &Piece{
					Kind:        DetonateBishop,
					Orientation: Black,
					Moved:       true,
				},
			},
		},
		PreviousFrom: AbsPoint{3, 3},
		PreviousTo:   AbsPoint{3, 4},
		From:         AbsPoint{1, 1},
		To:           AbsPoint{0, 2},
		Diff: PointSet{
			{
				AbsPoint: AbsPoint{1, 1},
			},
			{
				AbsPoint: AbsPoint{0, 2},
			},
			{
				AbsPoint: AbsPoint{1, 3},
			},
		},
	},
	{
		Name: "White Guard Rook Takes Moving Black Detonate Pawn, Surrounded",
		Initial: PointSet{
			{
				AbsPoint: AbsPoint{0, 0},
				Piece: &Piece{
					Kind:        GuardRook,
					Orientation: White,
					Moved:       false,
				},
			},
			WhiteLeftKnightStart,
			WhiteLeftBishopStart,
			WhiteKingStart,
			BlackKingStart,
			{
				AbsPoint: AbsPoint{0, 2},
				Piece: &Piece{
					Kind:        DetonatePawn,
					Orientation: Black,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{1, 2},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: Black,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{0, 3},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{2, 1},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
				},
			},
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
				AbsPoint: AbsPoint{1, 2},
			},
			{
				AbsPoint: AbsPoint{0, 0},
			},
			{
				AbsPoint: AbsPoint{1, 0},
			},
		},
	},
	{
		Name: "White Guard Rook Takes Moving Black Detonate Pawn",
		Initial: PointSet{
			{
				AbsPoint: AbsPoint{0, 0},
				Piece: &Piece{
					Kind:        GuardRook,
					Orientation: White,
					Moved:       false,
				},
			},
			WhiteKingStart,
			BlackKingStart,
			{
				AbsPoint: AbsPoint{0, 2},
				Piece: &Piece{
					Kind:        DetonatePawn,
					Orientation: Black,
					Moved:       true,
				},
			},
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
		},
	},
	{
		Name: "White Rook Takes Black Detonate Pawn",
		Initial: PointSet{
			WhiteLeftRookStart,
			WhiteKingStart,
			BlackKingStart,
			{
				AbsPoint: AbsPoint{0, 6},
				Piece: &Piece{
					Kind:        DetonatePawn,
					Orientation: Black,
					Moved:       false,
				},
			},
		},
		PreviousFrom: AbsPoint{3, 3},
		PreviousTo:   AbsPoint{3, 4},
		From:         AbsPoint{0, 0},
		To:           AbsPoint{0, 6},
		Diff: PointSet{
			{
				AbsPoint: AbsPoint{0, 0},
			},
			{
				AbsPoint: AbsPoint{0, 6},
			},
		},
	},
}
