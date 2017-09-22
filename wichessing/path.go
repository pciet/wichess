// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import (
	"fmt"
)

type Path []Point

// This function will truncate moves that leave the board and convert the defined relative points to actual board points but does not remove points with other pieces or other similar constraints.
func TruncatedAbsPathsForKind(the Kind, from AbsPoint, with Orientation) AbsPathSetMap {
	absmap := make(AbsPathSetMap)
	for movetype, paths := range RelPathMapForKind(the) {
		availablepaths := make(AbsPathSet)
		for path, _ := range paths {
			availablepath := AbsPath{
				Points: make([]AbsPoint, 0, len(*path)),
			}
			truncated := false
			for _, point := range *path {
				absfile := int8(from.File) + point.XOffset
				if (absfile > 7) || (absfile < 0) {
					truncated = true
					break
				}
				var absrank int8
				if with == White {
					absrank = int8(from.Rank) + point.YOffset
				} else {
					absrank = int8(from.Rank) - point.YOffset
				}
				if (absrank > 7) || (absrank < 0) {
					truncated = true
					break
				}
				availablepath.Points = append(availablepath.Points, AbsPoint{File: uint8(absfile), Rank: uint8(absrank)})
			}
			availablepath.Truncated = truncated
			if len(availablepath.Points) != 0 {
				availablepaths[&availablepath] = struct{}{}
			}
		}
		if len(availablepaths) != 0 {
			absmap[movetype] = availablepaths
		}
	}
	return absmap
}

func RelPathMapForKind(the Kind) RelPathSetMap {
	switch the {
	case King:
		return KingPathMap
	case Guard:
		return GuardPathMap
	case Queen:
		return QueenPathMap
	case Rook, Rally, Fortify:
		return RookPathMap
	case Bishop, Detonate, Ghost, Steal:
		return BishopPathMap
	case Knight, Swap, Lock, Recon:
		return KnightPathMap
	case Pawn:
		return PawnPathMap
	default:
		panic(fmt.Sprintf("wichessing: invalid kind %v", the))
	}
}

type RelPath []RelPoint

type AbsPath struct {
	Points    []AbsPoint
	Truncated bool
}

func (the AbsPath) Copy() *AbsPath {
	p := AbsPath{
		Points: make([]AbsPoint, 0, len(the.Points)),
	}
	for _, pt := range the.Points {
		p.Points = append(p.Points, pt)
	}
	return &p
}

// All relative paths for a piece, used to calculate actual paths for a board state.
type RelPathSet map[*RelPath]struct{}

type AbsPathSet map[*AbsPath]struct{}

