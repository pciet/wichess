// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

var CheckmateMovesCases = []AvailableMovesCase{
	{
		Name:      "Schulder-Boden 1853 Checkmate",
		Active:    White,
		Check:     true,
		Checkmate: true,
		Position: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 0},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
					Previous:    2,
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 3, Rank: 0},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: White,
					Moved:       true,
					Previous:    3,
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
			&WhitePawn0Start: {},
			&Point{
				AbsPoint: AbsPoint{File: 3, Rank: 1},
				Piece: &Piece{
					Kind:        Knight,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 3, Rank: 1}),
				},
			}: {},
			&WhitePawn5Start: {},
			&WhitePawn7Start: {},
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 2},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 0, Rank: 2}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 2},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 2, Rank: 2}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 4, Rank: 2},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 4, Rank: 2}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 5, Rank: 2},
				Piece: &Piece{
					Kind:        Queen,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 5, Rank: 2}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 5, Rank: 3},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 5, Rank: 3}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 3, Rank: 4},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 3, Rank: 4}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 5, Rank: 4},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 5, Rank: 4}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 5},
				Piece: &Piece{
					Kind:        Knight,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 2, Rank: 5}),
				},
			}: {},
			&BlackPawn0Start: {},
			&BlackPawn1Start: {},
			&BlackPawn2Start: {},
			&BlackPawn6Start: {},
			&BlackPawn7Start: {},
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 7},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 2, Rank: 7}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 4, Rank: 7},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 4, Rank: 7}),
				},
			}: {},
			&BlackRightRookStart: {},
		},
		Moves: map[AbsPoint]AbsPointSet{},
	},
	{
		Name:      "Cornered Queen Mate",
		Active:    White,
		Check:     true,
		Checkmate: true,
		Position: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 7, Rank: 0},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 7, Rank: 0}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 0},
				Piece: &Piece{
					Kind:        Queen,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 0, Rank: 1}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 6, Rank: 2},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 6, Rank: 2}),
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{},
	},
	{
		Name:      "Right Triangle Queen Mate",
		Active:    White,
		Check:     true,
		Checkmate: true,
		Position: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 7, Rank: 5},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 7, Rank: 5}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 7, Rank: 0},
				Piece: &Piece{
					Kind:        Queen,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 6, Rank: 0}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 5, Rank: 5},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 5, Rank: 5}),
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{},
	},
	{
		Name:      "Support Queen Mate",
		Active:    Black,
		Check:     true,
		Checkmate: true,
		Position: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 1, Rank: 5},
				Piece: &Piece{
					Kind:        Queen,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 1, Rank: 5}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 5},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 2, Rank: 5}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 5},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 0, Rank: 5}),
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{},
	},
	{
		Name:      "Box Mate",
		Active:    Black,
		Check:     true,
		Checkmate: true,
		Position: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 7},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: White,
					Moved:       true,
					Previous:    0,
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 3, Rank: 5},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 3, Rank: 5}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 3, Rank: 7},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 3, Rank: 7}),
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{},
	},
	{
		Name:      "Anderssen's Mate",
		Active:    Black,
		Check:     true,
		Checkmate: true,
		Position: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 5, Rank: 5},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 5, Rank: 4}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 6, Rank: 6},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 6, Rank: 6}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 7, Rank: 7},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 7, Rank: 7}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 6, Rank: 7},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 6, Rank: 7}),
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{},
	},
	{
		Name:      "Anastasia's Mate",
		Active:    Black,
		Check:     true,
		Checkmate: true,
		Position: PointSet{
			&WhiteKingStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 4, Rank: 6},
				Piece: &Piece{
					Kind:        Knight,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 4, Rank: 6}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 7, Rank: 4},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 6, Rank: 4}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 7, Rank: 6},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 7, Rank: 6}),
				},
			}: {},
			&BlackPawn6Start: {},
		},
		Moves: map[AbsPoint]AbsPointSet{},
	},
	{
		Name:      "Scholar's Mate",
		Active:    Black,
		Check:     true,
		Checkmate: true,
		Position: PointSet{
			&WhiteLeftRookStart:   {},
			&WhiteLeftKnightStart: {},
			&WhiteLeftBishopStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 5, Rank: 6},
				Piece: &Piece{
					Kind:        Queen,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 3, Rank: 4}),
				},
			}: {},
			&WhiteKingStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 3},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 2, Rank: 3}),
				},
			}: {},
			&WhiteRightKnightStart: {},
			&WhiteRightRookStart:   {},
			&WhitePawn0Start:       {},
			&WhitePawn1Start:       {},
			&WhitePawn2Start:       {},
			&WhitePawn3Start:       {},
			&Point{
				AbsPoint: AbsPoint{File: 4, Rank: 3},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 4, Rank: 3}),
				},
			}: {},
			&WhitePawn5Start:    {},
			&WhitePawn6Start:    {},
			&WhitePawn7Start:    {},
			&BlackLeftRookStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 5},
				Piece: &Piece{
					Kind:        Knight,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 2, Rank: 5}),
				},
			}: {},
			&BlackLeftBishopStart:  {},
			&BlackQueenStart:       {},
			&BlackKingStart:        {},
			&BlackRightBishopStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 5, Rank: 5},
				Piece: &Piece{
					Kind:        Knight,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 5, Rank: 5}),
				},
			}: {},
			&BlackRightRookStart: {},
			&BlackPawn0Start:     {},
			&BlackPawn1Start:     {},
			&BlackPawn2Start:     {},
			&BlackPawn3Start:     {},
			&Point{
				AbsPoint: AbsPoint{File: 4, Rank: 4},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 4, Rank: 4}),
				},
			}: {},
			&BlackPawn6Start: {},
			&BlackPawn7Start: {},
		},
		Moves: map[AbsPoint]AbsPointSet{},
	},
	{
		Name:      "Supposedly Published 1625 Gioachino Greco Vs. NN (Anonymous) Fast Checkmate",
		Active:    Black,
		Check:     true,
		Checkmate: true,
		Position: PointSet{
			&WhiteLeftRookStart:   {},
			&WhiteLeftKnightStart: {},
			&WhiteLeftBishopStart: {},
			&WhiteKingStart:       {},
			&Point{
				AbsPoint: AbsPoint{File: 6, Rank: 5},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 3, Rank: 2}),
				},
			}: {},
			&WhiteRightKnightStart: {},
			&WhiteRightRookStart:   {},
			&WhitePawn0Start:       {},
			&WhitePawn1Start:       {},
			&WhitePawn2Start:       {},
			&Point{
				AbsPoint: AbsPoint{File: 3, Rank: 3},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 3, Rank: 3}),
				},
			}: {},
			&WhitePawn5Start: {},
			&Point{
				AbsPoint: AbsPoint{File: 7, Rank: 6},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 7, Rank: 6}),
				},
			}: {},
			&WhitePawn7Start:      {},
			&BlackLeftRookStart:   {},
			&BlackLeftKnightStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 6, Rank: 1},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 6, Rank: 1}),
				},
			}: {},
			&BlackQueenStart:       {},
			&BlackKingStart:        {},
			&BlackRightBishopStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 7, Rank: 4},
				Piece: &Piece{
					Kind:        Knight,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 7, Rank: 4}),
				},
			}: {},
			&BlackRightRookStart: {},
			&BlackPawn0Start:     {},
			&Point{
				AbsPoint: AbsPoint{File: 1, Rank: 5},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 1, Rank: 5}),
				},
			}: {},
			&BlackPawn2Start: {},
			&BlackPawn3Start: {},
			&BlackPawn4Start: {},
		},
		Moves: map[AbsPoint]AbsPointSet{},
	},
	{
		Name:      "Fool's Mate",
		Active:    White,
		Check:     true,
		Checkmate: true,
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
			&Point{
				AbsPoint: AbsPoint{File: 5, Rank: 2},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 5, Rank: 2}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 6, Rank: 3},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 6, Rank: 3}),
				},
			}: {},
			&WhitePawn7Start:      {},
			&BlackLeftRookStart:   {},
			&BlackLeftKnightStart: {},
			&BlackLeftBishopStart: {},
			&Point{
				AbsPoint: AbsPoint{File: 7, Rank: 3},
				Piece: &Piece{
					Kind:        Queen,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 3, Rank: 7}),
				},
			}: {},
			&BlackKingStart:        {},
			&BlackRightBishopStart: {},
			&BlackRightKnightStart: {},
			&BlackRightRookStart:   {},
			&BlackPawn0Start:       {},
			&BlackPawn1Start:       {},
			&BlackPawn2Start:       {},
			&BlackPawn3Start:       {},
			&Point{
				AbsPoint: AbsPoint{File: 4, Rank: 4},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 4, Rank: 4}),
				},
			}: {},
			&BlackPawn5Start: {},
			&BlackPawn6Start: {},
			&BlackPawn7Start: {},
		},
		Moves: map[AbsPoint]AbsPointSet{},
	},
	{
		Name:      "Black King And Black Guard Bishop Checkmate On Corner White King",
		Active:    White,
		Check:     true,
		Checkmate: true,
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
				AbsPoint: AbsPoint{File: 2, Rank: 2},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 2, Rank: 2}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 0, Rank: 2},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 0, Rank: 3}),
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 1, Rank: 1},
				Piece: &Piece{
					Kind:        GuardBishop,
					Orientation: Black,
					Moved:       true,
					Previous:    AbsPointToIndex(AbsPoint{File: 1, Rank: 1}),
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{},
	},
}
