// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

func (b Board) AfterMove(from AbsPoint, to AbsPoint, turn Orientation) Board {
	board := b.Copy()
	for point, _ := range board.Move(from, to, turn) {
		board[point.Index()] = *point
	}
	return board
}

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
	// castling can only be done when in-between points are unoccupied and check isn't entered, as verified by MovesFromPoint above
	if (fromPoint.Kind == King) && (fromPoint.Moved == false) {
		// these to moves are only available when castling is available
		if (to.File == 2) && (to.Rank == 0) {
			set[&Point{
				AbsPoint: from,
			}] = struct{}{}
			set[&Point{
				AbsPoint: AbsPoint{
					File: 0,
					Rank: 0,
				},
			}] = struct{}{}
			set[&Point{
				AbsPoint: to,
				Piece:    fromPoint.Piece,
			}] = struct{}{}
			fromPoint.Moved = true
			set[&Point{
				AbsPoint: AbsPoint{
					File: 3,
					Rank: 0,
				},
				Piece: b[0].Piece,
			}] = struct{}{}
			b[0].Moved = true
			return set
		} else if (to.File == 6) && (to.Rank == 0) {
			set[&Point{
				AbsPoint: from,
			}] = struct{}{}
			set[&Point{
				AbsPoint: AbsPoint{
					File: 7,
					Rank: 0,
				},
			}] = struct{}{}
			set[&Point{
				AbsPoint: to,
				Piece:    fromPoint.Piece,
			}] = struct{}{}
			fromPoint.Moved = true
			set[&Point{
				AbsPoint: AbsPoint{
					File: 5,
					Rank: 0,
				},
				Piece: b[7].Piece,
			}] = struct{}{}
			b[7].Moved = true
			return set
		} else if (to.File == 2) && (to.Rank == 7) {
			set[&Point{
				AbsPoint: from,
			}] = struct{}{}
			set[&Point{
				AbsPoint: AbsPoint{
					File: 0,
					Rank: 7,
				},
			}] = struct{}{}
			set[&Point{
				AbsPoint: to,
				Piece:    fromPoint.Piece,
			}] = struct{}{}
			fromPoint.Moved = true
			set[&Point{
				AbsPoint: AbsPoint{
					File: 3,
					Rank: 7,
				},
				Piece: b[56].Piece,
			}] = struct{}{}
			b[56].Moved = true
			return set
		} else if (to.File == 6) && (to.Rank == 7) {
			set[&Point{
				AbsPoint: from,
			}] = struct{}{}
			set[&Point{
				AbsPoint: AbsPoint{
					File: 7,
					Rank: 7,
				},
			}] = struct{}{}
			set[&Point{
				AbsPoint: to,
				Piece:    fromPoint.Piece,
			}] = struct{}{}
			fromPoint.Moved = true
			set[&Point{
				AbsPoint: AbsPoint{
					File: 6,
					Rank: 7,
				},
				Piece: b[63].Piece,
			}] = struct{}{}
			b[63].Moved = true
			return set
		}
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
			// TODO: no side effects on the receiver
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
