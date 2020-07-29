package rules

import (
	"strings"

	"github.com/pciet/wichess/piece"
)

// TODO: define paths in package piece

// Most abilities and some rules (like en passant and castling) will cause changes or
// additions to these paths.

// Lookup a piece's relative path variations by keying PieceRelPaths with the PieceKind.
// If first, rally, or take are nil then the move paths are used.
// All piece kinds have an entry in this map.
var PieceRelPaths = func() map[piece.Kind]RelPathVariations {
	m := map[piece.Kind]RelPathVariations{
		piece.King: {
			NormalMove: KingPaths,
		},
		piece.Queen: {
			NormalMove: QueenPaths,
		},
		piece.Rook: {
			NormalMove: RookPaths,
			RallyMove:  KingPaths,
		},
		piece.Bishop: {
			NormalMove: BishopPaths,
			RallyMove:  KingPaths,
		},
		piece.Knight: {
			NormalMove: KnightPaths,
			RallyMove:  TripleKnightPaths,
		},
		piece.Pawn: {
			First: {
				{{0, 1}, {0, 2}},
			},
			NormalMove: {
				{{0, 1}},
			},
			RallyMove: {
				{{1, 0}},
				{{-1, 0}},
				{{1, 1}},
				{{-1, 1}},
			},
			Take: {
				{{1, 1}},
				{{-1, 1}},
			},
		},
		piece.War: {
			NormalMove: {
				{{0, 1}},
			},
			RallyMove: {
				{{1, 0}},
				{{-1, 0}},
				{{1, 1}},
				{{-1, 1}},
			},
			Take: {
				{{1, 1}},
				{{-1, 1}},
			},
		},
		piece.Original: {
			NormalMove: SingleBishopPaths,
			RallyMove:  TwoBishopPaths,
		},
		piece.Irrelevant: {
			NormalMove: FiveRookPaths,
			RallyMove:  KingPaths,
		},
		piece.Evident: {
			NormalMove: {
				{{0, 1}},
			},
			RallyMove: {
				{{1, 0}},
				{{-1, 0}},
				{{1, 1}},
				{{-1, 1}},
			},
			Take: {
				{{1, -1}},
				{{-1, -1}},
			},
		},
		piece.Line: {
			NormalMove: NoGhostForwardKnightPaths,
			RallyMove:  NoGhostKnightPaths,
		},
		piece.Impossible: {
			NormalMove: FourRookPaths,
			RallyMove:  KingPaths,
		},
		piece.Convenient: {
			NormalMove: TwoBishopPaths,
			RallyMove:  KingPaths,
		},
		piece.Appropriate: {
			NormalMove: NoGhostKnightPaths,
			RallyMove:  KingPaths,
		},
		piece.Warp: {
			NormalMove: FiveRookPaths,
			RallyMove:  KingPaths,
		},
		piece.Brilliant: {
			NormalMove: ForwardKnightPaths,
			RallyMove:  KnightPaths,
		},
		piece.Exit: {
			NormalMove: ThreeBishopPaths,
			RallyMove:  KingPaths,
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

	assign := func(v RelPathVariations, f []piece.Kind) {
		for _, k := range f {
			m[k] = v
		}
	}

	assign(m[piece.Pawn], []piece.Kind{piece.Form, piece.Confined})
	assign(m[piece.Knight], []piece.Kind{piece.Constructive})
	assign(m[piece.Rook], []piece.Kind{piece.Simple})

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
