// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

// TODO: pawn interactions, rooks, bishops

var BasicMovesCases = []AvailableMovesCase{
	{
		Name:   "Unmoved Kings And Taking Knights",
		Active: White,
		Position: PointSet{
			&Point{
				AbsPoint: AbsPoint{File: 4, Rank: 0},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       false,
					Previous:    4,
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 2, Rank: 2},
				Piece: &Piece{
					Kind:        Knight,
					Orientation: White,
					Moved:       true,
					Previous:    1,
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 6, Rank: 0},
				Piece: &Piece{
					Kind:        Knight,
					Orientation: White,
					Moved:       false,
					Previous:    6,
				},
			}: {},
			&Point{
				AbsPoint: AbsPoint{File: 4, Rank: 7},
				Piece: &Piece{
					Kind:        King,
					Orientation: Black,
					Moved:       false,
					Previous:    AbsPointToIndex(AbsPoint{File: 4, Rank: 7}),
				},
			}: {},
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
