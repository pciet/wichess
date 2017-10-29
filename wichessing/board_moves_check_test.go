// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

var CheckMovesCases = []AvailableMovesCase{
	{
		Name:   "False Check, Detonate Rook Vs Guard Bishop",
		Active: Black,
		Position: PointSet{
			&WhiteKingStart: {},
			{
				AbsPoint: AbsPoint{1, 3},
				Piece: &Piece{
					Kind:        GuardBishop,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPoint{1, 3}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{0, 7},
				Piece: &Piece{
					Kind:        DetonateRook,
					Orientation: Black,
					Moved:       false,
				},
			}: {},
			&BlackKingStart: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			{4, 0}: {
				{3, 0}: {},
				{3, 1}: {},
				{4, 1}: {},
				{5, 1}: {},
				{5, 0}: {},
			},
			{1, 3}: {
				{0, 2}: {},
				{2, 2}: {},
				{3, 1}: {},
				{0, 4}: {},
				{2, 4}: {},
				{3, 5}: {},
				{4, 6}: {},
				{5, 7}: {},
			},
			{0, 7}: {
				{0, 6}: {},
				{0, 5}: {},
				{0, 4}: {},
				{0, 3}: {},
				{0, 2}: {},
				{0, 1}: {},
				{0, 0}: {},
				{1, 7}: {},
				{2, 7}: {},
				{3, 7}: {},
			},
			{4, 7}: {
				{3, 7}: {},
				{3, 6}: {},
				{5, 6}: {},
				{2, 7}: {},
			},
		},
	},
	{
		Name:   "Black Detonate Pawn Versus White Guard Rook Adjacent White King",
		Active: White,
		Check:  true,
		Position: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 0},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
					Previous:    4,
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 3, Rank: 0},
				Piece: &Piece{
					Kind:        GuardRook,
					Orientation: White,
					Moved:       true,
					Previous:    0,
				},
			}: {},
			&WhitePawn6Start: {},
			&BlackKingStart:  {},
			&Point{
				AbsPoint: AbsPoint{File: 3, Rank: 2},
				Piece: &Piece{
					Kind:        DetonatePawn,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 3, Rank: 3}),
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			AbsPoint{File: 2, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 1, Rank: 0}: {},
				&AbsPoint{File: 1, Rank: 1}: {},
				&AbsPoint{File: 3, Rank: 1}: {},
			},
			AbsPoint{File: 3, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 1}: {},
				&AbsPoint{File: 3, Rank: 2}: {},
				&AbsPoint{File: 5, Rank: 0}: {},
				&AbsPoint{File: 6, Rank: 0}: {},
				&AbsPoint{File: 7, Rank: 0}: {},
			},
			AbsPoint{File: 4, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 7}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 7}: {},
			},
			AbsPoint{File: 3, Rank: 2}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 1}: {},
			},
		},
	},
	{
		Name:   "Black Detonate Pawn Adjacent White King",
		Active: White,
		Check:  false,
		Position: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 3, Rank: 2},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 3, Rank: 2}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 5, Rank: 1},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 5, Rank: 1}),
				},
			}: {},
			&WhitePawn6Start: {},
			&Point{
				AbsPoint: AbsPoint{File: 3, Rank: 3},
				Piece: &Piece{
					Kind:        DetonatePawn,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 3, Rank: 4}),
				},
			}: {},
			&BlackKingStart: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			AbsPoint{File: 3, Rank: 2}: AbsPointSet{
				&AbsPoint{File: 2, Rank: 3}: {},
				&AbsPoint{File: 4, Rank: 3}: {},
				&AbsPoint{File: 2, Rank: 1}: {},
				&AbsPoint{File: 3, Rank: 1}: {},
				&AbsPoint{File: 4, Rank: 1}: {},
			},
			AbsPoint{File: 5, Rank: 1}: AbsPointSet{
				&AbsPoint{File: 4, Rank: 0}: {},
				&AbsPoint{File: 6, Rank: 0}: {},
				&AbsPoint{File: 6, Rank: 2}: {},
				&AbsPoint{File: 7, Rank: 3}: {},
				&AbsPoint{File: 4, Rank: 2}: {},
			},
			AbsPoint{File: 6, Rank: 1}: AbsPointSet{
				&AbsPoint{File: 6, Rank: 2}: {},
				&AbsPoint{File: 6, Rank: 3}: {},
			},
			AbsPoint{File: 4, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 7}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 7}: {},
			},
		},
	},
	{
		Name:   "Black Detonate Pawn Checking White King",
		Active: White,
		Check:  true,
		Position: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 2},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 2, Rank: 2}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 5, Rank: 1},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 5, Rank: 1}),
				},
			}: {},
			&WhitePawn6Start: {},
			&Point{
				AbsPoint: AbsPoint{File: 3, Rank: 3},
				Piece: &Piece{
					Kind:        DetonatePawn,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 3, Rank: 4}),
				},
			}: {},
			&BlackKingStart: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			AbsPoint{File: 2, Rank: 2}: AbsPointSet{
				&AbsPoint{File: 2, Rank: 3}: {},
				&AbsPoint{File: 1, Rank: 3}: {},
				&AbsPoint{File: 1, Rank: 2}: {},
				&AbsPoint{File: 1, Rank: 1}: {},
				&AbsPoint{File: 2, Rank: 1}: {},
				&AbsPoint{File: 3, Rank: 1}: {},
				&AbsPoint{File: 3, Rank: 2}: {},
			},
			AbsPoint{File: 3, Rank: 3}: AbsPointSet{
				&AbsPoint{File: 2, Rank: 2}: {},
				&AbsPoint{File: 3, Rank: 2}: {},
			},
			AbsPoint{File: 4, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 7}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 7}: {},
			},
		},
	},
	{
		Name:   "Black Pawn Checking White King",
		Active: White,
		Check:  true,
		Position: PointSet{
			&WhiteKingStart:  {},
			&WhitePawn5Start: {},
			&BlackKingStart:  {},
			&Point{
				AbsPoint: AbsPoint{File: 3, Rank: 1},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 3, Rank: 2}),
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			AbsPoint{File: 4, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 0}: {},
				&AbsPoint{File: 3, Rank: 1}: {},
				&AbsPoint{File: 4, Rank: 1}: {},
				&AbsPoint{File: 5, Rank: 0}: {},
			},
			AbsPoint{File: 3, Rank: 1}: AbsPointSet{
				&AbsPoint{File: 4, Rank: 0}: {},
				&AbsPoint{File: 3, Rank: 0}: {},
			},
			AbsPoint{File: 4, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 7}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 7}: {},
			},
		},
	},
	{
		Name:   "Chess Initial Position No Pawns White Queen Checking",
		Active: Black,
		Check:  true,
		Position: PointSet{
			&WhiteLeftRookStart:   {},
			&WhiteLeftKnightStart: {},
			&WhiteLeftBishopStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 3, Rank: 6},
				Piece: &Piece{
					Kind:        Queen,
					Orientation: White,
					Moved:       true,
					Previous:    3,
				},
			}: {},
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
		Moves: map[AbsPoint]AbsPointSet{
			AbsPoint{File: 0, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 0, Rank: 1}: {},
				&AbsPoint{File: 0, Rank: 2}: {},
				&AbsPoint{File: 0, Rank: 3}: {},
				&AbsPoint{File: 0, Rank: 4}: {},
				&AbsPoint{File: 0, Rank: 5}: {},
				&AbsPoint{File: 0, Rank: 6}: {},
				&AbsPoint{File: 0, Rank: 7}: {},
			},
			AbsPoint{File: 1, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 0, Rank: 2}: {},
				&AbsPoint{File: 2, Rank: 2}: {},
				&AbsPoint{File: 3, Rank: 1}: {},
			},
			AbsPoint{File: 2, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 1, Rank: 1}: {},
				&AbsPoint{File: 0, Rank: 2}: {},
				&AbsPoint{File: 3, Rank: 1}: {},
				&AbsPoint{File: 4, Rank: 2}: {},
				&AbsPoint{File: 5, Rank: 3}: {},
				&AbsPoint{File: 6, Rank: 4}: {},
				&AbsPoint{File: 7, Rank: 5}: {},
			},
			AbsPoint{File: 3, Rank: 6}: AbsPointSet{
				&AbsPoint{File: 2, Rank: 7}: {},
				&AbsPoint{File: 3, Rank: 7}: {},
				&AbsPoint{File: 4, Rank: 7}: {},
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
				&AbsPoint{File: 6, Rank: 6}: {},
				&AbsPoint{File: 7, Rank: 6}: {},
				&AbsPoint{File: 4, Rank: 5}: {},
				&AbsPoint{File: 5, Rank: 4}: {},
				&AbsPoint{File: 6, Rank: 3}: {},
				&AbsPoint{File: 7, Rank: 2}: {},
				&AbsPoint{File: 3, Rank: 5}: {},
				&AbsPoint{File: 3, Rank: 4}: {},
				&AbsPoint{File: 3, Rank: 3}: {},
				&AbsPoint{File: 3, Rank: 2}: {},
				&AbsPoint{File: 3, Rank: 1}: {},
				&AbsPoint{File: 3, Rank: 0}: {},
				&AbsPoint{File: 2, Rank: 6}: {},
				&AbsPoint{File: 1, Rank: 6}: {},
				&AbsPoint{File: 0, Rank: 6}: {},
				&AbsPoint{File: 2, Rank: 5}: {},
				&AbsPoint{File: 1, Rank: 4}: {},
				&AbsPoint{File: 0, Rank: 3}: {},
			},
			AbsPoint{File: 4, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 4, Rank: 1}: {},
				&AbsPoint{File: 5, Rank: 1}: {},
				&AbsPoint{File: 3, Rank: 0}: {},
				&AbsPoint{File: 3, Rank: 1}: {},
			},
			AbsPoint{File: 5, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 4, Rank: 1}: {},
				&AbsPoint{File: 3, Rank: 2}: {},
				&AbsPoint{File: 2, Rank: 3}: {},
				&AbsPoint{File: 1, Rank: 4}: {},
				&AbsPoint{File: 0, Rank: 5}: {},
				&AbsPoint{File: 6, Rank: 1}: {},
				&AbsPoint{File: 7, Rank: 2}: {},
			},
			AbsPoint{File: 6, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 7, Rank: 2}: {},
				&AbsPoint{File: 5, Rank: 2}: {},
				&AbsPoint{File: 4, Rank: 1}: {},
			},
			AbsPoint{File: 7, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 7, Rank: 1}: {},
				&AbsPoint{File: 7, Rank: 2}: {},
				&AbsPoint{File: 7, Rank: 3}: {},
				&AbsPoint{File: 7, Rank: 4}: {},
				&AbsPoint{File: 7, Rank: 5}: {},
				&AbsPoint{File: 7, Rank: 6}: {},
				&AbsPoint{File: 7, Rank: 7}: {},
			},
			AbsPoint{File: 1, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 6}: {},
			},
			AbsPoint{File: 2, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 6}: {},
			},
			AbsPoint{File: 3, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 6}: {},
			},
			AbsPoint{File: 4, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 6}: {},
			},
		},
	},
}
