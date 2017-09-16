// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import (
	"fmt"
)

type Board [64]Point

// An empty PointSet return indicates no changes to the board - an invalid move.
// The board itself is not returned so no modifications are made to the receiver Board.
func (b Board) Move(from AbsPoint, to AbsPoint, turn Orientation) PointSet {
	if b[from.Index()].Piece == nil {
		return PointSet{}
	}
	if b[from.Index()].Orientation != turn {
		return PointSet{}
	}
	if b[to.Index()].Piece != nil {
		if (b[to.Index()].Orientation == turn) && (b[from.Index()].Swaps == false) {
			return PointSet{}
		}
	}
	set := make(PointSet)
	if b[to.Index()].Piece != nil {
		if b[from.Index()].Swaps && (b[to.Index()].Orientation == turn) {
			set[&Point{
				Piece:    b[to.Index()].Piece,
				AbsPoint: from,
			}] = struct{}{}
			b[to.Index()].Piece.Moved = true
			set[&Point{
				Piece:    b[from.Index()].Piece,
				AbsPoint: to,
			}] = struct{}{}
			b[from.Index()].Piece.Moved = true
			return set
		}
	}
	set[&Point{
		Piece:    nil,
		AbsPoint: from,
	}] = struct{}{}
	b[from.Index()].Piece.Moved = true
	set[&Point{
		Piece:    b[from.Index()].Piece,
		AbsPoint: to,
	}] = struct{}{}
	return set
}

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
	firstSet := make(AbsPointSet)
	moveSet := make(AbsPointSet)
	takeSet := make(AbsPointSet)
	for movetype, unfilteredpaths := range TruncatedAbsPathsForKind(the.Piece.Kind, the.AbsPoint, the.Piece.Orientation) {
		if (movetype == First) && (the.Piece.Moved == false) {
			for path, _ := range unfilteredpaths {
				if path.Truncated && the.Piece.MustEnd {
					continue
				}
				for i, point := range path.Points {
					if (b[point.Index()].Piece != nil) && (the.Piece.Ghost == false) {
						break
					} else if b[point.Index()].Piece != nil {
						if (b[point.Index()].Piece.Orientation != the.Piece.Orientation) || (the.Piece.Swaps == false) {
							continue
						}
					}
					if the.Piece.MustEnd {
						if len(path.Points) != i+1 {
							continue
						}
					}
					firstSet[&AbsPoint{
						File: point.File,
						Rank: point.Rank}] = struct{}{}
				}
			}
		} else if (movetype == Move) && (the.Piece.Moved == true) {
			for path, _ := range unfilteredpaths {
				if path.Truncated && the.Piece.MustEnd {
					continue
				}
				for i, point := range path.Points {
					if (b[point.Index()].Piece != nil) && (the.Piece.Ghost == false) {
						break
					} else if b[point.Index()].Piece != nil {
						if (b[point.Index()].Piece.Orientation != the.Piece.Orientation) || (the.Piece.Swaps == false) {
							continue
						}
					}
					if the.Piece.MustEnd {
						if len(path.Points) != i+1 {
							continue
						}
					}
					moveSet[&AbsPoint{
						File: point.File,
						Rank: point.Rank}] = struct{}{}
				}
			}
		} else if movetype == Take {
			for path, _ := range unfilteredpaths {
				if path.Truncated && the.Piece.MustEnd {
					continue
				}
				for ind, point := range path.Points {
					i := point.Index()
					if b[i].Piece == nil {
						continue
					}
					if the.Piece.MustEnd {
						if len(path.Points) != ind+1 {
							continue
						}
					}
					if b[i].Piece.Orientation != the.Piece.Orientation {
						takeSet[&AbsPoint{
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
	set := make(AbsPointSet)
	if (len(takeSet) == 0) || (the.Piece.MustTake == false) {
		set = set.Add(firstSet)
		set = set.Add(moveSet)
	}
	set = set.Add(takeSet)
	return set
}
