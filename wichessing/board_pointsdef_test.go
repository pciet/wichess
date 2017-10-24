// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

var (
	WhiteKingStart = Point{
		AbsPoint: AbsPoint{File: 4, Rank: 0},
		Piece: &Piece{
			Kind:        King,
			Orientation: White,
			Moved:       false,
			Previous:    4,
		},
	}
	BlackKingStart = Point{
		AbsPoint: AbsPoint{File: 4, Rank: 7},
		Piece: &Piece{
			Kind:        King,
			Orientation: Black,
			Moved:       false,
			Previous:    AbsPointToIndex(AbsPoint{File: 4, Rank: 7}),
		},
	}
	WhiteQueenStart = Point{
		AbsPoint: AbsPoint{File: 3, Rank: 0},
		Piece: &Piece{
			Kind:        Queen,
			Orientation: White,
			Moved:       false,
			Previous:    3,
		},
	}
	BlackQueenStart = Point{
		AbsPoint: AbsPoint{File: 3, Rank: 7},
		Piece: &Piece{
			Kind:        Queen,
			Orientation: Black,
			Moved:       false,
			Previous:    AbsPointToIndex(AbsPoint{File: 3, Rank: 7}),
		},
	}
	WhiteLeftRookStart = Point{
		AbsPoint: AbsPoint{File: 0, Rank: 0},
		Piece: &Piece{
			Kind:        Rook,
			Orientation: White,
			Moved:       false,
			Previous:    0,
		},
	}
	WhiteRightRookStart = Point{
		AbsPoint: AbsPoint{File: 7, Rank: 0},
		Piece: &Piece{
			Kind:        Rook,
			Orientation: White,
			Moved:       false,
			Previous:    7,
		},
	}
	BlackLeftRookStart = Point{
		AbsPoint: AbsPoint{File: 0, Rank: 7},
		Piece: &Piece{
			Kind:        Rook,
			Orientation: Black,
			Moved:       false,
			Previous:    AbsPointToIndex(AbsPoint{File: 0, Rank: 7}),
		},
	}
	BlackRightRookStart = Point{
		AbsPoint: AbsPoint{File: 7, Rank: 7},
		Piece: &Piece{
			Kind:        Rook,
			Orientation: Black,
			Moved:       false,
			Previous:    63,
		},
	}
	WhiteLeftKnightStart = Point{
		AbsPoint: AbsPoint{File: 1, Rank: 0},
		Piece: &Piece{
			Kind:        Knight,
			Orientation: White,
			Moved:       false,
			Previous:    1,
		},
	}
	WhiteRightKnightStart = Point{
		AbsPoint: AbsPoint{File: 6, Rank: 0},
		Piece: &Piece{
			Kind:        Knight,
			Orientation: White,
			Moved:       false,
			Previous:    6,
		},
	}
	BlackLeftKnightStart = Point{
		AbsPoint: AbsPoint{File: 1, Rank: 7},
		Piece: &Piece{
			Kind:        Knight,
			Orientation: Black,
			Moved:       false,
			Previous:    AbsPointToIndex(AbsPoint{File: 1, Rank: 7}),
		},
	}
	BlackRightKnightStart = Point{
		AbsPoint: AbsPoint{File: 6, Rank: 7},
		Piece: &Piece{
			Kind:        Knight,
			Orientation: Black,
			Moved:       false,
			Previous:    AbsPointToIndex(AbsPoint{File: 6, Rank: 7}),
		},
	}
	WhiteLeftBishopStart = Point{
		AbsPoint: AbsPoint{File: 2, Rank: 0},
		Piece: &Piece{
			Kind:        Bishop,
			Orientation: White,
			Moved:       false,
			Previous:    2,
		},
	}
	WhiteRightBishopStart = Point{
		AbsPoint: AbsPoint{File: 5, Rank: 0},
		Piece: &Piece{
			Kind:        Bishop,
			Orientation: White,
			Moved:       false,
			Previous:    5,
		},
	}
	BlackLeftBishopStart = Point{
		AbsPoint: AbsPoint{File: 2, Rank: 7},
		Piece: &Piece{
			Kind:        Bishop,
			Orientation: Black,
			Moved:       false,
			Previous:    AbsPointToIndex(AbsPoint{File: 2, Rank: 7}),
		},
	}
	BlackRightBishopStart = Point{
		AbsPoint: AbsPoint{File: 5, Rank: 7},
		Piece: &Piece{
			Kind:        Bishop,
			Orientation: Black,
			Moved:       false,
			Previous:    AbsPointToIndex(AbsPoint{File: 5, Rank: 7}),
		},
	}
)
