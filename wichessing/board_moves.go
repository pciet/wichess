// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import (
	"fmt"
)

// Moves, Check, Checkmate.
func (b Board) Moves(active Orientation) (map[AbsPoint]AbsPointSet, bool, bool) {
	if b.Checkmate(active) {
		return nil, true, true
	}
	sets := make(map[AbsPoint]AbsPointSet)
	if b.Check(active) {
		return b.CheckMoves(active), true, false
	}
	for _, point := range b {
		if point.Piece == nil {
			continue
		}
		moves := b.MovesFromPoint(point)
		if len(moves) == 0 {
			continue
		}
		for move, _ := range moves {
			if b.AfterMove(point.AbsPoint, *move, b[point.Index()].Orientation).Check(b[point.Index()].Orientation) {
				delete(moves, move)
			}
		}
		sets[point.AbsPoint] = moves
	}
	return sets, false, false
}

func (b Board) CheckMoves(active Orientation) map[AbsPoint]AbsPointSet {
	moves := make(map[AbsPoint]AbsPointSet)
	unfiltered := make(map[AbsPoint]AbsPointSet)
	for _, point := range b {
		if point.Piece == nil {
			continue
		}
		m := b.MovesFromPoint(point)
		if len(m) == 0 {
			continue
		}
		if point.Piece.Orientation == active {
			unfiltered[point.AbsPoint] = m
		} else {
			moves[point.AbsPoint] = m
		}
	}
	for pt, set := range unfiltered {
		allowed := make(AbsPointSet)
		for mv, _ := range set {
			if b.AfterMove(pt, *mv, active).Check(active) {
				continue
			}
			allowed[mv] = struct{}{}
		}
		if len(allowed) == 0 {
			continue
		}
		moves[pt] = allowed
	}
	return moves
}

func (b Board) AllMovesFor(player Orientation) map[AbsPoint]AbsPointSet {
	moves := make(map[AbsPoint]AbsPointSet)
	for _, pt := range b {
		if pt.Piece == nil {
			continue
		}
		if pt.Piece.Orientation != player {
			continue
		}
		moves[pt.AbsPoint] = b.MovesFromPoint(pt)
	}
	return moves
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
	// castling
	if (the.Kind == King) && (the.Moved == false) {
		if (the.File == 4) && (the.Rank == 0) {
			if (b[1].Piece == nil) && (b[2].Piece == nil) && (b[3].Piece == nil) && (b[0].Piece != nil) {
				if b[0].Moved == false {
					set[&AbsPoint{
						File: 2,
						Rank: 0,
					}] = struct{}{}
				}
			}
			if (b[5].Piece == nil) && (b[6].Piece == nil) && (b[7].Piece != nil) {
				if b[7].Moved == false {
					set[&AbsPoint{
						File: 6,
						Rank: 0,
					}] = struct{}{}
				}
			}
		} else if (the.File == 4) && (the.Rank == 7) {
			if (b[57].Piece == nil) && (b[58].Piece == nil) && (b[59].Piece == nil) && (b[56].Piece != nil) {
				if b[56].Moved == false {
					set[&AbsPoint{
						File: 2,
						Rank: 7,
					}] = struct{}{}
				}
			}
			if (b[61].Piece == nil) && (b[62].Piece == nil) && (b[63].Piece != nil) {
				if b[63].Moved == false {
					set[&AbsPoint{
						File: 6,
						Rank: 7,
					}] = struct{}{}
				}
			}
		}
	}
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
		FIRST_PATH_OUTER:
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
				if the.Kind == King {
					for pt, _ := range b.SurroundingPoints(actualPoint) {
						if pt.Piece == nil {
							continue
						}
						if pt.Orientation == the.Orientation {
							continue
						}
						if pt.Kind == Guard {
							break FIRST_PATH_OUTER
						}
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
		MOVE_PATH_OUTER:
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
				if the.Kind == King {
					for pt, _ := range b.SurroundingPoints(actualPoint) {
						if pt.Piece == nil {
							continue
						}
						if pt.Orientation == the.Orientation {
							continue
						}
						if pt.Kind == Guard {
							break MOVE_PATH_OUTER
						}
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
		TAKE_PATH_OUTER:
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
					if actualPoint.Kind == Detonate {
						if the.Kind == King {
							break
						}
						for pt, _ := range b.SurroundingPoints(actualPoint) {
							if pt.Piece == nil {
								continue
							}
							if pt.Orientation != the.Orientation {
								continue
							}
							if pt.Kind == King {
								break TAKE_PATH_OUTER
							}
						}
					}
					if the.Kind == King {
						for pt, _ := range b.SurroundingPoints(actualPoint) {
							if pt.Piece == nil {
								continue
							}
							if pt.Orientation == the.Orientation {
								continue
							}
							if pt.Kind == Guard {
								break TAKE_PATH_OUTER
							}
						}
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
		RALLY_PATH_OUTER:
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
				if the.Kind == King {
					for pt, _ := range b.SurroundingPoints(actualPoint) {
						if pt.Piece == nil {
							continue
						}
						if pt.Orientation == the.Orientation {
							continue
						}
						if pt.Kind == Guard {
							break RALLY_PATH_OUTER
						}
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
