// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

var DrawMovesCases = []AvailableMovesCase{
	{
		Name:   "Stalemate Not Caught 2",
		Active: White,
		Draw:   true,
		Position: PointSet{
			{
				AbsPoint: AbsPoint{5, 0},
				Piece: &Piece{
					Kind:        Queen,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPoint{5, 0}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{2, 1},
				Piece: &Piece{
					Kind:        Queen,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPoint{2, 1}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{4, 2},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPoint{4, 4}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{1, 3},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPoint{1, 3}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{2, 4},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPoint{2, 4}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{0, 5},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPoint{0, 5}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{3, 5},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPoint{3, 5}.Index(),
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
				{4, 1}: {},
				{3, 2}: {},
				{2, 3}: {},
				{1, 4}: {},
				{5, 1}: {},
				{5, 2}: {},
				{5, 3}: {},
				{5, 4}: {},
				{5, 5}: {},
				{5, 6}: {},
				{5, 7}: {},
				{6, 1}: {},
				{7, 2}: {},
				{6, 0}: {},
				{7, 0}: {},
			},
			{2, 1}: {
				{1, 0}: {},
				{2, 0}: {},
				{3, 0}: {},
				{1, 1}: {},
				{0, 1}: {},
				{3, 1}: {},
				{4, 1}: {},
				{5, 1}: {},
				{6, 1}: {},
				{7, 1}: {},
				{1, 2}: {},
				{0, 3}: {},
				{2, 2}: {},
				{2, 3}: {},
				{2, 4}: {},
				{3, 2}: {},
				{4, 3}: {},
				{5, 4}: {},
				{6, 5}: {},
				{7, 6}: {},
			},
			{4, 2}: {
				{4, 1}: {},
				{4, 0}: {},
				{3, 2}: {},
				{2, 2}: {},
				{1, 2}: {},
				{0, 2}: {},
				{4, 3}: {},
				{4, 4}: {},
				{4, 5}: {},
				{4, 6}: {},
				{4, 7}: {},
				{5, 2}: {},
				{6, 2}: {},
				{7, 2}: {},
			},
			{3, 5}: {
				{2, 4}: {},
				{2, 6}: {},
				{1, 7}: {},
				{4, 6}: {},
				{5, 7}: {},
				{4, 4}: {},
				{5, 3}: {},
				{6, 2}: {},
				{7, 1}: {},
			},
			{0, 5}: {
				{1, 5}: {},
				{1, 6}: {},
				{0, 6}: {},
			},
		},
	},
	{
		Name:   "Stalemate Not Caught",
		Active: Black,
		Draw:   true,
		Position: PointSet{
			&WhitePawn0Start: {},
			&WhitePawn2Start: {},
			{
				AbsPoint: AbsPoint{2, 2},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPoint{2, 2}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{2, 3},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPoint{2, 3}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{1, 4},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPoint{1, 4}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{3, 4},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPoint{3, 4}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{6, 4},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPoint{6, 4}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{3, 5},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPoint{3, 5}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{4, 5},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPoint{4, 5}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{2, 6},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPoint{2, 6}.Index(),
				},
			}: {},
			{
				AbsPoint: AbsPoint{6, 7},
				Piece: &Piece{
					Kind:        Queen,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPoint{6, 7}.Index(),
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			{0, 1}: {
				{0, 2}: {},
				{0, 3}: {},
			},
			{1, 4}: {
				{1, 3}: {},
				{1, 2}: {},
				{1, 1}: {},
				{1, 0}: {},
				{0, 4}: {},
				{1, 5}: {},
				{1, 6}: {},
				{1, 7}: {},
				{2, 4}: {},
			},
			{6, 4}: {
				{5, 3}: {},
				{4, 2}: {},
				{3, 1}: {},
				{2, 0}: {},
				{7, 3}: {},
				{5, 5}: {},
				{4, 6}: {},
				{3, 7}: {},
				{7, 5}: {},
			},
			{4, 5}: {
				{5, 4}: {},
				{5, 5}: {},
				{5, 6}: {},
				{4, 6}: {},
			},
			{6, 7}: {
				{6, 6}: {},
				{6, 5}: {},
				{7, 6}: {},
				{7, 7}: {},
				{5, 7}: {},
				{4, 7}: {},
				{3, 7}: {},
				{2, 7}: {},
				{1, 7}: {},
				{0, 7}: {},
				{5, 6}: {},
			},
		},
	},
	{
		Name:   "False Insufficient Material Draw - Bishops Of Different Color, Different Rows 2",
		Active: Black,
		Draw:   false,
		Position: PointSet{
			&WhiteKingStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 0},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: White,
					Moved:       true,
					Previous:    0,
				},
			}: {},
			&BlackKingStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 5, Rank: 2},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: Black,
					Moved:       true,
					Previous:    7,
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			AbsPoint{File: 4, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 1}: {},
				&AbsPoint{File: 5, Rank: 1}: {},
				&AbsPoint{File: 5, Rank: 0}: {},
			},
			AbsPoint{File: 0, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 1, Rank: 1}: {},
				&AbsPoint{File: 2, Rank: 2}: {},
				&AbsPoint{File: 3, Rank: 3}: {},
				&AbsPoint{File: 4, Rank: 4}: {},
				&AbsPoint{File: 5, Rank: 5}: {},
				&AbsPoint{File: 6, Rank: 6}: {},
				&AbsPoint{File: 7, Rank: 7}: {},
			},
			AbsPoint{File: 4, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 7}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 7}: {},
			},
			AbsPoint{File: 5, Rank: 2}: AbsPointSet{
				&AbsPoint{File: 6, Rank: 1}: {},
				&AbsPoint{File: 7, Rank: 0}: {},
				&AbsPoint{File: 4, Rank: 3}: {},
				&AbsPoint{File: 3, Rank: 4}: {},
				&AbsPoint{File: 2, Rank: 5}: {},
				&AbsPoint{File: 1, Rank: 6}: {},
				&AbsPoint{File: 0, Rank: 7}: {},
				&AbsPoint{File: 6, Rank: 3}: {},
				&AbsPoint{File: 7, Rank: 4}: {},
				&AbsPoint{File: 4, Rank: 1}: {},
				&AbsPoint{File: 3, Rank: 0}: {},
			},
		},
	},
	{
		Name:   "False Insufficient Material Draw - Bishops Of Different Color, Different Rows 1",
		Active: Black,
		Draw:   false,
		Position: PointSet{
			&WhiteKingStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 0},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: White,
					Moved:       true,
					Previous:    0,
				},
			}: {},
			&BlackKingStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 6, Rank: 1},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 6, Rank: 1}),
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			AbsPoint{File: 4, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 0}: {},
				&AbsPoint{File: 3, Rank: 1}: {},
				&AbsPoint{File: 4, Rank: 1}: {},
				&AbsPoint{File: 5, Rank: 1}: {},
			},
			AbsPoint{File: 0, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 1, Rank: 1}: {},
				&AbsPoint{File: 2, Rank: 2}: {},
				&AbsPoint{File: 3, Rank: 3}: {},
				&AbsPoint{File: 4, Rank: 4}: {},
				&AbsPoint{File: 5, Rank: 5}: {},
				&AbsPoint{File: 6, Rank: 6}: {},
				&AbsPoint{File: 7, Rank: 7}: {},
			},
			AbsPoint{File: 4, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 7}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 7}: {},
			},
			AbsPoint{File: 6, Rank: 1}: AbsPointSet{
				&AbsPoint{File: 7, Rank: 0}: {},
				&AbsPoint{File: 5, Rank: 2}: {},
				&AbsPoint{File: 4, Rank: 3}: {},
				&AbsPoint{File: 3, Rank: 4}: {},
				&AbsPoint{File: 2, Rank: 5}: {},
				&AbsPoint{File: 1, Rank: 6}: {},
				&AbsPoint{File: 0, Rank: 7}: {},
				&AbsPoint{File: 7, Rank: 2}: {},
				&AbsPoint{File: 5, Rank: 0}: {},
			},
		},
	},
	{
		Name:   "False Insufficient Material Draw - Bishops Of Different Color",
		Active: Black,
		Draw:   false,
		Position: PointSet{
			&WhiteKingStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 0},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: White,
					Moved:       true,
					Previous:    0,
				},
			}: {},
			&BlackKingStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 7, Rank: 0},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: Black,
					Moved:       true,
					Previous:    7,
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
			AbsPoint{File: 0, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 1, Rank: 1}: {},
				&AbsPoint{File: 2, Rank: 2}: {},
				&AbsPoint{File: 3, Rank: 3}: {},
				&AbsPoint{File: 4, Rank: 4}: {},
				&AbsPoint{File: 5, Rank: 5}: {},
				&AbsPoint{File: 6, Rank: 6}: {},
				&AbsPoint{File: 7, Rank: 7}: {},
			},
			AbsPoint{File: 4, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 7}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 7}: {},
			},
			AbsPoint{File: 7, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 6, Rank: 1}: {},
				&AbsPoint{File: 5, Rank: 2}: {},
				&AbsPoint{File: 4, Rank: 3}: {},
				&AbsPoint{File: 3, Rank: 4}: {},
				&AbsPoint{File: 2, Rank: 5}: {},
				&AbsPoint{File: 1, Rank: 6}: {},
				&AbsPoint{File: 0, Rank: 7}: {},
			},
		},
	},
	{
		Name:   "Insufficient Material Draw - King+Bishop Vs King+Bishop, Bishops Of Same Color Different Rows 2",
		Active: Black,
		Draw:   true,
		Position: PointSet{
			&WhiteKingStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 0},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: White,
					Moved:       true,
					Previous:    0,
				},
			}: {},
			&BlackKingStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 5, Rank: 5},
				Piece: &Piece{
					Kind:        Bishop,
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
			AbsPoint{File: 0, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 1, Rank: 1}: {},
				&AbsPoint{File: 2, Rank: 2}: {},
				&AbsPoint{File: 3, Rank: 3}: {},
				&AbsPoint{File: 4, Rank: 4}: {},
				&AbsPoint{File: 5, Rank: 5}: {},
			},
			AbsPoint{File: 4, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 7}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 7}: {},
			},
			AbsPoint{File: 5, Rank: 5}: AbsPointSet{
				&AbsPoint{File: 7, Rank: 7}: {},
				&AbsPoint{File: 6, Rank: 6}: {},
				&AbsPoint{File: 4, Rank: 4}: {},
				&AbsPoint{File: 3, Rank: 3}: {},
				&AbsPoint{File: 2, Rank: 2}: {},
				&AbsPoint{File: 1, Rank: 1}: {},
				&AbsPoint{File: 0, Rank: 0}: {},
				&AbsPoint{File: 6, Rank: 4}: {},
				&AbsPoint{File: 7, Rank: 3}: {},
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 3, Rank: 7}: {},
			},
		},
	},
	{
		Name:   "Insufficient Material Draw - King+Bishop Vs King+Bishop, Bishops Of Same Color Different Rows 1",
		Active: Black,
		Draw:   true,
		Position: PointSet{
			&WhiteKingStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 0},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: White,
					Moved:       true,
					Previous:    0,
				},
			}: {},
			&BlackKingStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 6, Rank: 6},
				Piece: &Piece{
					Kind:        Bishop,
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
			AbsPoint{File: 0, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 1, Rank: 1}: {},
				&AbsPoint{File: 2, Rank: 2}: {},
				&AbsPoint{File: 3, Rank: 3}: {},
				&AbsPoint{File: 4, Rank: 4}: {},
				&AbsPoint{File: 5, Rank: 5}: {},
				&AbsPoint{File: 6, Rank: 6}: {},
			},
			AbsPoint{File: 4, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 7}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 7}: {},
			},
			AbsPoint{File: 6, Rank: 6}: AbsPointSet{
				&AbsPoint{File: 7, Rank: 7}: {},
				&AbsPoint{File: 5, Rank: 5}: {},
				&AbsPoint{File: 4, Rank: 4}: {},
				&AbsPoint{File: 3, Rank: 3}: {},
				&AbsPoint{File: 2, Rank: 2}: {},
				&AbsPoint{File: 1, Rank: 1}: {},
				&AbsPoint{File: 0, Rank: 0}: {},
				&AbsPoint{File: 5, Rank: 7}: {},
				&AbsPoint{File: 7, Rank: 5}: {},
			},
		},
	},
	{
		Name:   "Insufficient Material Draw - King+Bishop Vs King+Bishop, Bishops Of Same Color",
		Active: White,
		Draw:   true,
		Position: PointSet{
			&WhiteKingStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 0},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: White,
					Moved:       true,
					Previous:    0,
				},
			}: {},
			&BlackKingStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 6, Rank: 0},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 5, Rank: 1}),
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
			AbsPoint{File: 0, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 1, Rank: 1}: {},
				&AbsPoint{File: 2, Rank: 2}: {},
				&AbsPoint{File: 3, Rank: 3}: {},
				&AbsPoint{File: 4, Rank: 4}: {},
				&AbsPoint{File: 5, Rank: 5}: {},
				&AbsPoint{File: 6, Rank: 6}: {},
				&AbsPoint{File: 7, Rank: 7}: {},
			},
			AbsPoint{File: 4, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 3, Rank: 7}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 7}: {},
			},
			AbsPoint{File: 6, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 7, Rank: 1}: {},
				&AbsPoint{File: 5, Rank: 1}: {},
				&AbsPoint{File: 4, Rank: 2}: {},
				&AbsPoint{File: 3, Rank: 3}: {},
				&AbsPoint{File: 2, Rank: 4}: {},
				&AbsPoint{File: 1, Rank: 5}: {},
				&AbsPoint{File: 0, Rank: 6}: {},
			},
		},
	},
	{
		Name:   "Insufficient Material Draw - King And Knight Vs King",
		Active: White,
		Draw:   true,
		Position: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 1, Rank: 2},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
					Previous:    1,
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 0},
				Piece: &Piece{
					Kind:        Knight,
					Orientation: White,
					Moved:       true,
					Previous:    0,
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 6, Rank: 6},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 6, Rank: 6}),
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			AbsPoint{File: 1, Rank: 2}: AbsPointSet{
				&AbsPoint{File: 0, Rank: 1}: {},
				&AbsPoint{File: 0, Rank: 2}: {},
				&AbsPoint{File: 0, Rank: 3}: {},
				&AbsPoint{File: 1, Rank: 3}: {},
				&AbsPoint{File: 2, Rank: 3}: {},
				&AbsPoint{File: 2, Rank: 2}: {},
				&AbsPoint{File: 2, Rank: 1}: {},
				&AbsPoint{File: 1, Rank: 1}: {},
			},
			AbsPoint{File: 0, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 2, Rank: 1}: {},
			},
			AbsPoint{File: 6, Rank: 6}: AbsPointSet{
				&AbsPoint{File: 5, Rank: 5}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 7}: {},
				&AbsPoint{File: 6, Rank: 7}: {},
				&AbsPoint{File: 7, Rank: 7}: {},
				&AbsPoint{File: 7, Rank: 6}: {},
				&AbsPoint{File: 7, Rank: 5}: {},
				&AbsPoint{File: 6, Rank: 5}: {},
			},
		},
	},
	{
		Name:   "Insufficient Material Draw - King And Bishop Vs King",
		Active: Black,
		Draw:   true,
		Position: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 1, Rank: 1},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
					Previous:    0,
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
				AbsPoint: AbsPoint{File: 7, Rank: 0},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: Black,
					Moved:       true,
					Previous:    7,
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			AbsPoint{File: 1, Rank: 1}: AbsPointSet{
				&AbsPoint{File: 0, Rank: 0}: {},
				&AbsPoint{File: 0, Rank: 1}: {},
				&AbsPoint{File: 0, Rank: 2}: {},
				&AbsPoint{File: 1, Rank: 2}: {},
				&AbsPoint{File: 2, Rank: 2}: {},
				&AbsPoint{File: 2, Rank: 1}: {},
				&AbsPoint{File: 2, Rank: 0}: {},
				&AbsPoint{File: 1, Rank: 0}: {},
			},
			AbsPoint{File: 7, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 6, Rank: 7}: {},
				&AbsPoint{File: 6, Rank: 6}: {},
				&AbsPoint{File: 7, Rank: 6}: {},
			},
			AbsPoint{File: 7, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 6, Rank: 1}: {},
				&AbsPoint{File: 5, Rank: 2}: {},
				&AbsPoint{File: 4, Rank: 3}: {},
				&AbsPoint{File: 3, Rank: 4}: {},
				&AbsPoint{File: 2, Rank: 5}: {},
				&AbsPoint{File: 1, Rank: 6}: {},
				&AbsPoint{File: 0, Rank: 7}: {},
			},
		},
	},
	{
		Name:   "Insufficient Material Draw - King Vs King",
		Active: White,
		Draw:   true,
		Position: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 0},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
					Previous:    1,
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
		},
		Moves: map[AbsPoint]AbsPointSet{
			AbsPoint{File: 0, Rank: 0}: AbsPointSet{
				&AbsPoint{File: 0, Rank: 1}: {},
				&AbsPoint{File: 1, Rank: 1}: {},
				&AbsPoint{File: 1, Rank: 0}: {},
			},
			AbsPoint{File: 7, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 6, Rank: 7}: {},
				&AbsPoint{File: 6, Rank: 6}: {},
				&AbsPoint{File: 7, Rank: 6}: {},
			},
		},
	},
	{
		Name:   "Silman Two Bishop Stalemate",
		Active: Black,
		Draw:   true,
		Position: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 1, Rank: 5},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 2, Rank: 5}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 6},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 2, Rank: 6}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 3},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 2, Rank: 3}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 7},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 0, Rank: 7}),
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			AbsPoint{File: 1, Rank: 5}: AbsPointSet{
				&AbsPoint{File: 0, Rank: 5}: {},
				&AbsPoint{File: 0, Rank: 4}: {},
				&AbsPoint{File: 1, Rank: 4}: {},
				&AbsPoint{File: 2, Rank: 4}: {},
				&AbsPoint{File: 2, Rank: 5}: {},
			},
			AbsPoint{File: 2, Rank: 6}: AbsPointSet{
				&AbsPoint{File: 1, Rank: 7}: {},
				&AbsPoint{File: 3, Rank: 7}: {},
				&AbsPoint{File: 3, Rank: 5}: {},
				&AbsPoint{File: 4, Rank: 4}: {},
				&AbsPoint{File: 5, Rank: 3}: {},
				&AbsPoint{File: 6, Rank: 2}: {},
				&AbsPoint{File: 7, Rank: 1}: {},
			},
			AbsPoint{File: 2, Rank: 3}: AbsPointSet{
				&AbsPoint{File: 1, Rank: 2}: {},
				&AbsPoint{File: 0, Rank: 1}: {},
				&AbsPoint{File: 3, Rank: 2}: {},
				&AbsPoint{File: 4, Rank: 1}: {},
				&AbsPoint{File: 5, Rank: 0}: {},
				&AbsPoint{File: 3, Rank: 4}: {},
				&AbsPoint{File: 4, Rank: 5}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
				&AbsPoint{File: 6, Rank: 7}: {},
				&AbsPoint{File: 1, Rank: 4}: {},
				&AbsPoint{File: 0, Rank: 5}: {},
			},
		},
	},
	{
		Name:   "Rook Stalemate Two",
		Active: White,
		Draw:   true,
		Position: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 7},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 0, Rank: 7}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 7},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 3, Rank: 7}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 7, Rank: 6},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 7, Rank: 6}),
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			AbsPoint{File: 2, Rank: 7}: AbsPointSet{
				&AbsPoint{File: 2, Rank: 6}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 3, Rank: 7}: {},
			},
			AbsPoint{File: 7, Rank: 6}: AbsPointSet{
				&AbsPoint{File: 7, Rank: 7}: {},
				&AbsPoint{File: 7, Rank: 5}: {},
				&AbsPoint{File: 7, Rank: 4}: {},
				&AbsPoint{File: 7, Rank: 3}: {},
				&AbsPoint{File: 7, Rank: 2}: {},
				&AbsPoint{File: 7, Rank: 1}: {},
				&AbsPoint{File: 7, Rank: 0}: {},
				&AbsPoint{File: 6, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 2, Rank: 6}: {},
				&AbsPoint{File: 1, Rank: 6}: {},
				&AbsPoint{File: 0, Rank: 6}: {},
			},
		},
	},
	{
		Name:   "Rook Stalemate One",
		Active: Black,
		Draw:   true,
		Position: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 5},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 3, Rank: 4}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 1, Rank: 6},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 1, Rank: 6}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 7},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 0, Rank: 7}),
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			AbsPoint{File: 2, Rank: 5}: AbsPointSet{
				&AbsPoint{File: 2, Rank: 4}: {},
				&AbsPoint{File: 1, Rank: 4}: {},
				&AbsPoint{File: 1, Rank: 5}: {},
				&AbsPoint{File: 2, Rank: 6}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 3, Rank: 5}: {},
				&AbsPoint{File: 3, Rank: 4}: {},
			},
			AbsPoint{File: 1, Rank: 6}: AbsPointSet{
				&AbsPoint{File: 1, Rank: 5}: {},
				&AbsPoint{File: 1, Rank: 4}: {},
				&AbsPoint{File: 1, Rank: 3}: {},
				&AbsPoint{File: 1, Rank: 2}: {},
				&AbsPoint{File: 1, Rank: 1}: {},
				&AbsPoint{File: 1, Rank: 0}: {},
				&AbsPoint{File: 0, Rank: 6}: {},
				&AbsPoint{File: 1, Rank: 7}: {},
				&AbsPoint{File: 2, Rank: 6}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
				&AbsPoint{File: 6, Rank: 6}: {},
				&AbsPoint{File: 7, Rank: 6}: {},
			},
		},
	},
	{
		Name:   "Queen Stalemate",
		Active: White,
		Draw:   true,
		Position: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 5},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 0, Rank: 5}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 4},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 2, Rank: 3}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 6},
				Piece: &Piece{
					Kind:        Queen,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 2, Rank: 3}),
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			AbsPoint{File: 2, Rank: 4}: AbsPointSet{
				&AbsPoint{File: 2, Rank: 3}: {},
				&AbsPoint{File: 3, Rank: 3}: {},
				&AbsPoint{File: 3, Rank: 4}: {},
				&AbsPoint{File: 3, Rank: 5}: {},
				&AbsPoint{File: 2, Rank: 5}: {},
				&AbsPoint{File: 1, Rank: 3}: {},
			},
			AbsPoint{File: 2, Rank: 6}: AbsPointSet{
				&AbsPoint{File: 2, Rank: 7}: {},
				&AbsPoint{File: 3, Rank: 7}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
				&AbsPoint{File: 6, Rank: 6}: {},
				&AbsPoint{File: 7, Rank: 6}: {},
				&AbsPoint{File: 3, Rank: 5}: {},
				&AbsPoint{File: 4, Rank: 4}: {},
				&AbsPoint{File: 5, Rank: 3}: {},
				&AbsPoint{File: 6, Rank: 2}: {},
				&AbsPoint{File: 7, Rank: 1}: {},
				&AbsPoint{File: 2, Rank: 5}: {},
				&AbsPoint{File: 1, Rank: 5}: {},
				&AbsPoint{File: 0, Rank: 4}: {},
				&AbsPoint{File: 1, Rank: 6}: {},
				&AbsPoint{File: 0, Rank: 6}: {},
				&AbsPoint{File: 1, Rank: 7}: {},
			},
		},
	},
	{
		Name:   "Corner Queen Stalemate",
		Active: Black,
		Draw:   true,
		Position: PointSet{
			&WhiteKingStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 6},
				Piece: &Piece{
					Kind:        Queen,
					Orientation: White,
					Moved:       true,
					Previous:    2,
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 7},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 0, Rank: 7}),
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
			AbsPoint{File: 2, Rank: 6}: AbsPointSet{
				&AbsPoint{File: 1, Rank: 6}: {},
				&AbsPoint{File: 0, Rank: 6}: {},
				&AbsPoint{File: 1, Rank: 7}: {},
				&AbsPoint{File: 2, Rank: 7}: {},
				&AbsPoint{File: 3, Rank: 7}: {},
				&AbsPoint{File: 3, Rank: 6}: {},
				&AbsPoint{File: 4, Rank: 6}: {},
				&AbsPoint{File: 5, Rank: 6}: {},
				&AbsPoint{File: 6, Rank: 6}: {},
				&AbsPoint{File: 7, Rank: 6}: {},
				&AbsPoint{File: 3, Rank: 5}: {},
				&AbsPoint{File: 4, Rank: 4}: {},
				&AbsPoint{File: 5, Rank: 3}: {},
				&AbsPoint{File: 6, Rank: 2}: {},
				&AbsPoint{File: 7, Rank: 1}: {},
				&AbsPoint{File: 2, Rank: 5}: {},
				&AbsPoint{File: 2, Rank: 4}: {},
				&AbsPoint{File: 2, Rank: 3}: {},
				&AbsPoint{File: 2, Rank: 2}: {},
				&AbsPoint{File: 2, Rank: 1}: {},
				&AbsPoint{File: 2, Rank: 0}: {},
				&AbsPoint{File: 1, Rank: 5}: {},
				&AbsPoint{File: 0, Rank: 4}: {},
			},
		},
	},
}
