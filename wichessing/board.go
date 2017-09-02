// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import (
	"fmt"
)

type Board [64]Point

func (b Board) Moves() map[AbsPoint]AbsPointSet {
	sets := make(map[AbsPoint]AbsPointSet)
	for _, point := range b {
		if point.Piece == nil {
			continue
		}
		moves := b.MovesFromPoint(point)
		if len(moves) == 0 {
			continue
		}
		sets[point.AbsPoint] = moves
	}
	return sets
}

func (b Board) MovesFromPoint(the Point) AbsPointSet {
	if the.Piece == nil {
		panic(fmt.Sprintf("wichessing: point (%v,%v) without piece", the.File, the.Rank))
	}
	set := make(AbsPointSet)
	for movetype, unfilteredpaths := range TruncatedAbsPathsForKind(the.Piece.Kind, the.AbsPoint, the.Piece.Orientation) {
		if (movetype == First) && (the.Piece.Moved == false) {
			for path, _ := range unfilteredpaths {
				for i, point := range *path {
					if (b[point.Index()].Piece != nil) && (the.Piece.Ghost == false) {
						break
					} else if b[point.Index()].Piece != nil {
						continue
					}
					if the.Piece.MustEnd {
						if len(*path) != i+1 {
							continue
						}
					}
					set[&AbsPoint{
						File: point.File,
						Rank: point.Rank}] = struct{}{}
				}
			}
		} else if (movetype == Move) && (the.Piece.Moved == true) {
			for path, _ := range unfilteredpaths {
				for i, point := range *path {
					if (b[point.Index()].Piece != nil) && (the.Piece.Ghost == false) {
						break
					} else if b[point.Index()].Piece != nil {
						continue
					}
					if the.Piece.MustEnd {
						if len(*path) != i+1 {
							continue
						}
					}
					set[&AbsPoint{
						File: point.File,
						Rank: point.Rank}] = struct{}{}
				}
			}
		} else if movetype == Take {
			for path, _ := range unfilteredpaths {
				for ind, point := range *path {
					i := point.Index()
					if b[i].Piece == nil {
						continue
					}
					if the.Piece.MustEnd {
						if len(*path) != ind+1 {
							continue
						}
					}
					if b[i].Piece.Orientation != the.Piece.Orientation {
						set[&AbsPoint{
							File: point.File,
							Rank: point.Rank}] = struct{}{}
						break
					} else {
						if the.Piece.Ghost == false {
							break
						}
					}
				}
			}
		}
	}
	return set
}
