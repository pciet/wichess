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
		piece := b[point.Index()].Piece
		// TODO: this remove castle moves logic is duplicated in Board.CheckMoves
		removeLeftCastle := false
		removeRightCastle := false
		if (piece.Kind == King) && (piece.Moved == false) {
			// are the skipped king squares threatened?
			if piece.Orientation == White {
				if b.AfterMove(point.AbsPoint, AbsPoint{3, 0}, piece.Orientation).Check(piece.Orientation) {
					removeLeftCastle = true
				}
				if b.AfterMove(point.AbsPoint, AbsPoint{5, 0}, active).Check(active) {
					removeRightCastle = true
				}
			} else {
				if b.AfterMove(point.AbsPoint, AbsPoint{3, 7}, piece.Orientation).Check(piece.Orientation) {
					removeLeftCastle = true
				}
				if b.AfterMove(point.AbsPoint, AbsPoint{5, 7}, piece.Orientation).Check(piece.Orientation) {
					removeRightCastle = true
				}
			}
		}
		for move, _ := range moves {
			if b.AfterMove(point.AbsPoint, *move, piece.Orientation).Check(piece.Orientation) {
				delete(moves, move)
			}
			if removeLeftCastle {
				if piece.Orientation == White {
					if (move.File == 2) && (move.Rank == 0) {
						delete(moves, move)
					}
				} else {
					if (move.File == 2) && (move.Rank == 7) {
						delete(moves, move)
					}
				}
			}
			if removeRightCastle {
				if piece.Orientation == White {
					if (move.File == 6) && (move.Rank == 0) {
						delete(moves, move)
					}
				} else {
					if (move.File == 6) && (move.Rank == 7) {
						delete(moves, move)
					}
				}
			}
		}
		if len(moves) > 0 {
			sets[point.AbsPoint] = moves
		}
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
		piece := b[pt.Index()]
		removeCastling := false
		removeLeftCastle := false
		removeRightCastle := false
		if (piece.Kind == King) && (piece.Moved == false) {
			if b.Check(active) {
				removeCastling = true
			} else {
				// are the skipped king squares threatened?
				if active == White {
					if b.AfterMove(pt, AbsPoint{3, 0}, active).Check(active) {
						removeLeftCastle = true
					} else if b.AfterMove(pt, AbsPoint{5, 0}, active).Check(active) {
						removeRightCastle = true
					}
				} else {
					if b.AfterMove(pt, AbsPoint{3, 7}, active).Check(active) {
						removeLeftCastle = true
					} else if b.AfterMove(pt, AbsPoint{5, 7}, active).Check(active) {
						removeRightCastle = true
					}
				}
			}
		}
		for mv, _ := range set {
			if removeCastling || (removeLeftCastle && removeRightCastle) {
				if active == White {
					if ((mv.File == 2) && (mv.Rank == 0)) || ((mv.File == 6) && (mv.Rank == 0)) {
						continue
					}
				} else {
					if ((mv.File == 2) && (mv.Rank == 7)) || ((mv.File == 6) && (mv.Rank == 7)) {
						continue
					}
				}
			} else if removeLeftCastle {
				if active == White {
					if (mv.File == 2) && (mv.Rank == 0) {
						continue
					}
				} else {
					if (mv.File == 2) && (mv.Rank == 7) {
						continue
					}
				}
			} else if removeRightCastle {
				if active == White {
					if (mv.File == 6) && (mv.Rank == 0) {
						continue
					}
				} else {
					if (mv.File == 6) && (mv.Rank == 7) {
						continue
					}
				}
			}
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

// Naive refers to moves that may put the King in check. The caller must filter those points that do put the King in check out from legal moves.
func (b Board) AllNaiveMovesFor(player Orientation) map[AbsPoint]AbsPointSet {
	moves := make(map[AbsPoint]AbsPointSet)
	for _, pt := range b {
		if pt.Piece == nil {
			continue
		}
		if pt.Piece.Orientation != player {
			continue
		}
		mfp := b.MovesFromPoint(pt)
		if len(mfp) > 0 {
			moves[pt.AbsPoint] = mfp
		}
	}
	return moves
}

func (b Board) MovesFromPoint(the Point) AbsPointSet {
	if the.Piece == nil {
		panic(fmt.Sprintf("wichessing: point (%v,%v) without piece", the.File, the.Rank))
	}
	for pt, _ := range b.SurroundingPoints(the) {
		if pt.Piece != nil {
			if (pt.Orientation != the.Orientation) && (pt.Locks) {
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
	if (the.Base == King) && (the.Moved == false) {
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
	set = set.Add(b.EnPassantTakeFromPoint(the))
	set = set.Reduce()
	return set
}

func (b Board) EnPassantTakeFromPoint(the Point) AbsPointSet {
	if the.Piece == nil {
		return AbsPointSet{}
	}
	if the.Base != Pawn {
		return AbsPointSet{}
	}
	if ((the.Orientation == White) && (the.Rank != 4)) || ((the.Orientation == Black) && (the.Rank != 3)) {
		return AbsPointSet{}
	}
	set := make(AbsPointSet)
	if the.Orientation == White {
		file := int(the.File) + 1
		if file < 8 {
			index := AbsPoint{
				File: uint8(file),
				Rank: 4,
			}.Index()
			piece := b[index].Piece
			if piece != nil {
				if (piece.Orientation != the.Orientation) && (piece.Base == Pawn) && (piece.Previous == AbsPoint{File: uint8(file), Rank: 6}.Index()) {
					set[&AbsPoint{
						File: uint8(file),
						Rank: 5,
					}] = struct{}{}
				}
			}
		}
		file = int(the.File) - 1
		if file >= 0 {
			index := AbsPoint{
				File: uint8(file),
				Rank: 4,
			}.Index()
			piece := b[index].Piece
			if piece != nil {
				if (piece.Orientation != the.Orientation) && (piece.Base == Pawn) && (piece.Previous == AbsPoint{File: uint8(file), Rank: 6}.Index()) {
					set[&AbsPoint{
						File: uint8(file),
						Rank: 5,
					}] = struct{}{}
				}
			}
		}
	} else {
		file := int(the.File) + 1
		if file < 8 {
			index := AbsPoint{
				File: uint8(file),
				Rank: 3,
			}.Index()
			piece := b[index].Piece
			if piece != nil {
				if (piece.Orientation != the.Orientation) && (piece.Base == Pawn) && (piece.Previous == AbsPoint{File: uint8(file), Rank: 1}.Index()) {
					set[&AbsPoint{
						File: uint8(file),
						Rank: 2,
					}] = struct{}{}
				}
			}
		}
		file = int(the.File) - 1
		if file >= 0 {
			index := AbsPoint{
				File: uint8(file),
				Rank: 3,
			}.Index()
			piece := b[index].Piece
			if piece != nil {
				if (piece.Orientation != the.Orientation) && (piece.Base == Pawn) && (piece.Previous == AbsPoint{File: uint8(file), Rank: 1}.Index()) {
					set[&AbsPoint{
						File: uint8(file),
						Rank: 2,
					}] = struct{}{}
				}
			}
		}
	}
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
				if actualPoint.Piece != nil {
					if the.MustEnd && (len(path.Points) != i+1) {
						if the.Ghost {
							continue
						} else {
							break
						}
					}
					if (actualPoint.Orientation == the.Orientation) && the.Swaps {
						filteredPath.Points = append(filteredPath.Points, point)
					}
					if the.Ghost {
						continue
					} else {
						break
					}
				}
				if the.MustEnd && (len(path.Points) != i+1) {
					continue
				}
				if the.Base == King {
					for pt, _ := range b.SurroundingPoints(actualPoint) {
						if pt.Piece == nil {
							continue
						}
						if pt.Orientation == the.Orientation {
							continue
						}
						if pt.Guards {
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
				if actualPoint.Piece != nil {
					if the.MustEnd && (len(path.Points) != i+1) {
						if the.Ghost {
							continue
						} else {
							break
						}
					}
					if (actualPoint.Orientation == the.Orientation) && the.Swaps {
						filteredPath.Points = append(filteredPath.Points, point)
					}
					if the.Ghost {
						continue
					} else {
						break
					}
				}
				if the.MustEnd && (len(path.Points) != i+1) {
					continue
				}
				if the.Base == King {
					for pt, _ := range b.SurroundingPoints(actualPoint) {
						if pt.Piece == nil {
							continue
						}
						if pt.Orientation == the.Orientation {
							continue
						}
						if pt.Guards {
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
					if (the.Base == Pawn) && actualPoint.Fortified {
						break
					}
					if actualPoint.Detonates {
						if the.Base == King {
							break
						}
						for pt, _ := range b.SurroundingPoints(actualPoint) {
							if pt.Piece == nil {
								continue
							}
							if pt.Orientation != the.Orientation {
								continue
							}
							if pt.Base == King {
								break TAKE_PATH_OUTER
							}
						}
					}
					if the.Base == King {
						for pt, _ := range b.SurroundingPoints(actualPoint) {
							if pt.Piece == nil {
								continue
							}
							if pt.Orientation == the.Orientation {
								continue
							}
							if pt.Guards {
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
				if actualPoint.Piece != nil {
					if the.MustEnd && (len(path.Points) != i+1) {
						if the.Ghost {
							continue
						} else {
							break
						}
					}
					if (actualPoint.Orientation == the.Orientation) && the.Swaps {
						filteredPath.Points = append(filteredPath.Points, point)
					}
					if the.Ghost {
						continue
					} else {
						break
					}
				}
				if the.MustEnd && (len(path.Points) != i+1) {
					continue
				}
				if the.Base == King {
					for pt, _ := range b.SurroundingPoints(actualPoint) {
						if pt.Piece == nil {
							continue
						}
						if pt.Orientation == the.Orientation {
							continue
						}
						if pt.Guards {
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

// TODO: refactor DetonationsFrom to something clearer than this recursive implementation

// if you know of no detonated points yet then provide nil for detonated
func (b Board) DetonationsFrom(the AbsPoint, detonated AbsPointSet) AbsPointSet {
	var set AbsPointSet
	if detonated == nil {
		set = make(AbsPointSet)
	} else {
		set = detonated
	}
	index := the.Index()
	if b[index].Piece == nil {
		return set
	}
	set[&the] = struct{}{}
	if b[index].Detonates == false {
		return set
	}
	for pt, _ := range b.SurroundingPoints(b[index]) {
		if set.Has(pt.AbsPoint) {
			continue
		}
		set = set.Add(b.DetonationsFrom(pt.AbsPoint, set))
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
		if the.Orientation == White {
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
