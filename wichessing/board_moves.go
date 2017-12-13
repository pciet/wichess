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
	for _, point := range b.Points {
		if point.Piece == nil {
			continue
		}
		moves := b.MovesFromPoint(point)
		if len(moves) == 0 {
			continue
		}
		piece := b.Points[point.Index()].Piece
		// TODO: this remove castle moves logic is duplicated in Board.CheckMoves
		removeLeftCastle := false
		removeRightCastle := false
		if (piece.Kind == King) && (piece.Moved == false) {
			// are the skipped king squares threatened?
			// if we have the castling move then the intermediate squares are empty
			if piece.Orientation == White {
				board := b.Copy()
				if moves.Has(AbsPoint{2, 0}) {
					board.Points[AbsPoint{3, 0}.Index()].Piece = piece
					board.Points[AbsPoint{4, 0}.Index()].Piece = nil
					if board.Check(piece.Orientation) {
						removeLeftCastle = true
					}
				}
				if moves.Has(AbsPoint{6, 0}) {
					board.Points[AbsPoint{3, 0}.Index()].Piece = nil
					board.Points[AbsPoint{5, 0}.Index()].Piece = piece
					if board.Check(active) {
						removeRightCastle = true
					}
				}
			} else {
				board := b.Copy()
				if moves.Has(AbsPoint{2, 7}) {
					board.Points[AbsPoint{3, 7}.Index()].Piece = piece
					board.Points[AbsPoint{4, 7}.Index()].Piece = nil
					if board.Check(piece.Orientation) {
						removeLeftCastle = true
					}
				}
				if moves.Has(AbsPoint{6, 7}) {
					board.Points[AbsPoint{3, 7}.Index()].Piece = nil
					board.Points[AbsPoint{5, 7}.Index()].Piece = piece
					if board.Check(piece.Orientation) {
						removeRightCastle = true
					}
				}
			}
		}
		for _, move := range moves {
			if b.AfterMove(point.AbsPoint, move, piece.Orientation).Check(piece.Orientation) {
				moves = moves.Remove(move)
			}
			if removeLeftCastle {
				if piece.Orientation == White {
					if (move.File == 2) && (move.Rank == 0) {
						moves = moves.Remove(move)
					}
				} else {
					if (move.File == 2) && (move.Rank == 7) {
						moves = moves.Remove(move)
					}
				}
			}
			if removeRightCastle {
				if piece.Orientation == White {
					if (move.File == 6) && (move.Rank == 0) {
						moves = moves.Remove(move)
					}
				} else {
					if (move.File == 6) && (move.Rank == 7) {
						moves = moves.Remove(move)
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
	for _, point := range b.Points {
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
	for pt, set := range moves {
		for _, mv := range set {
			if b.AfterMove(pt, mv, b.Points[pt.Index()].Orientation).Check(b.Points[pt.Index()].Orientation) {
				set = set.Remove(mv)
				if len(set) == 0 {
					delete(moves, pt)
				}
			}
		}
	}
	for pt, set := range unfiltered {
		allowed := make(AbsPointSet, 0, 8)
		piece := b.Points[pt.Index()].Piece
		removeCastling := false
		removeLeftCastle := false
		removeRightCastle := false
		if (piece.Kind == King) && (piece.Moved == false) {
			if b.Check(active) {
				removeCastling = true
			} else {
				// are the skipped king squares threatened?
				// if we have the castling move then the intermediate squares are empty
				if piece.Orientation == White {
					board := b.Copy()
					if set.Has(AbsPoint{2, 0}) {
						board.Points[AbsPoint{3, 0}.Index()].Piece = piece
						board.Points[AbsPoint{4, 0}.Index()].Piece = nil
						if board.Check(piece.Orientation) {
							removeLeftCastle = true
						}
					}
					if set.Has(AbsPoint{6, 0}) {
						board.Points[AbsPoint{3, 0}.Index()].Piece = nil
						board.Points[AbsPoint{5, 0}.Index()].Piece = piece
						if board.Check(active) {
							removeRightCastle = true
						}
					}
				} else {
					board := b.Copy()
					if set.Has(AbsPoint{2, 7}) {
						board.Points[AbsPoint{3, 7}.Index()].Piece = piece
						board.Points[AbsPoint{4, 7}.Index()].Piece = nil
						if board.Check(piece.Orientation) {
							removeLeftCastle = true
						}
					}
					if set.Has(AbsPoint{6, 7}) {
						board.Points[AbsPoint{3, 7}.Index()].Piece = nil
						board.Points[AbsPoint{5, 7}.Index()].Piece = piece
						if board.Check(piece.Orientation) {
							removeRightCastle = true
						}
					}
				}
			}
		}
		for _, mv := range set {
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
			if b.AfterMove(pt, mv, active).Check(active) {
				continue
			}
			allowed = allowed.Add(mv)
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
	for _, pt := range b.Points {
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
	for _, pt := range b.SurroundingPoints(the) {
		if pt.Piece != nil {
			if (pt.Orientation != the.Orientation) && (pt.Locks) {
				return AbsPointSet{}
			}
		}
	}
	firstSet := make(AbsPointSet, 0, 16)
	moveSet := make(AbsPointSet, 0, 16)
	takeSet := make(AbsPointSet, 0, 4)
	rallySet := make(AbsPointSet, 0, 16)
	for movetype, unfilteredpaths := range TruncatedAbsPathsForKind(the.Piece.Kind, the.AbsPoint, the.Piece.Orientation) {
		for _, path := range b.ActualPaths(the, movetype, unfilteredpaths) {
			for _, pt := range path.Points {
				switch movetype {
				case First:
					firstSet = firstSet.Add(pt)
				case Move:
					moveSet = moveSet.Add(pt)
				case Take:
					takeSet = takeSet.Add(pt)
				case RallyMove:
					rallySet = rallySet.Add(pt)
				}
			}
		}
	}
	set := make(AbsPointSet, 0, len(firstSet)+len(moveSet)+len(rallySet)+len(takeSet))
	if (len(takeSet) == 0) || (the.MustTake == false) {
		set = set.Combine(firstSet, moveSet, rallySet)
	}
	set = set.Combine(takeSet, b.ReconPointsFrom(the), b.EnPassantTakeFromPoint(the))
	// castling
	if (the.Base == King) && (the.Moved == false) {
		if (the.File == 4) && (the.Rank == 0) {
			if (b.Points[1].Piece == nil) && (b.Points[2].Piece == nil) && (b.Points[3].Piece == nil) && (b.Points[0].Piece != nil) {
				if b.Points[0].Moved == false {
					set = set.Add(AbsPoint{2, 0})
				}
			}
			if (b.Points[5].Piece == nil) && (b.Points[6].Piece == nil) && (b.Points[7].Piece != nil) {
				if b.Points[7].Moved == false {
					set = set.Add(AbsPoint{6, 0})
				}
			}
		} else if (the.File == 4) && (the.Rank == 7) {
			if (b.Points[57].Piece == nil) && (b.Points[58].Piece == nil) && (b.Points[59].Piece == nil) && (b.Points[56].Piece != nil) {
				if b.Points[56].Moved == false {
					set = set.Add(AbsPoint{2, 7})
				}
			}
			if (b.Points[61].Piece == nil) && (b.Points[62].Piece == nil) && (b.Points[63].Piece != nil) {
				if b.Points[63].Moved == false {
					set = set.Add(AbsPoint{6, 7})
				}
			}
		}
	}
	return set.Reduce()
}

func (b Board) EnPassantTakeFromPoint(the Point) AbsPointSet {
	// an fresh game starts with a previous from/to index of 64, which maps to AbsPoint{0, 8}
	if (b.PreviousFrom.Rank == 8) || (b.PreviousTo.Rank == 8) {
		return AbsPointSet{}
	}
	if the.Piece == nil {
		return AbsPointSet{}
	}
	if the.Base != Pawn {
		return AbsPointSet{}
	}
	if ((the.Orientation == White) && (the.Rank != 4)) || ((the.Orientation == Black) && (the.Rank != 3)) {
		return AbsPointSet{}
	}
	set := make(AbsPointSet, 0, 2)
	if the.Orientation == White {
		file := int(the.File) + 1
		if file < 8 {
			index := AbsPoint{
				File: uint8(file),
				Rank: 4,
			}
			piece := b.Points[index.Index()].Piece
			if piece != nil {
				if (piece.Orientation != the.Orientation) && (piece.Base == Pawn) && (b.PreviousFrom == AbsPoint{File: uint8(file), Rank: 6}) && (b.PreviousTo == index) {
					set = set.Add(AbsPoint{uint8(file), 5})
				}
			}
		}
		file = int(the.File) - 1
		if file >= 0 {
			index := AbsPoint{
				File: uint8(file),
				Rank: 4,
			}
			piece := b.Points[index.Index()].Piece
			if piece != nil {
				if (piece.Orientation != the.Orientation) && (piece.Base == Pawn) && (b.PreviousFrom == AbsPoint{File: uint8(file), Rank: 6}) && (b.PreviousTo == index) {
					set = set.Add(AbsPoint{uint8(file), 5})
				}
			}
		}
	} else {
		file := int(the.File) + 1
		if file < 8 {
			index := AbsPoint{
				File: uint8(file),
				Rank: 3,
			}
			piece := b.Points[index.Index()].Piece
			if piece != nil {
				if (piece.Orientation != the.Orientation) && (piece.Base == Pawn) && (b.PreviousFrom == AbsPoint{File: uint8(file), Rank: 1}) && (b.PreviousTo == index) {
					set = set.Add(AbsPoint{uint8(file), 2})
				}
			}
		}
		file = int(the.File) - 1
		if file >= 0 {
			index := AbsPoint{
				File: uint8(file),
				Rank: 3,
			}
			piece := b.Points[index.Index()].Piece
			if piece != nil {
				if (piece.Orientation != the.Orientation) && (piece.Base == Pawn) && (b.PreviousFrom == AbsPoint{File: uint8(file), Rank: 1}) && (b.PreviousTo == index) {
					set = set.Add(AbsPoint{uint8(file), 2})
				}
			}
		}
	}
	// remove cases where another piece is at the en passant capture point
	for _, point := range set {
		if b.Points[point.Index()].Piece != nil {
			set = set.Remove(point)
		}
	}
	return set
}

func (b Board) ActualPaths(the Point, movetype PathType, unfilteredpaths AbsPathSet) AbsPathSet {
	if the.Piece == nil {
		return AbsPathSet{}
	}
	actualPaths := make(AbsPathSet, 0, 4)
	switch movetype {
	case First:
		if the.Moved {
			break
		}
		for _, path := range unfilteredpaths {
			if path.Truncated && the.MustEnd {
				continue
			}
			filteredPath := AbsPath{
				Points: make([]AbsPoint, 0, len(path.Points)),
			}
		FIRST_PATH_OUTER:
			for i, point := range path.Points {
				actualPoint := b.Points[point.Index()]
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
					for _, pt := range b.SurroundingPoints(actualPoint) {
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
				actualPaths = actualPaths.Add(filteredPath)
			}
		}
	case Move:
		if the.Moved == false {
			break
		}
		for _, path := range unfilteredpaths {
			if path.Truncated && the.MustEnd {
				continue
			}
			filteredPath := AbsPath{
				Points: make([]AbsPoint, 0, len(path.Points)),
			}
		MOVE_PATH_OUTER:
			for i, point := range path.Points {
				actualPoint := b.Points[point.Index()]
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
					for _, pt := range b.SurroundingPoints(actualPoint) {
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
				actualPaths = actualPaths.Add(filteredPath)
			}
		}
	case Take:
		for _, path := range unfilteredpaths {
			if path.Truncated && the.MustEnd {
				continue
			}
			filteredPath := AbsPath{
				Points: make([]AbsPoint, 0, len(path.Points)),
			}
		TAKE_PATH_OUTER:
			for ind, point := range path.Points {
				actualPoint := b.Points[point.Index()]
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
						for _, pt := range b.SurroundingPoints(actualPoint) {
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
						for _, pt := range b.SurroundingPoints(actualPoint) {
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
				actualPaths = actualPaths.Add(filteredPath)
			}
		}
	case RallyMove:
		var rallied bool
		for _, point := range b.SurroundingPoints(the) {
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
		for _, path := range unfilteredpaths {
			if path.Truncated && the.MustEnd {
				continue
			}
			filteredPath := AbsPath{
				Points: make([]AbsPoint, 0, len(path.Points)),
			}
		RALLY_PATH_OUTER:
			for i, point := range path.Points {
				actualPoint := b.Points[point.Index()]
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
					for _, pt := range b.SurroundingPoints(actualPoint) {
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
				actualPaths = actualPaths.Add(filteredPath)
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
		set = make(AbsPointSet, 0, 1)
	} else {
		set = detonated
	}
	index := the.Index()
	if b.Points[index].Piece == nil {
		return set
	}
	set = set.Add(the)
	if b.Points[index].Detonates == false {
		return set
	}
	for _, pt := range b.SurroundingPoints(b.Points[index]) {
		if set.Has(pt.AbsPoint) {
			continue
		}
		set = set.Combine(b.DetonationsFrom(pt.AbsPoint, set))
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
	set := make(AbsPointSet, 0, 3)
	for f := -1; f <= 1; f++ {
		file := int(the.File) + f
		if (file < 0) || (file >= 8) {
			continue
		}
		index := IndexFromFileAndRank(uint8(file), uint8(rank))
		piece := b.Points[index].Piece
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
			nrank = int(b.Points[index].Rank) + 1
		} else {
			nrank = int(b.Points[index].Rank) - 1
		}
		if (nrank < 0) || (nrank >= 8) {
			continue
		}
		if b.Points[IndexFromFileAndRank(uint8(file), uint8(nrank))].Piece != nil {
			continue
		}
		set = set.Add(AbsPoint{b.Points[index].File, uint8(nrank)})
	}
	return set
}

func (b Board) SurroundingPoints(from Point) PointSet {
	set := make(PointSet, 0, 8)
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
			set = set.Add(b.Points[IndexFromFileAndRank(uint8(f), uint8(r))])
		}
	}
	return set
}
