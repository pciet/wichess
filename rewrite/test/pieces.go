package test

// Any test that uses these helper variables must
// call rules.Piece.ApplyCharacteristics for the
// rules moves methods to work correctly with them.

var (
	WLRook = AddressedSquare{
		Address{0, 0}, Square{Kind: Rook, Orientation: White},
	}
	WLKnight = AddressedSquare{
		Address{1, 0}, Square{Kind: Knight, Orientation: White},
	}
	WLBishop = AddressedSquare{
		Address{2, 0}, Square{Kind: Bishop, Orientation: White},
	}
	WSQueen = AddressedSquare{
		Address{3, 0}, Square{Kind: Queen, Orientation: White},
	}
	WSKing = AddressedSquare{
		Address{4, 0}, Square{Kind: King, Orientation: White},
	}
	WRBishop = AddressedSquare{
		Address{5, 0}, Square{Kind: Bishop, Orientation: White},
	}
	WRKnight = AddressedSquare{
		Address{6, 0}, Square{Kind: Knight, Orientation: White},
	}
	WRRook = AddressedSquare{
		Address{7, 0}, Square{Kind: Rook, Orientation: White},
	}
	WPawn0 = AddressedSquare{
		Address{0, 1}, Square{Kind: Pawn, Orientation: White},
	}
	WPawn1 = AddressedSquare{
		Address{1, 1}, Square{Kind: Pawn, Orientation: White},
	}
	WPawn2 = AddressedSquare{
		Address{2, 1}, Square{Kind: Pawn, Orientation: White},
	}
	WPawn3 = AddressedSquare{
		Address{3, 1}, Square{Kind: Pawn, Orientation: White},
	}
	WPawn4 = AddressedSquare{
		Address{4, 1}, Square{Kind: Pawn, Orientation: White},
	}
	WPawn5 = AddressedSquare{
		Address{5, 1}, Square{Kind: Pawn, Orientation: White},
	}
	WPawn6 = AddressedSquare{
		Address{6, 1}, Square{Kind: Pawn, Orientation: White},
	}
	WPawn7 = AddressedSquare{
		Address{7, 1}, Square{Kind: Pawn, Orientation: White},
	}

	BLRook = AddressedSquare{
		Address{0, 7}, Square{Kind: Rook, Orientation: Black},
	}
	BLKnight = AddressedSquare{
		Address{1, 7}, Square{Kind: Knight, Orientation: Black},
	}
	BLBishop = AddressedSquare{
		Address{2, 7}, Square{Kind: Bishop, Orientation: Black},
	}
	BSQueen = AddressedSquare{
		Address{3, 7}, Square{Kind: Queen, Orientation: Black},
	}
	BSKing = AddressedSquare{
		Address{4, 7}, Square{Kind: King, Orientation: Black},
	}
	BRBishop = AddressedSquare{
		Address{5, 7}, Square{Kind: Bishop, Orientation: Black},
	}
	BRKnight = AddressedSquare{
		Address{6, 7}, Square{Kind: Knight, Orientation: Black},
	}
	BRRook = AddressedSquare{
		Address{7, 7}, Square{Kind: Rook, Orientation: Black},
	}
	BPawn0 = AddressedSquare{
		Address{0, 6}, Square{Kind: Pawn, Orientation: Black},
	}
	BPawn1 = AddressedSquare{
		Address{1, 6}, Square{Kind: Pawn, Orientation: Black},
	}
	BPawn2 = AddressedSquare{
		Address{2, 6}, Square{Kind: Pawn, Orientation: Black},
	}
	BPawn3 = AddressedSquare{
		Address{3, 6}, Square{Kind: Pawn, Orientation: Black},
	}
	BPawn4 = AddressedSquare{
		Address{4, 6}, Square{Kind: Pawn, Orientation: Black},
	}
	BPawn5 = AddressedSquare{
		Address{5, 6}, Square{Kind: Pawn, Orientation: Black},
	}
	BPawn6 = AddressedSquare{
		Address{6, 6}, Square{Kind: Pawn, Orientation: Black},
	}
	BPawn7 = AddressedSquare{
		Address{7, 6}, Square{Kind: Pawn, Orientation: Black},
	}

	WRook   = Square{Kind: Rook, Orientation: White, Moved: true}
	WBishop = Square{Kind: Bishop, Orientation: White, Moved: true}
	WKnight = Square{Kind: Knight, Orientation: White, Moved: true}
	WQueen  = Square{Kind: Queen, Orientation: White, Moved: true}
	WKing   = Square{Kind: King, Orientation: White, Moved: true}
	WPawn   = Square{Kind: Pawn, Orientation: White, Moved: true}

	BRook   = Square{Kind: Rook, Orientation: Black, Moved: true}
	BBishop = Square{Kind: Bishop, Orientation: Black, Moved: true}
	BKnight = Square{Kind: Knight, Orientation: Black, Moved: true}
	BQueen  = Square{Kind: Queen, Orientation: Black, Moved: true}
	BKing   = Square{Kind: King, Orientation: Black, Moved: true}
	BPawn   = Square{Kind: Pawn, Orientation: Black, Moved: true}
)
