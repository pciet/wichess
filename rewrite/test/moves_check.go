package test

func init() {
	for _, c := range CheckMovesCases {
		MovesTestCases = append(MovesTestCases, c)
	}
}

var CheckMovesCases = []MovesTestCase{
	{
		"Failed Random Test",
		Black,
		Move{
			Address{3, 3},
			Address{3, 4},
		},
		Check,
		[]AddressedSquare{
			{Address{5, 0}, WRook},
			{Address{6, 0}, WRook},
			{Address{2, 1}, BPawn},
			{Address{6, 1}, BPawn},
			{Address{7, 1}, WKnight},
			{Address{1, 2}, WPawn},
			{Address{3, 2}, WKnight},
			{Address{4, 4}, BKing},
			{Address{6, 4}, WPawn},
			{Address{3, 5}, WPawn},
			{Address{5, 5}, BRook},
			{Address{6, 6}, WBishop},
			BLRook,
			{Address{3, 7}, WQueen},
		},
		[]MoveSet{
			{Address{4, 4}, []Address{
				{3, 4}, {4, 5}, {3, 3}, {4, 3},
			}},
		},
	},
}
