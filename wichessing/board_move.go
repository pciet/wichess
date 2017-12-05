// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import (
	"fmt"
)

func (b Board) AfterMove(from AbsPoint, to AbsPoint, turn Orientation) Board {
	board := b.Copy()
	diff, _ := board.Move(from, to, turn)
	for point, _ := range diff {
		board.Points[point.Index()] = *point
	}
	board.PreviousFrom = from
	board.PreviousTo = to
	return board
}

// An empty PointSet return indicates no changes to the board - an invalid move. Any taken pieces are returned as the second value.
// The board itself is not returned so no modifications are made to the receiver Board.
func (b Board) Move(from AbsPoint, to AbsPoint, turn Orientation) (PointSet, map[AbsPoint]*Piece) {
	fromPoint := b.Points[from.Index()]
	if fromPoint.Piece == nil {
		if debug {
			fmt.Println("b.Move: fromPoint.Piece == nil")
		}
		return PointSet{}, map[AbsPoint]*Piece{}
	}
	if fromPoint.Orientation != turn {
		if debug {
			fmt.Println("b.Move: fromPoint has wrong orientation")
		}
		return PointSet{}, map[AbsPoint]*Piece{}
	}
	promoting, _ := b.HasPawnToPromote()
	if promoting {
		if debug {
			fmt.Println("b.Move: pawn to promote")
		}
		return PointSet{}, map[AbsPoint]*Piece{}
	}
	toPoint := b.Points[to.Index()]
	if toPoint.Piece != nil {
		if (toPoint.Orientation == turn) && (fromPoint.Swaps == false) {
			if debug {
				fmt.Println("b.Move: moving onto friendly piece")
			}
			return PointSet{}, map[AbsPoint]*Piece{}
		}
	}
	for pt, _ := range b.SurroundingPoints(fromPoint) {
		if pt.Piece != nil {
			if (pt.Orientation != fromPoint.Orientation) && pt.Locks {
				if debug {
					fmt.Println("b.Move: piece is locked")
				}
				return PointSet{}, map[AbsPoint]*Piece{}
			}
		}
	}
	set := make(PointSet)
	pieceset := make(map[AbsPoint]*Piece)
	// this check is here so detonations happen before guard chaining occurs
	if toPoint.Piece != nil {
		if toPoint.Detonates && (toPoint.Orientation != turn) {
			set.SetPointPiece(from, nil)
			pieceset[from] = b.Points[from.Index()].Piece
			dset := b.DetonationsFrom(to, nil)
			if len(dset) > 0 {
				for pt, _ := range dset {
					if (pt.File == from.File) && (pt.Rank == from.Rank) {
						continue
					}
					set.SetPointPiece(*pt, nil)
					pieceset[*pt] = b.Points[(*pt).Index()].Piece
				}
				return set, pieceset
			}
		}
	}
	for pt, _ := range b.SurroundingPoints(toPoint) {
		if pt.Piece != nil {
			if (pt.Orientation != fromPoint.Orientation) && pt.Guards {
				if (pt.Base == Pawn) && fromPoint.Fortified {
					continue
				}
				set.SetPointPiece(fromPoint.AbsPoint, nil)
				set.SetPointPiece(pt.AbsPoint, nil)
				pieceset[from] = b.Points[from.Index()].Piece
				if fromPoint.Detonates {
					board := b.Copy()
					board.Points[fromPoint.AbsPoint.Index()].Piece = nil
					board.Points[pt.AbsPoint.Index()].Piece = nil
					board.Points[to.Index()].Piece = fromPoint.Piece
					dset := board.DetonationsFrom(to, nil)
					if len(dset) > 0 {
						for dpt, _ := range dset {
							if (dpt.File == to.File) && (dpt.Rank == to.Rank) {
								continue
							}
							set.SetPointPiece(*dpt, nil)
							pieceset[*dpt] = b.Points[(*dpt).Index()].Piece
						}
					}
				} else {
					// a detonator taken by a guard will remove all other possible chaining guards
					// but in these other cases we need to chain all of the other possible guards
					board := b.Copy()
					board.Points[fromPoint.AbsPoint.Index()].Piece = nil
					board.Points[pt.AbsPoint.Index()].Piece = nil
					board.Points[to.Index()].Piece = pt.Piece
					// in the diff we'll at least have the first piece here, maybe updated if another opponent guard is around
					set.SetPointPiece(to, pt.Piece)
					pieceset[from] = b.Points[from.Index()].Piece
					if b.Points[to.Index()].Piece != nil {
						pieceset[to] = b.Points[to.Index()].Piece
					}
				ALLGUARDS:
					for {
						for gpt, _ := range board.SurroundingPoints(toPoint) {
							if gpt.Piece != nil {
								if (gpt.Orientation != board.Points[to.Index()].Piece.Orientation) && gpt.Guards {
									set.SetPointPiece(gpt.AbsPoint, nil)
									set.SetPointPiece(to, gpt.Piece)
									pieceset[to] = board.Points[to.Index()].Piece
									board.Points[gpt.AbsPoint.Index()].Piece = nil
									board.Points[to.Index()].Piece = gpt.Piece
									continue ALLGUARDS
								}
							}
						}
						break
					}
				}
				return set, pieceset
			}
		}
	}
	if b.MovesFromPoint(fromPoint).Has(to) == false {
		if debug {
			fmt.Println("b.Move: no moves from")
		}
		return PointSet{}, pieceset
	}
	// en passant / in passing
	// previous points are marked as AbsPoint{0, 8} if the game has no moves made yet
	if (fromPoint.Base == Pawn) && (toPoint.Piece == nil) && (b.PreviousFrom.Rank != 8) && (b.PreviousTo.Rank != 8) {
		var piece *Piece
		var expectedToRank uint8
		if turn == Black {
			piece = b.Points[AbsPoint{File: to.File, Rank: 3}.Index()].Piece
			expectedToRank = 2
		} else {
			piece = b.Points[AbsPoint{File: to.File, Rank: 4}.Index()].Piece
			expectedToRank = 5
		}
		if (piece != nil) && (to.Rank == expectedToRank) {
			if (piece.Orientation != turn) && (piece.Base == Pawn) {
				if turn == Black {
					if (b.PreviousFrom == AbsPoint{File: to.File, Rank: 1}) && (b.PreviousTo == AbsPoint{File: to.File, Rank: 3}) {
						set.SetPointPiece(AbsPoint{to.File, 3}, nil)
						set.SetPointPiece(from, nil)
						set.SetPointPiece(to, fromPoint.Piece)
						pieceset[to] = b.Points[to.Index()].Piece
						return set, pieceset
					}
				} else {
					if (b.PreviousFrom == AbsPoint{File: to.File, Rank: 6}) && (b.PreviousTo == AbsPoint{File: to.File, Rank: 4}) {
						set.SetPointPiece(AbsPoint{to.File, 4}, nil)
						set.SetPointPiece(from, nil)
						set.SetPointPiece(to, fromPoint.Piece)
						pieceset[to] = b.Points[to.Index()].Piece
						return set, pieceset
					}
				}
			}
		}
	}
	// castling can only be done when in-between points are unoccupied and check isn't entered, as verified by MovesFromPoint above
	if (fromPoint.Base == King) && (fromPoint.Moved == false) {
		// these to moves are only available when castling is available
		if (to.File == 2) && (to.Rank == 0) {
			set.SetPointPiece(from, nil)
			set.SetPointPiece(AbsPoint{0, 0}, nil)
			set.SetPointPiece(to, fromPoint.Piece)
			fromPoint.Moved = true
			set.SetPointPiece(AbsPoint{3, 0}, b.Points[0].Piece)
			b.Points[0].Moved = true
			return set, pieceset
		} else if (to.File == 6) && (to.Rank == 0) {
			set.SetPointPiece(from, nil)
			set.SetPointPiece(AbsPoint{7, 0}, nil)
			set.SetPointPiece(to, fromPoint.Piece)
			fromPoint.Moved = true
			set.SetPointPiece(AbsPoint{5, 0}, b.Points[7].Piece)
			b.Points[7].Moved = true
			return set, pieceset
		} else if (to.File == 2) && (to.Rank == 7) {
			set.SetPointPiece(from, nil)
			set.SetPointPiece(AbsPoint{0, 7}, nil)
			set.SetPointPiece(to, fromPoint.Piece)
			fromPoint.Moved = true
			set.SetPointPiece(AbsPoint{3, 7}, b.Points[56].Piece)
			b.Points[56].Moved = true
			return set, pieceset
		} else if (to.File == 6) && (to.Rank == 7) {
			set.SetPointPiece(from, nil)
			set.SetPointPiece(AbsPoint{7, 7}, nil)
			set.SetPointPiece(to, fromPoint.Piece)
			fromPoint.Moved = true
			set.SetPointPiece(AbsPoint{5, 7}, b.Points[63].Piece)
			b.Points[63].Moved = true
			return set, pieceset
		}
	}
	if toPoint.Piece != nil {
		if fromPoint.Swaps && (toPoint.Orientation == turn) {
			set.SetPointPiece(from, toPoint.Piece)
			// TODO: no side effects on the receiver
			toPoint.Moved = true
			set.SetPointPiece(to, fromPoint.Piece)
			fromPoint.Moved = true
			return set, pieceset
		}
	}
	set.SetPointPiece(from, nil)
	fromPoint.Moved = true
	set.SetPointPiece(to, fromPoint.Piece)
	if toPoint.Piece != nil {
		pieceset[to] = toPoint.Piece
	}
	return set, pieceset
}
