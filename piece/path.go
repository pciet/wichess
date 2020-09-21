package piece

import (
	"log"
	"strings"
)

const (
	FirstPaths PathVariation = iota
	NormalPaths
	RallyPaths
	CapturePaths
	PathVariationCount // not a variation
)

type (
	// A PathAddress is a board square address from the perspective of a piece with the white
	// orientation. File is to the left (negative) or right (positive) and Rank is either
	// down (negative) or up (positive).
	PathAddress struct {
		File, Rank int
	}

	// A Path is an ordered list of PathAddress that represent a way a piece can move. Normally
	// pieces cannot move past another piece on a path.
	Path []PathAddress

	// PathVariation adds meaning to a Path. For example, some paths can only be used to capture
	// or are only available for the first move of the piece.
	PathVariation int

	// The PathVariations of a piece are an unordered set of paths for each PathVariation.
	PathVariations [PathVariationCount][]Path
)

// Paths defines the paths for every piece in Wisconsin Chess. If the first, rally, or capture
// variations are nil then the normal variation is used for them. Most characteristics and some
// chess rules, like en passant and castling, will cause changes or additions to these paths.
func Paths(of Kind) PathVariations { return paths[of] }

var paths = func() map[Kind]PathVariations {
	m := map[Kind]PathVariations{
		King: {
			NormalPaths: kingPaths,
		},
		Queen: {
			NormalPaths: queenPaths,
		},
		Rook: {
			NormalPaths: rookPaths,
			RallyPaths:  kingPaths,
		},
		Bishop: {
			NormalPaths: bishopPaths,
			RallyPaths:  kingPaths,
		},
		Knight: {
			NormalPaths: knightPaths,
			RallyPaths:  tripleKnightPaths,
		},
		Pawn: {
			FirstPaths: {
				{{0, 1}, {0, 2}},
			},
			NormalPaths: {
				{{0, 1}},
			},
			RallyPaths: {
				{{1, 0}},
				{{-1, 0}},
				{{1, 1}},
				{{-1, 1}},
			},
			CapturePaths: {
				{{1, 1}},
				{{-1, 1}},
			},
		},
		War: {
			NormalPaths: {
				{{0, 1}},
			},
			RallyPaths: {
				{{1, 0}},
				{{-1, 0}},
				{{1, 1}},
				{{-1, 1}},
			},
			CapturePaths: {
				{{1, 1}},
				{{-1, 1}},
			},
		},
		Original: {
			NormalPaths: singleBishopPaths,
			RallyPaths:  twoBishopPaths,
		},
		Irrelevant: {
			NormalPaths: fiveRookPaths,
			RallyPaths:  kingPaths,
		},
		Evident: {
			NormalPaths: {
				{{0, 1}},
			},
			RallyPaths: {
				{{1, 0}},
				{{-1, 0}},
				{{1, 1}},
				{{-1, 1}},
			},
			CapturePaths: {
				{{1, -1}},
				{{-1, -1}},
			},
		},
		Line: {
			NormalPaths: noGhostForwardKnightPaths,
			RallyPaths:  noGhostKnightPaths,
		},
		Impossible: {
			NormalPaths: fourRookPaths,
			RallyPaths:  kingPaths,
		},
		Convenient: {
			NormalPaths: twoBishopPaths,
			RallyPaths:  kingPaths,
		},
		Appropriate: {
			NormalPaths: noGhostKnightPaths,
			RallyPaths:  kingPaths,
		},
		Warp: {
			NormalPaths: fiveRookPaths,
			RallyPaths:  kingPaths,
		},
		Brilliant: {
			NormalPaths: forwardKnightPaths,
			RallyPaths:  knightPaths,
		},
		Exit: {
			NormalPaths: threeBishopPaths,
			RallyPaths:  kingPaths,
		},
		Imperfect: {
			FirstPaths: {
				{{0, 1}, {0, 2}},
			},
			NormalPaths: {
				{{0, 1}},
			},
			RallyPaths: {
				{{1, 0}},
				{{-1, 0}},
				{{1, 1}},
				{{-1, 1}},
			},
			CapturePaths: {
				{{0, -1}},
			},
		},
		Derange: {
			FirstPaths: {
				{{-1, 1}},
				{{0, 1}, {0, 2}, {1, 2}},
				{{0, 1}, {0, 2}, {-1, 2}},
				{{1, 1}},
			},
			NormalPaths: {
				{{-1, 1}},
				{{0, 1}},
				{{1, 1}},
			},
			RallyPaths: {
				{{1, 0}},
				{{-1, 0}},
				{{-1, -1}},
				{{1, -1}},
			},
			CapturePaths: {
				{{0, -1}},
			},
		},
	}

	for k, v := range m {
		if v[NormalPaths] == nil {
			log.Panicln("no basic move paths defined for piece", k)
		}
		if v[FirstPaths] == nil {
			v[FirstPaths] = v[NormalPaths]
		}
		if v[RallyPaths] == nil {
			v[RallyPaths] = v[NormalPaths]
		}
		if v[CapturePaths] == nil {
			v[CapturePaths] = v[NormalPaths]
		}
		m[k] = v
	}

	assign := func(f []Kind, v PathVariations) {
		for _, k := range f {
			m[k] = v
		}
	}

	assign([]Kind{Form, Confined}, m[Pawn])
	assign([]Kind{Constructive}, m[Knight])
	assign([]Kind{Simple}, m[Rook])

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

	s.WriteString("FirstPaths\n")
	writePaths(set[First])
	s.WriteString("Normal\n")
	writePaths(set[NormalPaths])
	s.WriteString("Rally\n")
	writePaths(set[RallyPaths])
	s.WriteString("Capture\n")
	writePaths(set[CapturePaths])

	return s.String()
}
