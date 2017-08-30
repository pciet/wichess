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
		sets[point.AbsPoint] = b.MovesFromPoint(point)
	}
	return sets
}

func (b Board) MovesFromPoint(the Point) AbsPointSet {
	if the.Piece == nil {
		panic(fmt.Sprintf("wichessing: point (%v,%v) without piece", the.File, the.Rank))
	}
	set := make(AbsPointSet)
	for movetype, unfilteredpaths := range TruncatedAbsPathsForKind(the.Piece.Kind, the.AbsPoint) {
		if (movetype == First) && (the.Piece.Moved == false) {
			for path, _ := range unfilteredpaths {
				for _, point := range *path {
					if (b[point.Index()].Piece != nil) && (the.Piece.Ghost == false) {
						break
					}
					set[&point] = struct{}{}
				}
			}
		} else if (movetype == Move) && (the.Piece.Moved == true) {
			for path, _ := range unfilteredpaths {
				for _, point := range *path {
					if (b[point.Index()].Piece != nil) && (the.Piece.Ghost == false) {
						break
					}
					set[&point] = struct{}{}
				}
			}
		} else if movetype == Take {
			for path, _ := range unfilteredpaths {
				for _, point := range *path {
					i := point.Index()
					if b[i].Piece == nil {
						continue
					}
					if b[i].Piece.Orientation != the.Piece.Orientation {
						set[&point] = struct{}{}
						break
					} else {
						if the.Piece.Ghost == false {
							break
						}
					}
				}
			}
		} else {
			panic(fmt.Sprintf("wichessing: unexpected move type %v", movetype))
		}
	}
	return set
}
