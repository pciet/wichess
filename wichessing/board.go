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
	fromPoint := b[from.Index()]
	if fromPoint.Piece == nil {
		return PointSet{}
	}
	if fromPoint.Orientation != turn {
		return PointSet{}
	}
	toPoint := b[to.Index()]
	if toPoint.Piece != nil {
		if (toPoint.Orientation == turn) && (fromPoint.Swaps == false) {
			return PointSet{}
		}
	}
	set := make(PointSet)
	for pt, _ := range b.SurroundingPoints(fromPoint) {
		if pt.Piece != nil {
			if (pt.Orientation != fromPoint.Orientation) && pt.Locks {
				return PointSet{}
			}
		}
	}
	for pt, _ := range b.SurroundingPoints(toPoint) {
		if pt.Piece != nil {
			if (pt.Orientation != fromPoint.Orientation) && pt.Guards {
				set[&Point{
					AbsPoint: fromPoint.AbsPoint,
				}] = struct{}{}
				set[&Point{
					AbsPoint: pt.AbsPoint,
				}] = struct{}{}
				set[&Point{
					Piece:    pt.Piece,
					AbsPoint: toPoint.AbsPoint,
				}] = struct{}{}
				return set
			}
		}
	}
	if b.MovesFromPoint(fromPoint).Has(to) == false {
		return PointSet{}
	}
	if toPoint.Piece != nil {
		if fromPoint.Steals && (toPoint.Orientation != turn) {
			toPoint.Orientation = turn
			set[&Point{
				Piece:    toPoint.Piece,
				AbsPoint: to,
			}] = struct{}{}
			adj := b.AdjacentPointOnPath(from, to)
			if (adj.File != from.File) || (adj.Rank != from.Rank) {
				set[&Point{
					Piece:    nil,
					AbsPoint: from,
				}] = struct{}{}
			}
			set[&Point{
				Piece:    fromPoint.Piece,
				AbsPoint: adj,
			}] = struct{}{}
			return set
		}
		if fromPoint.Swaps && (toPoint.Orientation == turn) {
			set[&Point{
				Piece:    toPoint.Piece,
				AbsPoint: from,
			}] = struct{}{}
			toPoint.Moved = true
			set[&Point{
				Piece:    fromPoint.Piece,
				AbsPoint: to,
			}] = struct{}{}
			fromPoint.Moved = true
			return set
		}
	}
	set[&Point{
		AbsPoint: from,
	}] = struct{}{}
	dset := b.DetonationsFrom(to)
	if len(dset) > 1 {
		for pt, _ := range dset {
			set[&Point{
				AbsPoint: *pt,
			}] = struct{}{}
		}
		return set
	}
	fromPoint.Moved = true
	set[&Point{
		Piece:    fromPoint.Piece,
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
	for pt, _ := range b.SurroundingPoints(the) {
		if pt.Piece != nil {
			if (pt.Piece.Orientation != the.Piece.Orientation) && (pt.Piece.Locks) {
				return AbsPointSet{}
			}
		}
	}
	firstSet := make(AbsPointSet)
	moveSet := make(AbsPointSet)
	takeSet := make(AbsPointSet)
	rallySet := make(AbsPointSet)
	for movetype, unfilteredpaths := range TruncatedAbsPathsForKind(the.Piece.Kind, the.AbsPoint, the.Piece.Orientation) {
		for path, _ := range b.ActualPaths(the, movetype, unfilteredpaths) {
			for _, pt := range path.Points {
				switch movetype {
				case First:
					firstSet[&AbsPoint{
						File: pt.File,
						Rank: pt.Rank,
					}] = struct{}{}
				case Move:
					moveSet[&AbsPoint{
						File: pt.File,
						Rank: pt.Rank,
					}] = struct{}{}
				case Take:
					takeSet[&AbsPoint{
						File: pt.File,
						Rank: pt.Rank,
					}] = struct{}{}
				case RallyMove:
					rallySet[&AbsPoint{
						File: pt.File,
						Rank: pt.Rank,
					}] = struct{}{}
				}
			}
		}
	}
	set := make(AbsPointSet)
	if (len(takeSet) == 0) || (the.MustTake == false) {
		set = set.Add(firstSet)
		set = set.Add(moveSet)
		set = set.Add(rallySet)
	}
	set = set.Add(takeSet)
	set = set.Add(b.ReconPointsFrom(the))
	set = set.Reduce()
	return set
}

func (b Board) ActualPaths(the Point, movetype PathType, unfilteredpaths AbsPathSet) AbsPathSet {
	if the.Piece == nil {
		return AbsPathSet{}
	}
	actualPaths := make(AbsPathSet)
	switch movetype {
	case First:
		if the.Moved {
			break
		}
		for path, _ := range unfilteredpaths {
			if path.Truncated && the.MustEnd {
				continue
			}
			filteredPath := AbsPath{
				Points: make([]AbsPoint, 0, len(path.Points)),
			}
			for i, point := range path.Points {
				actualPoint := b[point.Index()]
				if (actualPoint.Piece != nil) && (the.Ghost == false) {
					break
				} else if actualPoint.Piece != nil {
					if (actualPoint.Orientation != the.Orientation) || (the.Swaps == false) {
						continue
					}
				}
				if the.MustEnd {
					if len(path.Points) != i+1 {
						continue
					}
				}
				filteredPath.Points = append(filteredPath.Points, point)
			}
			if len(filteredPath.Points) > 0 {
				actualPaths[filteredPath.Copy()] = struct{}{}
			}
		}
	case Move:
		if the.Moved == false {
			break
		}
		for path, _ := range unfilteredpaths {
			if path.Truncated && the.MustEnd {
				continue
			}
			filteredPath := AbsPath{
				Points: make([]AbsPoint, 0, len(path.Points)),
			}
			for i, point := range path.Points {
				actualPoint := b[point.Index()]
				if (actualPoint.Piece != nil) && (the.Ghost == false) {
					break
				} else if actualPoint.Piece != nil {
					if (actualPoint.Orientation != the.Orientation) || (the.Swaps == false) {
						continue
					}
				}
				if the.MustEnd {
					if len(path.Points) != i+1 {
						continue
					}
				}
				filteredPath.Points = append(filteredPath.Points, point)
			}
			if len(filteredPath.Points) > 0 {
				actualPaths[filteredPath.Copy()] = struct{}{}
			}
		}
	case Take:
		for path, _ := range unfilteredpaths {
			if path.Truncated && the.MustEnd {
				continue
			}
			filteredPath := AbsPath{
				Points: make([]AbsPoint, 0, len(path.Points)),
			}
			for ind, point := range path.Points {
				actualPoint := b[point.Index()]
				if actualPoint.Piece == nil {
					continue
				}
				if the.MustEnd {
					if len(path.Points) != ind+1 {
						continue
					}
				}
				if actualPoint.Orientation != the.Orientation {
					if (the.Kind == Pawn) && actualPoint.Fortified {
						break
					}
					filteredPath.Points = append(filteredPath.Points, point)
					break
				} else {
					if the.Ghost == false {
						break
					}
				}
			}
			if len(filteredPath.Points) > 0 {
				actualPaths[filteredPath.Copy()] = struct{}{}
			}
		}
	case RallyMove:
		var rallied bool
		for point, _ := range b.SurroundingPoints(the) {
			if point.Piece != nil {
				if (point.Orientation == the.Orientation) && point.Rallies {
					rallied = true
					break
				}
			}
		}
		if rallied == false {
			break
		}
		for path, _ := range unfilteredpaths {
			if path.Truncated && the.MustEnd {
				continue
			}
			filteredPath := AbsPath{
				Points: make([]AbsPoint, 0, len(path.Points)),
			}
			for i, point := range path.Points {
				actualPoint := b[point.Index()]
				if (actualPoint.Piece != nil) && (the.Ghost == false) {
					break
				} else if actualPoint.Piece != nil {
					if (actualPoint.Orientation != the.Orientation) || (the.Swaps == false) {
						continue
					}
				}
				if the.MustEnd {
					if len(path.Points) != i+1 {
						continue
					}
				}
				filteredPath.Points = append(filteredPath.Points, point)
			}
			if len(filteredPath.Points) > 0 {
				actualPaths[filteredPath.Copy()] = struct{}{}
			}
		}
	}
	return actualPaths
}

func (b Board) AdjacentPointOnPath(from AbsPoint, to AbsPoint) AbsPoint {
	fromPoint := b[from.Index()]
	surrounding := b.SurroundingPoints(b[to.Index()])
	for movetype, unfilteredpaths := range TruncatedAbsPathsForKind(fromPoint.Kind, from, fromPoint.Orientation) {
		if movetype == Take {
			continue
		}
		for path, _ := range b.ActualPaths(fromPoint, movetype, unfilteredpaths) {
			for _, pt := range path.Points {
				for spt, _ := range surrounding {
					if (pt.File == spt.File) && (pt.Rank == spt.Rank) {
						return pt
					}
				}
			}
		}
	}
	return from
}

func (b Board) DetonationsFrom(the AbsPoint) AbsPointSet {
	index := the.Index()
	set := make(AbsPointSet)
	if b[index].Piece == nil {
		return set
	}

	set[&the] = struct{}{}
	if b[index].Piece.Detonates == false {
		return set
	}
	for pt, _ := range b.SurroundingPoints(b[index]) {
		set = set.Add(b.DetonationsFrom(pt.AbsPoint))
	}
	return set.Reduce()
}

func (b Board) ReconPointsFrom(the Point) AbsPointSet {
	if the.Piece == nil {
		return AbsPointSet{}
	}
	var rank int
	if the.Piece.Orientation == White {
		rank = int(the.Rank) + 1
	} else {
		rank = int(the.Rank) - 1
	}
	if (rank < 0) || (rank >= 8) {
		return AbsPointSet{}
	}
	set := make(AbsPointSet)
	for f := -1; f <= 1; f++ {
		file := int(the.File) + f
		if (file < 0) || (file >= 8) {
			continue
		}
		index := IndexFromFileAndRank(uint8(file), uint8(rank))
		piece := b[index].Piece
		if piece == nil {
			continue
		}
		if piece.Orientation != the.Piece.Orientation {
			continue
		}
		if piece.Recons == false {
			continue
		}
		var nrank int
		if the.Piece.Orientation == White {
			nrank = int(b[index].Rank) + 1
		} else {
			nrank = int(b[index].Rank) - 1
		}
		if (nrank < 0) || (nrank >= 8) {
			continue
		}
		if b[IndexFromFileAndRank(uint8(file), uint8(nrank))].Piece != nil {
			continue
		}
		set[&AbsPoint{
			File: b[index].File,
			Rank: uint8(nrank),
		}] = struct{}{}
	}
	return set
}

func (b Board) SurroundingPoints(from Point) PointSet {
	set := make(PointSet)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (i == 0) && (j == 0) {
				continue
			}
			f := i + int(from.File)
			r := j + int(from.Rank)
			if (f < 0) || (f >= 8) {
				continue
			}
			if (r < 0) || (r >= 8) {
				continue
			}
			set[&b[IndexFromFileAndRank(uint8(f), uint8(r))]] = struct{}{}
		}
	}
	return set
}
