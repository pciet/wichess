// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

var BasicMovesCases = []AvailableMovesCase{
	{
		Name:   "Failed Random Test",
		Active: Black,
		Position: PointSet{
			{
				AbsPoint: AbsPoint{5, 0},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPoint{5, 0}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{6, 0},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPoint{6, 0}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{2, 1},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPoint{2, 1}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{6, 1},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPoint{6, 1}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{7, 1},
				Piece: &Piece{
					Kind:        Knight,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPoint{7, 1}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{1, 2},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPoint{1, 2}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{3, 2},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPoint{3, 2}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{4, 4},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPoint{4, 4}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{6, 4},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPoint{6, 4}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{3, 5},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPoint{3, 5}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{5, 5},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPoint{5, 5}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{6, 6},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPoint{6, 6}.Index(),
				},
			}: {},
			&BlackLeftRookStart: {},
			{
				AbsPoint: AbsPoint{3, 7},
				Piece: &Piece{
					Kind:        Queen,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPoint{3, 7}.Index(),
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			{5, 0}: {
				{4, 0}: {},
				{3, 0}: {},
				{2, 0}: {},
				{1, 0}: {},
				{0, 0}: {},
				{5, 1}: {},
				{5, 2}: {},
				{5, 3}: {},
				{5, 4}: {},
				{5, 5}: {},
			},
			{6, 0}: {
				{7, 0}: {},
				{6, 1}: {},
			},
			{2, 1}: {
				{2, 0}: {},
			},
			{6, 1}: {
				{5, 0}: {},
			},
			{7, 1}: {
				{6, 3}: {},
				{5, 2}: {},
			},
			{1, 2}: {
				{1, 3}: {},
			},
			{3, 2}: {
				{2, 1}: {},
				{2, 2}: {},
				{2, 3}: {},
				{4, 2}: {},
				{4, 1}: {},
				{3, 1}: {},
			},
			{4, 4}: {
				{3, 4}: {},
				{4, 5}: {},
			},
			{6, 4}: {
				{5, 5}: {},
				{6, 5}: {},
			},
			{3, 5}: {
				{3, 6}: {},
			},
			{5, 5}: {
				{5, 4}: {},
				{5, 3}: {},
				{5, 2}: {},
				{5, 1}: {},
				{5, 0}: {},
				{6, 5}: {},
				{7, 5}: {},
				{5, 6}: {},
				{5, 7}: {},
				{4, 5}: {},
				{3, 5}: {},
			},
			{6, 6}: {
				{7, 7}: {},
				{7, 5}: {},
				{5, 7}: {},
				{5, 5}: {},
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
			{3, 7}: {
				{2, 7}: {},
				{1, 7}: {},
				{0, 7}: {},
				{4, 7}: {},
				{5, 7}: {},
				{6, 7}: {},
				{7, 7}: {},
				{2, 6}: {},
				{1, 5}: {},
				{0, 4}: {},
				{3, 6}: {},
				{4, 6}: {},
				{5, 5}: {},
			},
		},
	},
	{
		Name:   "Chess Initial Position No Pawns",
		Active: White,
		Position: PointSet{
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
			AbsPoint{File: 3, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 2, Rank: 1}: {},
				&AbsPoint{File: 1, Rank: 2}: {},
				&AbsPoint{File: 0, Rank: 3}: {},
				&AbsPoint{File: 3, Rank: 1}: {},
				&AbsPoint{File: 3, Rank: 2}: {},
				&AbsPoint{File: 3, Rank: 3}: {},
				&AbsPoint{File: 3, Rank: 4}: {},
				&AbsPoint{File: 3, Rank: 5}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 3, Rank: 7}: {},
				&AbsPoint{File: 4, Rank: 1}: {},
				&AbsPoint{File: 5, Rank: 2}: {},
				&AbsPoint{File: 6, Rank: 3}: {},
				&AbsPoint{File: 7, Rank: 4}: {},
			},
			AbsPoint{File: 4, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 4, Rank: 1}: {},
				&AbsPoint{File: 5, Rank: 1}: {},
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
			AbsPoint{File: 0, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 0, Rank: 6}: {},
				&AbsPoint{File: 0, Rank: 5}: {},
				&AbsPoint{File: 0, Rank: 4}: {},
				&AbsPoint{File: 0, Rank: 3}: {},
				&AbsPoint{File: 0, Rank: 2}: {},
				&AbsPoint{File: 0, Rank: 1}: {},
				&AbsPoint{File: 0, Rank: 0}: {},
			},
			AbsPoint{File: 1, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 0, Rank: 5}: {},
				&AbsPoint{File: 2, Rank: 5}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
			},
			AbsPoint{File: 2, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 1, Rank: 6}: {},
				&AbsPoint{File: 0, Rank: 5}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 4, Rank: 5}: {},
				&AbsPoint{File: 5, Rank: 4}: {},
				&AbsPoint{File: 6, Rank: 3}: {},
				&AbsPoint{File: 7, Rank: 2}: {},
			},
			AbsPoint{File: 3, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 2, Rank: 6}: {},
				&AbsPoint{File: 1, Rank: 5}: {},
				&AbsPoint{File: 0, Rank: 4}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 3, Rank: 5}: {},
				&AbsPoint{File: 3, Rank: 4}: {},
				&AbsPoint{File: 3, Rank: 3}: {},
				&AbsPoint{File: 3, Rank: 2}: {},
				&AbsPoint{File: 3, Rank: 1}: {},
				&AbsPoint{File: 3, Rank: 0}: {},
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 5}: {},
				&AbsPoint{File: 6, Rank: 4}: {},
				&AbsPoint{File: 7, Rank: 3}: {},
			},
			AbsPoint{File: 4, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
			},
			AbsPoint{File: 5, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 3, Rank: 5}: {},
				&AbsPoint{File: 2, Rank: 4}: {},
				&AbsPoint{File: 1, Rank: 3}: {},
				&AbsPoint{File: 0, Rank: 2}: {},
				&AbsPoint{File: 6, Rank: 6}: {},
				&AbsPoint{File: 7, Rank: 5}: {},
			},
			AbsPoint{File: 6, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 7, Rank: 5}: {},
				&AbsPoint{File: 5, Rank: 5}: {},
				&AbsPoint{File: 4, Rank: 6}: {},
			},
			AbsPoint{File: 7, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 7, Rank: 6}: {},
				&AbsPoint{File: 7, Rank: 5}: {},
				&AbsPoint{File: 7, Rank: 4}: {},
				&AbsPoint{File: 7, Rank: 3}: {},
				&AbsPoint{File: 7, Rank: 2}: {},
				&AbsPoint{File: 7, Rank: 1}: {},
				&AbsPoint{File: 7, Rank: 0}: {},
			},
		},
	},
	{
		Name:   "Chess Initial Position",
		Active: White,
		Position: PointSet{
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
		Moves: map[AbsPoint]AbsPointSet{
			AbsPoint{File: 1, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 0, Rank: 2}: {},
				&AbsPoint{File: 2, Rank: 2}: {},
			},
			AbsPoint{File: 6, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 5, Rank: 2}: {},
				&AbsPoint{File: 7, Rank: 2}: {},
			},
			AbsPoint{File: 0, Rank: 1}: AbsPointSet{
				&AbsPoint{File: 0, Rank: 2}: {},
				&AbsPoint{File: 0, Rank: 3}: {},
			},
			AbsPoint{File: 1, Rank: 1}: AbsPointSet{
				&AbsPoint{File: 1, Rank: 2}: {},
				&AbsPoint{File: 1, Rank: 3}: {},
			},
			AbsPoint{File: 2, Rank: 1}: AbsPointSet{
				&AbsPoint{File: 2, Rank: 2}: {},
				&AbsPoint{File: 2, Rank: 3}: {},
			},
			AbsPoint{File: 3, Rank: 1}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 2}: {},
				&AbsPoint{File: 3, Rank: 3}: {},
			},
			AbsPoint{File: 4, Rank: 1}: AbsPointSet{
				&AbsPoint{File: 4, Rank: 2}: {},
				&AbsPoint{File: 4, Rank: 3}: {},
			},
			AbsPoint{File: 5, Rank: 1}: AbsPointSet{
				&AbsPoint{File: 5, Rank: 2}: {},
				&AbsPoint{File: 5, Rank: 3}: {},
			},
			AbsPoint{File: 6, Rank: 1}: AbsPointSet{
				&AbsPoint{File: 6, Rank: 2}: {},
				&AbsPoint{File: 6, Rank: 3}: {},
			},
			AbsPoint{File: 7, Rank: 1}: AbsPointSet{
				&AbsPoint{File: 7, Rank: 2}: {},
				&AbsPoint{File: 7, Rank: 3}: {},
			},
			AbsPoint{File: 1, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 0, Rank: 5}: {},
				&AbsPoint{File: 2, Rank: 5}: {},
			},
			AbsPoint{File: 6, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 5, Rank: 5}: {},
				&AbsPoint{File: 7, Rank: 5}: {},
			},
			AbsPoint{File: 0, Rank: 6}: AbsPointSet{
				&AbsPoint{File: 0, Rank: 5}: {},
				&AbsPoint{File: 0, Rank: 4}: {},
			},
			AbsPoint{File: 1, Rank: 6}: AbsPointSet{
				&AbsPoint{File: 1, Rank: 5}: {},
				&AbsPoint{File: 1, Rank: 4}: {},
			},
			AbsPoint{File: 2, Rank: 6}: AbsPointSet{
				&AbsPoint{File: 2, Rank: 5}: {},
				&AbsPoint{File: 2, Rank: 4}: {},
			},
			AbsPoint{File: 3, Rank: 6}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 5}: {},
				&AbsPoint{File: 3, Rank: 4}: {},
			},
			AbsPoint{File: 4, Rank: 6}: AbsPointSet{
				&AbsPoint{File: 4, Rank: 5}: {},
				&AbsPoint{File: 4, Rank: 4}: {},
			},
			AbsPoint{File: 5, Rank: 6}: AbsPointSet{
				&AbsPoint{File: 5, Rank: 5}: {},
				&AbsPoint{File: 5, Rank: 4}: {},
			},
			AbsPoint{File: 6, Rank: 6}: AbsPointSet{
				&AbsPoint{File: 6, Rank: 5}: {},
				&AbsPoint{File: 6, Rank: 4}: {},
			},
			AbsPoint{File: 7, Rank: 6}: AbsPointSet{
				&AbsPoint{File: 7, Rank: 5}: {},
				&AbsPoint{File: 7, Rank: 4}: {},
			},
		},
	},
	{
		Name:   "Kings And Four Bishops",
		Active: White,
		Position: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 4, Rank: 1},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
					Previous:    4,
				},
			}: {},
			&WhiteLeftBishopStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 3, Rank: 2},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: White,
					Moved:       true,
					Previous:    5,
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 4, Rank: 6},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 4, Rank: 6}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 5},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 2, Rank: 7}),
				},
			}: {},
			&BlackRightBishopStart: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			AbsPoint{File: 4, Rank: 1}: AbsPointSet{
				&AbsPoint{File: 4, Rank: 0}: {},
				&AbsPoint{File: 5, Rank: 0}: {},
				&AbsPoint{File: 5, Rank: 1}: {},
				&AbsPoint{File: 5, Rank: 2}: {},
				&AbsPoint{File: 4, Rank: 2}: {},
				&AbsPoint{File: 3, Rank: 1}: {},
				&AbsPoint{File: 3, Rank: 0}: {},
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
			AbsPoint{File: 3, Rank: 2}: AbsPointSet{
				&AbsPoint{File: 2, Rank: 3}: {},
				&AbsPoint{File: 1, Rank: 4}: {},
				&AbsPoint{File: 0, Rank: 5}: {},
				// other moves put White King in check
			},
			AbsPoint{File: 4, Rank: 6}: AbsPointSet{
				&AbsPoint{File: 4, Rank: 5}: {},
				&AbsPoint{File: 5, Rank: 5}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
				&AbsPoint{File: 4, Rank: 7}: {},
				&AbsPoint{File: 3, Rank: 7}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 3, Rank: 5}: {},
			},
			AbsPoint{File: 0, Rank: 5}: AbsPointSet{
				&AbsPoint{File: 1, Rank: 6}: {},
				&AbsPoint{File: 2, Rank: 7}: {},
				&AbsPoint{File: 1, Rank: 4}: {},
				&AbsPoint{File: 2, Rank: 3}: {},
				&AbsPoint{File: 3, Rank: 2}: {},
			},
			AbsPoint{File: 5, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 6, Rank: 6}: {},
				&AbsPoint{File: 7, Rank: 5}: {},
				// Black King blocks other moves
			},
		},
	},
	{
		Name:   "Unmoved Kings And Four Moved Rooks",
		Active: White,
		Position: PointSet{
			&WhiteKingStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 1},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: White,
					Moved:       true,
					Previous:    2,
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 7, Rank: 0},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: White,
					Moved:       true,
					Previous:    7,
				},
			}: {},
			&BlackKingStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 6, Rank: 6},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 6, Rank: 7}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 7, Rank: 7},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: Black,
					Moved:       true,
					Previous:    63,
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			AbsPoint{File: 4, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 0}: {},
				&AbsPoint{File: 3, Rank: 1}: {},
				&AbsPoint{File: 4, Rank: 1}: {},
				&AbsPoint{File: 5, Rank: 1}: {},
				&AbsPoint{File: 5, Rank: 0}: {},
			},
			AbsPoint{File: 2, Rank: 1}: AbsPointSet{
				&AbsPoint{File: 1, Rank: 1}: {},
				&AbsPoint{File: 0, Rank: 1}: {},
				&AbsPoint{File: 2, Rank: 0}: {},
				&AbsPoint{File: 3, Rank: 1}: {},
				&AbsPoint{File: 4, Rank: 1}: {},
				&AbsPoint{File: 5, Rank: 1}: {},
				&AbsPoint{File: 6, Rank: 1}: {},
				&AbsPoint{File: 7, Rank: 1}: {},
				&AbsPoint{File: 2, Rank: 2}: {},
				&AbsPoint{File: 2, Rank: 3}: {},
				&AbsPoint{File: 2, Rank: 4}: {},
				&AbsPoint{File: 2, Rank: 5}: {},
				&AbsPoint{File: 2, Rank: 6}: {},
				&AbsPoint{File: 2, Rank: 7}: {},
			},
			AbsPoint{File: 7, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 6, Rank: 0}: {},
				&AbsPoint{File: 5, Rank: 0}: {},
				// White King at 4-0
				&AbsPoint{File: 7, Rank: 1}: {},
				&AbsPoint{File: 7, Rank: 2}: {},
				&AbsPoint{File: 7, Rank: 3}: {},
				&AbsPoint{File: 7, Rank: 4}: {},
				&AbsPoint{File: 7, Rank: 5}: {},
				&AbsPoint{File: 7, Rank: 6}: {},
				&AbsPoint{File: 7, Rank: 7}: {},
			},
			AbsPoint{File: 4, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 7}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 7}: {},
			},
			AbsPoint{File: 6, Rank: 6}: AbsPointSet{
				&AbsPoint{File: 6, Rank: 7}: {},
				&AbsPoint{File: 7, Rank: 6}: {},
				&AbsPoint{File: 6, Rank: 5}: {},
				&AbsPoint{File: 6, Rank: 4}: {},
				&AbsPoint{File: 6, Rank: 3}: {},
				&AbsPoint{File: 6, Rank: 2}: {},
				&AbsPoint{File: 6, Rank: 1}: {},
				&AbsPoint{File: 6, Rank: 0}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 2, Rank: 6}: {},
				&AbsPoint{File: 1, Rank: 6}: {},
				&AbsPoint{File: 0, Rank: 6}: {},
			},
			AbsPoint{File: 7, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 6, Rank: 7}: {},
				&AbsPoint{File: 5, Rank: 7}: {},
				// Black King at 4-7
				&AbsPoint{File: 7, Rank: 6}: {},
				&AbsPoint{File: 7, Rank: 5}: {},
				&AbsPoint{File: 7, Rank: 4}: {},
				&AbsPoint{File: 7, Rank: 3}: {},
				&AbsPoint{File: 7, Rank: 2}: {},
				&AbsPoint{File: 7, Rank: 1}: {},
				&AbsPoint{File: 7, Rank: 0}: {},
			},
		},
	},
	{
		Name:   "Unmoved Kings And Interacting Pawns",
		Active: White,
		Position: PointSet{
			&WhiteKingStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 3},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 0, Rank: 1}),
				},
			}: {},
			&BlackKingStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 1, Rank: 4},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 1, Rank: 6}),
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			AbsPoint{File: 4, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 0}: {},
				&AbsPoint{File: 3, Rank: 1}: {},
				&AbsPoint{File: 4, Rank: 1}: {},
				&AbsPoint{File: 5, Rank: 1}: {},
				&AbsPoint{File: 5, Rank: 0}: {},
			},
			AbsPoint{File: 0, Rank: 3}: AbsPointSet{
				&AbsPoint{File: 1, Rank: 4}: {},
				&AbsPoint{File: 0, Rank: 4}: {},
			},
			AbsPoint{File: 4, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 7}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 7}: {},
			},
			AbsPoint{File: 1, Rank: 4}: AbsPointSet{
				&AbsPoint{File: 1, Rank: 3}: {},
				&AbsPoint{File: 0, Rank: 3}: {},
			},
		},
	},
	{
		Name:   "Unmoved Kings, First Move Pawns, Black Moved Pawn",
		Active: White,
		Position: PointSet{
			&WhiteKingStart:  {},
			&WhitePawn5Start: {},
			&WhitePawn6Start: {},
			&BlackKingStart:  {},
			&BlackPawn0Start: {},
			&BlackPawn3Start: {},
			&Point{
				AbsPoint: AbsPoint{File: 7, Rank: 5},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 7, Rank: 6}),
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			AbsPoint{File: 4, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 0}: {},
				&AbsPoint{File: 3, Rank: 1}: {},
				&AbsPoint{File: 4, Rank: 1}: {},
				// Pawn at 5-1
				&AbsPoint{File: 5, Rank: 0}: {},
			},
			AbsPoint{File: 5, Rank: 1}: AbsPointSet{
				&AbsPoint{File: 5, Rank: 2}: {},
				&AbsPoint{File: 5, Rank: 3}: {},
			},
			AbsPoint{File: 6, Rank: 1}: AbsPointSet{
				&AbsPoint{File: 6, Rank: 2}: {},
				&AbsPoint{File: 6, Rank: 3}: {},
			},
			AbsPoint{File: 4, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 7}: {},
				// Pawn at 3-6
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 7}: {},
			},
			AbsPoint{File: 0, Rank: 6}: AbsPointSet{
				&AbsPoint{File: 0, Rank: 5}: {},
				&AbsPoint{File: 0, Rank: 4}: {},
			},
			AbsPoint{File: 3, Rank: 6}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 5}: {},
				&AbsPoint{File: 3, Rank: 4}: {},
			},
			AbsPoint{File: 7, Rank: 5}: AbsPointSet{
				&AbsPoint{File: 7, Rank: 4}: {},
			},
		},
	},
	{
		Name:   "Unmoved Kings And Taking Knights",
		Active: White,
		Position: PointSet{
			&WhiteKingStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 2},
				Piece: &Piece{
					Kind:        Knight,
					Orientation: White,
					Moved:       true,
					Previous:    1,
				},
			}: {},
			&WhiteRightKnightStart: {},
			&BlackKingStart:        {},
			&Point{
				AbsPoint: AbsPoint{File: 1, Rank: 4},
				Piece: &Piece{
					Kind:        Knight,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 1, Rank: 4}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 7, Rank: 2},
				Piece: &Piece{
					Kind:        Knight,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 7, Rank: 2}),
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			AbsPoint{File: 4, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 0}: {},
				&AbsPoint{File: 3, Rank: 1}: {},
				&AbsPoint{File: 4, Rank: 1}: {},
				// 5-1 is into check
				&AbsPoint{File: 5, Rank: 0}: {},
			},
			AbsPoint{File: 2, Rank: 2}: AbsPointSet{
				&AbsPoint{File: 1, Rank: 0}: {},
				&AbsPoint{File: 3, Rank: 0}: {},
				&AbsPoint{File: 0, Rank: 1}: {},
				&AbsPoint{File: 0, Rank: 3}: {},
				&AbsPoint{File: 3, Rank: 4}: {},
				&AbsPoint{File: 1, Rank: 4}: {},
				&AbsPoint{File: 4, Rank: 1}: {},
				&AbsPoint{File: 4, Rank: 3}: {},
			},
			AbsPoint{File: 6, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 7, Rank: 2}: {},
				&AbsPoint{File: 5, Rank: 2}: {},
				&AbsPoint{File: 4, Rank: 1}: {},
			},
			AbsPoint{File: 4, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 7}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 7}: {},
			},
			AbsPoint{File: 1, Rank: 4}: AbsPointSet{
				&AbsPoint{File: 0, Rank: 2}: {},
				&AbsPoint{File: 2, Rank: 2}: {},
				&AbsPoint{File: 3, Rank: 3}: {},
				&AbsPoint{File: 3, Rank: 5}: {},
				&AbsPoint{File: 2, Rank: 6}: {},
				&AbsPoint{File: 0, Rank: 6}: {},
			},
			AbsPoint{File: 7, Rank: 2}: AbsPointSet{
				&AbsPoint{File: 6, Rank: 0}: {},
				&AbsPoint{File: 5, Rank: 1}: {},
				&AbsPoint{File: 5, Rank: 3}: {},
				&AbsPoint{File: 6, Rank: 4}: {},
			},
		},
	},
	{
		Name:   "Adjacent Kings And Queens",
		Active: White,
		Position: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 1, Rank: 1},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 1, Rank: 1}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 2},
				Piece: &Piece{
					Kind:        Queen,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 2, Rank: 2}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 3, Rank: 3},
				Piece: &Piece{
					Kind:        Queen,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 3, Rank: 3}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 4, Rank: 4},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 4, Rank: 4}),
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			AbsPoint{File: 1, Rank: 1}: AbsPointSet{
				&AbsPoint{File: 0, Rank: 0}: {},
				&AbsPoint{File: 1, Rank: 0}: {},
				&AbsPoint{File: 2, Rank: 0}: {},
				&AbsPoint{File: 2, Rank: 1}: {},
				// White Queen at 2-2
				&AbsPoint{File: 1, Rank: 2}: {},
				&AbsPoint{File: 0, Rank: 2}: {},
				&AbsPoint{File: 0, Rank: 1}: {},
			},
			// the Queens can only take each other because any other move would leave their King in check
			AbsPoint{File: 2, Rank: 2}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 3}: {},
			},
			AbsPoint{File: 3, Rank: 3}: AbsPointSet{
				&AbsPoint{File: 2, Rank: 2}: {},
			},
			AbsPoint{File: 4, Rank: 4}: AbsPointSet{
				// Black Queen at 3-3
				&AbsPoint{File: 4, Rank: 3}: {},
				&AbsPoint{File: 5, Rank: 3}: {},
				&AbsPoint{File: 5, Rank: 4}: {},
				&AbsPoint{File: 5, Rank: 5}: {},
				&AbsPoint{File: 4, Rank: 5}: {},
				&AbsPoint{File: 3, Rank: 5}: {},
				&AbsPoint{File: 3, Rank: 4}: {},
			},
		},
	},
	{
		Name:   "Corner Kings And Two Pawns",
		Active: White,
		Position: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 0},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
					Previous:    0,
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 3, Rank: 2},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 3, Rank: 2}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 7, Rank: 7},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
					Previous:    63,
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 5, Rank: 4},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 5, Rank: 4}),
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			AbsPoint{File: 0, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 1, Rank: 0}: {},
				&AbsPoint{File: 1, Rank: 1}: {},
				&AbsPoint{File: 0, Rank: 1}: {},
			},
			AbsPoint{File: 3, Rank: 2}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 3}: {},
			},
			AbsPoint{File: 7, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 6, Rank: 7}: {},
				&AbsPoint{File: 6, Rank: 6}: {},
				&AbsPoint{File: 7, Rank: 6}: {},
			},
			AbsPoint{File: 5, Rank: 4}: AbsPointSet{
				&AbsPoint{File: 5, Rank: 3}: {},
			},
		},
	},
}
