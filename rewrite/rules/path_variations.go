package rules

import "strings"

// Most abilities and some rules (like en passant and castling) will cause changes or
// additions to these paths.

// Lookup a piece's relative path variations by keying PieceRelPaths with the PieceKind.
// If first, rally, or take are nil then the move paths are used.
// All piece kinds have an entry in this map.
var PieceRelPaths = func() map[PieceKind]RelPathVariations {
	m := map[PieceKind]RelPathVariations{
		King: {
			NormalMove: KingPaths,
		},
		Queen: {
			NormalMove: QueenPaths,
		},
		Rook: {
			NormalMove: RookPaths,
			RallyMove:  KingPaths,
		},
		Bishop: {
			NormalMove: BishopPaths,
			RallyMove:  KingPaths,
		},
		Knight: {
			NormalMove: KnightPaths,
			RallyMove:  TripleKnightPaths,
		},
		Pawn: {
			First: {
				{{0, 1}, {0, 2}},
			},
			NormalMove: {
				{{0, 1}},
			},
			RallyMove: {
				{{0, 1}, {0, 2}},
			},
			Take: {
				{{1, 1}},
				{{-1, 1}},
			},
		},
		War: {
			NormalMove: {
				{{0, 1}},
			},
			RallyMove: {
				{{0, 1}, {0, 2}},
			},
			Take: {
				{{1, 1}},
				{{-1, 1}},
			},
		},
	}

	for k, v := range m {
		if v[NormalMove] == nil {
			Panic("no basic move paths defined for piece", k)
		}
		if v[First] == nil {
			v[First] = v[NormalMove]
		}
		if v[RallyMove] == nil {
			v[RallyMove] = v[NormalMove]
		}
		if v[Take] == nil {
			v[Take] = v[NormalMove]
		}
		m[k] = v
	}

	assign := func(v RelPathVariations, f []PieceKind) {
		for _, k := range f {
			m[k] = v
		}
	}

	assign(m[Pawn], []PieceKind{Form})
	assign(m[Knight], []PieceKind{Constructive})

	return m
}()

func (set PathVariations) String() string {
	var s strings.Builder

	writePaths := func(paths []Path) {
		for _, path := range paths {
			s.WriteString(path.String())
			s.WriteRune('\n')
		}
	}

	s.WriteString("First\n")
	writePaths(set[First])
	s.WriteString("Normal\n")
	writePaths(set[NormalMove])
	s.WriteString("Rally\n")
	writePaths(set[RallyMove])
	s.WriteString("Take\n")
	writePaths(set[Take])

	return s.String()
}