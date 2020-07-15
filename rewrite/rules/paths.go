package rules

// Each piece has sets of possible paths that depend on the position.
// In the host these are represented as paths of relative addresses before application
// to a position.

var (
	KnightPaths = []RelPath{
		{{0, 1}, {0, 2}, {-1, 2}},
		{{0, 1}, {0, 2}, {1, 2}},
		{{1, 0}, {2, 0}, {2, 1}},
		{{1, 0}, {2, 0}, {2, -1}},
		{{0, -1}, {0, -2}, {1, -2}},
		{{0, -1}, {0, -2}, {-1, -2}},
		{{-1, 0}, {-2, 0}, {-2, 1}},
		{{-1, 0}, {-2, 0}, {-2, -1}},
	}

	TripleKnightPaths = []RelPath{
		{{0, 1}, {0, 2}, {0, 3}, {-1, 3}},
		{{0, 1}, {0, 2}, {0, 3}, {1, 3}},
		{{1, 0}, {2, 0}, {3, 0}, {3, 1}},
		{{1, 0}, {2, 0}, {3, 0}, {3, -1}},
		{{0, -1}, {0, -2}, {0, -3}, {1, -3}},
		{{0, -1}, {0, -2}, {0, -3}, {-1, -3}},
		{{-1, 0}, {-2, 0}, {-3, 0}, {-3, 1}},
		{{-1, 0}, {-2, 0}, {-3, 0}, {-3, -1}},
	}

	BishopPaths = []RelPath{
		{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}},
		{{-1, -1}, {-2, -2}, {-3, -3}, {-4, -4}, {-5, -5}, {-6, -6}, {-7, -7}},
		{{1, -1}, {2, -2}, {3, -3}, {4, -4}, {5, -5}, {6, -6}, {7, -7}},
		{{-1, 1}, {-2, 2}, {-3, 3}, {-4, 4}, {-5, 5}, {-6, 6}, {-7, 7}},
	}

	SingleBishopPaths = []RelPath{
		{{1, 1}},
		{{-1, -1}},
		{{1, -1}},
		{{-1, 1}},
	}

	TwoBishopPaths = []RelPath{
		{{1, 1}, {2, 2}},
		{{-1, -1}, {-2, -2}},
		{{1, -1}, {2, -2}},
		{{-1, 1}, {-2, 2}},
	}

	RookPaths = []RelPath{
		{{1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0}},
		{{-1, 0}, {-2, 0}, {-3, 0}, {-4, 0}, {-5, 0}, {-6, 0}, {-7, 0}},
		{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6}, {0, 7}},
		{{0, -1}, {0, -2}, {0, -3}, {0, -4}, {0, -5}, {0, -6}, {0, -7}},
	}

	FiveRookPaths = []RelPath{
		{{1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}},
		{{-1, 0}, {-2, 0}, {-3, 0}, {-4, 0}, {-5, 0}},
		{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}},
		{{0, -1}, {0, -2}, {0, -3}, {0, -4}, {0, -5}},
	}

	QueenPaths = []RelPath{
		{{1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0}},
		{{-1, 0}, {-2, 0}, {-3, 0}, {-4, 0}, {-5, 0}, {-6, 0}, {-7, 0}},
		{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6}, {0, 7}},
		{{0, -1}, {0, -2}, {0, -3}, {0, -4}, {0, -5}, {0, -6}, {0, -7}},
		{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}},
		{{-1, -1}, {-2, -2}, {-3, -3}, {-4, -4}, {-5, -5}, {-6, -6}, {-7, -7}},
		{{1, -1}, {2, -2}, {3, -3}, {4, -4}, {5, -5}, {6, -6}, {7, -7}},
		{{-1, 1}, {-2, 2}, {-3, 3}, {-4, 4}, {-5, 5}, {-6, 6}, {-7, 7}},
	}

	KingPaths = []RelPath{
		{{0, 1}},
		{{1, 1}},
		{{1, 0}},
		{{1, -1}},
		{{0, -1}},
		{{-1, -1}},
		{{-1, 0}},
		{{-1, 1}},
	}

	DoubleKingPaths = []RelPath{
		{{0, 1}, {0, 2}},
		{{1, 1}, {2, 2}},
		{{1, 0}, {2, 0}},
		{{1, -1}, {2, -2}},
		{{0, -1}, {0, -2}},
		{{-1, -1}, {-2, -2}},
		{{-1, 0}, {-2, 0}},
		{{-1, 1}, {-2, 2}},
	}

	ExtendedBishopPaths = CombineRelPathSlices(KingPaths, BishopPaths)
	ExtendedRookPaths   = CombineRelPathSlices(KingPaths, RookPaths)

	ExtendedKnightRallyPaths = CombineRelPathSlices(TripleKnightPaths, KnightPaths)
	ExtendedBishopRallyPaths = CombineRelPathSlices(ExtendedBishopPaths, DoubleKingPaths)
	ExtendedRookRallyPaths   = CombineRelPathSlices(ExtendedRookPaths, DoubleKingPaths)
)
