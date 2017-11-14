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
		},
	}
	BlackKingStart = Point{
		AbsPoint: AbsPoint{File: 4, Rank: 7},
		Piece: &Piece{
			Kind:        King,
			Orientation: Black,
			Moved:       false,
		},
	}
	WhiteQueenStart = Point{
		AbsPoint: AbsPoint{File: 3, Rank: 0},
		Piece: &Piece{
			Kind:        Queen,
			Orientation: White,
			Moved:       false,
		},
	}
	BlackQueenStart = Point{
		AbsPoint: AbsPoint{File: 3, Rank: 7},
		Piece: &Piece{
			Kind:        Queen,
			Orientation: Black,
			Moved:       false,
		},
	}
	WhiteLeftRookStart = Point{
		AbsPoint: AbsPoint{File: 0, Rank: 0},
		Piece: &Piece{
			Kind:        Rook,
			Orientation: White,
			Moved:       false,
		},
	}
	WhiteRightRookStart = Point{
		AbsPoint: AbsPoint{File: 7, Rank: 0},
		Piece: &Piece{
			Kind:        Rook,
			Orientation: White,
			Moved:       false,
		},
	}
	BlackLeftRookStart = Point{
		AbsPoint: AbsPoint{File: 0, Rank: 7},
		Piece: &Piece{
			Kind:        Rook,
			Orientation: Black,
			Moved:       false,
		},
	}
	BlackRightRookStart = Point{
		AbsPoint: AbsPoint{File: 7, Rank: 7},
		Piece: &Piece{
			Kind:        Rook,
			Orientation: Black,
			Moved:       false,
		},
	}
	WhiteLeftKnightStart = Point{
		AbsPoint: AbsPoint{File: 1, Rank: 0},
		Piece: &Piece{
			Kind:        Knight,
			Orientation: White,
			Moved:       false,
		},
	}
	WhiteRightKnightStart = Point{
		AbsPoint: AbsPoint{File: 6, Rank: 0},
		Piece: &Piece{
			Kind:        Knight,
			Orientation: White,
			Moved:       false,
		},
	}
	BlackLeftKnightStart = Point{
		AbsPoint: AbsPoint{File: 1, Rank: 7},
		Piece: &Piece{
			Kind:        Knight,
			Orientation: Black,
			Moved:       false,
		},
	}
	BlackRightKnightStart = Point{
		AbsPoint: AbsPoint{File: 6, Rank: 7},
		Piece: &Piece{
			Kind:        Knight,
			Orientation: Black,
			Moved:       false,
		},
	}
	WhiteLeftBishopStart = Point{
		AbsPoint: AbsPoint{File: 2, Rank: 0},
		Piece: &Piece{
			Kind:        Bishop,
			Orientation: White,
			Moved:       false,
		},
	}
	WhiteRightBishopStart = Point{
		AbsPoint: AbsPoint{File: 5, Rank: 0},
		Piece: &Piece{
			Kind:        Bishop,
			Orientation: White,
			Moved:       false,
		},
	}
	BlackLeftBishopStart = Point{
		AbsPoint: AbsPoint{File: 2, Rank: 7},
		Piece: &Piece{
			Kind:        Bishop,
			Orientation: Black,
			Moved:       false,
		},
	}
	BlackRightBishopStart = Point{
		AbsPoint: AbsPoint{File: 5, Rank: 7},
		Piece: &Piece{
			Kind:        Bishop,
			Orientation: Black,
			Moved:       false,
		},
	}
	WhitePawn0Start = Point{
		AbsPoint: AbsPoint{File: 0, Rank: 1},
		Piece: &Piece{
			Kind:        Pawn,
			Orientation: White,
			Moved:       false,
		},
	}
	WhitePawn1Start = Point{
		AbsPoint: AbsPoint{File: 1, Rank: 1},
		Piece: &Piece{
			Kind:        Pawn,
			Orientation: White,
			Moved:       false,
		},
	}
	WhitePawn2Start = Point{
		AbsPoint: AbsPoint{File: 2, Rank: 1},
		Piece: &Piece{
			Kind:        Pawn,
			Orientation: White,
			Moved:       false,
		},
	}
	WhitePawn3Start = Point{
		AbsPoint: AbsPoint{File: 3, Rank: 1},
		Piece: &Piece{
			Kind:        Pawn,
			Orientation: White,
			Moved:       false,
		},
	}
	WhitePawn4Start = Point{
		AbsPoint: AbsPoint{File: 4, Rank: 1},
		Piece: &Piece{
			Kind:        Pawn,
			Orientation: White,
			Moved:       false,
		},
	}
	WhitePawn5Start = Point{
		AbsPoint: AbsPoint{File: 5, Rank: 1},
		Piece: &Piece{
			Kind:        Pawn,
			Orientation: White,
			Moved:       false,
		},
	}
	WhitePawn6Start = Point{
		AbsPoint: AbsPoint{File: 6, Rank: 1},
		Piece: &Piece{
			Kind:        Pawn,
			Orientation: White,
			Moved:       false,
		},
	}
	WhitePawn7Start = Point{
		AbsPoint: AbsPoint{File: 7, Rank: 1},
		Piece: &Piece{
			Kind:        Pawn,
			Orientation: White,
			Moved:       false,
		},
	}
	BlackPawn0Start = Point{
		AbsPoint: AbsPoint{File: 0, Rank: 6},
		Piece: &Piece{
			Kind:        Pawn,
			Orientation: Black,
			Moved:       false,
		},
	}
	BlackPawn1Start = Point{
		AbsPoint: AbsPoint{File: 1, Rank: 6},
		Piece: &Piece{
			Kind:        Pawn,
			Orientation: Black,
			Moved:       false,
		},
	}
	BlackPawn2Start = Point{
		AbsPoint: AbsPoint{File: 2, Rank: 6},
		Piece: &Piece{
			Kind:        Pawn,
			Orientation: Black,
			Moved:       false,
		},
	}
	BlackPawn3Start = Point{
		AbsPoint: AbsPoint{File: 3, Rank: 6},
		Piece: &Piece{
			Kind:        Pawn,
			Orientation: Black,
			Moved:       false,
		},
	}
	BlackPawn4Start = Point{
		AbsPoint: AbsPoint{File: 4, Rank: 6},
		Piece: &Piece{
			Kind:        Pawn,
			Orientation: Black,
			Moved:       false,
		},
	}
	BlackPawn5Start = Point{
		AbsPoint: AbsPoint{File: 5, Rank: 6},
		Piece: &Piece{
			Kind:        Pawn,
			Orientation: Black,
			Moved:       false,
		},
	}
	BlackPawn6Start = Point{
		AbsPoint: AbsPoint{File: 6, Rank: 6},
		Piece: &Piece{
			Kind:        Pawn,
			Orientation: Black,
			Moved:       false,
		},
	}
	BlackPawn7Start = Point{
		AbsPoint: AbsPoint{File: 7, Rank: 6},
		Piece: &Piece{
			Kind:        Pawn,
			Orientation: Black,
			Moved:       false,
		},
	}
)
