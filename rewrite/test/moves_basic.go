package test

func init() {
	for _, c := range BasicMovesCases {
		MovesTestCases = append(MovesTestCases, c)
	}
}

var BasicMovesCases = []MovesTestCase{
	{
		"Failed Random Test 2",
		White,
		Move{
			Address{3, 3},
			Address{3, 2}},
		Normal,
		[]AddressedSquare{
			WLRook, WSKing, WRRook,
			WPawn0, WPawn2,
			{Address{4, 1}, WKnight},
			{Address{6, 1}, WBishop},
			{Address{0, 2}, WKnight},
			{Address{5, 2}, WPawn},
			{Address{7, 2}, WPawn},
			{Address{5, 3}, WPawn},
			{Address{7, 3}, BPawn},
			{Address{2, 4}, BPawn},
			{Address{4, 4}, WPawn},
			{Address{5, 4}, BPawn},
			{Address{0, 5}, BKnight},
			{Address{1, 5}, BQueen},
			{Address{3, 5}, WPawn},
			{Address{4, 5}, BPawn},
			{Address{5, 5}, WBishop},
			BPawn1, BPawn3,
			{Address{4, 6}, BBishop},
			{Address{1, 7}, BRook},
			BLBishop,
			{Address{5, 7}, BKing},
			BRRook,
		},
		[]MoveSet{
			{Address{0, 0}, []Address{
				{1, 0}, {2, 0}, {3, 0},
			}},
			{Address{4, 0}, []Address{
				{3, 0}, {3, 1}, {5, 1},
				{5, 0}, {2, 0}, {6, 0},
			}},
			{Address{7, 0}, []Address{
				{6, 0}, {5, 0}, {7, 1},
			}},
			{Address{2, 1}, []Address{
				{2, 2}, {2, 3},
			}},
			{Address{4, 1}, []Address{
				{2, 0}, {2, 2}, {3, 3},
				{6, 0}, {6, 2},
			}},
			{Address{6, 1}, []Address{{5, 0}}},
			{Address{0, 2}, []Address{
				{1, 0}, {2, 3}, {1, 4},
			}},
			{Address{3, 5}, []Address{{4, 6}}},
			{Address{5, 5}, []Address{
				{6, 4}, {6, 6}, {4, 6},
				{7, 7}, {7, 3},
			}},
		},
	},
	{
		"Chess Initial Position No Pawns White",
		White,
		Move{
			Address{0, 8},
			Address{0, 8},
		},
		Normal,
		[]AddressedSquare{
			WLRook, WLKnight, WLBishop,
			WSQueen, WSKing,
			WRBishop, WRKnight, WRRook,
			BLRook, BLKnight, BLBishop,
			BSQueen, BSKing,
			BRBishop, BRKnight, BRRook,
		},
		[]MoveSet{
			{Address{0, 0}, []Address{
				{0, 1}, {0, 2}, {0, 3}, {0, 4},
				{0, 5}, {0, 6}, {0, 7},
			}},
			{Address{1, 0}, []Address{
				{0, 2}, {2, 2}, {3, 1},
			}},
			{Address{2, 0}, []Address{
				{1, 1}, {0, 2}, {3, 1}, {4, 2},
				{5, 3}, {6, 4}, {7, 5},
			}},
			{Address{3, 0}, []Address{
				{2, 1}, {1, 2}, {0, 3}, {3, 1},
				{3, 2}, {3, 3}, {3, 4}, {3, 5},
				{3, 6}, {3, 7}, {4, 1}, {5, 2},
				{6, 3}, {7, 4},
			}},
			{Address{4, 0}, []Address{
				{4, 1}, {5, 1},
			}},
			{Address{5, 0}, []Address{
				{4, 1}, {3, 2}, {2, 3}, {1, 4},
				{0, 5}, {6, 1}, {7, 2},
			}},
			{Address{6, 0}, []Address{
				{7, 2}, {5, 2}, {4, 1},
			}},
			{Address{7, 0}, []Address{
				{7, 1}, {7, 2}, {7, 3}, {7, 4},
				{7, 5}, {7, 6}, {7, 7},
			}},
		},
	},
	{
		"Chess Initial Position No Pawns Black",
		Black,
		Move{
			Address{0, 8},
			Address{0, 8},
		},
		Normal,
		[]AddressedSquare{
			WLRook, WLKnight, WLBishop,
			WSQueen, WSKing,
			WRBishop, WRKnight, WRRook,
			BLRook, BLKnight, BLBishop,
			BSQueen, BSKing,
			BRBishop, BRKnight, BRRook,
		},
		[]MoveSet{
			{Address{0, 7}, []Address{
				{0, 6}, {0, 5}, {0, 4}, {0, 3},
				{0, 2}, {0, 1}, {0, 0},
			}},
			{Address{1, 7}, []Address{
				{0, 5}, {2, 5}, {3, 6},
			}},
			{Address{2, 7}, []Address{
				{1, 6}, {0, 5}, {3, 6}, {4, 5},
				{5, 4}, {6, 3}, {7, 2},
			}},
			{Address{3, 7}, []Address{
				{2, 6}, {1, 5}, {0, 4}, {3, 6},
				{3, 5}, {3, 4}, {3, 3}, {3, 2},
				{3, 1}, {3, 0}, {4, 6}, {5, 5},
				{6, 4}, {7, 3},
			}},
			{Address{4, 7}, []Address{
				{4, 6}, {5, 6},
			}},
			{Address{5, 7}, []Address{
				{4, 6}, {3, 5}, {2, 4}, {1, 3},
				{0, 2}, {6, 6}, {7, 5},
			}},
			{Address{6, 7}, []Address{
				{7, 5}, {5, 5}, {4, 6},
			}},
			{Address{7, 7}, []Address{
				{7, 6}, {7, 5}, {7, 4}, {7, 3},
				{7, 2}, {7, 1}, {7, 0},
			}},
		},
	},
}
