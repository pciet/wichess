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
			{
				AbsPoint: AbsPoint{File: 2, Rank: 0},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 3, Rank: 0},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 7, Rank: 0},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: White,
					Moved:       true,
				},
			},
			WhitePawn0Start,
			{
				AbsPoint: AbsPoint{File: 3, Rank: 1},
				Piece: &Piece{
					Kind:        Knight,
					Orientation: White,
					Moved:       true,
				},
			},
			WhitePawn5Start,
			WhitePawn7Start,
			{
				AbsPoint: AbsPoint{File: 0, Rank: 2},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: Black,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 2, Rank: 2},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 4, Rank: 2},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 5, Rank: 2},
				Piece: &Piece{
					Kind:        Queen,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 5, Rank: 3},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 3, Rank: 4},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 5, Rank: 4},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: Black,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 2, Rank: 5},
				Piece: &Piece{
					Kind:        Knight,
					Orientation: Black,
					Moved:       true,
				},
			},
			BlackPawn0Start,
			BlackPawn1Start,
			BlackPawn2Start,
			BlackPawn6Start,
			BlackPawn7Start,
			{
				AbsPoint: AbsPoint{File: 2, Rank: 7},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 4, Rank: 7},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: Black,
					Moved:       true,
				},
			},
			BlackRightRookStart,
		},
		PreviousFrom: AbsPoint{3, 3},
		PreviousTo:   AbsPoint{3, 2},
		Moves:        map[AbsPoint]AbsPointSet{},
	},
	{
		Name:      "Cornered Queen Mate",
		Active:    White,
		Check:     true,
		Checkmate: true,
		Position: PointSet{
			{
				AbsPoint: AbsPoint{File: 7, Rank: 0},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 0, Rank: 0},
				Piece: &Piece{
					Kind:        Queen,
					Orientation: Black,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 6, Rank: 2},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
				},
			},
		},
		PreviousFrom: AbsPoint{3, 3},
		PreviousTo:   AbsPoint{3, 4},
		Moves:        map[AbsPoint]AbsPointSet{},
	},
	{
		Name:      "Right Triangle Queen Mate",
		Active:    White,
		Check:     true,
		Checkmate: true,
		Position: PointSet{
			{
				AbsPoint: AbsPoint{File: 7, Rank: 5},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 7, Rank: 0},
				Piece: &Piece{
					Kind:        Queen,
					Orientation: Black,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 5, Rank: 5},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
				},
			},
		},
		PreviousFrom: AbsPoint{3, 3},
		PreviousTo:   AbsPoint{3, 4},
		Moves:        map[AbsPoint]AbsPointSet{},
	},
	{
		Name:      "Support Queen Mate",
		Active:    Black,
		Check:     true,
		Checkmate: true,
		Position: PointSet{
			{
				AbsPoint: AbsPoint{File: 1, Rank: 5},
				Piece: &Piece{
					Kind:        Queen,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 2, Rank: 5},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 0, Rank: 5},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
				},
			},
		},
		PreviousFrom: AbsPoint{3, 3},
		PreviousTo:   AbsPoint{3, 4},
		Moves:        map[AbsPoint]AbsPointSet{},
	},
	{
		Name:      "Box Mate",
		Active:    Black,
		Check:     true,
		Checkmate: true,
		Position: PointSet{
			{
				AbsPoint: AbsPoint{File: 0, Rank: 7},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 3, Rank: 5},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 3, Rank: 7},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
				},
			},
		},
		PreviousFrom: AbsPoint{3, 3},
		PreviousTo:   AbsPoint{3, 4},
		Moves:        map[AbsPoint]AbsPointSet{},
	},
	{
		Name:      "Anderssen's Mate",
		Active:    Black,
		Check:     true,
		Checkmate: true,
		Position: PointSet{
			{
				AbsPoint: AbsPoint{File: 5, Rank: 5},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 6, Rank: 6},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 7, Rank: 7},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 6, Rank: 7},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
				},
			},
		},
		PreviousFrom: AbsPoint{3, 3},
		PreviousTo:   AbsPoint{3, 4},
		Moves:        map[AbsPoint]AbsPointSet{},
	},
	{
		Name:      "Anastasia's Mate",
		Active:    Black,
		Check:     true,
		Checkmate: true,
		Position: PointSet{
			WhiteKingStart,
			{
				AbsPoint: AbsPoint{File: 4, Rank: 6},
				Piece: &Piece{
					Kind:        Knight,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 7, Rank: 4},
				Piece: &Piece{
					Kind:        Rook,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 7, Rank: 6},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
				},
			},
			BlackPawn6Start,
		},
		PreviousFrom: AbsPoint{3, 3},
		PreviousTo:   AbsPoint{3, 4},
		Moves:        map[AbsPoint]AbsPointSet{},
	},
	{
		Name:      "Scholar's Mate",
		Active:    Black,
		Check:     true,
		Checkmate: true,
		Position: PointSet{
			WhiteLeftRookStart,
			WhiteLeftKnightStart,
			WhiteLeftBishopStart,
			{
				AbsPoint: AbsPoint{File: 5, Rank: 6},
				Piece: &Piece{
					Kind:        Queen,
					Orientation: White,
					Moved:       true,
				},
			},
			WhiteKingStart,
			{
				AbsPoint: AbsPoint{File: 2, Rank: 3},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: White,
					Moved:       true,
				},
			},
			WhiteRightKnightStart,
			WhiteRightRookStart,
			WhitePawn0Start,
			WhitePawn1Start,
			WhitePawn2Start,
			WhitePawn3Start,
			{
				AbsPoint: AbsPoint{File: 4, Rank: 3},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
				},
			},
			WhitePawn5Start,
			WhitePawn6Start,
			WhitePawn7Start,
			BlackLeftRookStart,
			{
				AbsPoint: AbsPoint{File: 2, Rank: 5},
				Piece: &Piece{
					Kind:        Knight,
					Orientation: Black,
					Moved:       true,
				},
			},
			BlackLeftBishopStart,
			BlackQueenStart,
			BlackKingStart,
			BlackRightBishopStart,
			{
				AbsPoint: AbsPoint{File: 5, Rank: 5},
				Piece: &Piece{
					Kind:        Knight,
					Orientation: Black,
					Moved:       true,
				},
			},
			BlackRightRookStart,
			BlackPawn0Start,
			BlackPawn1Start,
			BlackPawn2Start,
			BlackPawn3Start,
			{
				AbsPoint: AbsPoint{File: 4, Rank: 4},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
				},
			},
			BlackPawn6Start,
			BlackPawn7Start,
		},
		PreviousFrom: AbsPoint{3, 3},
		PreviousTo:   AbsPoint{3, 4},
		Moves:        map[AbsPoint]AbsPointSet{},
	},
	{
		Name:      "Supposedly Published 1625 Gioachino Greco Vs. NN (Anonymous) Fast Checkmate",
		Active:    Black,
		Check:     true,
		Checkmate: true,
		Position: PointSet{
			WhiteLeftRookStart,
			WhiteLeftKnightStart,
			WhiteLeftBishopStart,
			WhiteKingStart,
			{
				AbsPoint: AbsPoint{File: 6, Rank: 5},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: White,
					Moved:       true,
				},
			},
			WhiteRightKnightStart,
			WhiteRightRookStart,
			WhitePawn0Start,
			WhitePawn1Start,
			WhitePawn2Start,
			{
				AbsPoint: AbsPoint{File: 3, Rank: 3},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
				},
			},
			WhitePawn5Start,
			{
				AbsPoint: AbsPoint{File: 7, Rank: 6},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
				},
			},
			WhitePawn7Start,
			BlackLeftRookStart,
			BlackLeftKnightStart,
			{
				AbsPoint: AbsPoint{File: 6, Rank: 1},
				Piece: &Piece{
					Kind:        Bishop,
					Orientation: Black,
					Moved:       true,
				},
			},
			BlackQueenStart,
			BlackKingStart,
			BlackRightBishopStart,
			{
				AbsPoint: AbsPoint{File: 7, Rank: 4},
				Piece: &Piece{
					Kind:        Knight,
					Orientation: Black,
					Moved:       true,
				},
			},
			BlackRightRookStart,
			BlackPawn0Start,
			{
				AbsPoint: AbsPoint{File: 1, Rank: 5},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
				},
			},
			BlackPawn2Start,
			BlackPawn3Start,
			BlackPawn4Start,
		},
		PreviousFrom: AbsPoint{5, 3},
		PreviousTo:   AbsPoint{3, 4},
		Moves:        map[AbsPoint]AbsPointSet{},
	},
	{
		Name:      "Fool's Mate",
		Active:    White,
		Check:     true,
		Checkmate: true,
		Position: PointSet{
			WhiteLeftRookStart,
			WhiteLeftKnightStart,
			WhiteLeftBishopStart,
			WhiteQueenStart,
			WhiteKingStart,
			WhiteRightBishopStart,
			WhiteRightKnightStart,
			WhiteRightRookStart,
			WhitePawn0Start,
			WhitePawn1Start,
			WhitePawn2Start,
			WhitePawn3Start,
			WhitePawn4Start,
			{
				AbsPoint: AbsPoint{File: 5, Rank: 2},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 6, Rank: 3},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
				},
			},
			WhitePawn7Start,
			BlackLeftRookStart,
			BlackLeftKnightStart,
			BlackLeftBishopStart,
			{
				AbsPoint: AbsPoint{File: 7, Rank: 3},
				Piece: &Piece{
					Kind:        Queen,
					Orientation: Black,
					Moved:       true,
				},
			},
			BlackKingStart,
			BlackRightBishopStart,
			BlackRightKnightStart,
			BlackRightRookStart,
			BlackPawn0Start,
			BlackPawn1Start,
			BlackPawn2Start,
			BlackPawn3Start,
			{
				AbsPoint: AbsPoint{File: 4, Rank: 4},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
				},
			},
			BlackPawn5Start,
			BlackPawn6Start,
			BlackPawn7Start,
		},
		PreviousFrom: AbsPoint{3, 3},
		PreviousTo:   AbsPoint{3, 4},
		Moves:        map[AbsPoint]AbsPointSet{},
	},
	{
		Name:      "Black King And Black Guard Bishop Checkmate On Corner White King",
		Active:    White,
		Check:     true,
		Checkmate: true,
		Position: PointSet{
			{
				AbsPoint: AbsPoint{File: 0, Rank: 0},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 2, Rank: 2},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 0, Rank: 2},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{File: 1, Rank: 1},
				Piece: &Piece{
					Kind:        GuardBishop,
					Orientation: Black,
					Moved:       true,
				},
			},
		},
		PreviousFrom: AbsPoint{3, 3},
		PreviousTo:   AbsPoint{3, 4},
		Moves:        map[AbsPoint]AbsPointSet{},
	},
}