var (
	KnightPathSet = RelPathSet{
		&RelPath{{0, 1}, {0, 2}, {-1, 2}}:    {},
		&RelPath{{0, 1}, {0, 2}, {1, 2}}:     {},
		&RelPath{{1, 0}, {2, 0}, {2, 1}}:     {},
		&RelPath{{1, 0}, {2, 0}, {2, -1}}:    {},
		&RelPath{{0, -1}, {0, -2}, {1, -2}}:  {},
		&RelPath{{0, -1}, {0, -2}, {-1, -2}}: {},
		&RelPath{{-1, 0}, {-2, 0}, {-2, 1}}:  {},
		&RelPath{{-1, 0}, {-2, 0}, {-2, -1}}: {},
	}
	TripleKnightPathSet = RelPathSet{
		&RelPath{{0, 1}, {0, 2}, {0, 3}, {-1, 3}}:     {},
		&RelPath{{0, 1}, {0, 2}, {0, 3}, {1, 3}}:      {},
		&RelPath{{1, 0}, {2, 0}, {3, 0}, {3, 1}}:      {},
		&RelPath{{1, 0}, {2, 0}, {3, 0}, {3, -1}}:     {},
		&RelPath{{0, -1}, {0, -2}, {0, -3}, {1, -3}}:  {},
		&RelPath{{0, -1}, {0, -2}, {0, -3}, {-1, -3}}: {},
		&RelPath{{-1, 0}, {-2, 0}, {-3, 0}, {-3, 1}}:  {},
		&RelPath{{-1, 0}, {-2, 0}, {-3, 0}, {-3, -1}}: {},
	}
	BishopPathSet = RelPathSet{
		&RelPath{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}}:               {},
		&RelPath{{-1, -1}, {-2, -2}, {-3, -3}, {-4, -4}, {-5, -5}, {-6, -6}, {-7, -7}}: {},
		&RelPath{{1, -1}, {2, -2}, {3, -3}, {4, -4}, {5, -5}, {6, -6}, {7, -7}}:        {},
		&RelPath{{-1, 1}, {-2, 2}, {-3, 3}, {-4, 4}, {-5, 5}, {-6, 6}, {-7, 7}}:        {},
	}
	RookPathSet = RelPathSet{
		&RelPath{{1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0}}:        {},
		&RelPath{{-1, 0}, {-2, 0}, {-3, 0}, {-4, 0}, {-5, 0}, {-6, 0}, {-7, 0}}: {},
		&RelPath{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6}, {0, 7}}:        {},
		&RelPath{{0, -1}, {0, -2}, {0, -3}, {0, -4}, {0, -5}, {0, -6}, {0, -7}}: {},
	}
	QueenPathSet = RelPathSet{
		&RelPath{{1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0}}:               {},
		&RelPath{{-1, 0}, {-2, 0}, {-3, 0}, {-4, 0}, {-5, 0}, {-6, 0}, {-7, 0}}:        {},
		&RelPath{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6}, {0, 7}}:               {},
		&RelPath{{0, -1}, {0, -2}, {0, -3}, {0, -4}, {0, -5}, {0, -6}, {0, -7}}:        {},
		&RelPath{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}}:               {},
		&RelPath{{-1, -1}, {-2, -2}, {-3, -3}, {-4, -4}, {-5, -5}, {-6, -6}, {-7, -7}}: {},
		&RelPath{{1, -1}, {2, -2}, {3, -3}, {4, -4}, {5, -5}, {6, -6}, {7, -7}}:        {},
		&RelPath{{-1, 1}, {-2, 2}, {-3, 3}, {-4, 4}, {-5, 5}, {-6, 6}, {-7, 7}}:        {},
	}
	KingPathSet = RelPathSet{
		&RelPath{{0, 1}}:   {},
		&RelPath{{1, 1}}:   {},
		&RelPath{{1, 0}}:   {},
		&RelPath{{1, -1}}:  {},
		&RelPath{{0, -1}}:  {},
		&RelPath{{-1, -1}}: {},
		&RelPath{{-1, 0}}:  {},
		&RelPath{{-1, 1}}:  {},
	}
	DoubleKingPathSet = RelPathSet{
		&RelPath{{0, 1}, {0, 2}}:     {},
		&RelPath{{1, 1}, {2, 2}}:     {},
		&RelPath{{1, 0}, {2, 0}}:     {},
		&RelPath{{1, -1}, {2, -2}}:   {},
		&RelPath{{0, -1}, {0, -2}}:   {},
		&RelPath{{-1, -1}, {-2, -2}}: {},
		&RelPath{{-1, 0}, {-2, 0}}:   {},
		&RelPath{{-1, 1}, {-2, 2}}:   {},
	}
)

// The PathType is used for pieces with varying moves; the pawn is the base chess example with different first move, take moves, and regular moves.
type PathType int

const (
	First PathType = iota
	Move
	Take
	RallyMove
)

// These are all of the movement paths available for any possible board state.
type RelPathSetMap map[PathType]RelPathSet

type AbsPathSetMap map[PathType]AbsPathSet

var (
	PawnPathMap = RelPathSetMap{
		First: RelPathSet{
			&RelPath{{0, 1}, {0, 2}}: {},
		},
		Move: RelPathSet{
			&RelPath{{0, 1}}: {},
		},
		Take: RelPathSet{
			&RelPath{{1, 1}}:  {},
			&RelPath{{-1, 1}}: {},
		},
		RallyMove: RelPathSet{
			&RelPath{{0, 1}, {0, 2}}: {},
		},
	}
	RookPathMap = RelPathSetMap{
		First:     RookPathSet,
		Move:      RookPathSet,
		Take:      RookPathSet,
		RallyMove: KingPathSet,
	}
	KnightPathMap = RelPathSetMap{
		First:     KnightPathSet,
		Move:      KnightPathSet,
		Take:      KnightPathSet,
		RallyMove: TripleKnightPathSet,
	}
	BishopPathMap = RelPathSetMap{
		First:     BishopPathSet,
		Move:      BishopPathSet,
		Take:      BishopPathSet,
		RallyMove: KingPathSet,
	}
	QueenPathMap = RelPathSetMap{
		First: QueenPathSet,
		Move:  QueenPathSet,
		Take:  QueenPathSet,
	}
	KingPathMap = RelPathSetMap{
		First: KingPathSet,
		Move:  KingPathSet,
		Take:  KingPathSet,
	}
	GuardPathMap = RelPathSetMap{
		First:     KingPathSet,
		Move:      KingPathSet,
		RallyMove: DoubleKingPathSet,
	}
)
